// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/stratumn/alice/grpc/chat/chat.proto

/*
Package chat is a generated protocol buffer package.

It is generated from these files:
	github.com/stratumn/alice/grpc/chat/chat.proto

It has these top-level messages:
	MessageReq
	Ack
*/
package chat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/stratumn/alice/grpc/ext"

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The message request.
type MessageReq struct {
	PeerId  []byte `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
}

func (m *MessageReq) Reset()                    { *m = MessageReq{} }
func (m *MessageReq) String() string            { return proto.CompactTextString(m) }
func (*MessageReq) ProtoMessage()               {}
func (*MessageReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MessageReq) GetPeerId() []byte {
	if m != nil {
		return m.PeerId
	}
	return nil
}

func (m *MessageReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// An empty ack.
type Ack struct {
}

func (m *Ack) Reset()                    { *m = Ack{} }
func (m *Ack) String() string            { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()               {}
func (*Ack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*MessageReq)(nil), "stratumn.alice.grpc.chat.MessageReq")
	proto.RegisterType((*Ack)(nil), "stratumn.alice.grpc.chat.Ack")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Chat service

type ChatClient interface {
	// Sends a message to a peer.
	Message(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*Ack, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Message(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.chat.Chat/Message", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Chat service

type ChatServer interface {
	// Sends a message to a peer.
	Message(context.Context, *MessageReq) (*Ack, error)
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_Message_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Message(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.chat.Chat/Message",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Message(ctx, req.(*MessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stratumn.alice.grpc.chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Message",
			Handler:    _Chat_Message_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/stratumn/alice/grpc/chat/chat.proto",
}

func init() { proto.RegisterFile("github.com/stratumn/alice/grpc/chat/chat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x2e, 0x29, 0x4a, 0x2c, 0x29, 0xcd, 0xcd, 0xd3,
	0x4f, 0xcc, 0xc9, 0x4c, 0x4e, 0xd5, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x4f, 0xce, 0x48, 0x2c, 0x01,
	0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x12, 0x30, 0x45, 0x7a, 0x60, 0x45, 0x7a, 0x20,
	0x45, 0x7a, 0x20, 0x79, 0x29, 0x1d, 0x02, 0x26, 0xa5, 0x56, 0x94, 0x80, 0x30, 0xc4, 0x1c, 0xa5,
	0x44, 0x2e, 0x2e, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0xd4, 0xa0, 0xd4, 0x42, 0x21, 0x1d, 0x2e,
	0xf6, 0x82, 0xd4, 0xd4, 0xa2, 0xf8, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x1e, 0x27, 0xe1,
	0x45, 0xbb, 0x25, 0xd8, 0x03, 0x52, 0x53, 0x8b, 0x14, 0x3c, 0x5d, 0x56, 0xec, 0x96, 0x60, 0xfc,
	0xb0, 0x5b, 0x82, 0x31, 0x88, 0x0d, 0xa4, 0xc6, 0x33, 0x45, 0x48, 0x95, 0x8b, 0x3d, 0x39, 0x3f,
	0xaf, 0x24, 0x35, 0xaf, 0x44, 0x82, 0x49, 0x81, 0x51, 0x83, 0xd3, 0x89, 0x1b, 0xa4, 0x1a, 0x66,
	0x1c, 0x4c, 0x4e, 0x89, 0x95, 0x8b, 0xd9, 0x31, 0x39, 0xdb, 0xa8, 0x84, 0x8b, 0xc5, 0x39, 0x23,
	0xb1, 0x44, 0x28, 0x87, 0x0b, 0xa6, 0x44, 0x48, 0x45, 0x0f, 0x97, 0x2f, 0xf4, 0x10, 0x8e, 0x92,
	0x92, 0xc5, 0xad, 0xca, 0x31, 0x39, 0x5b, 0x49, 0xa6, 0x69, 0xab, 0x84, 0x44, 0x70, 0x6a, 0x5e,
	0x8a, 0x42, 0xa2, 0x42, 0x2e, 0x44, 0x97, 0x42, 0x49, 0xbe, 0x42, 0xa2, 0x02, 0xc8, 0x95, 0x4e,
	0x5a, 0x51, 0x1a, 0x44, 0x84, 0xac, 0x35, 0x88, 0x48, 0x62, 0x03, 0x07, 0x89, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x95, 0x9a, 0xc3, 0x1a, 0x8c, 0x01, 0x00, 0x00,
}
