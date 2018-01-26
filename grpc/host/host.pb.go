// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stratumn/alice/grpc/host/host.proto

/*
	Package host is a generated protocol buffer package.

	It is generated from these files:
		github.com/stratumn/alice/grpc/host/host.proto

	It has these top-level messages:
		IdReq
		AddressesReq
		ConnectReq
		HostId
		Address
		Connection
*/
package host

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

// The ID request message.
type IdReq struct {
}

func (m *IdReq) Reset()                    { *m = IdReq{} }
func (m *IdReq) String() string            { return proto.CompactTextString(m) }
func (*IdReq) ProtoMessage()               {}
func (*IdReq) Descriptor() ([]byte, []int) { return fileDescriptorHost, []int{0} }

// The address request message.
type AddressesReq struct {
}

func (m *AddressesReq) Reset()                    { *m = AddressesReq{} }
func (m *AddressesReq) String() string            { return proto.CompactTextString(m) }
func (*AddressesReq) ProtoMessage()               {}
func (*AddressesReq) Descriptor() ([]byte, []int) { return fileDescriptorHost, []int{1} }

// The connect request message.
type ConnectReq struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *ConnectReq) Reset()                    { *m = ConnectReq{} }
func (m *ConnectReq) String() string            { return proto.CompactTextString(m) }
func (*ConnectReq) ProtoMessage()               {}
func (*ConnectReq) Descriptor() ([]byte, []int) { return fileDescriptorHost, []int{2} }

func (m *ConnectReq) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

// The host ID message containing the ID of the host.
type HostId struct {
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *HostId) Reset()                    { *m = HostId{} }
func (m *HostId) String() string            { return proto.CompactTextString(m) }
func (*HostId) ProtoMessage()               {}
func (*HostId) Descriptor() ([]byte, []int) { return fileDescriptorHost, []int{3} }

func (m *HostId) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

// The address message containing a multiaddress.
type Address struct {
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptorHost, []int{4} }

func (m *Address) GetAddress() []byte {
	if m != nil {
		return m.Address
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
func (*Connection) Descriptor() ([]byte, []int) { return fileDescriptorHost, []int{5} }

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
	proto.RegisterType((*IdReq)(nil), "stratumn.alice.grpc.host.IdReq")
	proto.RegisterType((*AddressesReq)(nil), "stratumn.alice.grpc.host.AddressesReq")
	proto.RegisterType((*ConnectReq)(nil), "stratumn.alice.grpc.host.ConnectReq")
	proto.RegisterType((*HostId)(nil), "stratumn.alice.grpc.host.HostId")
	proto.RegisterType((*Address)(nil), "stratumn.alice.grpc.host.Address")
	proto.RegisterType((*Connection)(nil), "stratumn.alice.grpc.host.Connection")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Host service

type HostClient interface {
	// Returns the host ID.
	ID(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*HostId, error)
	// Streams the host addresses.
	Addresses(ctx context.Context, in *AddressesReq, opts ...grpc.CallOption) (Host_AddressesClient, error)
	// Connects to a multiaddress.
	Connect(ctx context.Context, in *ConnectReq, opts ...grpc.CallOption) (Host_ConnectClient, error)
}

type hostClient struct {
	cc *grpc.ClientConn
}

func NewHostClient(cc *grpc.ClientConn) HostClient {
	return &hostClient{cc}
}

func (c *hostClient) ID(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*HostId, error) {
	out := new(HostId)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.host.Host/ID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostClient) Addresses(ctx context.Context, in *AddressesReq, opts ...grpc.CallOption) (Host_AddressesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Host_serviceDesc.Streams[0], c.cc, "/stratumn.alice.grpc.host.Host/Addresses", opts...)
	if err != nil {
		return nil, err
	}
	x := &hostAddressesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Host_AddressesClient interface {
	Recv() (*Address, error)
	grpc.ClientStream
}

type hostAddressesClient struct {
	grpc.ClientStream
}

func (x *hostAddressesClient) Recv() (*Address, error) {
	m := new(Address)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hostClient) Connect(ctx context.Context, in *ConnectReq, opts ...grpc.CallOption) (Host_ConnectClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Host_serviceDesc.Streams[1], c.cc, "/stratumn.alice.grpc.host.Host/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &hostConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Host_ConnectClient interface {
	Recv() (*Connection, error)
	grpc.ClientStream
}

type hostConnectClient struct {
	grpc.ClientStream
}

func (x *hostConnectClient) Recv() (*Connection, error) {
	m := new(Connection)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Host service

type HostServer interface {
	// Returns the host ID.
	ID(context.Context, *IdReq) (*HostId, error)
	// Streams the host addresses.
	Addresses(*AddressesReq, Host_AddressesServer) error
	// Connects to a multiaddress.
	Connect(*ConnectReq, Host_ConnectServer) error
}

func RegisterHostServer(s *grpc.Server, srv HostServer) {
	s.RegisterService(&_Host_serviceDesc, srv)
}

func _Host_ID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServer).ID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.host.Host/ID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServer).ID(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Host_Addresses_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AddressesReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HostServer).Addresses(m, &hostAddressesServer{stream})
}

