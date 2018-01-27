// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stratumn/alice/pb/coin/coin.proto

/*
	Package coin is a generated protocol buffer package.

	It is generated from these files:
		github.com/stratumn/alice/pb/coin/coin.proto

	It has these top-level messages:
		Gossip
		Transaction
		Signature
		Header
		Block
*/
package coin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/stratumn/alice/grpc/ext"
import google_protobuf1 "github.com/gogo/protobuf/types"

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

// Types of digital keys supported.
type KeyType int32

const (
	KeyType_RSA       KeyType = 0
	KeyType_ECDSA     KeyType = 1
	KeyType_Ed25519   KeyType = 2
	KeyType_Secp256k1 KeyType = 3
)

var KeyType_name = map[int32]string{
	0: "RSA",
	1: "ECDSA",
	2: "Ed25519",
	3: "Secp256k1",
}
var KeyType_value = map[string]int32{
	"RSA":       0,
	"ECDSA":     1,
	"Ed25519":   2,
	"Secp256k1": 3,
}

func (x KeyType) String() string {
	return proto.EnumName(KeyType_name, int32(x))
}
func (KeyType) EnumDescriptor() ([]byte, []int) { return fileDescriptorCoin, []int{0} }

// The consensus engine exchanges Gossip messages between peers.
// They are used to forward transactions to all peers, notify of a new block, etc.
type Gossip struct {
	// Types that are valid to be assigned to Msg:
	//	*Gossip_Tx
	//	*Gossip_Block
	Msg isGossip_Msg `protobuf_oneof:"msg"`
}

func (m *Gossip) Reset()                    { *m = Gossip{} }
func (m *Gossip) String() string            { return proto.CompactTextString(m) }
func (*Gossip) ProtoMessage()               {}
func (*Gossip) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{0} }

type isGossip_Msg interface {
	isGossip_Msg()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Gossip_Tx struct {
	Tx *Transaction `protobuf:"bytes,1,opt,name=tx,oneof"`
}
type Gossip_Block struct {
	Block *Block `protobuf:"bytes,2,opt,name=block,oneof"`
}

func (*Gossip_Tx) isGossip_Msg()    {}
func (*Gossip_Block) isGossip_Msg() {}

func (m *Gossip) GetMsg() isGossip_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *Gossip) GetTx() *Transaction {
	if x, ok := m.GetMsg().(*Gossip_Tx); ok {
		return x.Tx
	}
	return nil
}

func (m *Gossip) GetBlock() *Block {
	if x, ok := m.GetMsg().(*Gossip_Block); ok {
		return x.Block
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Gossip) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Gossip_OneofMarshaler, _Gossip_OneofUnmarshaler, _Gossip_OneofSizer, []interface{}{
		(*Gossip_Tx)(nil),
		(*Gossip_Block)(nil),
	}
}

func _Gossip_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Gossip)
	// msg
	switch x := m.Msg.(type) {
	case *Gossip_Tx:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tx); err != nil {
			return err
		}
	case *Gossip_Block:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Block); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Gossip.Msg has unexpected type %T", x)
	}
	return nil
}

