// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stratumn/go-indigonode/core/app/swarm/grpc/swarm.proto

package grpc // import "github.com/stratumn/go-indigonode/core/app/swarm/grpc"

import (
	"context"
	fmt "fmt"
	io "io"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/stratumn/go-indigonode/cli/grpc/ext"

	google_golang_org_grpc "google.golang.org/grpc"
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

// The local peer request message.
type LocalPeerReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalPeerReq) Reset()         { *m = LocalPeerReq{} }
func (m *LocalPeerReq) String() string { return proto.CompactTextString(m) }
func (*LocalPeerReq) ProtoMessage()    {}
func (*LocalPeerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_swarm_02a9d2aa1756b24b, []int{0}
}
func (m *LocalPeerReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LocalPeerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LocalPeerReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *LocalPeerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalPeerReq.Merge(dst, src)
}
func (m *LocalPeerReq) XXX_Size() int {
	return m.Size()
}
func (m *LocalPeerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalPeerReq.DiscardUnknown(m)
}

var xxx_messageInfo_LocalPeerReq proto.InternalMessageInfo

// The peers request message.
type PeersReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeersReq) Reset()         { *m = PeersReq{} }
func (m *PeersReq) String() string { return proto.CompactTextString(m) }
func (*PeersReq) ProtoMessage()    {}
func (*PeersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_swarm_02a9d2aa1756b24b, []int{1}
}
func (m *PeersReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeersReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PeersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeersReq.Merge(dst, src)
}
func (m *PeersReq) XXX_Size() int {
	return m.Size()
}
func (m *PeersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PeersReq.DiscardUnknown(m)
}

var xxx_messageInfo_PeersReq proto.InternalMessageInfo

// The connections request message.
type ConnectionsReq struct {
	PeerId               []byte   `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionsReq) Reset()         { *m = ConnectionsReq{} }
func (m *ConnectionsReq) String() string { return proto.CompactTextString(m) }
func (*ConnectionsReq) ProtoMessage()    {}
func (*ConnectionsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_swarm_02a9d2aa1756b24b, []int{2}
}
func (m *ConnectionsReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConnectionsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConnectionsReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ConnectionsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionsReq.Merge(dst, src)
}
func (m *ConnectionsReq) XXX_Size() int {
	return m.Size()
}
func (m *ConnectionsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionsReq proto.InternalMessageInfo

func (m *ConnectionsReq) GetPeerId() []byte {
	if m != nil {
		return m.PeerId
	}
	return nil
}

// The peer message containing the ID of the peer.
type Peer struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_swarm_02a9d2aa1756b24b, []int{3}
}
func (m *Peer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(dst, src)
}
func (m *Peer) XXX_Size() int {
	return m.Size()
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

// The connection message containing the peer ID and addresses.
type Connection struct {
	PeerId               []byte   `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	LocalAddress         []byte   `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	RemoteAddress        []byte   `protobuf:"bytes,3,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_swarm_02a9d2aa1756b24b, []int{4}
}
func (m *Connection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Connection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Connection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Connection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connection.Merge(dst, src)
}
func (m *Connection) XXX_Size() int {
	return m.Size()
}
func (m *Connection) XXX_DiscardUnknown() {
	xxx_messageInfo_Connection.DiscardUnknown(m)
}

var xxx_messageInfo_Connection proto.InternalMessageInfo

func (m *Connection) GetPeerId() []byte {
	if m != nil {
		return m.PeerId
	}
	return nil
}

func (m *Connection) GetLocalAddress() []byte {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *Connection) GetRemoteAddress() []byte {
	if m != nil {
		return m.RemoteAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*LocalPeerReq)(nil), "stratumn.indigonode.core.app.swarm.grpc.LocalPeerReq")
	proto.RegisterType((*PeersReq)(nil), "stratumn.indigonode.core.app.swarm.grpc.PeersReq")
	proto.RegisterType((*ConnectionsReq)(nil), "stratumn.indigonode.core.app.swarm.grpc.ConnectionsReq")
	proto.RegisterType((*Peer)(nil), "stratumn.indigonode.core.app.swarm.grpc.Peer")
	proto.RegisterType((*Connection)(nil), "stratumn.indigonode.core.app.swarm.grpc.Connection")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ google_golang_org_grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = google_golang_org_grpc.SupportPackageIsVersion4

// Client API for Swarm service

