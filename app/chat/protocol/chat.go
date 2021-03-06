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

// Package protocol defines types for the chat protocol.
package protocol

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	pbchat "github.com/stratumn/go-node/app/chat/grpc"
	"github.com/stratumn/go-node/app/chat/pb"
	pbevent "github.com/stratumn/go-node/core/app/event/grpc"
	event "github.com/stratumn/go-node/core/app/event/service"
	"github.com/stratumn/go-node/core/monitoring"

	protobuf "github.com/multiformats/go-multicodec/protobuf"

	ihost "github.com/libp2p/go-libp2p-host"

	inet "github.com/libp2p/go-libp2p-net"
	peer "github.com/libp2p/go-libp2p-peer"
	protocol "github.com/libp2p/go-libp2p-protocol"
)

// Host represents a Stratumn Node host.
type Host = ihost.Host

// ProtocolID is the protocol ID of the service.
var ProtocolID = protocol.ID("/stratumn/node/chat/v1.0.0")

// Chat implements the chat protocol.
type Chat struct {
	host          Host
	eventEmitter  event.Emitter
	msgReceivedCh chan *pbchat.DatedMessage
}

// NewChat creates a new chat server.
func NewChat(host Host, eventEmitter event.Emitter, msgReceivedCh chan *pbchat.DatedMessage) *Chat {
	return &Chat{
		host:          host,
		eventEmitter:  eventEmitter,
		msgReceivedCh: msgReceivedCh,
	}
}

// StreamHandler handles incoming messages from peers.
func (c *Chat) StreamHandler(ctx context.Context, stream inet.Stream) {
	ctx, span := monitoring.StartSpan(ctx, "chat", "StreamHandler", monitoring.SpanOptionPeerID(stream.Conn().RemotePeer()))
	defer span.End()

	c.receive(ctx, stream)
	if err := inet.FullClose(stream); err != nil {
		span.Annotate(ctx, "close_err", err.Error())
	}
}

// receive reads a message from an incoming stream.
func (c *Chat) receive(ctx context.Context, stream inet.Stream) {
	from := stream.Conn().RemotePeer()
	ctx, span := monitoring.StartSpan(ctx, "chat", "receive", monitoring.SpanOptionPeerID(from))
	defer span.End()

	msgReceived.Record(ctx, 1)

	dec := protobuf.Multicodec(nil).Decoder(stream)
	var message pb.Message
	err := dec.Decode(&message)
	if err != nil {
		span.SetUnknownError(err)
		msgError.Record(ctx, 1)
		return
	}

	go func() {
		c.msgReceivedCh <- pbchat.NewDatedMessageReceived(from, message.Message)
	}()

	chatEvent := &pbevent.Event{
		Message: fmt.Sprintf("[%s] %s", from.Pretty(), message.Message),
		Level:   pbevent.Level_INFO,
		Topic:   fmt.Sprintf("chat.%s", from.Pretty()),
	}

	c.eventEmitter.Emit(chatEvent)
}

// Send sends a message to a peer.
func (c *Chat) Send(ctx context.Context, pid peer.ID, message string) error {
	ctx, span := monitoring.StartSpan(ctx, "chat", "Send", monitoring.SpanOptionPeerID(pid))
	defer span.End()

	msgSent.Record(ctx, 1)

	successCh := make(chan struct{}, 1)
	errCh := make(chan error, 1)

	go func() {
		stream, err := c.host.NewStream(ctx, pid, ProtocolID)
		if err != nil {
			errCh <- errors.WithStack(err)
			return
		}

		enc := protobuf.Multicodec(nil).Encoder(stream)
		err = enc.Encode(&pb.Message{Message: message})
		if err != nil {
			errCh <- errors.WithStack(err)
			return
		}

		successCh <- struct{}{}
		err = inet.FullClose(stream)
		if err != nil {
			span.Annotate(ctx, "close_err", err.Error())
		}
	}()

	select {
	case <-ctx.Done():
		return errors.WithStack(ctx.Err())
	case <-successCh:
		return nil
	case err := <-errCh:
		span.SetUnknownError(err)
		msgError.Record(ctx, 1)
		return err
	}
}