func _Gossip_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Gossip)
	switch tag {
	case 1: // msg.tx
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Transaction)
		err := b.DecodeMessage(msg)
		m.Msg = &Gossip_Tx{msg}
		return true, err
	case 2: // msg.block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Block)
		err := b.DecodeMessage(msg)
		m.Msg = &Gossip_Block{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Gossip_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Gossip)
	// msg
	switch x := m.Msg.(type) {
	case *Gossip_Tx:
		s := proto.Size(x.Tx)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Gossip_Block:
		s := proto.Size(x.Block)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A coin transaction.
type Transaction struct {
	From      []byte     `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To        []byte     `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Value     int64      `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	Nonce     int64      `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Signature *Signature `protobuf:"bytes,5,opt,name=signature" json:"signature,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{1} }

func (m *Transaction) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Transaction) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Transaction) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Transaction) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

// A digital signature.
type Signature struct {
	KeyType   KeyType `protobuf:"varint,1,opt,name=key_type,json=keyType,proto3,enum=stratumn.alice.pb.coin.KeyType" json:"key_type,omitempty"`
	PublicKey []byte  `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Signature []byte  `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{2} }

func (m *Signature) GetKeyType() KeyType {
	if m != nil {
		return m.KeyType
	}
	return KeyType_RSA
}

func (m *Signature) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// A block header.
type Header struct {
	Version      int32                       `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	BlockNumber  int64                       `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	PreviousHash []byte                      `protobuf:"bytes,3,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"`
	MerkleRoot   []byte                      `protobuf:"bytes,4,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	Timestamp    *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=timestamp" json:"timestamp,omitempty"`
	Nonce        int64                       `protobuf:"varint,6,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{3} }

func (m *Header) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Header) GetBlockNumber() int64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *Header) GetPreviousHash() []byte {
	if m != nil {
		return m.PreviousHash
	}
	return nil
}

func (m *Header) GetMerkleRoot() []byte {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

func (m *Header) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Header) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

