// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stratumn/alice/grpc/swarm/swarm.proto

/*
	Package swarm is a generated protocol buffer package.

	It is generated from these files:
		github.com/stratumn/alice/grpc/swarm/swarm.proto

	It has these top-level messages:
		LocalPeerReq
		PeersReq
		ConnectionsReq
		Peer
		Connection
*/
package swarm

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

// The local peer request message.
type LocalPeerReq struct {
}

func (m *LocalPeerReq) Reset()                    { *m = LocalPeerReq{} }
func (m *LocalPeerReq) String() string            { return proto.CompactTextString(m) }
func (*LocalPeerReq) ProtoMessage()               {}
func (*LocalPeerReq) Descriptor() ([]byte, []int) { return fileDescriptorSwarm, []int{0} }

// The peers request message.
type PeersReq struct {
}

func (m *PeersReq) Reset()                    { *m = PeersReq{} }
func (m *PeersReq) String() string            { return proto.CompactTextString(m) }
func (*PeersReq) ProtoMessage()               {}
func (*PeersReq) Descriptor() ([]byte, []int) { return fileDescriptorSwarm, []int{1} }

// The connections request message.
type ConnectionsReq struct {
	PeerId []byte `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
}

func (m *ConnectionsReq) Reset()                    { *m = ConnectionsReq{} }
func (m *ConnectionsReq) String() string            { return proto.CompactTextString(m) }
func (*ConnectionsReq) ProtoMessage()               {}
func (*ConnectionsReq) Descriptor() ([]byte, []int) { return fileDescriptorSwarm, []int{2} }

func (m *ConnectionsReq) GetPeerId() []byte {
	if m != nil {
		return m.PeerId
	}
	return nil
}

// The peer message containing the ID of the peer.
type Peer struct {
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptorSwarm, []int{3} }

func (m *Peer) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

// The connection message containing the peer ID and addresses.
type Connection struct {
	PeerId        []byte `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	LocalAddress  []byte `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	RemoteAddress []byte `protobuf:"bytes,3,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
}

func (m *Connection) Reset()                    { *m = Connection{} }
func (m *Connection) String() string            { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()               {}
func (*Connection) Descriptor() ([]byte, []int) { return fileDescriptorSwarm, []int{4} }

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
	proto.RegisterType((*LocalPeerReq)(nil), "stratumn.alice.grpc.swarm.LocalPeerReq")
	proto.RegisterType((*PeersReq)(nil), "stratumn.alice.grpc.swarm.PeersReq")
	proto.RegisterType((*ConnectionsReq)(nil), "stratumn.alice.grpc.swarm.ConnectionsReq")
	proto.RegisterType((*Peer)(nil), "stratumn.alice.grpc.swarm.Peer")
	proto.RegisterType((*Connection)(nil), "stratumn.alice.grpc.swarm.Connection")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Swarm service

type SwarmClient interface {
	// Returns the local peer.
	LocalPeer(ctx context.Context, in *LocalPeerReq, opts ...grpc.CallOption) (*Peer, error)
	// Streams the connected peers.
	Peers(ctx context.Context, in *PeersReq, opts ...grpc.CallOption) (Swarm_PeersClient, error)
	// Streams connections.
	Connections(ctx context.Context, in *ConnectionsReq, opts ...grpc.CallOption) (Swarm_ConnectionsClient, error)
}

type swarmClient struct {
	cc *grpc.ClientConn
}

func NewSwarmClient(cc *grpc.ClientConn) SwarmClient {
	return &swarmClient{cc}
}

func (c *swarmClient) LocalPeer(ctx context.Context, in *LocalPeerReq, opts ...grpc.CallOption) (*Peer, error) {
	out := new(Peer)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.swarm.Swarm/LocalPeer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swarmClient) Peers(ctx context.Context, in *PeersReq, opts ...grpc.CallOption) (Swarm_PeersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Swarm_serviceDesc.Streams[0], c.cc, "/stratumn.alice.grpc.swarm.Swarm/Peers", opts...)
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
	grpc.ClientStream
}

