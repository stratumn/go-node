// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stratumn/alice/app/clock/grpc/clock.proto

/*
	Package grpc is a generated protocol buffer package.

	It is generated from these files:
		github.com/stratumn/alice/app/clock/grpc/clock.proto

	It has these top-level messages:
		LocalReq
		RemoteReq
		Time
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/stratumn/alice/cli/grpc/ext"

import context "context"
import grpc1 "google.golang.org/grpc"

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

// The Local request message.
type LocalReq struct {
}

func (m *LocalReq) Reset()                    { *m = LocalReq{} }
func (m *LocalReq) String() string            { return proto.CompactTextString(m) }
func (*LocalReq) ProtoMessage()               {}
func (*LocalReq) Descriptor() ([]byte, []int) { return fileDescriptorClock, []int{0} }

// The Remote request message.
type RemoteReq struct {
	PeerId []byte `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
}

func (m *RemoteReq) Reset()                    { *m = RemoteReq{} }
func (m *RemoteReq) String() string            { return proto.CompactTextString(m) }
func (*RemoteReq) ProtoMessage()               {}
func (*RemoteReq) Descriptor() ([]byte, []int) { return fileDescriptorClock, []int{1} }

func (m *RemoteReq) GetPeerId() []byte {
	if m != nil {
		return m.PeerId
	}
	return nil
}

// The time message containing a Unix nano timestamp.
type Time struct {
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *Time) Reset()                    { *m = Time{} }
func (m *Time) String() string            { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()               {}
func (*Time) Descriptor() ([]byte, []int) { return fileDescriptorClock, []int{2} }

func (m *Time) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*LocalReq)(nil), "stratumn.alice.grpc.clock.LocalReq")
	proto.RegisterType((*RemoteReq)(nil), "stratumn.alice.grpc.clock.RemoteReq")
	proto.RegisterType((*Time)(nil), "stratumn.alice.grpc.clock.Time")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for Clock service

type ClockClient interface {
	// Returns the local time.
	Local(ctx context.Context, in *LocalReq, opts ...grpc1.CallOption) (*Time, error)
	// Returns a peer's remote time.
	Remote(ctx context.Context, in *RemoteReq, opts ...grpc1.CallOption) (*Time, error)
}

type clockClient struct {
	cc *grpc1.ClientConn
}

func NewClockClient(cc *grpc1.ClientConn) ClockClient {
	return &clockClient{cc}
}

func (c *clockClient) Local(ctx context.Context, in *LocalReq, opts ...grpc1.CallOption) (*Time, error) {
	out := new(Time)
	err := grpc1.Invoke(ctx, "/stratumn.alice.grpc.clock.Clock/Local", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clockClient) Remote(ctx context.Context, in *RemoteReq, opts ...grpc1.CallOption) (*Time, error) {
	out := new(Time)
	err := grpc1.Invoke(ctx, "/stratumn.alice.grpc.clock.Clock/Remote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Clock service

type ClockServer interface {
	// Returns the local time.
	Local(context.Context, *LocalReq) (*Time, error)
	// Returns a peer's remote time.
	Remote(context.Context, *RemoteReq) (*Time, error)
}

func RegisterClockServer(s *grpc1.Server, srv ClockServer) {
	s.RegisterService(&_Clock_serviceDesc, srv)
}

func _Clock_Local_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockServer).Local(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.clock.Clock/Local",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockServer).Local(ctx, req.(*LocalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clock_Remote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockServer).Remote(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.grpc.clock.Clock/Remote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockServer).Remote(ctx, req.(*RemoteReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Clock_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "stratumn.alice.grpc.clock.Clock",
	HandlerType: (*ClockServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "Local",
			Handler:    _Clock_Local_Handler,
		},
		{
			MethodName: "Remote",
			Handler:    _Clock_Remote_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "github.com/stratumn/alice/app/clock/grpc/clock.proto",
}

func (m *LocalReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *RemoteReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoteReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PeerId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClock(dAtA, i, uint64(len(m.PeerId)))
		i += copy(dAtA[i:], m.PeerId)
	}
	return i, nil
}

func (m *Time) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Time) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintClock(dAtA, i, uint64(m.Timestamp))
	}
	return i, nil
}

func encodeVarintClock(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LocalReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *RemoteReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovClock(uint64(l))
	}
	return n
}

func (m *Time) Size() (n int) {
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovClock(uint64(m.Timestamp))
	}
	return n
}

func sovClock(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClock(x uint64) (n int) {
	return sovClock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LocalReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClock
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
			return fmt.Errorf("proto: LocalReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipClock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClock
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
func (m *RemoteReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClock
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
			return fmt.Errorf("proto: RemoteReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoteReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClock
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
				return ErrInvalidLengthClock
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
			skippy, err := skipClock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClock
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
func (m *Time) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClock
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
			return fmt.Errorf("proto: Time: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Time: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClock
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
func skipClock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClock
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
					return 0, ErrIntOverflowClock
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
					return 0, ErrIntOverflowClock
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
				return 0, ErrInvalidLengthClock
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClock
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
				next, err := skipClock(dAtA[start:])
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
	ErrInvalidLengthClock = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClock   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/stratumn/alice/app/clock/grpc/clock.proto", fileDescriptorClock)
}

var fileDescriptorClock = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x2e, 0x29, 0x4a, 0x2c, 0x29, 0xcd, 0xcd, 0xd3,
	0x4f, 0xcc, 0xc9, 0x4c, 0x4e, 0xd5, 0x4f, 0x2c, 0x28, 0xd0, 0x4f, 0xce, 0xc9, 0x4f, 0xce, 0xd6,
	0x4f, 0x2f, 0x2a, 0x48, 0x86, 0x30, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x24, 0x61, 0x4a,
	0xf5, 0xc0, 0x4a, 0xf5, 0x40, 0x0a, 0xf4, 0xc0, 0x0a, 0xa4, 0x74, 0x70, 0x1b, 0x08, 0x36, 0x26,
	0xb5, 0xa2, 0x04, 0x84, 0x21, 0x06, 0x29, 0x71, 0x71, 0x71, 0xf8, 0xe4, 0x27, 0x27, 0xe6, 0x04,
	0xa5, 0x16, 0x2a, 0x59, 0x72, 0x71, 0x06, 0xa5, 0xe6, 0xe6, 0x97, 0xa4, 0x06, 0xa5, 0x16, 0x0a,
	0xe9, 0x70, 0xb1, 0x17, 0xa4, 0xa6, 0x16, 0xc5, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0,
	0x38, 0x09, 0x2f, 0xda, 0x2d, 0xc1, 0x1e, 0x90, 0x9a, 0x5a, 0xa4, 0xe0, 0xe9, 0xb2, 0x62, 0xb7,
	0x04, 0xe3, 0x87, 0xdd, 0x12, 0x8c, 0x41, 0x6c, 0x20, 0x35, 0x9e, 0x29, 0x4a, 0x5a, 0x5c, 0x2c,
	0x21, 0x99, 0xb9, 0xa9, 0x42, 0x4a, 0x5c, 0x9c, 0x25, 0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9,
	0x05, 0x60, 0x7d, 0xcc, 0x4e, 0x2c, 0x07, 0xf6, 0x48, 0x30, 0x06, 0x21, 0x84, 0x8d, 0xde, 0x31,
	0x72, 0xb1, 0x3a, 0x83, 0x9c, 0x2a, 0x54, 0xc0, 0xc5, 0x0a, 0xb6, 0x5c, 0x48, 0x59, 0x0f, 0xa7,
	0x7f, 0xf4, 0x60, 0xce, 0x93, 0x92, 0xc7, 0xa3, 0x08, 0x64, 0xb9, 0x92, 0x62, 0xd3, 0x56, 0x09,
	0x59, 0x97, 0xcc, 0xe2, 0x82, 0x9c, 0xc4, 0x4a, 0x85, 0x92, 0x8c, 0x54, 0x85, 0xbc, 0xfc, 0x94,
	0x54, 0xf5, 0x62, 0x85, 0x1c, 0x90, 0x09, 0x0a, 0x20, 0x07, 0x08, 0x15, 0x73, 0xb1, 0x41, 0xbc,
	0x28, 0xa4, 0x82, 0xc7, 0x34, 0x78, 0x28, 0x10, 0xb6, 0x53, 0xa9, 0x69, 0xab, 0x84, 0x1c, 0xb2,
	0x9d, 0xa0, 0x00, 0x51, 0x2f, 0x56, 0x28, 0x02, 0x1b, 0x01, 0xb6, 0xd4, 0xc9, 0xfe, 0xc4, 0x23,
	0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x4a,
	0x97, 0xd8, 0x48, 0xb7, 0x06, 0x11, 0x49, 0x6c, 0xe0, 0xb8, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x2d, 0x7f, 0x2f, 0x0e, 0x2c, 0x02, 0x00, 0x00,
}
