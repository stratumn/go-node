// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stratumn/alice/grpc/manager/manager.proto

/*
	Package manager is a generated protocol buffer package.

	It is generated from these files:
		github.com/stratumn/alice/grpc/manager/manager.proto

	It has these top-level messages:
		ListReq
		InfoReq
		StartReq
		StopReq
		PruneReq
		Service
*/
package manager

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/stratumn/alice/grpc/ext"

import (
	context "context"
	grpc "google.golang.org/grpc"
)

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

type Service_Status int32

const (
	Service_STOPPED  Service_Status = 0
	Service_STARTING Service_Status = 1
	Service_RUNNING  Service_Status = 2
	Service_STOPPING Service_Status = 3
	Service_ERRORED  Service_Status = 4
)

var Service_Status_name = map[int32]string{
	0: "STOPPED",
	1: "STARTING",
	2: "RUNNING",
	3: "STOPPING",
	4: "ERRORED",
}
var Service_Status_value = map[string]int32{
	"STOPPED":  0,
	"STARTING": 1,
	"RUNNING":  2,
	"STOPPING": 3,
	"ERRORED":  4,
}

func (x Service_Status) String() string {
	return proto.EnumName(Service_Status_name, int32(x))
}
func (Service_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptorManager, []int{5, 0} }

// The list request message.
type ListReq struct {
}

func (m *ListReq) Reset()                    { *m = ListReq{} }
func (m *ListReq) String() string            { return proto.CompactTextString(m) }
func (*ListReq) ProtoMessage()               {}
func (*ListReq) Descriptor() ([]byte, []int) { return fileDescriptorManager, []int{0} }

// The info request message.
type InfoReq struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *InfoReq) Reset()                    { *m = InfoReq{} }
func (m *InfoReq) String() string            { return proto.CompactTextString(m) }
func (*InfoReq) ProtoMessage()               {}
func (*InfoReq) Descriptor() ([]byte, []int) { return fileDescriptorManager, []int{1} }

func (m *InfoReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// The start request message.
type StartReq struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *StartReq) Reset()                    { *m = StartReq{} }
func (m *StartReq) String() string            { return proto.CompactTextString(m) }
func (*StartReq) ProtoMessage()               {}
func (*StartReq) Descriptor() ([]byte, []int) { return fileDescriptorManager, []int{2} }

