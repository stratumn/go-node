package circle

import (
	"context"

	"github.com/pkg/errors"
	pb "github.com/stratumn/alice/pb/circle"

	inet "gx/ipfs/QmQm7WmgYCa4RSz76tKEYpRjApjnRw8ZTUVQC15b8JM4a2/go-libp2p-net"
	protobuf "gx/ipfs/QmRDePEiL4Yupq5EkcK3L3ko3iMgYaqUdLu7xc1kqs7dnV/go-multicodec/protobuf"
	logging "gx/ipfs/QmSpJByNKFX1sCsHBEp3R73FL4NF6FnQTEGyNAXHm2GS52/go-log"
	protocol "gx/ipfs/QmZNkThpqfVXs9GNbexPrfBbXSLNYeKrE7jwFM2oqHbyqN/go-libp2p-protocol"
	peer "gx/ipfs/Qma7H6RW8wRrfZpNSXwxYGcd1E149s42FpWNpDNieSVrnU/go-libp2p-peer"
	ihost "gx/ipfs/QmfCtHMCd9xFvehvHeVxtKVXJTMVTuHhyPRVHEXetn87vL/go-libp2p-host"
)

type host = ihost.Host

var DiscoverProtocol = protocol.ID("/alice/circle-discover/v1.0.0")
var CircleProtocol = protocol.ID("/alice/circle/v1.0.0")

type netProcess struct {
	host               host
	errorChan          chan error
	callDiscoverChan   <-chan hubCallDiscover
	callPeersChan      chan<- hubCallPeers
	msgNetToCircleChan chan<- pb.Internode
	msgCircleToNetChan <-chan pb.Internode
}

// NetProcess exports internal net process structure
type NetProcess interface {
	Run(ctx context.Context) error
	CircleHandler(stream inet.Stream)
	DiscoverHandler(stream inet.Stream)
}

// NewNetProcess creates a runnable instance of netProcess
func NewNetProcess(host host, hub *Hub) NetProcess {

	n := &netProcess{host: host, errorChan: make(chan error)}

	hub.bindNet(n)

	return n
}

// Run starts net process
func (n *netProcess) Run(ctx context.Context) error {

	for {
		select {
		case err := <-n.errorChan:
			log.Event(ctx, "errorNetProcess", logging.Metadata{
				"error": err.Error(),
			})
		case msg := <-n.callDiscoverChan:
			go n.discoverClient(ctx, msg)
		case msg := <-n.msgCircleToNetChan:
			go n.circleClient(ctx, msg)
		case <-ctx.Done():
			// TODO: close all streams?
			return nil
		}
	}
}

func (n *netProcess) closeStream(stream inet.Stream) {
	if err := stream.Close(); err != nil {
		n.errorChan <- err
	}
}

// DiscoverHandler handles discover requests
func (n *netProcess) DiscoverHandler(stream inet.Stream) {
	defer n.closeStream(stream)

	enc := protobuf.Multicodec(nil).Encoder(stream)

	peersChan := make(chan pb.Peer)
	n.callPeersChan <- hubCallPeers{
		PeersChan: peersChan,
	}

	var peers pb.Peers
	for peer := range peersChan {
		peers.Peers = append(peers.Peers, &pb.Peer{Id: peer.Id, Address: peer.Address})
	}

	err := enc.Encode(&peers)
	if err != nil {
		n.errorChan <- err
		return
	}

}

func (n *netProcess) discoverClient(ctx context.Context, callDiscover hubCallDiscover) {

	defer close(callDiscover.PeersChan)

	pid, err := peer.IDFromBytes(callDiscover.PeerID.Address)
	if err != nil {
		n.errorChan <- err
		return
	}

	stream, err := n.host.NewStream(ctx, pid, DiscoverProtocol)
	if err != nil {
		n.errorChan <- err
		return
	}
	defer n.closeStream(stream)

	dec := protobuf.Multicodec(nil).Decoder(stream)

	var peers pb.Peers

	err = dec.Decode(&peers)
	if err != nil {
		n.errorChan <- err
		return
	}

	for _, peer := range peers.Peers {
		callDiscover.PeersChan <- *peer
	}
}

// CircleHandler handles requests from other nodes
func (n *netProcess) CircleHandler(stream inet.Stream) {

	defer n.closeStream(stream)
	dec := protobuf.Multicodec(nil).Decoder(stream)

	var message pb.Internode

	err := dec.Decode(&message)
	if err != nil {
		n.errorChan <- errors.WithStack(err)
		return
	}

	n.msgNetToCircleChan <- message

}

func (n *netProcess) circleClient(ctx context.Context, msgCircle pb.Internode) {

	pid, err := peer.IDFromBytes(msgCircle.PeerId.Address)
	if err != nil {
		n.errorChan <- errors.WithStack(err)
		return
	}

	stream, err := n.host.NewStream(ctx, pid, CircleProtocol)
	if err != nil {
		n.errorChan <- errors.WithStack(err)
		return
	}
	defer n.closeStream(stream)

	enc := protobuf.Multicodec(nil).Encoder(stream)
	if err != nil {
		n.errorChan <- errors.WithStack(err)
		return
	}

	err = enc.Encode(&msgCircle)
	if err != nil {
		n.errorChan <- errors.WithStack(err)
		return
	}

}
