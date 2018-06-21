// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stratumn/alice/core/app/grpcapi/grpc/grpcapi.proto

/*
	Package grpc is a generated protocol buffer package.

	It is generated from these files:
		github.com/stratumn/alice/core/app/grpcapi/grpc/grpcapi.proto

	It has these top-level messages:
		InformReq
		Info
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/stratumn/alice/grpc/ext"

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

// The inform request message.
type InformReq struct {
}

func (m *InformReq) Reset()                    { *m = InformReq{} }
func (m *InformReq) String() string            { return proto.CompactTextString(m) }
func (*InformReq) ProtoMessage()               {}
func (*InformReq) Descriptor() ([]byte, []int) { return fileDescriptorGrpcapi, []int{0} }

// The info message containing information about the API.
type Info struct {
	Protocol  string `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Version   string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	GitCommit []byte `protobuf:"bytes,3,opt,name=git_commit,json=gitCommit,proto3" json:"git_commit,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptorGrpcapi, []int{1} }

func (m *Info) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Info) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Info) GetGitCommit() []byte {
	if m != nil {
		return m.GitCommit
	}
	return nil
}

func init() {
	proto.RegisterType((*InformReq)(nil), "stratumn.alice.core.app.grpcapi.grpc.InformReq")
	proto.RegisterType((*Info)(nil), "stratumn.alice.core.app.grpcapi.grpc.Info")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	// Returns information about the API.
	Inform(ctx context.Context, in *InformReq, opts ...grpc1.CallOption) (*Info, error)
}

type aPIClient struct {
	cc *grpc1.ClientConn
}

func NewAPIClient(cc *grpc1.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Inform(ctx context.Context, in *InformReq, opts ...grpc1.CallOption) (*Info, error) {
	out := new(Info)
	err := grpc1.Invoke(ctx, "/stratumn.alice.core.app.grpcapi.grpc.API/Inform", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	// Returns information about the API.
	Inform(context.Context, *InformReq) (*Info, error)
}

func RegisterAPIServer(s *grpc1.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Inform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(InformReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Inform(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.core.app.grpcapi.grpc.API/Inform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Inform(ctx, req.(*InformReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "stratumn.alice.core.app.grpcapi.grpc.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "Inform",
			Handler:    _API_Inform_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "github.com/stratumn/alice/core/app/grpcapi/grpc/grpcapi.proto",
}

func (m *InformReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InformReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *Info) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Info) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Protocol) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGrpcapi(dAtA, i, uint64(len(m.Protocol)))
		i += copy(dAtA[i:], m.Protocol)
	}
	if len(m.Version) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGrpcapi(dAtA, i, uint64(len(m.Version)))
		i += copy(dAtA[i:], m.Version)
	}
	if len(m.GitCommit) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGrpcapi(dAtA, i, uint64(len(m.GitCommit)))
		i += copy(dAtA[i:], m.GitCommit)
	}
	return i, nil
}

func encodeVarintGrpcapi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *InformReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *Info) Size() (n int) {
	var l int
	_ = l
	l = len(m.Protocol)
	if l > 0 {
		n += 1 + l + sovGrpcapi(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovGrpcapi(uint64(l))
	}
	l = len(m.GitCommit)
	if l > 0 {
		n += 1 + l + sovGrpcapi(uint64(l))
	}
	return n
}

func sovGrpcapi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGrpcapi(x uint64) (n int) {
	return sovGrpcapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InformReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpcapi
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
			return fmt.Errorf("proto: InformReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InformReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGrpcapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpcapi
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
func (m *Info) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpcapi
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
			return fmt.Errorf("proto: Info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcapi
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
				return ErrInvalidLengthGrpcapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Protocol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcapi
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
				return ErrInvalidLengthGrpcapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GitCommit", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcapi
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
				return ErrInvalidLengthGrpcapi
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GitCommit = append(m.GitCommit[:0], dAtA[iNdEx:postIndex]...)
			if m.GitCommit == nil {
				m.GitCommit = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpcapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpcapi
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
func skipGrpcapi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGrpcapi
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
					return 0, ErrIntOverflowGrpcapi
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
					return 0, ErrIntOverflowGrpcapi
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
				return 0, ErrInvalidLengthGrpcapi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGrpcapi
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
				next, err := skipGrpcapi(dAtA[start:])
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
	ErrInvalidLengthGrpcapi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGrpcapi   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/stratumn/alice/core/app/grpcapi/grpc/grpcapi.proto", fileDescriptorGrpcapi)
}

var fileDescriptorGrpcapi = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x2e, 0x29, 0x4a, 0x2c, 0x29, 0xcd, 0xcd, 0xd3,
	0x4f, 0xcc, 0xc9, 0x4c, 0x4e, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x2c, 0x28, 0xd0, 0x4f,
	0x2f, 0x2a, 0x48, 0x4e, 0x2c, 0xc8, 0x04, 0xd3, 0x30, 0x8e, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe,
	0x90, 0x0a, 0x4c, 0x8f, 0x1e, 0x58, 0x8f, 0x1e, 0x48, 0x8f, 0x5e, 0x62, 0x41, 0x81, 0x1e, 0x4c,
	0x19, 0x88, 0x96, 0xd2, 0xc1, 0x6d, 0x09, 0xd8, 0xcc, 0xd4, 0x8a, 0x12, 0x10, 0x86, 0x98, 0xa9,
	0xc4, 0xcd, 0xc5, 0xe9, 0x99, 0x97, 0x96, 0x5f, 0x94, 0x1b, 0x94, 0x5a, 0xa8, 0x14, 0xcd, 0xc5,
	0x02, 0xe2, 0x08, 0x49, 0x71, 0x71, 0x80, 0x65, 0x93, 0xf3, 0x73, 0x24, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0xe0, 0x7c, 0x21, 0x09, 0x2e, 0xf6, 0xb2, 0xd4, 0xa2, 0xe2, 0xcc, 0xfc, 0x3c, 0x09,
	0x26, 0xb0, 0x14, 0x8c, 0x2b, 0x24, 0xcb, 0xc5, 0x95, 0x9e, 0x59, 0x12, 0x9f, 0x9c, 0x9f, 0x9b,
	0x9b, 0x59, 0x22, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x13, 0xc4, 0x99, 0x9e, 0x59, 0xe2, 0x0c, 0x16,
	0x30, 0x9a, 0xc0, 0xc8, 0xc5, 0xec, 0x18, 0xe0, 0x29, 0xd4, 0xc1, 0xc8, 0xc5, 0x06, 0xb1, 0x52,
	0x48, 0x5f, 0x8f, 0x18, 0x1f, 0xe9, 0xc1, 0x1d, 0x28, 0xa5, 0x45, 0xbc, 0x06, 0x25, 0xc5, 0xa6,
	0xad, 0x12, 0xb2, 0xee, 0xa9, 0x25, 0x0a, 0x99, 0x60, 0xed, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x0a,
	0x89, 0x49, 0xf9, 0xa5, 0x25, 0x0a, 0x25, 0x19, 0xa9, 0x0a, 0x8e, 0x01, 0x9e, 0x4e, 0x6e, 0x27,
	0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c,
	0x51, 0x26, 0x24, 0xc6, 0x90, 0x35, 0x88, 0x48, 0x62, 0x03, 0x87, 0x8e, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x72, 0xc9, 0x73, 0x93, 0xe0, 0x01, 0x00, 0x00,
}