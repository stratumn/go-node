// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stratumn/alice/grpc/coin/coin.proto

/*
	Package coin is a generated protocol buffer package.

	It is generated from these files:
		github.com/stratumn/alice/grpc/coin/coin.proto

	It has these top-level messages:
		AccountReq
		BlockchainReq
		BlockchainResp
		TransactionResp
		AccountTransactionsReq
*/
package coin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/stratumn/alice/grpc/ext"
import stratumn_alice_pb_coin "github.com/stratumn/alice/pb/coin"

import context "context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request to get an account.
type AccountReq struct {
	PeerId []byte `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
}

func (m *AccountReq) Reset()                    { *m = AccountReq{} }
func (m *AccountReq) String() string            { return proto.CompactTextString(m) }
func (*AccountReq) ProtoMessage()               {}
func (*AccountReq) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{0} }

func (m *AccountReq) GetPeerId() []byte {
	if m != nil {
		return m.PeerId
	}
	return nil
}

// The request to get part of the blockchain.
type BlockchainReq struct {
	BlockNumber uint32 `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	HeaderHash  []byte `protobuf:"bytes,2,opt,name=header_hash,json=headerHash,proto3" json:"header_hash,omitempty"`
	Count       uint32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *BlockchainReq) Reset()                    { *m = BlockchainReq{} }
func (m *BlockchainReq) String() string            { return proto.CompactTextString(m) }
func (*BlockchainReq) ProtoMessage()               {}
func (*BlockchainReq) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{1} }

func (m *BlockchainReq) GetBlockNumber() uint32 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *BlockchainReq) GetHeaderHash() []byte {
	if m != nil {
		return m.HeaderHash
	}
	return nil
}

func (m *BlockchainReq) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

// The response to a request for blockchain blocks.
type BlockchainResp struct {
	Blocks []*stratumn_alice_pb_coin.Block `protobuf:"bytes,1,rep,name=blocks" json:"blocks,omitempty"`
}

func (m *BlockchainResp) Reset()                    { *m = BlockchainResp{} }
func (m *BlockchainResp) String() string            { return proto.CompactTextString(m) }
func (*BlockchainResp) ProtoMessage()               {}
func (*BlockchainResp) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{2} }

func (m *BlockchainResp) GetBlocks() []*stratumn_alice_pb_coin.Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

// The response to a request to do a coin transaction.
type TransactionResp struct {
	TxHash []byte `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
}

func (m *TransactionResp) Reset()                    { *m = TransactionResp{} }
func (m *TransactionResp) String() string            { return proto.CompactTextString(m) }
func (*TransactionResp) ProtoMessage()               {}
func (*TransactionResp) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{3} }

func (m *TransactionResp) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

// The request to get an account transaction history.
type AccountTransactionsReq struct {
	PeerId []byte `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
}

func (m *AccountTransactionsReq) Reset()                    { *m = AccountTransactionsReq{} }
func (m *AccountTransactionsReq) String() string            { return proto.CompactTextString(m) }
func (*AccountTransactionsReq) ProtoMessage()               {}
func (*AccountTransactionsReq) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{4} }