type swarmPeersClient struct {
	grpc.ClientStream
}

func (x *swarmPeersClient) Recv() (*Peer, error) {
	m := new(Peer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *swarmClient) Connections(ctx context.Context, in *ConnectionsReq, opts ...grpc.CallOption) (Swarm_ConnectionsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Swarm_serviceDesc.Streams[1], c.cc, "/stratumn.alice.grpc.swarm.Swarm/Connections", opts...)
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
	grpc.ClientStream
}

type swarmConnectionsClient struct {
	grpc.ClientStream
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

func RegisterSwarmServer(s *grpc.Server, srv SwarmServer) {
	s.RegisterService(&_Swarm_serviceDesc, srv)
}

func _Swarm_LocalPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocalPeerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwarmServer).LocalPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.swarm.Swarm/LocalPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwarmServer).LocalPeer(ctx, req.(*LocalPeerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Swarm_Peers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PeersReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SwarmServer).Peers(m, &swarmPeersServer{stream})
}

type Swarm_PeersServer interface {
	Send(*Peer) error
	grpc.ServerStream
}

type swarmPeersServer struct {
	grpc.ServerStream
}

func (x *swarmPeersServer) Send(m *Peer) error {
	return x.ServerStream.SendMsg(m)
}

func _Swarm_Connections_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectionsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SwarmServer).Connections(m, &swarmConnectionsServer{stream})
}

type Swarm_ConnectionsServer interface {
	Send(*Connection) error
	grpc.ServerStream
}

type swarmConnectionsServer struct {
	grpc.ServerStream
}

func (x *swarmConnectionsServer) Send(m *Connection) error {
	return x.ServerStream.SendMsg(m)
}

var _Swarm_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stratumn.alice.grpc.swarm.Swarm",
	HandlerType: (*SwarmServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LocalPeer",
			Handler:    _Swarm_LocalPeer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
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
	Metadata: "github.com/stratumn/alice/grpc/swarm/swarm.proto",
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
	return n
}

func (m *PeersReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ConnectionsReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovSwarm(uint64(l))
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
	proto.RegisterFile("github.com/stratumn/alice/grpc/swarm/swarm.proto", fileDescriptorSwarm)
}