type Host_AddressesServer interface {
	Send(*Address) error
	grpc.ServerStream
}

type hostAddressesServer struct {
	grpc.ServerStream
}

func (x *hostAddressesServer) Send(m *Address) error {
	return x.ServerStream.SendMsg(m)
}

func _Host_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HostServer).Connect(m, &hostConnectServer{stream})
}

type Host_ConnectServer interface {
	Send(*Connection) error
	grpc.ServerStream
}

type hostConnectServer struct {
	grpc.ServerStream
}

func (x *hostConnectServer) Send(m *Connection) error {
	return x.ServerStream.SendMsg(m)
}

var _Host_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stratumn.alice.grpc.host.Host",
	HandlerType: (*HostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ID",
			Handler:    _Host_ID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Addresses",
			Handler:       _Host_Addresses_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Connect",
			Handler:       _Host_Connect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/stratumn/alice/grpc/host/host.proto",
}

func (m *IdReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *AddressesReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddressesReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ConnectReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHost(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	return i, nil
}

func (m *HostId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostId) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHost(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	return i, nil
}

func (m *Address) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Address) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHost(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
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
		i = encodeVarintHost(dAtA, i, uint64(len(m.PeerId)))
		i += copy(dAtA[i:], m.PeerId)
	}
	if len(m.LocalAddress) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHost(dAtA, i, uint64(len(m.LocalAddress)))
		i += copy(dAtA[i:], m.LocalAddress)
	}
	if len(m.RemoteAddress) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHost(dAtA, i, uint64(len(m.RemoteAddress)))
		i += copy(dAtA[i:], m.RemoteAddress)
	}
	return i, nil
}

func encodeVarintHost(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *IdReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *AddressesReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ConnectReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	return n
}

func (m *HostId) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	return n
}

func (m *Address) Size() (n int) {
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	return n
}

func (m *Connection) Size() (n int) {
	var l int
	_ = l
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	l = len(m.LocalAddress)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	l = len(m.RemoteAddress)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	return n
}