func (m *AccountTransactionsReq) GetPeerId() []byte {
	if m != nil {
		return m.PeerId
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountReq)(nil), "stratumn.alice.grpc.coin.AccountReq")
	proto.RegisterType((*BlockchainReq)(nil), "stratumn.alice.grpc.coin.BlockchainReq")
	proto.RegisterType((*BlockchainResp)(nil), "stratumn.alice.grpc.coin.BlockchainResp")
	proto.RegisterType((*TransactionResp)(nil), "stratumn.alice.grpc.coin.TransactionResp")
	proto.RegisterType((*AccountTransactionsReq)(nil), "stratumn.alice.grpc.coin.AccountTransactionsReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Coin service

type CoinClient interface {
	// Get an account.
	Account(ctx context.Context, in *AccountReq, opts ...grpc.CallOption) (*stratumn_alice_pb_coin.Account, error)
	// Send a coin transaction to the consensus engine.
	Transaction(ctx context.Context, in *stratumn_alice_pb_coin.Transaction, opts ...grpc.CallOption) (*TransactionResp, error)
	// Get all the past transactions of an account.
	AccountTransactions(ctx context.Context, in *AccountTransactionsReq, opts ...grpc.CallOption) (Coin_AccountTransactionsClient, error)
	// Get blocks from the blockchain.
	Blockchain(ctx context.Context, in *BlockchainReq, opts ...grpc.CallOption) (*BlockchainResp, error)
}

type coinClient struct {
	cc *grpc.ClientConn
}

func NewCoinClient(cc *grpc.ClientConn) CoinClient {
	return &coinClient{cc}
}

func (c *coinClient) Account(ctx context.Context, in *AccountReq, opts ...grpc.CallOption) (*stratumn_alice_pb_coin.Account, error) {
	out := new(stratumn_alice_pb_coin.Account)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.coin.Coin/Account", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinClient) Transaction(ctx context.Context, in *stratumn_alice_pb_coin.Transaction, opts ...grpc.CallOption) (*TransactionResp, error) {
	out := new(TransactionResp)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.coin.Coin/Transaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinClient) AccountTransactions(ctx context.Context, in *AccountTransactionsReq, opts ...grpc.CallOption) (Coin_AccountTransactionsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Coin_serviceDesc.Streams[0], c.cc, "/stratumn.alice.grpc.coin.Coin/AccountTransactions", opts...)
	if err != nil {
		return nil, err
	}
	x := &coinAccountTransactionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Coin_AccountTransactionsClient interface {
	Recv() (*stratumn_alice_pb_coin.Transaction, error)
	grpc.ClientStream
}

type coinAccountTransactionsClient struct {
	grpc.ClientStream
}

func (x *coinAccountTransactionsClient) Recv() (*stratumn_alice_pb_coin.Transaction, error) {
	m := new(stratumn_alice_pb_coin.Transaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coinClient) Blockchain(ctx context.Context, in *BlockchainReq, opts ...grpc.CallOption) (*BlockchainResp, error) {
	out := new(BlockchainResp)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.coin.Coin/Blockchain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Coin service

type CoinServer interface {
	// Get an account.
	Account(context.Context, *AccountReq) (*stratumn_alice_pb_coin.Account, error)
	// Send a coin transaction to the consensus engine.
	Transaction(context.Context, *stratumn_alice_pb_coin.Transaction) (*TransactionResp, error)
	// Get all the past transactions of an account.
	AccountTransactions(*AccountTransactionsReq, Coin_AccountTransactionsServer) error
	// Get blocks from the blockchain.
	Blockchain(context.Context, *BlockchainReq) (*BlockchainResp, error)
}

func RegisterCoinServer(s *grpc.Server, srv CoinServer) {
	s.RegisterService(&_Coin_serviceDesc, srv)
}

func _Coin_Account_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinServer).Account(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.coin.Coin/Account",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinServer).Account(ctx, req.(*AccountReq))
	}
	return interceptor(ctx, in, info, handler)
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

func _Coin_AccountTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AccountTransactionsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CoinServer).AccountTransactions(m, &coinAccountTransactionsServer{stream})
}

type Coin_AccountTransactionsServer interface {
	Send(*stratumn_alice_pb_coin.Transaction) error
	grpc.ServerStream
}

type coinAccountTransactionsServer struct {
	grpc.ServerStream
}

func (x *coinAccountTransactionsServer) Send(m *stratumn_alice_pb_coin.Transaction) error {
	return x.ServerStream.SendMsg(m)
}

