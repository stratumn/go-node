// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stratumn/alice/app/storage/grpc/storage.proto

/*
	Package grpc is a generated protocol buffer package.

	It is generated from these files:
		github.com/stratumn/alice/app/storage/grpc/storage.proto

	It has these top-level messages:
		SessionFileChunk
		UploadReq
		UploadSession
		AuthRequest
		DownloadRequest
		Ack
		UploadAck
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/stratumn/alice/cli/grpc/ext"
import stratumn_alice_app_storage_pb "github.com/stratumn/alice/app/storage/pb"

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

type SessionFileChunk struct {
	// file name is required only on first message.
	Id   []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SessionFileChunk) Reset()                    { *m = SessionFileChunk{} }
func (m *SessionFileChunk) String() string            { return proto.CompactTextString(m) }
func (*SessionFileChunk) ProtoMessage()               {}
func (*SessionFileChunk) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{0} }

func (m *SessionFileChunk) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SessionFileChunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type UploadReq struct {
	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (m *UploadReq) Reset()                    { *m = UploadReq{} }
func (m *UploadReq) String() string            { return proto.CompactTextString(m) }
func (*UploadReq) ProtoMessage()               {}
func (*UploadReq) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{1} }

func (m *UploadReq) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

type UploadSession struct {
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *UploadSession) Reset()                    { *m = UploadSession{} }
func (m *UploadSession) String() string            { return proto.CompactTextString(m) }
func (*UploadSession) ProtoMessage()               {}
func (*UploadSession) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{2} }

func (m *UploadSession) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type AuthRequest struct {
	PeerIds  [][]byte `protobuf:"bytes,1,rep,name=peer_ids,json=peerIds" json:"peer_ids,omitempty"`
	FileHash []byte   `protobuf:"bytes,2,opt,name=file_hash,json=fileHash,proto3" json:"file_hash,omitempty"`
}

func (m *AuthRequest) Reset()                    { *m = AuthRequest{} }
func (m *AuthRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()               {}
func (*AuthRequest) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{3} }

func (m *AuthRequest) GetPeerIds() [][]byte {
	if m != nil {
		return m.PeerIds
	}
	return nil
}

func (m *AuthRequest) GetFileHash() []byte {
	if m != nil {
		return m.FileHash
	}
	return nil
}

type DownloadRequest struct {
	PeerId   []byte `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	FileHash []byte `protobuf:"bytes,2,opt,name=file_hash,json=fileHash,proto3" json:"file_hash,omitempty"`
}

func (m *DownloadRequest) Reset()                    { *m = DownloadRequest{} }
func (m *DownloadRequest) String() string            { return proto.CompactTextString(m) }
func (*DownloadRequest) ProtoMessage()               {}
func (*DownloadRequest) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{4} }

func (m *DownloadRequest) GetPeerId() []byte {
	if m != nil {
		return m.PeerId
	}
	return nil
}

func (m *DownloadRequest) GetFileHash() []byte {
	if m != nil {
		return m.FileHash
	}
	return nil
}

// An empty ack.
type Ack struct {
}

func (m *Ack) Reset()                    { *m = Ack{} }
func (m *Ack) String() string            { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()               {}
func (*Ack) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{5} }

type UploadAck struct {
	FileHash []byte `protobuf:"bytes,1,opt,name=file_hash,json=fileHash,proto3" json:"file_hash,omitempty"`
}

func (m *UploadAck) Reset()                    { *m = UploadAck{} }
func (m *UploadAck) String() string            { return proto.CompactTextString(m) }
func (*UploadAck) ProtoMessage()               {}
func (*UploadAck) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{6} }

func (m *UploadAck) GetFileHash() []byte {
	if m != nil {
		return m.FileHash
	}
	return nil
}

func init() {
	proto.RegisterType((*SessionFileChunk)(nil), "stratumn.alice.app.storage.grpc.SessionFileChunk")
	proto.RegisterType((*UploadReq)(nil), "stratumn.alice.app.storage.grpc.UploadReq")
	proto.RegisterType((*UploadSession)(nil), "stratumn.alice.app.storage.grpc.UploadSession")
	proto.RegisterType((*AuthRequest)(nil), "stratumn.alice.app.storage.grpc.AuthRequest")
	proto.RegisterType((*DownloadRequest)(nil), "stratumn.alice.app.storage.grpc.DownloadRequest")
	proto.RegisterType((*Ack)(nil), "stratumn.alice.app.storage.grpc.Ack")
	proto.RegisterType((*UploadAck)(nil), "stratumn.alice.app.storage.grpc.UploadAck")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for Storage service

type StorageClient interface {
	// Upload a file to the alice node.
	Upload(ctx context.Context, opts ...grpc1.CallOption) (Storage_UploadClient, error)
	// Sends a request to the server to start a new upload.
	StartUpload(ctx context.Context, in *UploadReq, opts ...grpc1.CallOption) (*UploadSession, error)
	// Upload a single chunk of a file to the server.
	UploadChunk(ctx context.Context, in *SessionFileChunk, opts ...grpc1.CallOption) (*Ack, error)
	// Notifies the server that the session's file has been entirely sent.
	EndUpload(ctx context.Context, in *UploadSession, opts ...grpc1.CallOption) (*UploadAck, error)
	// Give peers access to a file.
	AuthorizePeers(ctx context.Context, in *AuthRequest, opts ...grpc1.CallOption) (*Ack, error)
	// Download downloads a file from a peer.
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc1.CallOption) (Storage_DownloadClient, error)
}

type storageClient struct {
	cc *grpc1.ClientConn
}

func NewStorageClient(cc *grpc1.ClientConn) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Upload(ctx context.Context, opts ...grpc1.CallOption) (Storage_UploadClient, error) {
	stream, err := grpc1.NewClientStream(ctx, &_Storage_serviceDesc.Streams[0], c.cc, "/stratumn.alice.app.storage.grpc.Storage/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageUploadClient{stream}
	return x, nil
}

type Storage_UploadClient interface {
	Send(*stratumn_alice_app_storage_pb.FileChunk) error
	CloseAndRecv() (*UploadAck, error)
	grpc1.ClientStream
}

type storageUploadClient struct {
	grpc1.ClientStream
}

func (x *storageUploadClient) Send(m *stratumn_alice_app_storage_pb.FileChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageUploadClient) CloseAndRecv() (*UploadAck, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageClient) StartUpload(ctx context.Context, in *UploadReq, opts ...grpc1.CallOption) (*UploadSession, error) {
	out := new(UploadSession)
	err := grpc1.Invoke(ctx, "/stratumn.alice.app.storage.grpc.Storage/StartUpload", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) UploadChunk(ctx context.Context, in *SessionFileChunk, opts ...grpc1.CallOption) (*Ack, error) {
	out := new(Ack)
	err := grpc1.Invoke(ctx, "/stratumn.alice.app.storage.grpc.Storage/UploadChunk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) EndUpload(ctx context.Context, in *UploadSession, opts ...grpc1.CallOption) (*UploadAck, error) {
	out := new(UploadAck)
	err := grpc1.Invoke(ctx, "/stratumn.alice.app.storage.grpc.Storage/EndUpload", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) AuthorizePeers(ctx context.Context, in *AuthRequest, opts ...grpc1.CallOption) (*Ack, error) {
	out := new(Ack)
	err := grpc1.Invoke(ctx, "/stratumn.alice.app.storage.grpc.Storage/AuthorizePeers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc1.CallOption) (Storage_DownloadClient, error) {
	stream, err := grpc1.NewClientStream(ctx, &_Storage_serviceDesc.Streams[1], c.cc, "/stratumn.alice.app.storage.grpc.Storage/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Storage_DownloadClient interface {
	Recv() (*stratumn_alice_app_storage_pb.FileChunk, error)
	grpc1.ClientStream
}

type storageDownloadClient struct {
	grpc1.ClientStream
}

func (x *storageDownloadClient) Recv() (*stratumn_alice_app_storage_pb.FileChunk, error) {
	m := new(stratumn_alice_app_storage_pb.FileChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Storage service

type StorageServer interface {
	// Upload a file to the alice node.
	Upload(Storage_UploadServer) error
	// Sends a request to the server to start a new upload.
	StartUpload(context.Context, *UploadReq) (*UploadSession, error)
	// Upload a single chunk of a file to the server.
	UploadChunk(context.Context, *SessionFileChunk) (*Ack, error)
	// Notifies the server that the session's file has been entirely sent.
	EndUpload(context.Context, *UploadSession) (*UploadAck, error)
	// Give peers access to a file.
	AuthorizePeers(context.Context, *AuthRequest) (*Ack, error)
	// Download downloads a file from a peer.
	Download(*DownloadRequest, Storage_DownloadServer) error
}

func RegisterStorageServer(s *grpc1.Server, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_Upload_Handler(srv interface{}, stream grpc1.ServerStream) error {
	return srv.(StorageServer).Upload(&storageUploadServer{stream})
}

type Storage_UploadServer interface {
	SendAndClose(*UploadAck) error
	Recv() (*stratumn_alice_app_storage_pb.FileChunk, error)
	grpc1.ServerStream
}

type storageUploadServer struct {
	grpc1.ServerStream
}

func (x *storageUploadServer) SendAndClose(m *UploadAck) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageUploadServer) Recv() (*stratumn_alice_app_storage_pb.FileChunk, error) {
	m := new(stratumn_alice_app_storage_pb.FileChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Storage_StartUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).StartUpload(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.app.storage.grpc.Storage/StartUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).StartUpload(ctx, req.(*UploadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_UploadChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionFileChunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).UploadChunk(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.app.storage.grpc.Storage/UploadChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).UploadChunk(ctx, req.(*SessionFileChunk))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_EndUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadSession)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).EndUpload(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.app.storage.grpc.Storage/EndUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).EndUpload(ctx, req.(*UploadSession))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_AuthorizePeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).AuthorizePeers(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.alice.app.storage.grpc.Storage/AuthorizePeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).AuthorizePeers(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Download_Handler(srv interface{}, stream grpc1.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServer).Download(m, &storageDownloadServer{stream})
}

type Storage_DownloadServer interface {
	Send(*stratumn_alice_app_storage_pb.FileChunk) error
	grpc1.ServerStream
}

type storageDownloadServer struct {
	grpc1.ServerStream
}

func (x *storageDownloadServer) Send(m *stratumn_alice_app_storage_pb.FileChunk) error {
	return x.ServerStream.SendMsg(m)
}

var _Storage_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "stratumn.alice.app.storage.grpc.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "StartUpload",
			Handler:    _Storage_StartUpload_Handler,
		},
		{
			MethodName: "UploadChunk",
			Handler:    _Storage_UploadChunk_Handler,
		},
		{
			MethodName: "EndUpload",
			Handler:    _Storage_EndUpload_Handler,
		},
		{
			MethodName: "AuthorizePeers",
			Handler:    _Storage_AuthorizePeers_Handler,
		},
	},
	Streams: []grpc1.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Storage_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Download",
			Handler:       _Storage_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/stratumn/alice/app/storage/grpc/storage.proto",
}

func (m *SessionFileChunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionFileChunk) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func (m *UploadReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FileName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.FileName)))
		i += copy(dAtA[i:], m.FileName)
	}
	return i, nil
}

func (m *UploadSession) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadSession) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	return i, nil
}

func (m *AuthRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PeerIds) > 0 {
		for _, b := range m.PeerIds {
			dAtA[i] = 0xa
			i++
			i = encodeVarintStorage(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if len(m.FileHash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.FileHash)))
		i += copy(dAtA[i:], m.FileHash)
	}
	return i, nil
}

func (m *DownloadRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DownloadRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PeerId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.PeerId)))
		i += copy(dAtA[i:], m.PeerId)
	}
	if len(m.FileHash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.FileHash)))
		i += copy(dAtA[i:], m.FileHash)
	}
	return i, nil
}

func (m *Ack) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ack) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *UploadAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FileHash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.FileHash)))
		i += copy(dAtA[i:], m.FileHash)
	}
	return i, nil
}

func encodeVarintStorage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SessionFileChunk) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func (m *UploadReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.FileName)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func (m *UploadSession) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func (m *AuthRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.PeerIds) > 0 {
		for _, b := range m.PeerIds {
			l = len(b)
			n += 1 + l + sovStorage(uint64(l))
		}
	}
	l = len(m.FileHash)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func (m *DownloadRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.FileHash)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func (m *Ack) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *UploadAck) Size() (n int) {
	var l int
	_ = l
	l = len(m.FileHash)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func sovStorage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStorage(x uint64) (n int) {
	return sovStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SessionFileChunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: SessionFileChunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionFileChunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
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
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
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
func (m *UploadReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: UploadReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
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
func (m *UploadSession) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: UploadSession: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadSession: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
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
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
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
func (m *AuthRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: AuthRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerIds", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerIds = append(m.PeerIds, make([]byte, postIndex-iNdEx))
			copy(m.PeerIds[len(m.PeerIds)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileHash = append(m.FileHash[:0], dAtA[iNdEx:postIndex]...)
			if m.FileHash == nil {
				m.FileHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
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
func (m *DownloadRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: DownloadRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DownloadRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
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
				return fmt.Errorf("proto: wrong wireType = %d for field FileHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileHash = append(m.FileHash[:0], dAtA[iNdEx:postIndex]...)
			if m.FileHash == nil {
				m.FileHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
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
func (m *Ack) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: Ack: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ack: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
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
func (m *UploadAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: UploadAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileHash = append(m.FileHash[:0], dAtA[iNdEx:postIndex]...)
			if m.FileHash == nil {
				m.FileHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
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
func skipStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStorage
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
					return 0, ErrIntOverflowStorage
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
					return 0, ErrIntOverflowStorage
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
				return 0, ErrInvalidLengthStorage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStorage
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
				next, err := skipStorage(dAtA[start:])
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
	ErrInvalidLengthStorage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStorage   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/stratumn/alice/app/storage/grpc/storage.proto", fileDescriptorStorage)
}

var fileDescriptorStorage = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x8e, 0xd2, 0x40,
	0x1c, 0xa7, 0xac, 0xcb, 0xc7, 0xb0, 0xea, 0x66, 0x34, 0x71, 0xd2, 0x03, 0x3b, 0x69, 0x3c, 0x90,
	0x0d, 0x69, 0xd1, 0x8d, 0xc6, 0xc4, 0x13, 0x75, 0xfd, 0xe0, 0x62, 0x0c, 0xe8, 0xc5, 0xcb, 0x66,
	0x68, 0x07, 0x5a, 0x81, 0xb6, 0x74, 0xa6, 0xae, 0xeb, 0xd5, 0x97, 0xf0, 0xdc, 0x93, 0xef, 0xd0,
	0x17, 0xf0, 0xe8, 0x23, 0x18, 0x7c, 0x01, 0x1f, 0xc1, 0xcc, 0x74, 0x8a, 0x2e, 0x51, 0x29, 0x7b,
	0x80, 0x40, 0xfb, 0xfb, 0xfa, 0x4f, 0x7e, 0xff, 0x01, 0x8f, 0xa6, 0x3e, 0xf7, 0x92, 0xb1, 0xe9,
	0x84, 0x0b, 0x8b, 0xf1, 0x98, 0xf0, 0x64, 0x11, 0x58, 0x64, 0xee, 0x3b, 0xd4, 0x22, 0x51, 0x64,
	0x31, 0x1e, 0xc6, 0x64, 0x4a, 0xad, 0x69, 0x1c, 0x39, 0xc5, 0x1f, 0x33, 0x8a, 0x43, 0x1e, 0xc2,
	0xa3, 0x02, 0x6e, 0x4a, 0xb8, 0x49, 0xa2, 0xc8, 0x2c, 0x10, 0x02, 0xae, 0x77, 0xff, 0x2d, 0x2d,
	0xe5, 0xe8, 0x07, 0x2e, 0x3e, 0xb9, 0x9c, 0xfe, 0xb0, 0x5c, 0x90, 0x68, 0x7c, 0x39, 0x86, 0x71,
	0x01, 0x0e, 0x47, 0x94, 0x31, 0x3f, 0x0c, 0x9e, 0xf9, 0x73, 0xfa, 0xc4, 0x4b, 0x82, 0x19, 0xec,
	0x81, 0xaa, 0xef, 0x22, 0x0d, 0x6b, 0x9d, 0x03, 0x1b, 0xa7, 0x19, 0x42, 0x03, 0x17, 0x87, 0x13,
	0xcc, 0x3d, 0x8a, 0x93, 0x68, 0x1e, 0x12, 0x17, 0xb3, 0x9c, 0xf2, 0x25, 0x43, 0xda, 0xb0, 0xea,
	0xbb, 0xf0, 0x04, 0x5c, 0x73, 0x09, 0x27, 0xa8, 0x2a, 0x39, 0x47, 0x69, 0x86, 0xee, 0xd8, 0x17,
	0x9c, 0x32, 0x41, 0x0b, 0x03, 0x8a, 0x27, 0xfe, 0x9c, 0x62, 0x47, 0x88, 0x4b, 0x8a, 0x04, 0x1b,
	0x36, 0x68, 0xbe, 0x91, 0x72, 0x43, 0xba, 0x84, 0x0f, 0x40, 0x53, 0x80, 0xce, 0x02, 0xb2, 0xa0,
	0xd2, 0xba, 0x69, 0xa3, 0x34, 0x43, 0x87, 0x2f, 0xc9, 0x82, 0x16, 0xe6, 0x02, 0x20, 0xf9, 0x0d,
	0xf1, 0x4b, 0xbc, 0x31, 0xfa, 0xe0, 0x7a, 0xae, 0xa1, 0x86, 0xd8, 0x3d, 0xbb, 0x91, 0x80, 0x56,
	0x3f, 0xe1, 0xde, 0x90, 0x2e, 0x13, 0xca, 0x38, 0xb4, 0x40, 0x23, 0xa2, 0x34, 0x3e, 0xf3, 0x5d,
	0x86, 0x34, 0xbc, 0xd7, 0x39, 0xb0, 0x6f, 0xa7, 0x19, 0x6a, 0xbc, 0xa2, 0x34, 0xc6, 0x83, 0x53,
	0x26, 0x68, 0x3f, 0x05, 0xb5, 0x2e, 0x50, 0x03, 0x97, 0xad, 0x93, 0x7b, 0x84, 0x79, 0xea, 0x00,
	0x64, 0xf2, 0x17, 0x84, 0x79, 0x7f, 0x4f, 0x2e, 0xde, 0x18, 0xef, 0xc1, 0xcd, 0xd3, 0xf0, 0x3c,
	0x50, 0xf3, 0x4b, 0xeb, 0x2e, 0xa8, 0x2b, 0x6b, 0x35, 0xc0, 0xad, 0x34, 0x43, 0x75, 0xe5, 0xbc,
	0x36, 0xae, 0xe5, 0xc6, 0x57, 0xf5, 0xdd, 0x07, 0x7b, 0x7d, 0x67, 0x66, 0xbc, 0x2e, 0x0e, 0xbf,
	0xef, 0xcc, 0xe0, 0xf3, 0x3f, 0xa5, 0x72, 0xeb, 0xe3, 0x34, 0x43, 0xc6, 0xa6, 0x14, 0xe6, 0x1e,
	0xe1, 0xf8, 0x9c, 0x30, 0x75, 0x96, 0xd4, 0xbd, 0x2c, 0x7e, 0xff, 0xd3, 0x3e, 0xa8, 0x8f, 0xf2,
	0x7e, 0xc1, 0x31, 0xa8, 0xe5, 0x0e, 0xb0, 0x63, 0xfe, 0xa7, 0xeb, 0xd1, 0xd8, 0x5c, 0x37, 0x4f,
	0x3f, 0x36, 0xb7, 0x6c, 0x85, 0xb9, 0x0e, 0x6d, 0x54, 0x3a, 0x1a, 0x9c, 0x81, 0xd6, 0x88, 0x93,
	0x98, 0x2b, 0xa3, 0xb2, 0xf4, 0x21, 0x5d, 0xea, 0x66, 0x49, 0xac, 0x2a, 0x96, 0x51, 0x81, 0x1e,
	0x68, 0xe5, 0x8f, 0xf2, 0x2d, 0xb9, 0xb7, 0x55, 0x60, 0x73, 0xb1, 0xf4, 0xbb, 0x5b, 0x29, 0x72,
	0x30, 0xe8, 0x83, 0xe6, 0xd3, 0xc0, 0x55, 0x43, 0xed, 0x18, 0x74, 0xb7, 0x33, 0x84, 0x13, 0x70,
	0x43, 0xb4, 0x3f, 0x8c, 0xfd, 0x8f, 0x54, 0x34, 0x8d, 0xc1, 0xee, 0xf6, 0x90, 0xbf, 0xd7, 0xa5,
	0xf4, 0x48, 0xef, 0x40, 0xa3, 0xa8, 0x3b, 0xec, 0x6d, 0xe5, 0x6c, 0x6c, 0x86, 0x5e, 0xba, 0x41,
	0x46, 0xa5, 0xa7, 0xd9, 0xfd, 0xaf, 0xab, 0xb6, 0xf6, 0x6d, 0xd5, 0xd6, 0xbe, 0xaf, 0xda, 0xda,
	0xe7, 0x1f, 0xed, 0xca, 0x5b, 0xab, 0xfc, 0x35, 0xfd, 0x58, 0x7c, 0x8d, 0x6b, 0xf2, 0x76, 0x3c,
	0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0x35, 0x34, 0x4b, 0xc5, 0xe0, 0x05, 0x00, 0x00,
}