// A block.
type Block struct {
	Header       *Header        `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{4} }

func (m *Block) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto.RegisterType((*Gossip)(nil), "stratumn.alice.pb.coin.Gossip")
	proto.RegisterType((*Transaction)(nil), "stratumn.alice.pb.coin.Transaction")
	proto.RegisterType((*Signature)(nil), "stratumn.alice.pb.coin.Signature")
	proto.RegisterType((*Header)(nil), "stratumn.alice.pb.coin.Header")
	proto.RegisterType((*Block)(nil), "stratumn.alice.pb.coin.Block")
	proto.RegisterEnum("stratumn.alice.pb.coin.KeyType", KeyType_name, KeyType_value)
}
func (m *Gossip) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Gossip) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Msg != nil {
		nn1, err := m.Msg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *Gossip_Tx) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Tx != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCoin(dAtA, i, uint64(m.Tx.Size()))
		n2, err := m.Tx.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *Gossip_Block) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Block != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCoin(dAtA, i, uint64(m.Block.Size()))
		n3, err := m.Block.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *Transaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transaction) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.From) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCoin(dAtA, i, uint64(len(m.From)))
		i += copy(dAtA[i:], m.From)
	}
	if len(m.To) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCoin(dAtA, i, uint64(len(m.To)))
		i += copy(dAtA[i:], m.To)
	}
	if m.Value != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCoin(dAtA, i, uint64(m.Value))
	}
	if m.Nonce != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintCoin(dAtA, i, uint64(m.Nonce))
	}
	if m.Signature != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCoin(dAtA, i, uint64(m.Signature.Size()))
		n4, err := m.Signature.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *Signature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Signature) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.KeyType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCoin(dAtA, i, uint64(m.KeyType))
	}
	if len(m.PublicKey) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCoin(dAtA, i, uint64(len(m.PublicKey)))
		i += copy(dAtA[i:], m.PublicKey)
	}
	if len(m.Signature) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCoin(dAtA, i, uint64(len(m.Signature)))
		i += copy(dAtA[i:], m.Signature)
	}
	return i, nil
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCoin(dAtA, i, uint64(m.Version))
	}
	if m.BlockNumber != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCoin(dAtA, i, uint64(m.BlockNumber))
	}
	if len(m.PreviousHash) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCoin(dAtA, i, uint64(len(m.PreviousHash)))
		i += copy(dAtA[i:], m.PreviousHash)
	}
	if len(m.MerkleRoot) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCoin(dAtA, i, uint64(len(m.MerkleRoot)))
		i += copy(dAtA[i:], m.MerkleRoot)
	}
	if m.Timestamp != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCoin(dAtA, i, uint64(m.Timestamp.Size()))
		n5, err := m.Timestamp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Nonce != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintCoin(dAtA, i, uint64(m.Nonce))
	}
	return i, nil
}

func (m *Block) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Block) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCoin(dAtA, i, uint64(m.Header.Size()))
		n6, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if len(m.Transactions) > 0 {
		for _, msg := range m.Transactions {
			dAtA[i] = 0x12
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

func encodeVarintCoin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Gossip) Size() (n int) {
	var l int
	_ = l
	if m.Msg != nil {
		n += m.Msg.Size()
	}
	return n
}

func (m *Gossip_Tx) Size() (n int) {
	var l int
	_ = l
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovCoin(uint64(l))
	}
	return n
}
func (m *Gossip_Block) Size() (n int) {
	var l int
	_ = l
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovCoin(uint64(l))
	}
	return n
}
func (m *Transaction) Size() (n int) {
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovCoin(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovCoin(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovCoin(uint64(m.Value))
	}
	if m.Nonce != 0 {
		n += 1 + sovCoin(uint64(m.Nonce))
	}
	if m.Signature != nil {
		l = m.Signature.Size()
		n += 1 + l + sovCoin(uint64(l))
	}
	return n
}

func (m *Signature) Size() (n int) {
	var l int
	_ = l
	if m.KeyType != 0 {
		n += 1 + sovCoin(uint64(m.KeyType))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovCoin(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovCoin(uint64(l))
	}
	return n
}

func (m *Header) Size() (n int) {
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovCoin(uint64(m.Version))
	}
	if m.BlockNumber != 0 {
		n += 1 + sovCoin(uint64(m.BlockNumber))
	}
	l = len(m.PreviousHash)
	if l > 0 {
		n += 1 + l + sovCoin(uint64(l))
	}
	l = len(m.MerkleRoot)
	if l > 0 {
		n += 1 + l + sovCoin(uint64(l))
	}
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovCoin(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovCoin(uint64(m.Nonce))
	}
	return n
}

func (m *Block) Size() (n int) {
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovCoin(uint64(l))
	}
	if len(m.Transactions) > 0 {
		for _, e := range m.Transactions {
			l = e.Size()
			n += 1 + l + sovCoin(uint64(l))
		}
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
func (m *Gossip) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Gossip: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Gossip: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
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
			v := &Transaction{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Msg = &Gossip_Tx{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
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
			v := &Block{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Msg = &Gossip_Block{v}
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
func (m *Transaction) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Transaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = append(m.From[:0], dAtA[iNdEx:postIndex]...)
			if m.From == nil {
				m.From = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
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
			m.To = append(m.To[:0], dAtA[iNdEx:postIndex]...)
			if m.To == nil {
				m.To = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
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
			if m.Signature == nil {
				m.Signature = &Signature{}
			}
			if err := m.Signature.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Signature) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Signature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Signature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyType", wireType)
			}
			m.KeyType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyType |= (KeyType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
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
			m.PublicKey = append(m.PublicKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PublicKey == nil {
				m.PublicKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
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
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
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
func (m *Header) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
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
				m.BlockNumber |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousHash", wireType)
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
			m.PreviousHash = append(m.PreviousHash[:0], dAtA[iNdEx:postIndex]...)
			if m.PreviousHash == nil {
				m.PreviousHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MerkleRoot", wireType)
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
			m.MerkleRoot = append(m.MerkleRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.MerkleRoot == nil {
				m.MerkleRoot = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
			if m.Timestamp == nil {
				m.Timestamp = &google_protobuf1.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= (int64(b) & 0x7F) << shift
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
func (m *Block) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Block: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Block: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
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
			if m.Header == nil {
				m.Header = &Header{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transactions", wireType)
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
			m.Transactions = append(m.Transactions, &Transaction{})
			if err := m.Transactions[len(m.Transactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

func init() { proto.RegisterFile("github.com/stratumn/alice/pb/coin/coin.proto", fileDescriptorCoin) }

var fileDescriptorCoin = []byte{
	// 856 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x5e, 0xc7, 0x4d, 0x42, 0x26, 0xd9, 0x12, 0x86, 0xbf, 0xd1, 0x56, 0xcd, 0x4e, 0xb7, 0x12,
	0x2d, 0x7f, 0x8e, 0xba, 0x28, 0x40, 0xa9, 0xb8, 0xa8, 0xdb, 0x42, 0xa0, 0xb0, 0x42, 0x49, 0xc4,
	0x05, 0x37, 0xd1, 0xc4, 0x39, 0x9b, 0x58, 0x71, 0x3c, 0xd6, 0xcc, 0x38, 0x5d, 0x5f, 0x70, 0xc7,
	0x43, 0x70, 0xed, 0x2b, 0x9e, 0x01, 0xbf, 0x00, 0x97, 0x20, 0x71, 0x0f, 0x5a, 0x5e, 0x80, 0x47,
	0x40, 0x33, 0x13, 0xef, 0x1a, 0x75, 0x03, 0xbd, 0x58, 0x6b, 0xe7, 0xf8, 0xfb, 0xbe, 0x73, 0xfc,
	0x9d, 0x73, 0x26, 0xe8, 0xbd, 0x45, 0xa8, 0x96, 0xe9, 0xcc, 0x0b, 0xf8, 0xba, 0x2f, 0x95, 0x60,
	0x2a, 0x5d, 0xc7, 0x7d, 0x16, 0x85, 0x01, 0xf4, 0x93, 0x59, 0x3f, 0xe0, 0x61, 0x6c, 0x1e, 0x5e,
	0x22, 0xb8, 0xe2, 0xf8, 0x8d, 0x12, 0xe2, 0x19, 0x88, 0x97, 0x68, 0x62, 0x18, 0x1f, 0xfc, 0x87,
	0xca, 0x42, 0x24, 0x41, 0x1f, 0xce, 0x94, 0xfe, 0xb3, 0x2a, 0x07, 0x87, 0x0b, 0xce, 0x17, 0x11,
	0xf4, 0xcd, 0x69, 0x96, 0x9e, 0xf6, 0x55, 0xb8, 0x06, 0xa9, 0xd8, 0x3a, 0xb1, 0x80, 0xa3, 0xef,
	0x51, 0xe3, 0x73, 0x2e, 0x65, 0x98, 0xe0, 0x01, 0xaa, 0xa9, 0x33, 0xe2, 0x50, 0xe7, 0x6e, 0xfb,
	0xf8, 0xb6, 0x77, 0x75, 0x76, 0x6f, 0x22, 0x58, 0x2c, 0x59, 0xa0, 0x42, 0x1e, 0x0f, 0xf7, 0x46,
	0x35, 0x75, 0x86, 0x07, 0xa8, 0x3e, 0x8b, 0x78, 0xb0, 0x22, 0x35, 0xc3, 0xbc, 0xb9, 0x8b, 0xe9,
	0x6b, 0xd0, 0x70, 0x6f, 0x64, 0xd1, 0x7e, 0x1d, 0xb9, 0x6b, 0xb9, 0x38, 0xfa, 0xa3, 0x86, 0xda,
	0x15, 0x4d, 0xec, 0xa1, 0x6b, 0xa7, 0x82, 0xaf, 0x4d, 0x19, 0x1d, 0xff, 0x20, 0x2f, 0xc8, 0xf5,
	0x31, 0xc4, 0x73, 0x10, 0xf4, 0x1b, 0x00, 0x41, 0xbf, 0x78, 0xfc, 0x53, 0x41, 0x9c, 0xbf, 0x0b,
	0xe2, 0x8c, 0x0c, 0x0e, 0xbf, 0x8f, 0x6a, 0x8a, 0x9b, 0xd4, 0x1d, 0xff, 0x66, 0x5e, 0x90, 0x57,
	0x46, 0x10, 0x84, 0x49, 0x08, 0xb1, 0x7a, 0x8e, 0x50, 0x53, 0x1c, 0xbf, 0x8b, 0xea, 0x1b, 0x16,
	0xa5, 0x40, 0x5c, 0xea, 0xdc, 0x75, 0xfd, 0xd7, 0xf3, 0x82, 0xec, 0x7f, 0xab, 0x03, 0x54, 0x71,
	0x2a, 0x21, 0x9e, 0x6b, 0xf4, 0xc8, 0x62, 0xf0, 0xa7, 0xa8, 0x1e, 0xf3, 0x38, 0x00, 0x72, 0xcd,
	0x80, 0xef, 0xe4, 0x05, 0x39, 0x3c, 0xd1, 0x01, 0x0d, 0x4e, 0x04, 0x6c, 0x74, 0x16, 0x01, 0x49,
	0xc4, 0x32, 0xca, 0x94, 0x62, 0xc1, 0x4a, 0x5a, 0xba, 0x61, 0xe1, 0x1f, 0x1c, 0xd4, 0x92, 0xe1,
	0x22, 0x66, 0x2a, 0x15, 0x40, 0xea, 0xc6, 0x9d, 0x5b, 0xbb, 0xdc, 0x19, 0x97, 0x40, 0xff, 0x51,
	0x5e, 0x90, 0x07, 0x17, 0x47, 0xca, 0x4f, 0xa9, 0x5a, 0x02, 0x55, 0x97, 0x1e, 0xd1, 0x44, 0xf0,
	0x4d, 0x18, 0x2f, 0x28, 0x7f, 0x16, 0x83, 0x90, 0xcb, 0x30, 0x29, 0x41, 0x5a, 0xc9, 0x96, 0x70,
	0x99, 0xf8, 0xe8, 0x77, 0x07, 0xb5, 0x2e, 0xe4, 0xf0, 0x04, 0xbd, 0xb4, 0x82, 0x6c, 0xaa, 0xb2,
	0x04, 0x8c, 0xc7, 0xd7, 0x8f, 0x0f, 0x77, 0x95, 0xf4, 0x14, 0xb2, 0x49, 0x96, 0x80, 0x4f, 0xf2,
	0x82, 0x74, 0xf5, 0x7f, 0x3a, 0xcd, 0x0a, 0x32, 0x9a, 0x4a, 0xb0, 0x3e, 0x35, 0x57, 0x16, 0x82,
	0x3f, 0x41, 0x28, 0x49, 0x67, 0x51, 0x18, 0x4c, 0x57, 0x90, 0x6d, 0xbb, 0x71, 0x23, 0x2f, 0xc8,
	0xab, 0x3a, 0x31, 0x88, 0x3b, 0x92, 0xda, 0xd7, 0x9a, 0x6e, 0xeb, 0xb3, 0xe7, 0xa7, 0x90, 0xe1,
	0x41, 0xd5, 0x25, 0xd7, 0x50, 0xdf, 0xcc, 0x0b, 0xf2, 0xf2, 0xa5, 0x05, 0xb3, 0x4c, 0xc1, 0x73,
	0x9f, 0xf5, 0x9b, 0x8b, 0x1a, 0x43, 0x60, 0x73, 0x10, 0xb8, 0x8f, 0x9a, 0x1b, 0x10, 0x32, 0xe4,
	0xb1, 0xf9, 0xa4, 0xba, 0x6d, 0xab, 0x99, 0x37, 0xba, 0x7d, 0x61, 0xcb, 0xdd, 0x1e, 0xf0, 0x47,
	0xa8, 0x63, 0x86, 0x70, 0x1a, 0xa7, 0xeb, 0x19, 0x08, 0x53, 0xb0, 0xeb, 0xbf, 0x96, 0x17, 0xa4,
	0x63, 0x59, 0x36, 0x6e, 0x48, 0x6d, 0x83, 0x3c, 0x31, 0x01, 0x3c, 0x44, 0xfb, 0xba, 0xef, 0x21,
	0x4f, 0xe5, 0x74, 0xc9, 0xe4, 0x72, 0x5b, 0xef, 0xed, 0xbc, 0x20, 0x37, 0x86, 0x4c, 0x2e, 0xcb,
	0x46, 0x94, 0x20, 0xba, 0x34, 0x15, 0x1a, 0xa1, 0x4e, 0x19, 0xd4, 0x40, 0xfc, 0x25, 0x6a, 0xaf,
	0x41, 0xac, 0x22, 0x98, 0x0a, 0xce, 0x95, 0x99, 0xb0, 0x8e, 0xff, 0x76, 0x5e, 0x90, 0x5b, 0x5f,
	0x9b, 0x30, 0xd5, 0x61, 0x2d, 0x67, 0xd2, 0x57, 0xdb, 0x6f, 0x9d, 0x40, 0x96, 0x3d, 0xe2, 0x5c,
	0x61, 0x86, 0x5a, 0x17, 0x5b, 0xbd, 0x9d, 0xb3, 0x03, 0xcf, 0xee, 0xbd, 0x57, 0xee, 0xbd, 0x37,
	0x29, 0x11, 0x76, 0x8e, 0x2f, 0x8e, 0x65, 0xc9, 0x36, 0x4f, 0x20, 0x80, 0xa9, 0xd2, 0xaf, 0x4b,
	0x55, 0x3c, 0x2f, 0x57, 0xa1, 0x61, 0xac, 0x3a, 0xc9, 0x0b, 0xf2, 0xd9, 0x64, 0x09, 0xd4, 0x04,
	0xcd, 0x40, 0xe8, 0x9d, 0x58, 0x40, 0x0c, 0x82, 0x29, 0xa8, 0x08, 0x2a, 0x4e, 0x59, 0x14, 0xf1,
	0x67, 0x74, 0xc3, 0x44, 0x68, 0xa4, 0x65, 0x99, 0xb1, 0xe2, 0x8d, 0x15, 0x3f, 0xfa, 0xd9, 0x41,
	0x75, 0xd3, 0x00, 0xfc, 0x15, 0x6a, 0xd8, 0xd7, 0xdb, 0xfb, 0xa8, 0xb7, 0x6b, 0x48, 0xed, 0x08,
	0x54, 0x7b, 0x57, 0x91, 0xdd, 0x6a, 0xe0, 0x00, 0x75, 0xaa, 0x06, 0x92, 0x1a, 0x75, 0x5f, 0xf0,
	0x8e, 0x33, 0x37, 0x10, 0xf6, 0xaf, 0xee, 0xc1, 0xbf, 0x44, 0xdf, 0xb9, 0x8f, 0x9a, 0xdb, 0x8d,
	0xc1, 0x4d, 0xe4, 0x8e, 0xc6, 0x0f, 0xbb, 0x7b, 0xb8, 0x85, 0xea, 0x4f, 0x1e, 0x3d, 0x1e, 0x3f,
	0xec, 0x3a, 0xb8, 0x8d, 0x9a, 0x4f, 0xe6, 0xc7, 0x83, 0xc1, 0xbd, 0xfb, 0xdd, 0x1a, 0xde, 0x47,
	0xad, 0x31, 0x04, 0xc9, 0xf1, 0xe0, 0xc3, 0xd5, 0xbd, 0xae, 0xeb, 0x7f, 0xfc, 0xcb, 0x79, 0xcf,
	0xf9, 0xf5, 0xbc, 0xe7, 0xfc, 0x79, 0xde, 0x73, 0x7e, 0xfc, 0xab, 0xb7, 0xf7, 0xdd, 0x5b, 0xff,
	0xfb, 0x53, 0xf1, 0x40, 0x3f, 0x66, 0x0d, 0xd3, 0xdf, 0x0f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x1a, 0xb7, 0x87, 0x16, 0x5b, 0x06, 0x00, 0x00,
}