func _Coin_Blockchain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockchainReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinServer).Blockchain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.coin.Coin/Blockchain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinServer).Blockchain(ctx, req.(*BlockchainReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Coin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stratumn.alice.grpc.coin.Coin",
	HandlerType: (*CoinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Account",
			Handler:    _Coin_Account_Handler,
		},
		{
			MethodName: "Transaction",
			Handler:    _Coin_Transaction_Handler,
		},
		{
			MethodName: "Blockchain",
			Handler:    _Coin_Blockchain_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AccountTransactions",
			Handler:       _Coin_AccountTransactions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/stratumn/alice/grpc/coin/coin.proto",
}

func (m *AccountReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PeerId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCoin(dAtA, i, uint64(len(m.PeerId)))
		i += copy(dAtA[i:], m.PeerId)
	}
	return i, nil
}

func (m *BlockchainReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockchainReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BlockNumber != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCoin(dAtA, i, uint64(m.BlockNumber))
	}
	if len(m.HeaderHash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCoin(dAtA, i, uint64(len(m.HeaderHash)))
		i += copy(dAtA[i:], m.HeaderHash)
	}
	if m.Count != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCoin(dAtA, i, uint64(m.Count))
	}
	return i, nil
}

func (m *BlockchainResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockchainResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Blocks) > 0 {
		for _, msg := range m.Blocks {
			dAtA[i] = 0xa
			i++
			i = encodeVarintCoin(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *TransactionResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TxHash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCoin(dAtA, i, uint64(len(m.TxHash)))
		i += copy(dAtA[i:], m.TxHash)
	}
	return i, nil
}

func (m *AccountTransactionsReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountTransactionsReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PeerId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCoin(dAtA, i, uint64(len(m.PeerId)))
		i += copy(dAtA[i:], m.PeerId)
	}
	return i, nil
}

func encodeVarintCoin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AccountReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovCoin(uint64(l))
	}
	return n
}

func (m *BlockchainReq) Size() (n int) {
	var l int
	_ = l
	if m.BlockNumber != 0 {
		n += 1 + sovCoin(uint64(m.BlockNumber))
	}
	l = len(m.HeaderHash)
	if l > 0 {
		n += 1 + l + sovCoin(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovCoin(uint64(m.Count))
	}
	return n
}

func (m *BlockchainResp) Size() (n int) {
	var l int
	_ = l
	if len(m.Blocks) > 0 {
		for _, e := range m.Blocks {
			l = e.Size()
			n += 1 + l + sovCoin(uint64(l))
		}
	}
	return n
}

func (m *TransactionResp) Size() (n int) {
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovCoin(uint64(l))
	}
	return n
}

func (m *AccountTransactionsReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovCoin(uint64(l))
	}
	return n
}

