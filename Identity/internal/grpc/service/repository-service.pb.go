package service

import (
	domain "github.com/yudafatah/golang-microservices-boilerplate/tree/main/Identity/internal/grpc/domain"
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddRepositoryResponse struct {
	AddedRepository      *domain.Repository `protobuf:"bytes,1,opt,name=addedRepository,proto3" json:"addedRepository,omitempty"`
	Error                *Error             `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AddRepositoryResponse) Reset()         { *m = AddRepositoryResponse{} }
func (m *AddRepositoryResponse) String() string { return proto.CompactTextString(m) }
func (*AddRepositoryResponse) ProtoMessage()    {}
func (*AddRepositoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06728efaff680198, []int{0}
}

func (m *AddRepositoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRepositoryResponse.Unmarshal(m, b)
}
func (m *AddRepositoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRepositoryResponse.Marshal(b, m, deterministic)
}
func (m *AddRepositoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRepositoryResponse.Merge(m, src)
}
func (m *AddRepositoryResponse) XXX_Size() int {
	return xxx_messageInfo_AddRepositoryResponse.Size(m)
}
func (m *AddRepositoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRepositoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddRepositoryResponse proto.InternalMessageInfo

func (m *AddRepositoryResponse) GetAddedRepository() *domain.Repository {
	if m != nil {
		return m.AddedRepository
	}
	return nil
}

func (m *AddRepositoryResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_06728efaff680198, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*AddRepositoryResponse)(nil), "service.AddRepositoryResponse")
	proto.RegisterType((*Error)(nil), "service.Error")
}

func init() {
	proto.RegisterFile("bitbucket-repository-management-service/internal/proto-files/service/repository-service.proto", fileDescriptor_06728efaff680198)
}

var fileDescriptor_06728efaff680198 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0xa5, 0x6a, 0x2d, 0x1d, 0x41, 0x71, 0x40, 0x08, 0x3d, 0x88, 0x14, 0x0f, 0x5e, 0x92, 0x40,
	0xc5, 0x83, 0xa0, 0x88, 0x82, 0x57, 0x0f, 0xab, 0x27, 0xc1, 0xc3, 0x26, 0x3b, 0x86, 0xc5, 0x66,
	0x27, 0xcc, 0xae, 0x82, 0xf8, 0xf3, 0xe2, 0xa6, 0xa9, 0x41, 0x72, 0x91, 0xde, 0xf2, 0xde, 0x9b,
	0xf7, 0x5e, 0x66, 0x16, 0x5e, 0x0a, 0x1b, 0x8a, 0xf7, 0xf2, 0x8d, 0x42, 0x2a, 0xd4, 0xb0, 0xb7,
	0x81, 0xe5, 0x33, 0xad, 0xb5, 0xd3, 0x15, 0xd5, 0xe4, 0x42, 0xea, 0x49, 0x3e, 0x6c, 0x49, 0xb9,
	0x75, 0x81, 0xc4, 0xe9, 0x65, 0xde, 0x08, 0x07, 0x4e, 0x5f, 0xed, 0x92, 0x7c, 0xde, 0x89, 0x3d,
	0xeb, 0x8a, 0xca, 0xe2, 0x18, 0x4e, 0x56, 0x70, 0xf6, 0xb4, 0x51, 0x8f, 0xe1, 0x5a, 0x5b, 0xd7,
	0xab, 0x69, 0xe3, 0xe7, 0x5f, 0x70, 0x74, 0x6b, 0x8c, 0x5a, 0xd3, 0x8a, 0x7c, 0xc3, 0xce, 0x13,
	0x5e, 0xc1, 0x81, 0x36, 0x86, 0x7a, 0x52, 0x32, 0x3a, 0x19, 0x9d, 0xed, 0x2d, 0x30, 0x6b, 0xb3,
	0xb2, 0x9e, 0xe9, 0xef, 0x28, 0x9e, 0xc2, 0x98, 0x44, 0x58, 0x92, 0xad, 0xe8, 0xd9, 0xcf, 0xba,
	0xa5, 0xee, 0x7f, 0x58, 0xd5, 0x8a, 0xf3, 0x0b, 0x18, 0x47, 0x8c, 0x08, 0x3b, 0x25, 0x1b, 0x8a,
	0x0d, 0x53, 0x15, 0xbf, 0x31, 0x81, 0x49, 0x4d, 0xde, 0xeb, 0x8a, 0x62, 0xc8, 0x54, 0x75, 0x70,
	0xf1, 0x00, 0x87, 0xbf, 0x55, 0x8f, 0x6d, 0x30, 0x5e, 0xc2, 0xb6, 0x36, 0x06, 0x07, 0xfe, 0x6e,
	0x76, 0xbc, 0x6e, 0x1f, 0x5c, 0xf5, 0xee, 0xe6, 0xf9, 0xfa, 0xdf, 0xb7, 0xad, 0xa4, 0x29, 0xbb,
	0xc7, 0x2b, 0x76, 0xe3, 0x2d, 0xcf, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x15, 0xd6, 0x8f,
	0x0b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RepositoryServiceClient is the client API for RepositoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RepositoryServiceClient interface {
	Add(ctx context.Context, in *domain.Repository, opts ...grpc.CallOption) (*AddRepositoryResponse, error)
}

type repositoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRepositoryServiceClient(cc *grpc.ClientConn) RepositoryServiceClient {
	return &repositoryServiceClient{cc}
}

func (c *repositoryServiceClient) Add(ctx context.Context, in *domain.Repository, opts ...grpc.CallOption) (*AddRepositoryResponse, error) {
	out := new(AddRepositoryResponse)
	err := c.cc.Invoke(ctx, "/service.RepositoryService/add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryServiceServer is the server API for RepositoryService service.
type RepositoryServiceServer interface {
	Add(context.Context, *domain.Repository) (*AddRepositoryResponse, error)
}

func RegisterRepositoryServiceServer(s *grpc.Server, srv RepositoryServiceServer) {
	s.RegisterService(&_RepositoryService_serviceDesc, srv)
}

func _RepositoryService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.Repository)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RepositoryService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).Add(ctx, req.(*domain.Repository))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepositoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.RepositoryService",
	HandlerType: (*RepositoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _RepositoryService_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bitbucket-repository-management-service/internal/proto-files/service/repository-service.proto",
}