var fileDescriptorSwarm = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x8e, 0xd2, 0x40,
	0x18, 0xc7, 0x9d, 0xca, 0xae, 0xfa, 0xc9, 0x72, 0x98, 0x18, 0xd3, 0x6d, 0x76, 0xeb, 0xa4, 0xc6,
	0xe8, 0xae, 0xa6, 0x25, 0x7a, 0xd4, 0x8b, 0xe0, 0xc5, 0x84, 0x83, 0x81, 0x78, 0xf1, 0x42, 0x4a,
	0x3b, 0xa1, 0x93, 0xb4, 0x9d, 0x32, 0x33, 0x04, 0xb9, 0x19, 0x7c, 0x09, 0xcf, 0xbe, 0x03, 0x27,
	0x5e, 0xc0, 0xa3, 0x8f, 0x60, 0xf0, 0x05, 0x38, 0x7a, 0x34, 0x33, 0x45, 0xa8, 0x24, 0x0b, 0x1c,
	0x3a, 0x99, 0xf4, 0xfb, 0xf5, 0xff, 0xff, 0xf2, 0xff, 0xbe, 0x42, 0x73, 0xc8, 0x54, 0x32, 0x1e,
	0xf8, 0x11, 0xcf, 0x02, 0xa9, 0x44, 0xa8, 0xc6, 0x59, 0x1e, 0x84, 0x29, 0x8b, 0x68, 0x30, 0x14,
	0x45, 0x14, 0xc8, 0x49, 0x28, 0xb2, 0xf2, 0xf4, 0x0b, 0xc1, 0x15, 0xc7, 0xe7, 0xff, 0x30, 0xdf,
	0x60, 0xbe, 0xc6, 0x7c, 0x03, 0x38, 0x2f, 0x0e, 0x88, 0xd1, 0xcf, 0x4a, 0x3f, 0xa5, 0x90, 0xd7,
	0x80, 0x7a, 0x87, 0x47, 0x61, 0xfa, 0x81, 0x52, 0xd1, 0xa5, 0x23, 0x0f, 0xe0, 0xae, 0xbe, 0x4a,
	0x7d, 0xff, 0x08, 0x8d, 0x36, 0xcf, 0x73, 0x1a, 0x29, 0xc6, 0x73, 0xfd, 0x06, 0xb7, 0xe1, 0x4e,
	0x41, 0xa9, 0xe8, 0xb3, 0xd8, 0x46, 0x04, 0x3d, 0xab, 0xb7, 0xae, 0xbf, 0x2f, 0x6c, 0xaf, 0x97,
	0xf0, 0x09, 0xe1, 0x79, 0x3a, 0x25, 0xd1, 0x16, 0x27, 0x8a, 0x13, 0x95, 0x30, 0x49, 0xf4, 0x07,
	0xab, 0x85, 0x8d, 0xba, 0xa7, 0xfa, 0xf6, 0x3e, 0xf6, 0x2e, 0xa0, 0xa6, 0x2d, 0xf0, 0x03, 0xb0,
	0x36, 0x3a, 0x35, 0x43, 0x58, 0x2c, 0xf6, 0xbe, 0x22, 0x80, 0xad, 0x2b, 0xbe, 0xdc, 0x75, 0xac,
	0x55, 0xb5, 0xf0, 0x15, 0x9c, 0xa5, 0xba, 0xfd, 0x7e, 0x18, 0xc7, 0x82, 0x4a, 0x69, 0x5b, 0x25,
	0xf4, 0x47, 0x43, 0x75, 0x53, 0x7a, 0x5b, 0x56, 0xf0, 0x73, 0x68, 0x08, 0x9a, 0x71, 0x45, 0x37,
	0xec, 0xed, 0x0a, 0x7b, 0x56, 0xd6, 0xd6, 0xf0, 0xcb, 0x95, 0x05, 0x27, 0x3d, 0x1d, 0x27, 0x1e,
	0xc1, 0xbd, 0x4d, 0x40, 0xf8, 0xa9, 0x7f, 0x63, 0xee, 0x7e, 0x35, 0x46, 0xe7, 0xd1, 0x1e, 0x50,
	0x33, 0x9e, 0x33, 0x9b, 0xdb, 0x0f, 0xdf, 0x31, 0x59, 0xa4, 0xe1, 0x94, 0xa8, 0x84, 0x12, 0xd3,
	0xab, 0x09, 0x0b, 0x8f, 0xe0, 0xc4, 0xcc, 0x00, 0x3f, 0x3e, 0xa0, 0x22, 0x8f, 0xb2, 0x22, 0xb3,
	0xb9, 0x7d, 0xd1, 0x61, 0x52, 0x91, 0x30, 0x4d, 0x8d, 0xd7, 0x7a, 0x4c, 0x34, 0x36, 0x7e, 0xb2,
	0x89, 0xf0, 0x17, 0x04, 0xf7, 0x2b, 0xb3, 0xc6, 0x57, 0x7b, 0x44, 0xff, 0xdf, 0x09, 0xe7, 0xc9,
	0x51, 0xa8, 0x77, 0x39, 0x9b, 0xdb, 0xe7, 0xa6, 0x8b, 0x9d, 0x25, 0x59, 0xb7, 0xd0, 0x7a, 0xf3,
	0x63, 0xe9, 0xa2, 0x9f, 0x4b, 0x17, 0xfd, 0x5a, 0xba, 0xe8, 0xdb, 0x6f, 0xf7, 0xd6, 0xa7, 0xeb,
	0x63, 0x7e, 0x8b, 0xd7, 0xe6, 0x1c, 0x9c, 0x9a, 0x75, 0x7e, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xdc, 0x2d, 0xd6, 0x84, 0x4b, 0x03, 0x00, 0x00,
}