func (m *StartReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// The stop request message.
type StopReq struct {
	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Prune bool   `protobuf:"varint,2,opt,name=prune,proto3" json:"prune,omitempty"`
}

func (m *StopReq) Reset()                    { *m = StopReq{} }
func (m *StopReq) String() string            { return proto.CompactTextString(m) }
func (*StopReq) ProtoMessage()               {}
func (*StopReq) Descriptor() ([]byte, []int) { return fileDescriptorManager, []int{3} }

func (m *StopReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StopReq) GetPrune() bool {
	if m != nil {
		return m.Prune
	}
	return false
}

// The prune request message.
type PruneReq struct {
}

func (m *PruneReq) Reset()                    { *m = PruneReq{} }
func (m *PruneReq) String() string            { return proto.CompactTextString(m) }
func (*PruneReq) ProtoMessage()               {}
func (*PruneReq) Descriptor() ([]byte, []int) { return fileDescriptorManager, []int{4} }

// The service message containing information about a service.
type Service struct {
	Id        string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status    Service_Status `protobuf:"varint,2,opt,name=status,proto3,enum=stratumn.alice.grpc.manager.Service_Status" json:"status,omitempty"`
	Needs     []string       `protobuf:"bytes,3,rep,name=needs" json:"needs,omitempty"`
	Stoppable bool           `protobuf:"varint,4,opt,name=stoppable,proto3" json:"stoppable,omitempty"`
	Prunable  bool           `protobuf:"varint,5,opt,name=prunable,proto3" json:"prunable,omitempty"`
	Name      string         `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Desc      string         `protobuf:"bytes,7,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptorManager, []int{5} }

func (m *Service) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Service) GetStatus() Service_Status {
	if m != nil {
		return m.Status
	}
	return Service_STOPPED
}

func (m *Service) GetNeeds() []string {
	if m != nil {
		return m.Needs
	}
	return nil
}

func (m *Service) GetStoppable() bool {
	if m != nil {
		return m.Stoppable
	}
	return false
}

func (m *Service) GetPrunable() bool {
	if m != nil {
		return m.Prunable
	}
	return false
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func init() {
	proto.RegisterType((*ListReq)(nil), "stratumn.alice.grpc.manager.ListReq")
	proto.RegisterType((*InfoReq)(nil), "stratumn.alice.grpc.manager.InfoReq")
	proto.RegisterType((*StartReq)(nil), "stratumn.alice.grpc.manager.StartReq")
	proto.RegisterType((*StopReq)(nil), "stratumn.alice.grpc.manager.StopReq")
	proto.RegisterType((*PruneReq)(nil), "stratumn.alice.grpc.manager.PruneReq")
	proto.RegisterType((*Service)(nil), "stratumn.alice.grpc.manager.Service")
	proto.RegisterEnum("stratumn.alice.grpc.manager.Service_Status", Service_Status_name, Service_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Manager service

type ManagerClient interface {
	// Streams the registered services.
	List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (Manager_ListClient, error)
	// Returns information about a service.
	Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*Service, error)
	// Starts a service.
	Start(ctx context.Context, in *StartReq, opts ...grpc.CallOption) (*Service, error)
	// Stops a service.
	Stop(ctx context.Context, in *StopReq, opts ...grpc.CallOption) (*Service, error)
	// Prunes services.
	Prune(ctx context.Context, in *PruneReq, opts ...grpc.CallOption) (Manager_PruneClient, error)
}

type managerClient struct {
	cc *grpc.ClientConn
}

func NewManagerClient(cc *grpc.ClientConn) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (Manager_ListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Manager_serviceDesc.Streams[0], c.cc, "/stratumn.alice.grpc.manager.Manager/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &managerListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Manager_ListClient interface {
	Recv() (*Service, error)
	grpc.ClientStream
}

type managerListClient struct {
	grpc.ClientStream
}

func (x *managerListClient) Recv() (*Service, error) {
	m := new(Service)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *managerClient) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.manager.Manager/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Start(ctx context.Context, in *StartReq, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.manager.Manager/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Stop(ctx context.Context, in *StopReq, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := grpc.Invoke(ctx, "/stratumn.alice.grpc.manager.Manager/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Prune(ctx context.Context, in *PruneReq, opts ...grpc.CallOption) (Manager_PruneClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Manager_serviceDesc.Streams[1], c.cc, "/stratumn.alice.grpc.manager.Manager/Prune", opts...)
	if err != nil {
		return nil, err
	}
	x := &managerPruneClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Manager_PruneClient interface {
	Recv() (*Service, error)
	grpc.ClientStream
}

type managerPruneClient struct {
	grpc.ClientStream
}

func (x *managerPruneClient) Recv() (*Service, error) {
	m := new(Service)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Manager service

type ManagerServer interface {
	// Streams the registered services.
	List(*ListReq, Manager_ListServer) error
	// Returns information about a service.
	Info(context.Context, *InfoReq) (*Service, error)
	// Starts a service.
	Start(context.Context, *StartReq) (*Service, error)
	// Stops a service.
	Stop(context.Context, *StopReq) (*Service, error)
	// Prunes services.
	Prune(*PruneReq, Manager_PruneServer) error
}

func RegisterManagerServer(s *grpc.Server, srv ManagerServer) {
	s.RegisterService(&_Manager_serviceDesc, srv)
}

func _Manager_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManagerServer).List(m, &managerListServer{stream})
}

type Manager_ListServer interface {
	Send(*Service) error
	grpc.ServerStream
}

type managerListServer struct {
	grpc.ServerStream
}

func (x *managerListServer) Send(m *Service) error {
	return x.ServerStream.SendMsg(m)
}

func _Manager_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.manager.Manager/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Info(ctx, req.(*InfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.manager.Manager/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Start(ctx, req.(*StartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.manager.Manager/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Stop(ctx, req.(*StopReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Prune_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PruneReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManagerServer).Prune(m, &managerPruneServer{stream})
}

type Manager_PruneServer interface {
	Send(*Service) error
	grpc.ServerStream
}

type managerPruneServer struct {
	grpc.ServerStream
}

func (x *managerPruneServer) Send(m *Service) error {
	return x.ServerStream.SendMsg(m)
}

var _Manager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stratumn.alice.grpc.manager.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Manager_Info_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Manager_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Manager_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Manager_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Prune",
			Handler:       _Manager_Prune_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/stratumn/alice/grpc/manager/manager.proto",
}

func (m *ListReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *InfoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfoReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintManager(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	return i, nil
}

func (m *StartReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StartReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintManager(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	return i, nil
}

func (m *StopReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StopReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintManager(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Prune {
		dAtA[i] = 0x10
		i++
		if m.Prune {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *PruneReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PruneReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *Service) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Service) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintManager(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Status != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintManager(dAtA, i, uint64(m.Status))
	}
	if len(m.Needs) > 0 {
		for _, s := range m.Needs {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.Stoppable {
		dAtA[i] = 0x20
		i++
		if m.Stoppable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Prunable {
		dAtA[i] = 0x28
		i++
		if m.Prunable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintManager(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Desc) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintManager(dAtA, i, uint64(len(m.Desc)))
		i += copy(dAtA[i:], m.Desc)
	}
	return i, nil
}

func encodeVarintManager(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ListReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *InfoReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovManager(uint64(l))
	}
	return n
}

func (m *StartReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovManager(uint64(l))
	}
	return n
}

func (m *StopReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovManager(uint64(l))
	}
	if m.Prune {
		n += 2
	}
	return n
}

func (m *PruneReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *Service) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovManager(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovManager(uint64(m.Status))
	}
	if len(m.Needs) > 0 {
		for _, s := range m.Needs {
			l = len(s)
			n += 1 + l + sovManager(uint64(l))
		}
	}
	if m.Stoppable {
		n += 2
	}
	if m.Prunable {
		n += 2
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovManager(uint64(l))
	}
	l = len(m.Desc)
	if l > 0 {
		n += 1 + l + sovManager(uint64(l))
	}
	return n
}

func sovManager(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozManager(x uint64) (n int) {
	return sovManager(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManager
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
			return fmt.Errorf("proto: ListReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManager
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
func (m *InfoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManager
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
			return fmt.Errorf("proto: InfoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManager
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
func (m *StartReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManager
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
			return fmt.Errorf("proto: StartReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StartReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManager
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
func (m *StopReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManager
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
			return fmt.Errorf("proto: StopReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StopReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prune", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
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
			m.Prune = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManager
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
func (m *PruneReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManager
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
			return fmt.Errorf("proto: PruneReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PruneReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManager
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
func (m *Service) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManager
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
			return fmt.Errorf("proto: Service: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Service: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (Service_Status(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Needs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Needs = append(m.Needs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stoppable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
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
			m.Stoppable = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prunable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
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
			m.Prunable = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Desc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManager
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
func skipManager(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowManager
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
					return 0, ErrIntOverflowManager
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
					return 0, ErrIntOverflowManager
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
				return 0, ErrInvalidLengthManager
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowManager
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
				next, err := skipManager(dAtA[start:])
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
	ErrInvalidLengthManager = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowManager   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/stratumn/alice/grpc/manager/manager.proto", fileDescriptorManager)
}

var fileDescriptorManager = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x41, 0x8f, 0xd2, 0x4e,
	0x14, 0xdf, 0xb2, 0x85, 0xc2, 0xfb, 0xff, 0xb3, 0x92, 0x51, 0x93, 0x11, 0x37, 0x38, 0xa9, 0x6b,
	0x82, 0xd1, 0x2d, 0x66, 0xf5, 0xe6, 0x45, 0x11, 0xb2, 0x21, 0x2a, 0x90, 0x29, 0x5e, 0xbc, 0x0d,
	0x30, 0xb0, 0x4d, 0x60, 0x5a, 0xda, 0xe9, 0x46, 0xaf, 0x7c, 0x0a, 0xcf, 0xeb, 0xc5, 0x2f, 0xc0,
	0x89, 0x8b, 0x47, 0x8f, 0x7e, 0x04, 0x83, 0x5f, 0xc4, 0xcc, 0xb4, 0xb0, 0x9c, 0xa0, 0x87, 0xa6,
	0x7d, 0xef, 0xfd, 0x7e, 0xef, 0xf7, 0xde, 0xcc, 0x7b, 0x85, 0x57, 0x13, 0x4f, 0x5e, 0xc5, 0x03,
	0x67, 0xe8, 0xcf, 0xea, 0x91, 0x0c, 0x99, 0x8c, 0x67, 0xa2, 0xce, 0xa6, 0xde, 0x90, 0xd7, 0x27,
	0x61, 0x30, 0xac, 0xcf, 0x98, 0x60, 0x13, 0x1e, 0x6e, 0xde, 0x4e, 0x10, 0xfa, 0xd2, 0x47, 0x0f,
	0x37, 0x50, 0x47, 0x43, 0x1d, 0x05, 0x75, 0x52, 0x48, 0xe5, 0xf9, 0x81, 0x94, 0xfc, 0x8b, 0x54,
	0x4f, 0x92, 0xca, 0x2e, 0x81, 0xf5, 0xc1, 0x8b, 0x24, 0xe5, 0x73, 0xfb, 0x1c, 0xac, 0xb6, 0x18,
	0xfb, 0x94, 0xcf, 0x91, 0x0d, 0x39, 0x6f, 0x84, 0x0d, 0x62, 0xd4, 0x4a, 0x0d, 0x74, 0xb3, 0xc2,
	0xe0, 0xf2, 0xf0, 0xda, 0x1b, 0x72, 0xd2, 0x6e, 0xfe, 0x58, 0x61, 0x83, 0xe6, 0xbc, 0x91, 0xed,
	0x40, 0xd1, 0x95, 0x2c, 0x94, 0x59, 0xf1, 0x21, 0x58, 0xae, 0xf4, 0x83, 0x8c, 0x70, 0xf4, 0x0e,
	0xf2, 0x41, 0x18, 0x0b, 0x8e, 0x73, 0xc4, 0xa8, 0x15, 0x1b, 0xe7, 0x37, 0x2b, 0xfc, 0xb4, 0xa7,
	0x1c, 0x24, 0x4a, 0xc0, 0x11, 0x61, 0x63, 0xc9, 0x43, 0x12, 0x49, 0x3f, 0x08, 0x3c, 0x31, 0x21,
	0xf2, 0x6a, 0x1b, 0xa3, 0x09, 0xd7, 0x06, 0x28, 0x6a, 0x8e, 0x6a, 0xef, 0x7b, 0x0e, 0xac, 0x54,
	0x07, 0x9d, 0xdc, 0x16, 0x90, 0x8a, 0x15, 0x22, 0xc9, 0x64, 0x1c, 0x69, 0xb5, 0x93, 0x8b, 0x67,
	0xce, 0x9e, 0x13, 0x76, 0xd2, 0x2c, 0x8e, 0xab, 0x29, 0x34, 0xa5, 0xa2, 0x7b, 0x90, 0x17, 0x9c,
	0x8f, 0x22, 0x7c, 0x4c, 0x8e, 0x6b, 0x25, 0x9a, 0x18, 0xe8, 0x14, 0x4a, 0xba, 0x42, 0x36, 0x98,
	0x72, 0x6c, 0xaa, 0x5e, 0xe8, 0xad, 0x03, 0x55, 0xa0, 0xa8, 0x2a, 0xd5, 0xc1, 0xbc, 0x0e, 0x6e,
	0x6d, 0x84, 0xc0, 0x14, 0x6c, 0xc6, 0x71, 0x41, 0x97, 0xa9, 0xbf, 0x95, 0x6f, 0xc4, 0xa3, 0x21,
	0xb6, 0x12, 0x9f, 0xfa, 0xb6, 0xdf, 0x43, 0x21, 0xa9, 0x04, 0xfd, 0x07, 0x96, 0xdb, 0xef, 0xf6,
	0x7a, 0xad, 0x66, 0xf9, 0x08, 0xfd, 0x0f, 0x45, 0xb7, 0xff, 0x96, 0xf6, 0xdb, 0x9d, 0xcb, 0xb2,
	0xa1, 0x42, 0xf4, 0x53, 0xa7, 0xa3, 0x8c, 0x5c, 0x12, 0xea, 0xf6, 0x7a, 0xca, 0x3a, 0x56, 0xa1,
	0x16, 0xa5, 0x5d, 0xda, 0x6a, 0x96, 0xcd, 0x8b, 0x9f, 0x26, 0x58, 0x1f, 0x93, 0x3e, 0xd1, 0x57,
	0x30, 0xd5, 0x6c, 0xa0, 0xb3, 0xbd, 0xa7, 0x91, 0x8e, 0x4f, 0xe5, 0x2c, 0xcb, 0x99, 0xd9, 0x8f,
	0x17, 0x4b, 0xfc, 0x48, 0x51, 0x08, 0x9b, 0x4e, 0xf5, 0x95, 0xb1, 0x6b, 0xe6, 0x4d, 0x55, 0xc3,
	0xdb, 0x8b, 0x7d, 0x61, 0x20, 0x09, 0xa6, 0x9a, 0xc5, 0x03, 0xd2, 0xe9, 0xb8, 0x66, 0x94, 0x26,
	0x8b, 0x25, 0x3e, 0xbd, 0xe4, 0x92, 0x78, 0x62, 0xec, 0x87, 0x33, 0x26, 0x3d, 0x5f, 0x10, 0x5f,
	0x10, 0xb6, 0x11, 0x46, 0x1e, 0xe4, 0xf5, 0x48, 0xa3, 0x27, 0xfb, 0x13, 0xa6, 0x63, 0x9f, 0x51,
	0xf7, 0xee, 0x62, 0x89, 0xef, 0x68, 0xce, 0x8e, 0xd4, 0x18, 0x4c, 0xb5, 0x0d, 0x07, 0x1a, 0x4c,
	0x17, 0x26, 0xa3, 0x10, 0x5a, 0x2c, 0xf1, 0x89, 0xa2, 0xec, 0xe8, 0xcc, 0x21, 0xaf, 0x37, 0xe0,
	0x40, 0x4b, 0x9b, 0x2d, 0xc9, 0xa8, 0xf4, 0x60, 0xb1, 0xc4, 0xf7, 0x93, 0x6d, 0x8c, 0x45, 0x1c,
	0xf1, 0xd1, 0xce, 0xdd, 0x35, 0xde, 0xfc, 0x5a, 0x57, 0x8d, 0xdf, 0xeb, 0xaa, 0xf1, 0x67, 0x5d,
	0x35, 0xbe, 0xfd, 0xad, 0x1e, 0x7d, 0x76, 0xb2, 0xfd, 0xe5, 0x5e, 0xa7, 0xef, 0x41, 0x41, 0xff,
	0x9b, 0x5e, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xa8, 0x70, 0x32, 0x1e, 0x05, 0x00, 0x00,
}
