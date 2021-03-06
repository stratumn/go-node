// Copyright © 2017-2018 Stratumn SAS
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package gossip

import (
	"context"
	"encoding/binary"
	"sync"

	"github.com/pkg/errors"

	"github.com/stratumn/go-node/app/coin/pb"
	"github.com/stratumn/go-node/app/coin/protocol/chain"
	"github.com/stratumn/go-node/app/coin/protocol/state"
	"github.com/stratumn/go-node/app/coin/protocol/validator"

	logging "github.com/ipfs/go-log"
	host "github.com/libp2p/go-libp2p-host"
	peer "github.com/libp2p/go-libp2p-peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

const (
	// TxTopicName is the topic name for transactions.
	TxTopicName = "coin.tx"

	// BlockTopicName is the topic name for blocks.
	BlockTopicName = "coin.block"
)

// log is the logger for the coin gossip.
var log = logging.Logger("coin.gossip")

// Gossip is an interface used to gossip transactions and blocks.
type Gossip interface {
	// ListenTx passes incoming transactions to a callback.
	// Only valid transactions will be passed to the callback.
	ListenTx(ctx context.Context, processTx func(*pb.Transaction) error) error
	// ListenBlock passes incoming blocks to a callback.
	// Only valid blocks will be passed to the callback.
	// A synchronous callback is given to sync the local chain.
	ListenBlock(ctx context.Context, processBlock func(*pb.Block) error, sync func([]byte) error) error
	// PublishTx sends a new transaction to the gossip.
	PublishTx(tx *pb.Transaction) error
	// PublishBlock sends a new block to the gossip.
	PublishBlock(block *pb.Block) error

	// AddBlockListener returns a channel on which new valid
	// block headers received will be published.
	AddBlockListener() chan *pb.Header

	// Close closes the gossip layer.
	Close() error
}

// Gossip handles the gossiping of blocks and transactions.
type gossip struct {
	host      host.Host
	pubsub    *pubsub.PubSub
	state     state.Reader
	chain     chain.Reader
	validator validator.Validator

	listenersMu    sync.RWMutex
	blockListeners []chan *pb.Header

	mu   sync.RWMutex
	subs map[string]*pubsub.Subscription
}

// NewGossip returns gossip.
func NewGossip(
	h host.Host,
	p *pubsub.PubSub,
	s state.Reader,
	c chain.Reader,
	v validator.Validator,
) Gossip {
	return &gossip{
		host:      h,
		pubsub:    p,
		state:     s,
		chain:     c,
		validator: v,
		subs:      make(map[string]*pubsub.Subscription),
	}
}

// AddBlockListener returns a channel on which new valid
// block headers received will be published.
func (g *gossip) AddBlockListener() chan *pb.Header {
	g.listenersMu.Lock()
	defer g.listenersMu.Unlock()

	c := make(chan *pb.Header)
	g.blockListeners = append(g.blockListeners, c)

	return c
}

// Close closes subscriptions and listeners channels.
func (g *gossip) Close() error {
	g.mu.Lock()
	defer g.mu.Unlock()
	for _, s := range g.subs {

		err := g.pubsub.UnregisterTopicValidator(s.Topic())
		if err != nil {
			return errors.WithStack(err)
		}
		s.Cancel()
	}

	g.listenersMu.Lock()
	defer g.listenersMu.Unlock()
	for _, c := range g.blockListeners {
		close(c)
	}

	g.blockListeners = nil

	return nil
}

// ListenTx subscribes to transaction topic and listens to incoming transactions.
func (g *gossip) ListenTx(ctx context.Context, processTx func(*pb.Transaction) error) error {
	if err := g.subscribeTx(); err != nil {
		return err
	}

	return g.listen(ctx, TxTopicName, func(msg *pubsub.Message) error {
		tx := &pb.Transaction{}
		if err := tx.Unmarshal(msg.GetData()); err != nil {
			return err
		}

		log.Event(context.Background(), "TransactionReceived", &logging.Metadata{"tx": tx.Loggable()})

		return processTx(tx)
	})
}

// ListenBlock subscribes to block topic and listens to incoming blocks.
func (g *gossip) ListenBlock(ctx context.Context, processBlock func(*pb.Block) error, sync func([]byte) error) error {
	if err := g.subscribeBlock(sync); err != nil {
		return err
	}

	return g.listen(ctx, BlockTopicName, func(msg *pubsub.Message) error {
		block := &pb.Block{}
		if err := block.Unmarshal(msg.GetData()); err != nil {
			return err
		}

		log.Event(context.Background(), "BlockReceived", &logging.Metadata{"block": block.Loggable()})

		if err := processBlock(block); err != nil {
			return err
		}

		g.listenersMu.RLock()
		defer g.listenersMu.RUnlock()

		for _, c := range g.blockListeners {
			go func(c chan *pb.Header) {
				for {
					select {
					case <-ctx.Done():
						return
					case c <- block.Header:
						return
					}
				}
			}(c)
		}

		return nil
	})
}

// PublishTx publishes a transaction message.
func (g *gossip) PublishTx(tx *pb.Transaction) error {
	log.Event(context.Background(), "PublishTx", &logging.Metadata{"tx": tx.Loggable()})
	txData, err := tx.Marshal()
	if err != nil {
		return err
	}

	return g.publish(TxTopicName, txData)
}

// PublishBlock publishes a block message.
func (g *gossip) PublishBlock(block *pb.Block) error {
	log.Event(context.Background(), "PublishBlock", &logging.Metadata{"block": block.Loggable()})
	blockData, err := block.Marshal()
	if err != nil {
		return err
	}

	return g.publish(BlockTopicName, blockData)
}

// subscribeTx subscribes to the transaction topic.
func (g *gossip) subscribeTx() error {
	return g.subscribe(TxTopicName, func(ctx context.Context, peerID peer.ID, m *pubsub.Message) bool {
		tx := &pb.Transaction{}
		if err := tx.Unmarshal(m.GetData()); err != nil {
			log.Event(ctx, "InvalidTransactionFormat", logging.Metadata{"error": err.Error()})
			return false
		}

		if err := g.validator.ValidateTx(tx, g.state); err != nil {
			log.Event(ctx, "InvalidTransaction", logging.Metadata{
				"error": err.Error(),
				"tx":    tx.Loggable(),
			})
			return false
		}

		return true
	})
}

// subscribeBlock subscribes to the block topic.
func (g *gossip) subscribeBlock(sync func([]byte) error) error {
	return g.subscribe(BlockTopicName, func(ctx context.Context, peerID peer.ID, m *pubsub.Message) bool {
		block := &pb.Block{}
		if err := block.Unmarshal(m.GetData()); err != nil {
			log.Event(ctx, "InvalidBlockFormat", logging.Metadata{"error": err.Error()})
			return false
		}

		// First thing is to sync with the given branch if we don't have it.
		// TODO: we could choose to sync with it only if it is longer than ours.
		_, err := g.chain.GetBlockByHash(block.PreviousHash())
		if err == chain.ErrBlockNotFound {
			// Synchronize.
			err := sync(block.PreviousHash())
			if err != nil {
				return false
			}
		} else if err != nil {
			log.Event(ctx, "ChainUnavailable", logging.Metadata{"error": err.Error()})
			return false
		}

		if err := g.validator.ValidateBlock(block, g.state); err != nil {
			log.Event(ctx, "InvalidBlock", logging.Metadata{
				"error": err.Error(),
				"block": block.Loggable(),
			})
			return false
		}

		return true
	})
}

func (g *gossip) subscribe(topic string, validator pubsub.Validator) error {
	if g.isSubscribed(topic) {
		return nil
	}

	sub, err := g.pubsub.Subscribe(topic)
	if err != nil {
		return err
	}

	g.subs[topic] = sub

	err = g.pubsub.RegisterTopicValidator(topic, validator)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (g *gossip) listen(ctx context.Context, topic string, callback func(msg *pubsub.Message) error) error {
	if !g.isSubscribed(topic) {
		return errors.Errorf("not subscribed to %v topic", topic)
	}

	sub := g.subs[topic]

	go func() {
		msg, errIncoming := sub.Next(ctx)
		for errIncoming == nil {
			if g.host.ID() != msg.GetFrom() {
				if err := callback(msg); err != nil {
					seqno := binary.BigEndian.Uint64(msg.GetSeqno())
					log.Event(ctx, "UnableToProcessMessage", logging.Metadata{
						"topic": topic,
						"from":  msg.GetFrom().Pretty(),
						"seqno": seqno,
						"error": err.Error(),
					})
				}
			}
			msg, errIncoming = sub.Next(ctx)
		}
	}()

	return nil
}

func (g *gossip) publish(topic string, data []byte) error {
	g.mu.RLock()
	defer g.mu.RUnlock()

	err := g.pubsub.Publish(topic, data)
	if err != nil {
		log.Event(context.Background(), "UnableToPublish", logging.Metadata{"topic": topic, "error": err.Error()})
	}

	return err
}

func (g *gossip) isSubscribed(topic string) bool {
	isSubscribed := false
	topics := g.pubsub.GetTopics()

	for _, t := range topics {
		if t == topic {
			isSubscribed = true
			break
		}
	}

	return isSubscribed
}