type SwarmClient interface {
	// Returns the local peer.
	LocalPeer(ctx context.Context, in *LocalPeerReq, opts ...google_golang_org_grpc.CallOption) (*Peer, error)
	// Streams the connected peers.
	Peers(ctx context.Context, in *PeersReq, opts ...google_golang_org_grpc.CallOption) (Swarm_PeersClient, error)
	// Streams connections.
	Connections(ctx context.Context, in *ConnectionsReq, opts ...google_golang_org_grpc.CallOption) (Swarm_ConnectionsClient, error)
}

type swarmClient struct {
	cc *google_golang_org_grpc.ClientConn
}

func NewSwarmClient(cc *google_golang_org_grpc.ClientConn) SwarmClient {
	return &swarmClient{cc}
}

func (c *swarmClient) LocalPeer(ctx context.Context, in *LocalPeerReq, opts ...google_golang_org_grpc.CallOption) (*Peer, error) {
	out := new(Peer)
	err := c.cc.Invoke(ctx, "/stratumn.indigonode.core.app.swarm.grpc.Swarm/LocalPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swarmClient) Peers(ctx context.Context, in *PeersReq, opts ...google_golang_org_grpc.CallOption) (Swarm_PeersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Swarm_serviceDesc.Streams[0], "/stratumn.indigonode.core.app.swarm.grpc.Swarm/Peers", opts...)
	if err != nil {
		return nil, err
	}
	x := &swarmPeersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Swarm_PeersClient interface {
	Recv() (*Peer, error)
	google_golang_org_grpc.ClientStream
}

type swarmPeersClient struct {
	google_golang_org_grpc.ClientStream
}

func (x *swarmPeersClient) Recv() (*Peer, error) {
	m := new(Peer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *swarmClient) Connections(ctx context.Context, in *ConnectionsReq, opts ...google_golang_org_grpc.CallOption) (Swarm_ConnectionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Swarm_serviceDesc.Streams[1], "/stratumn.indigonode.core.app.swarm.grpc.Swarm/Connections", opts...)
	if err != nil {
		return nil, err
	}
	x := &swarmConnectionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Swarm_ConnectionsClient interface {
	Recv() (*Connection, error)
	google_golang_org_grpc.ClientStream
}

type swarmConnectionsClient struct {
	google_golang_org_grpc.ClientStream
}

func (x *swarmConnectionsClient) Recv() (*Connection, error) {
	m := new(Connection)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Swarm service

type SwarmServer interface {
	// Returns the local peer.
	LocalPeer(context.Context, *LocalPeerReq) (*Peer, error)
	// Streams the connected peers.
	Peers(*PeersReq, Swarm_PeersServer) error
	// Streams connections.
	Connections(*ConnectionsReq, Swarm_ConnectionsServer) error
}

func RegisterSwarmServer(s *google_golang_org_grpc.Server, srv SwarmServer) {
	s.RegisterService(&_Swarm_serviceDesc, srv)
}

func _Swarm_LocalPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor google_golang_org_grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocalPeerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwarmServer).LocalPeer(ctx, in)
	}
	info := &google_golang_org_grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.indigonode.core.app.swarm.grpc.Swarm/LocalPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwarmServer).LocalPeer(ctx, req.(*LocalPeerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Swarm_Peers_Handler(srv interface{}, stream google_golang_org_grpc.ServerStream) error {
	m := new(PeersReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SwarmServer).Peers(m, &swarmPeersServer{stream})
}

type Swarm_PeersServer interface {
	Send(*Peer) error
	google_golang_org_grpc.ServerStream
}

type swarmPeersServer struct {
	google_golang_org_grpc.ServerStream
}

func (x *swarmPeersServer) Send(m *Peer) error {
	return x.ServerStream.SendMsg(m)
}

func _Swarm_Connections_Handler(srv interface{}, stream google_golang_org_grpc.ServerStream) error {
	m := new(ConnectionsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SwarmServer).Connections(m, &swarmConnectionsServer{stream})
}

type Swarm_ConnectionsServer interface {
	Send(*Connection) error
	google_golang_org_grpc.ServerStream
}

type swarmConnectionsServer struct {
	google_golang_org_grpc.ServerStream
}

func (x *swarmConnectionsServer) Send(m *Connection) error {
	return x.ServerStream.SendMsg(m)
}

var _Swarm_serviceDesc = google_golang_org_grpc.ServiceDesc{
	ServiceName: "stratumn.indigonode.core.app.swarm.grpc.Swarm",
	HandlerType: (*SwarmServer)(nil),
	Methods: []google_golang_org_grpc.MethodDesc{
		{
			MethodName: "LocalPeer",
			Handler:    _Swarm_LocalPeer_Handler,
		},
	},
	Streams: []google_golang_org_grpc.StreamDesc{
		{
			StreamName:    "Peers",
			Handler:       _Swarm_Peers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Connections",
			Handler:       _Swarm_Connections_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/stratumn/go-indigonode/core/app/swarm/grpc/swarm.proto",
}

func (m *LocalPeerReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalPeerReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PeersReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeersReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ConnectionsReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectionsReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PeerId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSwarm(dAtA, i, uint64(len(m.PeerId)))
		i += copy(dAtA[i:], m.PeerId)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
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
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSwarm(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Connection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Connection) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PeerId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSwarm(dAtA, i, uint64(len(m.PeerId)))
		i += copy(dAtA[i:], m.PeerId)
	}
	if len(m.LocalAddress) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSwarm(dAtA, i, uint64(len(m.LocalAddress)))
		i += copy(dAtA[i:], m.LocalAddress)
	}
	if len(m.RemoteAddress) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSwarm(dAtA, i, uint64(len(m.RemoteAddress)))
		i += copy(dAtA[i:], m.RemoteAddress)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSwarm(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LocalPeerReq) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PeersReq) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConnectionsReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovSwarm(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Peer) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSwarm(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Connection) Size() (n int) {
	var l int
	_ = l
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovSwarm(uint64(l))
	}
	l = len(m.LocalAddress)
	if l > 0 {
		n += 1 + l + sovSwarm(uint64(l))
	}
	l = len(m.RemoteAddress)
	if l > 0 {
		n += 1 + l + sovSwarm(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSwarm(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSwarm(x uint64) (n int) {
	return sovSwarm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LocalPeerReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwarm
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
			return fmt.Errorf("proto: LocalPeerReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalPeerReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipSwarm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSwarm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PeersReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwarm
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
			return fmt.Errorf("proto: PeersReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeersReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipSwarm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSwarm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConnectionsReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwarm
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
			return fmt.Errorf("proto: ConnectionsReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectionsReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwarm
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
				return ErrInvalidLengthSwarm
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
			skippy, err := skipSwarm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSwarm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
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
				return ErrIntOverflowSwarm
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwarm
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
				return ErrInvalidLengthSwarm
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwarm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSwarm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Connection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwarm
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
			return fmt.Errorf("proto: Connection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Connection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwarm
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
				return ErrInvalidLengthSwarm
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwarm
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
				return ErrInvalidLengthSwarm
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LocalAddress = append(m.LocalAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.LocalAddress == nil {
				m.LocalAddress = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwarm
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
				return ErrInvalidLengthSwarm
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemoteAddress = append(m.RemoteAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.RemoteAddress == nil {
				m.RemoteAddress = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwarm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSwarm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSwarm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSwarm
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
					return 0, ErrIntOverflowSwarm
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
					return 0, ErrIntOverflowSwarm
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
				return 0, ErrInvalidLengthSwarm
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSwarm
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
				next, err := skipSwarm(dAtA[start:])
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
	ErrInvalidLengthSwarm = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSwarm   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/stratumn/go-indigonode/core/app/swarm/grpc/swarm.proto", fileDescriptor_swarm_02a9d2aa1756b24b)
}

var fileDescriptor_swarm_02a9d2aa1756b24b = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x3d, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0xbd, 0xde, 0x01, 0xa6, 0xd7, 0xc1, 0x42, 0x28, 0x44, 0x77, 0x91, 0x95, 0x85, 0x2f,
	0x9d, 0x03, 0x9c, 0x10, 0x08, 0xa6, 0xbb, 0x63, 0x41, 0xea, 0x80, 0xee, 0xc4, 0xc2, 0x72, 0x72,
	0x63, 0x2b, 0xb1, 0x94, 0xc4, 0xc6, 0x76, 0x55, 0xba, 0x96, 0x99, 0x9d, 0x81, 0x89, 0xff, 0xd0,
	0xa9, 0x3b, 0x62, 0xe4, 0x27, 0xa0, 0xf2, 0x07, 0x18, 0x19, 0x91, 0x9d, 0xd2, 0x86, 0x4e, 0xa1,
	0x43, 0xa2, 0x57, 0x79, 0x9f, 0xe7, 0x7d, 0x9e, 0xbc, 0x1f, 0xf0, 0x24, 0x13, 0x36, 0x1f, 0x0d,
	0x49, 0x2a, 0xcb, 0xc4, 0x58, 0x4d, 0xed, 0xa8, 0xac, 0x92, 0x4c, 0x1e, 0x89, 0x8a, 0x89, 0x4c,
	0x56, 0x92, 0xf1, 0x24, 0x95, 0x9a, 0x27, 0x54, 0xa9, 0xc4, 0x8c, 0xa9, 0x2e, 0x93, 0x4c, 0xab,
	0xb4, 0x0e, 0x89, 0xd2, 0xd2, 0x4a, 0x74, 0xe7, 0x2f, 0x8f, 0xac, 0x49, 0xc4, 0x91, 0x08, 0x55,
	0x8a, 0xd4, 0x48, 0x47, 0x0a, 0x9f, 0xb5, 0xd0, 0x2a, 0x44, 0x2d, 0xc0, 0xdf, 0x5b, 0xf7, 0xd4,
	0x12, 0x71, 0x1f, 0xf6, 0x06, 0x32, 0xa5, 0xc5, 0x6b, 0xce, 0xf5, 0x39, 0x7f, 0x17, 0x43, 0x78,
	0xcd, 0x85, 0xc6, 0xc5, 0x6f, 0x60, 0xff, 0x4c, 0x56, 0x15, 0x4f, 0xad, 0x90, 0x95, 0xfb, 0x82,
	0xce, 0xe0, 0x55, 0xc5, 0xb9, 0xbe, 0x14, 0x2c, 0x00, 0x18, 0xdc, 0xed, 0x9d, 0xde, 0xff, 0x32,
	0x0f, 0xe2, 0x8b, 0x5c, 0x8e, 0xb1, 0xac, 0x8a, 0x09, 0x4e, 0xd7, 0x70, 0x6c, 0x25, 0xb6, 0xb9,
	0x30, 0xd8, 0x11, 0x7e, 0xcd, 0x03, 0x70, 0xbe, 0xe7, 0xa2, 0x57, 0x2c, 0x3e, 0x80, 0x5d, 0x27,
	0x81, 0x6e, 0xc2, 0xce, 0xaa, 0x4e, 0xd7, 0x23, 0x3a, 0x82, 0xc5, 0x1f, 0x00, 0x84, 0x6b, 0x55,
	0x74, 0xb8, 0xa9, 0xd8, 0x6d, 0xd6, 0x42, 0xf7, 0xe0, 0x7e, 0xe1, 0xec, 0x5f, 0x52, 0xc6, 0x34,
	0x37, 0x26, 0xe8, 0xd4, 0xa0, 0xdf, 0x0e, 0xd4, 0xf3, 0xa9, 0x93, 0x3a, 0x83, 0x1e, 0xc0, 0xbe,
	0xe6, 0xa5, 0xb4, 0x7c, 0x85, 0xdd, 0x69, 0x60, 0xf7, 0xeb, 0xdc, 0x12, 0xfc, 0xf8, 0xeb, 0x0e,
	0xdc, 0xbd, 0x70, 0xfd, 0x45, 0x1f, 0x01, 0xbc, 0xbe, 0xea, 0x10, 0x7a, 0x42, 0x5a, 0x8e, 0x84,
	0x34, 0xbb, 0x1a, 0x1e, 0xb5, 0xa6, 0x39, 0x46, 0x1c, 0x4e, 0x67, 0xc1, 0xad, 0x97, 0xc2, 0xa8,
	0x82, 0x4e, 0xb0, 0xcd, 0x39, 0xf6, 0x3f, 0xe2, 0x3b, 0xe9, 0xfc, 0xec, 0xfa, 0x09, 0xa1, 0x47,
	0xff, 0x55, 0xd4, 0x6c, 0xe1, 0x03, 0x4f, 0x67, 0xc1, 0xc1, 0x40, 0x18, 0x8b, 0x69, 0x51, 0x78,
	0x23, 0xcb, 0x01, 0x73, 0xe6, 0xcd, 0x98, 0x87, 0x00, 0x7d, 0x06, 0xf0, 0x46, 0x63, 0x4b, 0xd0,
	0xd3, 0xd6, 0x12, 0xff, 0xee, 0x56, 0x78, 0xbc, 0x05, 0x31, 0x3e, 0x9c, 0xce, 0x82, 0xdb, 0xde,
	0xe1, 0xc6, 0xea, 0x2d, 0xed, 0x9d, 0x0e, 0xbe, 0x2d, 0x22, 0xf0, 0x7d, 0x11, 0x81, 0x1f, 0x8b,
	0x08, 0x7c, 0xfa, 0x19, 0x5d, 0x79, 0xfb, 0x7c, 0xab, 0xbb, 0x7c, 0xe1, 0x5e, 0xc3, 0x3d, 0x7f,
	0x34, 0xc7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x92, 0xf4, 0x3a, 0xa7, 0xdc, 0x03, 0x00, 0x00,
}