func sovCoin(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCoin(x uint64) (n int) {
	return sovCoin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AccountReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCoin
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerId = append(m.PeerId[:0], dAtA[iNdEx:postIndex]...)
			if m.PeerId == nil {
				m.PeerId = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCoin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockchainReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockchainReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockchainReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCoin
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderHash = append(m.HeaderHash[:0], dAtA[iNdEx:postIndex]...)
			if m.HeaderHash == nil {
				m.HeaderHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCoin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockchainResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockchainResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockchainResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocks = append(m.Blocks, &stratumn_alice_pb_coin.Block{})
			if err := m.Blocks[len(m.Blocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCoin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TransactionResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TransactionResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCoin
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCoin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AccountTransactionsReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AccountTransactionsReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountTransactionsReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCoin
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerId = append(m.PeerId[:0], dAtA[iNdEx:postIndex]...)
			if m.PeerId == nil {
				m.PeerId = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCoin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCoin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoin
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCoin
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCoin
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCoin(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCoin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoin   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("github.com/stratumn/alice/grpc/coin/coin.proto", fileDescriptorCoin) }

var fileDescriptorCoin = []byte{
	// 611 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0x36, 0xbb, 0x6b, 0x17, 0xa6, 0xbb, 0x0a, 0x23, 0x96, 0x31, 0x68, 0x1d, 0xa2, 0x60, 0x17,
	0xd6, 0xe9, 0xb2, 0x22, 0x88, 0x2b, 0xa8, 0xdd, 0x05, 0xdb, 0x83, 0x22, 0x51, 0x2f, 0x5e, 0xca,
	0x24, 0x99, 0x36, 0x61, 0xd3, 0x4c, 0xcc, 0x4c, 0xa0, 0xf5, 0xd8, 0x8b, 0x57, 0x8f, 0x5e, 0xbc,
	0xe4, 0xe4, 0x59, 0xe8, 0x29, 0x7f, 0xc0, 0xa3, 0x3f, 0x41, 0xea, 0x1f, 0xf0, 0x27, 0xc8, 0x4c,
	0x5a, 0x5a, 0x83, 0xdd, 0x2d, 0x7a, 0x68, 0x0a, 0x6f, 0xbe, 0xf7, 0xbd, 0xef, 0x7d, 0xef, 0xcd,
	0x00, 0xd2, 0x0f, 0xa4, 0x9f, 0x3a, 0xc4, 0xe5, 0x83, 0xa6, 0x90, 0x09, 0x95, 0xe9, 0x20, 0x6a,
	0xd2, 0x30, 0x70, 0x59, 0xb3, 0x9f, 0xc4, 0x6e, 0xd3, 0xe5, 0x41, 0xa4, 0x3f, 0x24, 0x4e, 0xb8,
	0xe4, 0x10, 0xcd, 0x41, 0x44, 0x83, 0x88, 0x02, 0x11, 0x75, 0x6e, 0xee, 0x9f, 0xc3, 0xc4, 0x86,
	0x52, 0xfd, 0x0a, 0x9e, 0xb3, 0xd0, 0xb1, 0x53, 0xae, 0x6a, 0x75, 0x00, 0x78, 0xea, 0xba, 0x3c,
	0x8d, 0xa4, 0xcd, 0xde, 0xc1, 0x23, 0xb0, 0x1d, 0x33, 0x96, 0x74, 0x03, 0x0f, 0x19, 0xd8, 0x68,
	0xec, 0xb4, 0xac, 0x2c, 0x47, 0xb5, 0x97, 0x8c, 0x25, 0xb8, 0x73, 0x82, 0x79, 0x0f, 0x4b, 0x9f,
	0x61, 0x5a, 0xe0, 0xbf, 0xe4, 0xc8, 0xf8, 0x95, 0x23, 0xc3, 0xae, 0xa8, 0x94, 0x8e, 0x67, 0x7d,
	0xdd, 0x00, 0xbb, 0xad, 0x90, 0xbb, 0xa7, 0xae, 0x4f, 0x83, 0x48, 0xd1, 0xf5, 0xc1, 0x8e, 0xa3,
	0x02, 0xdd, 0x28, 0x1d, 0x38, 0x2c, 0xd1, 0x9c, 0xbb, 0xad, 0x93, 0x2c, 0x47, 0x4f, 0x34, 0x10,
	0x17, 0x71, 0x2c, 0x39, 0x0e, 0x39, 0x3f, 0x4d, 0x63, 0x82, 0x3b, 0x3d, 0xcc, 0x07, 0x83, 0x40,
	0x4a, 0xe6, 0xed, 0x63, 0x8f, 0xf5, 0x68, 0x1a, 0x4a, 0xa1, 0x00, 0xaa, 0x72, 0x48, 0x85, 0xc4,
	0x9a, 0x91, 0xd8, 0x55, 0xfd, 0xff, 0x42, 0x13, 0x40, 0x0f, 0x54, 0x7d, 0x46, 0x3d, 0x96, 0x74,
	0x7d, 0x2a, 0x7c, 0xb4, 0xa1, 0xb5, 0x1f, 0x67, 0x39, 0x7a, 0xdc, 0xd6, 0x61, 0xac, 0xc2, 0xff,
	0x54, 0x06, 0x14, 0xbc, 0x6d, 0x2a, 0x7c, 0x68, 0x83, 0x8b, 0xba, 0x73, 0xb4, 0xa9, 0xfb, 0x78,
	0x94, 0xe5, 0xe8, 0x41, 0x21, 0x40, 0x59, 0xa3, 0x13, 0x04, 0x4e, 0x98, 0x4c, 0x93, 0x88, 0x79,
	0x67, 0xd4, 0xe0, 0x11, 0x23, 0x76, 0x41, 0x65, 0x75, 0xc1, 0xa5, 0x65, 0xcf, 0x44, 0x0c, 0x9f,
	0x83, 0x4a, 0x41, 0x85, 0x0c, 0xbc, 0xd9, 0xa8, 0x1e, 0xde, 0x20, 0xa5, 0xc5, 0x88, 0x1d, 0xbd,
	0x16, 0x44, 0xe7, 0xb5, 0x6a, 0x59, 0x8e, 0xe0, 0x82, 0x62, 0x26, 0x83, 0xd8, 0x33, 0x12, 0xab,
	0x0d, 0x2e, 0xbf, 0x4e, 0x68, 0x24, 0xa8, 0x2b, 0x03, 0x5e, 0x54, 0xb8, 0x0f, 0xb6, 0xe5, 0xb0,
	0x70, 0xaa, 0x98, 0xf2, 0xf5, 0x2c, 0x47, 0x48, 0xb5, 0x38, 0x1f, 0xb1, 0x5c, 0x64, 0x10, 0xbb,
	0x22, 0x87, 0xea, 0xcc, 0x7a, 0x03, 0x6a, 0xb3, 0x55, 0x59, 0x22, 0x14, 0xff, 0xbb, 0x36, 0x87,
	0x1f, 0xb6, 0xc0, 0xd6, 0x31, 0x0f, 0x22, 0xe8, 0x83, 0xed, 0x19, 0x3f, 0xbc, 0x4d, 0x56, 0x5d,
	0x06, 0xb2, 0xd8, 0x56, 0xf3, 0xe6, 0x2a, 0x67, 0x66, 0x18, 0xab, 0x36, 0x9e, 0x20, 0x78, 0x12,
	0x88, 0x38, 0xa4, 0x23, 0x4c, 0xa3, 0xb9, 0x04, 0x38, 0x02, 0xd5, 0xa5, 0x16, 0xe0, 0xad, 0x55,
	0x3c, 0x4b, 0x20, 0x73, 0x6f, 0xb5, 0xa4, 0x92, 0xbf, 0x96, 0xa9, 0xca, 0xbe, 0x62, 0x91, 0x87,
	0xe9, 0xb2, 0x93, 0x1f, 0x27, 0xc8, 0x80, 0x9f, 0x0d, 0x70, 0xe5, 0x2f, 0x2e, 0xc2, 0x83, 0x73,
	0x3b, 0x2e, 0x99, 0x6e, 0xae, 0xa3, 0xda, 0xba, 0x3b, 0x9e, 0xa0, 0xbd, 0x67, 0x4c, 0x96, 0xa7,
	0x8a, 0xfd, 0x40, 0x48, 0x9e, 0x8c, 0xd4, 0x70, 0x16, 0xc6, 0x1c, 0x18, 0xf0, 0x3d, 0x00, 0x8b,
	0x65, 0x82, 0x77, 0x56, 0xab, 0xfa, 0xe3, 0xa6, 0x9b, 0x8d, 0xf5, 0x80, 0x22, 0xb6, 0xae, 0x8d,
	0x27, 0xe8, 0xaa, 0x52, 0xe4, 0x94, 0x77, 0xb6, 0xf5, 0xf0, 0xdb, 0xb4, 0x6e, 0x7c, 0x9f, 0xd6,
	0x8d, 0x1f, 0xd3, 0xba, 0xf1, 0xe9, 0x67, 0xfd, 0xc2, 0xdb, 0xc6, 0x1a, 0x6f, 0xe8, 0x91, 0xfa,
	0x38, 0x15, 0xfd, 0x9c, 0xdd, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x69, 0x5a, 0xeb, 0x76,
	0x05, 0x00, 0x00,
}
