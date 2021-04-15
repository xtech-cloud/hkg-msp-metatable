// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/metatable/vocabulary.proto

package metatable

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

// 词汇实体
type VocabularyEntity struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Label                []string `protobuf:"bytes,3,rep,name=label,proto3" json:"label,omitempty"`
	Schema               string   `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VocabularyEntity) Reset()         { *m = VocabularyEntity{} }
func (m *VocabularyEntity) String() string { return proto.CompactTextString(m) }
func (*VocabularyEntity) ProtoMessage()    {}
func (*VocabularyEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_78db5907983eb399, []int{0}
}

func (m *VocabularyEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VocabularyEntity.Unmarshal(m, b)
}
func (m *VocabularyEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VocabularyEntity.Marshal(b, m, deterministic)
}
func (m *VocabularyEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VocabularyEntity.Merge(m, src)
}
func (m *VocabularyEntity) XXX_Size() int {
	return xxx_messageInfo_VocabularyEntity.Size(m)
}
func (m *VocabularyEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_VocabularyEntity.DiscardUnknown(m)
}

var xxx_messageInfo_VocabularyEntity proto.InternalMessageInfo

func (m *VocabularyEntity) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *VocabularyEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VocabularyEntity) GetLabel() []string {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *VocabularyEntity) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

type VocabularyListResponse struct {
	Status               *Status             `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                int64               `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Entity               []*VocabularyEntity `protobuf:"bytes,3,rep,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VocabularyListResponse) Reset()         { *m = VocabularyListResponse{} }
func (m *VocabularyListResponse) String() string { return proto.CompactTextString(m) }
func (*VocabularyListResponse) ProtoMessage()    {}
func (*VocabularyListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78db5907983eb399, []int{1}
}

func (m *VocabularyListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VocabularyListResponse.Unmarshal(m, b)
}
func (m *VocabularyListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VocabularyListResponse.Marshal(b, m, deterministic)
}
func (m *VocabularyListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VocabularyListResponse.Merge(m, src)
}
func (m *VocabularyListResponse) XXX_Size() int {
	return xxx_messageInfo_VocabularyListResponse.Size(m)
}
func (m *VocabularyListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VocabularyListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VocabularyListResponse proto.InternalMessageInfo

func (m *VocabularyListResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VocabularyListResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *VocabularyListResponse) GetEntity() []*VocabularyEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type VocabularyFindResponse struct {
	Status               *Status             `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Entity               []*VocabularyEntity `protobuf:"bytes,2,rep,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VocabularyFindResponse) Reset()         { *m = VocabularyFindResponse{} }
func (m *VocabularyFindResponse) String() string { return proto.CompactTextString(m) }
func (*VocabularyFindResponse) ProtoMessage()    {}
func (*VocabularyFindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78db5907983eb399, []int{2}
}

func (m *VocabularyFindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VocabularyFindResponse.Unmarshal(m, b)
}
func (m *VocabularyFindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VocabularyFindResponse.Marshal(b, m, deterministic)
}
func (m *VocabularyFindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VocabularyFindResponse.Merge(m, src)
}
func (m *VocabularyFindResponse) XXX_Size() int {
	return xxx_messageInfo_VocabularyFindResponse.Size(m)
}
func (m *VocabularyFindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VocabularyFindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VocabularyFindResponse proto.InternalMessageInfo

func (m *VocabularyFindResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VocabularyFindResponse) GetEntity() []*VocabularyEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*VocabularyEntity)(nil), "metatable.VocabularyEntity")
	proto.RegisterType((*VocabularyListResponse)(nil), "metatable.VocabularyListResponse")
	proto.RegisterType((*VocabularyFindResponse)(nil), "metatable.VocabularyFindResponse")
}

func init() {
	proto.RegisterFile("proto/metatable/vocabulary.proto", fileDescriptor_78db5907983eb399)
}

