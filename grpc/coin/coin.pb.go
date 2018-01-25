// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/stratumn/alice/grpc/coin/coin.proto

/*
Package coin is a generated protocol buffer package.

It is generated from these files:
	github.com/stratumn/alice/grpc/coin/coin.proto

It has these top-level messages:
	TransactionResp
*/
package coin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/stratumn/alice/grpc/ext"
import stratumn_alice_pb_coin "github.com/stratumn/alice/pb/coin"

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

// The response to a request to do a coin transaction.
type TransactionResp struct {
	TxHash []byte `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
}

func (m *TransactionResp) Reset()                    { *m = TransactionResp{} }
func (m *TransactionResp) String() string            { return proto.CompactTextString(m) }
func (*TransactionResp) ProtoMessage()               {}
func (*TransactionResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TransactionResp) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionResp)(nil), "stratumn.alice.grpc.coin.TransactionResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Coin service

type CoinClient interface {
	// Send a coin transaction to the consensus engine.
	Transaction(ctx context.Context, in *stratumn_alice_pb_coin.Transaction, opts ...grpc.CallOption) (*TransactionResp, error)
}

type coinClient struct {
	cc *grpc.ClientConn
}

func NewCoinClient(cc *grpc.ClientConn) CoinClient {
	return &coinClient{cc}
}

func (c *coinClient) Transaction(ctx context.Context, in *stratumn_alice_pb_coin.Transaction, opts ...grpc.CallOption) (*TransactionResp, error) {
	out := new(TransactionResp)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.coin.Coin/Transaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Coin service

type CoinServer interface {
	// Send a coin transaction to the consensus engine.
	Transaction(context.Context, *stratumn_alice_pb_coin.Transaction) (*TransactionResp, error)
}

func RegisterCoinServer(s *grpc.Server, srv CoinServer) {
	s.RegisterService(&_Coin_serviceDesc, srv)
}

func _Coin_Transaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(stratumn_alice_pb_coin.Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinServer).Transaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.coin.Coin/Transaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinServer).Transaction(ctx, req.(*stratumn_alice_pb_coin.Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _Coin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stratumn.alice.grpc.coin.Coin",
	HandlerType: (*CoinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transaction",
			Handler:    _Coin_Transaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/stratumn/alice/grpc/coin/coin.proto",
}

func init() { proto.RegisterFile("github.com/stratumn/alice/grpc/coin/coin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x2e, 0x29, 0x4a, 0x2c, 0x29, 0xcd, 0xcd, 0xd3,
	0x4f, 0xcc, 0xc9, 0x4c, 0x4e, 0xd5, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x4f, 0xce, 0xcf, 0xcc, 0x03,
	0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x12, 0x30, 0x45, 0x7a, 0x60, 0x45, 0x7a, 0x20,
	0x45, 0x7a, 0x20, 0x79, 0x29, 0x1d, 0x02, 0x26, 0xa5, 0x56, 0x94, 0x80, 0x30, 0xc4, 0x1c, 0x7c,
	0xaa, 0x0b, 0x92, 0xd0, 0x6d, 0x55, 0xaa, 0xe4, 0xe2, 0x0f, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4c,
	0x2e, 0xc9, 0xcc, 0xcf, 0x0b, 0x4a, 0x2d, 0x2e, 0x10, 0x4a, 0xe3, 0x62, 0x2f, 0xa9, 0x88, 0xcf,
	0x48, 0x2c, 0xce, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x71, 0xf2, 0x5d, 0xb4, 0x5b, 0xc2, 0xd3,
	0x23, 0xb1, 0x38, 0x43, 0x21, 0x3f, 0x4d, 0xa1, 0x24, 0x23, 0x55, 0xa1, 0x04, 0xa1, 0x43, 0x4f,
	0xc1, 0x39, 0x31, 0x4f, 0x21, 0x29, 0x55, 0xa1, 0xb4, 0x38, 0x35, 0x45, 0xa1, 0x24, 0x5f, 0xa1,
	0x28, 0x35, 0x2d, 0xb5, 0x28, 0x35, 0x2f, 0x39, 0x15, 0x5d, 0xa5, 0x7a, 0xb1, 0x42, 0x68, 0x48,
	0x84, 0xbf, 0x5e, 0x10, 0x5b, 0x49, 0x05, 0xc8, 0x30, 0xa3, 0x46, 0x46, 0x2e, 0x16, 0xe7, 0xfc,
	0xcc, 0x3c, 0xa1, 0x4a, 0x2e, 0x6e, 0x24, 0x37, 0x08, 0x29, 0xeb, 0xa1, 0x85, 0x44, 0x41, 0x12,
	0x38, 0x1c, 0xf4, 0x90, 0x14, 0x49, 0x69, 0xea, 0xe1, 0x0a, 0x2e, 0x3d, 0x34, 0xff, 0x28, 0x49,
	0x35, 0x6d, 0x95, 0x10, 0x0a, 0x4e, 0xcd, 0x4b, 0x51, 0x48, 0x44, 0x76, 0xcf, 0x84, 0xad, 0x12,
	0x8c, 0x4e, 0x5a, 0x51, 0x1a, 0x44, 0x44, 0x93, 0x35, 0x88, 0x48, 0x62, 0x03, 0x87, 0x98, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xfb, 0xe3, 0x37, 0xd9, 0x01, 0x00, 0x00,
}
