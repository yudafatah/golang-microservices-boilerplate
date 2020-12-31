package domain

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Repository struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	IsPrivate            bool     `protobuf:"varint,4,opt,name=isPrivate,proto3" json:"isPrivate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repository) Reset()         { *m = Repository{} }
func (m *Repository) String() string { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()    {}
func (*Repository) Descriptor() ([]byte, []int) {
	return fileDescriptor_0293ad53cd466a6d, []int{0}
}

func (m *Repository) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repository.Unmarshal(m, b)
}
func (m *Repository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repository.Marshal(b, m, deterministic)
}
func (m *Repository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repository.Merge(m, src)
}
func (m *Repository) XXX_Size() int {
	return xxx_messageInfo_Repository.Size(m)
}
func (m *Repository) XXX_DiscardUnknown() {
	xxx_messageInfo_Repository.DiscardUnknown(m)
}

var xxx_messageInfo_Repository proto.InternalMessageInfo

func (m *Repository) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Repository) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Repository) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Repository) GetIsPrivate() bool {
	if m != nil {
		return m.IsPrivate
	}
	return false
}

func init() {
	proto.RegisterType((*Repository)(nil), "domain.Repository")
}

func init() {
	proto.RegisterFile("bitbucket-repository-management-service/internal/proto-files/domain/repository.proto", fileDescriptor_0293ad53cd466a6d)
}

var fileDescriptor_0293ad53cd466a6d = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0x31, 0x4b, 0xc0, 0x30,
	0x10, 0x85, 0x49, 0x5b, 0x8a, 0xcd, 0xe0, 0x90, 0x41, 0x32, 0x38, 0x14, 0xa7, 0x2e, 0x31, 0x83,
	0xab, 0x38, 0xb8, 0xb9, 0x49, 0x70, 0x72, 0x4b, 0x9b, 0x6b, 0x39, 0x6c, 0x92, 0x72, 0x49, 0x0b,
	0xfe, 0x7b, 0x21, 0x16, 0x3b, 0xbb, 0xdd, 0xdd, 0x77, 0x0f, 0xde, 0xc7, 0x3f, 0x46, 0xcc, 0xe3,
	0x3e, 0x7d, 0x41, 0x56, 0x04, 0x5b, 0x4c, 0x98, 0x23, 0x7d, 0x2b, 0x6f, 0x83, 0x5d, 0xc0, 0x43,
	0xc8, 0x2a, 0x01, 0x1d, 0x38, 0x81, 0xc6, 0x90, 0x81, 0x82, 0x5d, 0xf5, 0x46, 0x31, 0x47, 0x35,
	0xe3, 0x0a, 0x49, 0xbb, 0xe8, 0x2d, 0x06, 0x7d, 0x25, 0x1f, 0x0b, 0x15, 0xed, 0x2f, 0x78, 0x98,
	0x39, 0x37, 0x7f, 0x4c, 0xdc, 0xf2, 0x0a, 0x9d, 0x64, 0x3d, 0x1b, 0x6a, 0x53, 0xa1, 0x13, 0x82,
	0x37, 0xc1, 0x7a, 0x90, 0x55, 0xcf, 0x86, 0xce, 0x94, 0x59, 0xdc, 0xf1, 0x76, 0x4f, 0x40, 0x6f,
	0x4e, 0xd6, 0xe5, 0xef, 0xdc, 0xc4, 0x3d, 0xef, 0x30, 0xbd, 0x13, 0x1e, 0x36, 0x83, 0x6c, 0x7a,
	0x36, 0xdc, 0x98, 0xeb, 0xf0, 0xfa, 0xf2, 0xf9, 0xfc, 0x6f, 0x8f, 0x85, 0xb6, 0xe9, 0x14, 0x18,
	0xdb, 0x52, 0xfb, 0xe9, 0x27, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xaa, 0x34, 0xc5, 0x0e, 0x01, 0x00,
	0x00,
}