var fileDescriptor_78db5907983eb399 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0xed, 0xbf, 0x5f, 0xa0, 0xd3, 0xcb, 0xcf, 0x41, 0x4a, 0xa8, 0x3d, 0xd4, 0x9c, 0xea, 0xa5,
	0x85, 0xf6, 0xee, 0xc1, 0x3f, 0x05, 0xc1, 0xd3, 0x0a, 0x82, 0xc7, 0x49, 0x3b, 0xd0, 0xe0, 0x26,
	0xa9, 0xd9, 0x89, 0xd8, 0x2f, 0xe1, 0x47, 0xf0, 0xb3, 0x4a, 0x76, 0x6b, 0xb3, 0x16, 0x91, 0xe2,
	0x6d, 0x67, 0xde, 0xdb, 0xf7, 0xde, 0xbe, 0x04, 0x46, 0x9b, 0x22, 0x97, 0x7c, 0x9a, 0xb2, 0x90,
	0x50, 0xac, 0x79, 0xfa, 0x9a, 0x2f, 0x29, 0x2e, 0x35, 0x15, 0xdb, 0x89, 0x85, 0xb0, 0xbb, 0xc7,
	0x06, 0xc3, 0x43, 0xb2, 0x59, 0x53, 0xc1, 0x2b, 0x47, 0x8c, 0xd6, 0xf0, 0xff, 0x71, 0x7f, 0xf9,
	0x36, 0x93, 0x44, 0xb6, 0x88, 0xd0, 0x29, 0xcb, 0x64, 0x15, 0x36, 0x47, 0xcd, 0x71, 0x57, 0xd9,
	0x73, 0xb5, 0xcb, 0x28, 0xe5, 0xb0, 0xe5, 0x76, 0xd5, 0x19, 0x4f, 0xe1, 0x9f, 0xa6, 0x98, 0x75,
	0xd8, 0x1e, 0xb5, 0xc7, 0x5d, 0xe5, 0x06, 0xec, 0x43, 0x60, 0x96, 0x6b, 0x4e, 0x29, 0xec, 0x58,
	0xee, 0x6e, 0x8a, 0xde, 0x9b, 0xd0, 0xaf, 0xad, 0xee, 0x13, 0x23, 0x8a, 0xcd, 0x26, 0xcf, 0x0c,
	0xe3, 0x05, 0x04, 0x46, 0x48, 0x4a, 0x63, 0x2d, 0x7b, 0xb3, 0x93, 0xc9, 0x3e, 0xed, 0xe4, 0xc1,
	0x02, 0x6a, 0x47, 0xa8, 0x3c, 0x25, 0x17, 0xd2, 0x36, 0x48, 0x5b, 0xb9, 0x01, 0xe7, 0x10, 0xb0,
	0xcd, 0x6e, 0xa3, 0xf4, 0x66, 0x67, 0x9e, 0xc0, 0xe1, 0xf3, 0xd4, 0x8e, 0x1a, 0xbd, 0xf9, 0x79,
	0x16, 0x49, 0xb6, 0xfa, 0x4b, 0x9e, 0xda, 0xb9, 0x75, 0xb4, 0xf3, 0xec, 0xa3, 0x05, 0x50, 0x83,
	0xb8, 0x00, 0xb8, 0x4b, 0x37, 0x79, 0x21, 0x4f, 0x94, 0x6a, 0x1c, 0x7a, 0x0a, 0xf5, 0x5a, 0xf1,
	0x4b, 0xc9, 0x46, 0x06, 0xa1, 0x87, 0x5e, 0x69, 0xca, 0x9e, 0xbf, 0x42, 0x47, 0x0d, 0xbc, 0x86,
	0x4e, 0x55, 0x2b, 0xf6, 0x3d, 0x8e, 0xeb, 0xd9, 0xdd, 0x3d, 0xff, 0x31, 0x9b, 0xff, 0x25, 0x9c,
	0x48, 0xd5, 0xc5, 0x37, 0x11, 0x57, 0xce, 0x6f, 0x22, 0x7e, 0x7d, 0x51, 0x03, 0x2f, 0x21, 0xb8,
	0x61, 0xcd, 0xc2, 0xe8, 0xe7, 0x75, 0xab, 0x23, 0x5e, 0x12, 0x07, 0xf6, 0xe7, 0x9c, 0x7f, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x79, 0xd0, 0x7e, 0xa7, 0xe9, 0x02, 0x00, 0x00,
}
