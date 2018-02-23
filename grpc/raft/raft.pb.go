// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stratumn/alice/grpc/raft/raft.proto

/*
	Package raft is a generated protocol buffer package.

	It is generated from these files:
		github.com/stratumn/alice/grpc/raft/raft.proto

	It has these top-level messages:
		Empty
		Peer
		PeerID
		StatusInfo
		Proposal
		Entry
*/
package raft

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/stratumn/alice/grpc/ext"

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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptorRaft, []int{0} }

type Peer struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptorRaft, []int{1} }

func (m *Peer) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Peer) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type PeerID struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *PeerID) Reset()                    { *m = PeerID{} }
func (m *PeerID) String() string            { return proto.CompactTextString(m) }
func (*PeerID) ProtoMessage()               {}
func (*PeerID) Descriptor() ([]byte, []int) { return fileDescriptorRaft, []int{2} }

func (m *PeerID) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type StatusInfo struct {
	Running bool   `protobuf:"varint,1,opt,name=running,proto3" json:"running,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *StatusInfo) Reset()                    { *m = StatusInfo{} }
func (m *StatusInfo) String() string            { return proto.CompactTextString(m) }
func (*StatusInfo) ProtoMessage()               {}
func (*StatusInfo) Descriptor() ([]byte, []int) { return fileDescriptorRaft, []int{3} }

func (m *StatusInfo) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

func (m *StatusInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Proposal struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Proposal) Reset()                    { *m = Proposal{} }
func (m *Proposal) String() string            { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()               {}
func (*Proposal) Descriptor() ([]byte, []int) { return fileDescriptorRaft, []int{4} }

func (m *Proposal) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Entry struct {
	Index uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Data  []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptorRaft, []int{5} }

func (m *Entry) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Entry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "stratumn.alice.grpc.raft.Empty")
	proto.RegisterType((*Peer)(nil), "stratumn.alice.grpc.raft.Peer")
	proto.RegisterType((*PeerID)(nil), "stratumn.alice.grpc.raft.PeerID")
	proto.RegisterType((*StatusInfo)(nil), "stratumn.alice.grpc.raft.StatusInfo")
	proto.RegisterType((*Proposal)(nil), "stratumn.alice.grpc.raft.Proposal")
	proto.RegisterType((*Entry)(nil), "stratumn.alice.grpc.raft.Entry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Raft service

type RaftClient interface {
	Start(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Status(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StatusInfo, error)
	Peers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Raft_PeersClient, error)
	Discover(ctx context.Context, in *PeerID, opts ...grpc.CallOption) (Raft_DiscoverClient, error)
	Invite(ctx context.Context, in *PeerID, opts ...grpc.CallOption) (*Empty, error)
	Join(ctx context.Context, in *PeerID, opts ...grpc.CallOption) (*Empty, error)
	Expel(ctx context.Context, in *PeerID, opts ...grpc.CallOption) (*Empty, error)
	Propose(ctx context.Context, in *Proposal, opts ...grpc.CallOption) (*Empty, error)
	Log(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Raft_LogClient, error)
}

type raftClient struct {
	cc *grpc.ClientConn
}

func NewRaftClient(cc *grpc.ClientConn) RaftClient {
	return &raftClient{cc}
}

func (c *raftClient) Start(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.raft.Raft/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.raft.Raft/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) Status(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StatusInfo, error) {
	out := new(StatusInfo)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.raft.Raft/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) Peers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Raft_PeersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Raft_serviceDesc.Streams[0], c.cc, "/stratumn.alice.grpc.raft.Raft/Peers", opts...)
	if err != nil {
		return nil, err
	}
	x := &raftPeersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Raft_PeersClient interface {
	Recv() (*Peer, error)
	grpc.ClientStream
}

type raftPeersClient struct {
	grpc.ClientStream
}

func (x *raftPeersClient) Recv() (*Peer, error) {
	m := new(Peer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *raftClient) Discover(ctx context.Context, in *PeerID, opts ...grpc.CallOption) (Raft_DiscoverClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Raft_serviceDesc.Streams[1], c.cc, "/stratumn.alice.grpc.raft.Raft/Discover", opts...)
	if err != nil {
		return nil, err
	}
	x := &raftDiscoverClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Raft_DiscoverClient interface {
	Recv() (*Peer, error)
	grpc.ClientStream
}

type raftDiscoverClient struct {
	grpc.ClientStream
}

func (x *raftDiscoverClient) Recv() (*Peer, error) {
	m := new(Peer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *raftClient) Invite(ctx context.Context, in *PeerID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.raft.Raft/Invite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) Join(ctx context.Context, in *PeerID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.raft.Raft/Join", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) Expel(ctx context.Context, in *PeerID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.raft.Raft/Expel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) Propose(ctx context.Context, in *Proposal, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.raft.Raft/Propose", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) Log(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Raft_LogClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Raft_serviceDesc.Streams[2], c.cc, "/stratumn.alice.grpc.raft.Raft/Log", opts...)
	if err != nil {
		return nil, err
	}
	x := &raftLogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Raft_LogClient interface {
	Recv() (*Entry, error)
	grpc.ClientStream
}

type raftLogClient struct {
	grpc.ClientStream
}

func (x *raftLogClient) Recv() (*Entry, error) {
	m := new(Entry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Raft service

type RaftServer interface {
	Start(context.Context, *Empty) (*Empty, error)
	Stop(context.Context, *Empty) (*Empty, error)
	Status(context.Context, *Empty) (*StatusInfo, error)
	Peers(*Empty, Raft_PeersServer) error
	Discover(*PeerID, Raft_DiscoverServer) error
	Invite(context.Context, *PeerID) (*Empty, error)
	Join(context.Context, *PeerID) (*Empty, error)
	Expel(context.Context, *PeerID) (*Empty, error)
	Propose(context.Context, *Proposal) (*Empty, error)
	Log(*Empty, Raft_LogServer) error
}

func RegisterRaftServer(s *grpc.Server, srv RaftServer) {
	s.RegisterService(&_Raft_serviceDesc, srv)
}

func _Raft_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.raft.Raft/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).Start(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.raft.Raft/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).Stop(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.raft.Raft/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).Status(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_Peers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RaftServer).Peers(m, &raftPeersServer{stream})
}

type Raft_PeersServer interface {
	Send(*Peer) error
	grpc.ServerStream
}

type raftPeersServer struct {
	grpc.ServerStream
}

func (x *raftPeersServer) Send(m *Peer) error {
	return x.ServerStream.SendMsg(m)
}

func _Raft_Discover_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PeerID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RaftServer).Discover(m, &raftDiscoverServer{stream})
}

type Raft_DiscoverServer interface {
	Send(*Peer) error
	grpc.ServerStream
}

type raftDiscoverServer struct {
	grpc.ServerStream
}

func (x *raftDiscoverServer) Send(m *Peer) error {
	return x.ServerStream.SendMsg(m)
}

func _Raft_Invite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).Invite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.raft.Raft/Invite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).Invite(ctx, req.(*PeerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.raft.Raft/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).Join(ctx, req.(*PeerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_Expel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).Expel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.raft.Raft/Expel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).Expel(ctx, req.(*PeerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_Propose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Proposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).Propose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.raft.Raft/Propose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).Propose(ctx, req.(*Proposal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_Log_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RaftServer).Log(m, &raftLogServer{stream})
}

type Raft_LogServer interface {
	Send(*Entry) error
	grpc.ServerStream
}

type raftLogServer struct {
	grpc.ServerStream
}

func (x *raftLogServer) Send(m *Entry) error {
	return x.ServerStream.SendMsg(m)
}

var _Raft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stratumn.alice.grpc.raft.Raft",
	HandlerType: (*RaftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Raft_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Raft_Stop_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Raft_Status_Handler,
		},
		{
			MethodName: "Invite",
			Handler:    _Raft_Invite_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _Raft_Join_Handler,
		},
		{
			MethodName: "Expel",
			Handler:    _Raft_Expel_Handler,
		},
		{
			MethodName: "Propose",
			Handler:    _Raft_Propose_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Peers",
			Handler:       _Raft_Peers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Discover",
			Handler:       _Raft_Discover_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Log",
			Handler:       _Raft_Log_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/stratumn/alice/grpc/raft/raft.proto",
}

func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *Peer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Peer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRaft(dAtA, i, uint64(m.Id))
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRaft(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	return i, nil
}

func (m *PeerID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeerID) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRaft(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	return i, nil
}

func (m *StatusInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Running {
		dAtA[i] = 0x8
		i++
		if m.Running {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Id != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRaft(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *Proposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proposal) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRaft(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func (m *Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRaft(dAtA, i, uint64(m.Index))
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRaft(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func encodeVarintRaft(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Empty) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *Peer) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovRaft(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRaft(uint64(l))
	}
	return n
}

func (m *PeerID) Size() (n int) {
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRaft(uint64(l))
	}
	return n
}

func (m *StatusInfo) Size() (n int) {
	var l int
	_ = l
	if m.Running {
		n += 2
	}
	if m.Id != 0 {
		n += 1 + sovRaft(uint64(m.Id))
	}
	return n
}

func (m *Proposal) Size() (n int) {
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovRaft(uint64(l))
	}
	return n
}

func (m *Entry) Size() (n int) {
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovRaft(uint64(m.Index))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovRaft(uint64(l))
	}
	return n
}

func sovRaft(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRaft(x uint64) (n int) {
	return sovRaft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
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
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func (m *Peer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
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
			return fmt.Errorf("proto: Peer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Peer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
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
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func (m *PeerID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
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
			return fmt.Errorf("proto: PeerID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeerID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
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
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func (m *StatusInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
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
			return fmt.Errorf("proto: StatusInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Running", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Running = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func (m *Proposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
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
			return fmt.Errorf("proto: Proposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
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
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func (m *Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
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
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
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
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func skipRaft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRaft
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
					return 0, ErrIntOverflowRaft
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
					return 0, ErrIntOverflowRaft
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
				return 0, ErrInvalidLengthRaft
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRaft
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
				next, err := skipRaft(dAtA[start:])
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
	ErrInvalidLengthRaft = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRaft   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("github.com/stratumn/alice/grpc/raft/raft.proto", fileDescriptorRaft) }

var fileDescriptorRaft = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xfe, 0xd3, 0x3f, 0x69, 0xcb, 0x61, 0x12, 0xc2, 0x68, 0x93, 0x37, 0xb6, 0xce, 0x04, 0x10,
	0x43, 0x1a, 0xe9, 0x04, 0x17, 0x48, 0x70, 0xc5, 0xd4, 0x81, 0x82, 0x10, 0xaa, 0xb2, 0x3b, 0xc4,
	0x8d, 0xd7, 0xb8, 0x99, 0x51, 0x6a, 0x07, 0xdb, 0x9d, 0xba, 0xdb, 0x3c, 0x05, 0x77, 0x48, 0xbb,
	0xe2, 0x05, 0x72, 0xd5, 0x17, 0xe0, 0x92, 0x47, 0x40, 0xe3, 0x05, 0x78, 0x04, 0xe4, 0x24, 0x65,
	0x4c, 0x62, 0xeb, 0x24, 0x7a, 0x11, 0x5f, 0x9c, 0xf3, 0x9d, 0xef, 0x3b, 0xf6, 0x77, 0xec, 0x40,
	0x90, 0x70, 0x73, 0x38, 0x3e, 0x08, 0x06, 0x72, 0xd4, 0xd5, 0x46, 0x51, 0x33, 0x1e, 0x89, 0x2e,
	0x4d, 0xf9, 0x80, 0x75, 0x13, 0x95, 0x0d, 0xba, 0x8a, 0x0e, 0x4d, 0xb9, 0x04, 0x99, 0x92, 0x46,
	0x22, 0x3c, 0x03, 0x05, 0x25, 0x28, 0xb0, 0xa0, 0xc0, 0xe6, 0xd7, 0xb6, 0xe7, 0x30, 0xb1, 0x89,
	0xb1, 0x5f, 0xc5, 0xe3, 0xb7, 0xc0, 0xdb, 0x1b, 0x65, 0xe6, 0xd8, 0x7f, 0x0b, 0x6e, 0x9f, 0x31,
	0x85, 0x6e, 0x43, 0x83, 0xc7, 0xd8, 0x21, 0xce, 0x96, 0xbb, 0x7b, 0xfd, 0x64, 0x8a, 0x5b, 0x11,
	0x1d, 0x1a, 0x12, 0xf6, 0xa2, 0x06, 0x8f, 0xd1, 0x43, 0x68, 0xd1, 0x38, 0x56, 0x4c, 0x6b, 0xdc,
	0x20, 0xce, 0xd6, 0xd2, 0xee, 0x0d, 0x8b, 0xb0, 0x75, 0x24, 0xec, 0xfd, 0x9c, 0x62, 0x27, 0x9a,
	0xe5, 0xfd, 0xa7, 0xd0, 0xb4, 0xf1, 0xb0, 0x87, 0x1e, 0x9d, 0x15, 0x39, 0x65, 0xd1, 0xad, 0x3f,
	0x8a, 0xbe, 0x4c, 0xb1, 0x73, 0xbe, 0xb0, 0x0f, 0xb0, 0x6f, 0xa8, 0x19, 0xeb, 0x50, 0x0c, 0x25,
	0xba, 0x0f, 0x2d, 0x35, 0x16, 0x82, 0x8b, 0xa4, 0x2c, 0x6e, 0xd7, 0x3d, 0x55, 0xa1, 0x68, 0x96,
	0xab, 0xbb, 0x6e, 0xfc, 0xb5, 0x6b, 0x7f, 0x1b, 0xda, 0x7d, 0x25, 0x33, 0xa9, 0x69, 0x8a, 0x08,
	0xb8, 0x31, 0x35, 0xb4, 0xee, 0x64, 0xe9, 0x64, 0x8a, 0xdd, 0x1e, 0x35, 0xd4, 0xb6, 0x11, 0x95,
	0x19, 0xff, 0x25, 0x78, 0x7b, 0xc2, 0xa8, 0x63, 0xb4, 0x09, 0x1e, 0x17, 0x31, 0x9b, 0xd4, 0x87,
	0x71, 0xed, 0x64, 0x8a, 0xbd, 0xd0, 0x06, 0xa2, 0x2a, 0x8e, 0xd6, 0x6b, 0xae, 0xea, 0x28, 0xda,
	0x33, 0xae, 0x8a, 0xe7, 0xf1, 0xe7, 0x36, 0xb8, 0xb6, 0x0b, 0x94, 0x82, 0xb7, 0x6f, 0xa8, 0x32,
	0x68, 0x33, 0xb8, 0xc8, 0xb4, 0xa0, 0xf4, 0x60, 0x6d, 0x1e, 0xc0, 0xbf, 0x93, 0x17, 0x78, 0xa3,
	0x24, 0x23, 0x9a, 0x8b, 0x24, 0x65, 0x44, 0xc8, 0x98, 0x11, 0x8b, 0x20, 0x9a, 0xa9, 0x23, 0xa6,
	0xd0, 0x47, 0x70, 0xf7, 0x8d, 0xcc, 0x16, 0x20, 0xf6, 0x20, 0x2f, 0xf0, 0x5d, 0xcb, 0x45, 0xcc,
	0x21, 0xab, 0xe9, 0x09, 0x15, 0x31, 0x51, 0x6c, 0x24, 0x8f, 0x58, 0x19, 0xb5, 0x3b, 0x45, 0x43,
	0x68, 0x56, 0x8e, 0xcd, 0x17, 0xbd, 0x77, 0x31, 0xe0, 0xcc, 0x74, 0x7f, 0x39, 0x2f, 0xf0, 0xcd,
	0x57, 0x6c, 0xb6, 0x27, 0xa2, 0x2b, 0x76, 0x0a, 0x9e, 0x9d, 0x9a, 0x2b, 0xc8, 0x74, 0x2e, 0x06,
	0x58, 0x86, 0x33, 0x81, 0x41, 0x3a, 0xd6, 0x86, 0x29, 0x92, 0x59, 0xde, 0x1d, 0x07, 0x71, 0x68,
	0xf7, 0xb8, 0x1e, 0x48, 0x7b, 0x92, 0xe4, 0x72, 0x92, 0xb0, 0x37, 0x57, 0x66, 0x35, 0x2f, 0xf0,
	0xf2, 0x8c, 0x8f, 0x94, 0x23, 0x59, 0x0b, 0xee, 0x38, 0xe8, 0x03, 0x34, 0x43, 0x71, 0xc4, 0x0d,
	0xbb, 0x82, 0xd0, 0x5c, 0xaf, 0x36, 0xf2, 0x02, 0xaf, 0xbe, 0x88, 0x63, 0x42, 0xab, 0x89, 0x30,
	0xb2, 0xb4, 0xa7, 0x56, 0x43, 0x0c, 0xdc, 0xd7, 0x92, 0x8b, 0x45, 0x28, 0xad, 0xe5, 0x05, 0x5e,
	0xb1, 0x64, 0x84, 0x4d, 0xb8, 0xe6, 0x86, 0x8b, 0xe4, 0xb7, 0x4c, 0x0a, 0xde, 0xde, 0x24, 0x63,
	0xe9, 0x22, 0x74, 0x48, 0x5e, 0xe0, 0xf5, 0xa8, 0x9a, 0x34, 0xeb, 0x0e, 0x19, 0x2a, 0x39, 0x3a,
	0xb7, 0xa9, 0xf7, 0xd0, 0xaa, 0xae, 0x35, 0x43, 0xfe, 0x25, 0x7a, 0xf5, 0xcd, 0x9f, 0xaf, 0xb8,
	0x94, 0x17, 0xb8, 0xdd, 0x1f, 0x9b, 0x6a, 0xa8, 0x29, 0xfc, 0xff, 0x46, 0x26, 0xff, 0x76, 0x8d,
	0xec, 0x33, 0xe2, 0xaf, 0xe4, 0x05, 0x46, 0x76, 0xd6, 0x52, 0x99, 0x10, 0x39, 0x24, 0x4c, 0x18,
	0xc5, 0x99, 0xde, 0x71, 0x76, 0x9f, 0x7d, 0x3d, 0xed, 0x38, 0xdf, 0x4e, 0x3b, 0xce, 0xf7, 0xd3,
	0x8e, 0xf3, 0xe9, 0x47, 0xe7, 0xbf, 0x77, 0x5b, 0x57, 0xf8, 0x0b, 0x3c, 0xb7, 0xcb, 0x41, 0xb3,
	0x7c, 0xbe, 0x9f, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x43, 0x62, 0x55, 0x9c, 0x38, 0x06, 0x00,
	0x00,
}