func sovHost(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHost(x uint64) (n int) {
	return sovHost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IdReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
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
			return fmt.Errorf("proto: IdReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
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
func (m *AddressesReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
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
			return fmt.Errorf("proto: AddressesReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddressesReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
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
func (m *ConnectReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
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
			return fmt.Errorf("proto: ConnectReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
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
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
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
func (m *HostId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
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
			return fmt.Errorf("proto: HostId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
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
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
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
func (m *Address) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
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
			return fmt.Errorf("proto: Address: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Address: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
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
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
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
				return ErrIntOverflowHost
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
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
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
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
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
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
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
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
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
func skipHost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHost
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
					return 0, ErrIntOverflowHost
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
					return 0, ErrIntOverflowHost
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
				return 0, ErrInvalidLengthHost
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHost
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
				next, err := skipHost(dAtA[start:])
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
	ErrInvalidLengthHost = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHost   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("github.com/stratumn/alice/grpc/host/host.proto", fileDescriptorHost) }

var fileDescriptorHost = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0x87, 0x71, 0x18, 0x8d, 0x78, 0xd5, 0xed, 0x60, 0x81, 0x64, 0x2a, 0x2d, 0x78, 0x11, 0x82,
	0x55, 0x80, 0x3b, 0x8d, 0x1b, 0x5c, 0x60, 0x14, 0x89, 0x48, 0x70, 0xe9, 0x91, 0xcb, 0xe4, 0x26,
	0xd6, 0x6a, 0x29, 0x8d, 0x33, 0xdb, 0x91, 0x40, 0xdc, 0x8a, 0xc4, 0x67, 0xe0, 0xcc, 0x89, 0x2f,
	0xd0, 0x53, 0xbf, 0x00, 0x47, 0x3e, 0x02, 0x2a, 0x5f, 0x60, 0x47, 0x8e, 0xc8, 0xf9, 0xb7, 0x82,
	0x88, 0xba, 0x43, 0x2c, 0xc5, 0xbf, 0x27, 0xcf, 0x9b, 0xbc, 0xaf, 0x03, 0xec, 0x4c, 0xda, 0x59,
	0x31, 0x65, 0xb1, 0x9a, 0x8f, 0x8c, 0xd5, 0xdc, 0x16, 0xf3, 0x6c, 0xc4, 0x53, 0x19, 0x8b, 0xd1,
	0x99, 0xce, 0xe3, 0xd1, 0x4c, 0x19, 0x5b, 0x2e, 0x2c, 0xd7, 0xca, 0x2a, 0x4c, 0x1a, 0x88, 0x95,
	0x10, 0x73, 0x10, 0x73, 0xf9, 0xe0, 0xd1, 0x16, 0x93, 0x78, 0x6f, 0xdd, 0x55, 0x79, 0x42, 0x1f,
	0x6e, 0x44, 0xc9, 0x44, 0x9c, 0x87, 0x7b, 0xd0, 0x7f, 0x91, 0x24, 0x5a, 0x18, 0x23, 0x8c, 0xbb,
	0x7f, 0x0e, 0xf0, 0x52, 0x65, 0x99, 0x88, 0xed, 0x44, 0x9c, 0xe3, 0x63, 0xf0, 0x79, 0x95, 0x12,
	0x44, 0xd1, 0x61, 0xff, 0x84, 0x7c, 0x5d, 0x91, 0xfe, 0xdb, 0x22, 0xb5, 0xb2, 0xde, 0xff, 0xb6,
	0x22, 0xe8, 0xf7, 0x8a, 0xa0, 0x49, 0x03, 0x86, 0x01, 0xf4, 0x5e, 0x2b, 0x63, 0xa3, 0x04, 0xdf,
	0x02, 0x4f, 0x26, 0xf5, 0x83, 0x3b, 0x17, 0x0e, 0xf2, 0x64, 0x12, 0x0e, 0xc1, 0xaf, 0x2b, 0xe2,
	0xe0, 0x52, 0xef, 0x55, 0xd4, 0xdf, 0xaa, 0x4f, 0xa8, 0x7d, 0x1b, 0xa9, 0x32, 0xbc, 0x0f, 0x7e,
	0x2e, 0x84, 0x3e, 0xfd, 0x47, 0xda, 0x73, 0x9b, 0x51, 0x82, 0x87, 0xb0, 0x9b, 0xaa, 0x98, 0xa7,
	0xa7, 0xff, 0x73, 0xf6, 0xcb, 0xa8, 0x29, 0xfc, 0x10, 0xf6, 0xb4, 0x98, 0x2b, 0x2b, 0x5a, 0xf6,
	0xfa, 0x06, 0xbb, 0x5b, 0x65, 0x35, 0x7c, 0x7c, 0xe1, 0xc1, 0x8e, 0xfb, 0x22, 0x3c, 0x05, 0x2f,
	0x1a, 0xe3, 0xbb, 0xac, 0x6b, 0x06, 0xac, 0x6c, 0xe9, 0x80, 0x76, 0x03, 0x55, 0x63, 0xc2, 0x3b,
	0x8b, 0x25, 0xb9, 0x3d, 0x96, 0x26, 0x4f, 0xf9, 0x07, 0x6a, 0x67, 0x82, 0xba, 0xf0, 0x81, 0xa1,
	0xd1, 0x18, 0x7f, 0x84, 0x9b, 0xed, 0x3c, 0xf0, 0xfd, 0x6e, 0xd3, 0xe6, 0xd0, 0x06, 0x07, 0x5b,
	0xb9, 0xf0, 0x60, 0xb1, 0x24, 0xfb, 0x6f, 0xa4, 0xb1, 0x94, 0xa7, 0xe9, 0x66, 0x4d, 0xde, 0x88,
	0x8e, 0x10, 0xfe, 0x8c, 0xc0, 0xaf, 0xfb, 0x8d, 0xef, 0x75, 0x3b, 0x2f, 0x0f, 0xc8, 0x60, 0x3b,
	0x25, 0x55, 0x16, 0x3e, 0x5e, 0x2c, 0xc9, 0xf0, 0x55, 0x66, 0x0a, 0x2d, 0x5c, 0x69, 0x2d, 0xa8,
	0x34, 0x94, 0xd3, 0xb8, 0x25, 0xa8, 0x55, 0xd4, 0xce, 0xa4, 0xa1, 0x6e, 0x94, 0x47, 0xe8, 0xe4,
	0xe9, 0xf7, 0x75, 0x80, 0x7e, 0xac, 0x03, 0xf4, 0x73, 0x1d, 0xa0, 0x2f, 0xbf, 0x82, 0x6b, 0xef,
	0x0e, 0xaf, 0xf0, 0xa3, 0x3c, 0x73, 0xcb, 0xb4, 0x57, 0x9e, 0xf0, 0x27, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x21, 0x35, 0xb5, 0x13, 0x5b, 0x03, 0x00, 0x00,
}
