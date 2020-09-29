// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/rbac/v3alpha/rbac.proto

package envoy_config_filter_http_rbac_v3alpha

import (
	fmt "fmt"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/rbac/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type RBAC struct {
	Rules                *v3alpha.RBAC `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	ShadowRules          *v3alpha.RBAC `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules,proto3" json:"shadow_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f8a7c98dce62dc2, []int{0}
}

func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetRules() *v3alpha.RBAC {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *RBAC) GetShadowRules() *v3alpha.RBAC {
	if m != nil {
		return m.ShadowRules
	}
	return nil
}

type RBACPerRoute struct {
	Rbac                 *RBAC    `protobuf:"bytes,2,opt,name=rbac,proto3" json:"rbac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RBACPerRoute) Reset()         { *m = RBACPerRoute{} }
func (m *RBACPerRoute) String() string { return proto.CompactTextString(m) }
func (*RBACPerRoute) ProtoMessage()    {}
func (*RBACPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f8a7c98dce62dc2, []int{1}
}

func (m *RBACPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBACPerRoute.Unmarshal(m, b)
}
func (m *RBACPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBACPerRoute.Marshal(b, m, deterministic)
}
func (m *RBACPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBACPerRoute.Merge(m, src)
}
func (m *RBACPerRoute) XXX_Size() int {
	return xxx_messageInfo_RBACPerRoute.Size(m)
}
func (m *RBACPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_RBACPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_RBACPerRoute proto.InternalMessageInfo

func (m *RBACPerRoute) GetRbac() *RBAC {
	if m != nil {
		return m.Rbac
	}
	return nil
}

func init() {
	proto.RegisterType((*RBAC)(nil), "envoy.config.filter.http.rbac.v3alpha.RBAC")
	proto.RegisterType((*RBACPerRoute)(nil), "envoy.config.filter.http.rbac.v3alpha.RBACPerRoute")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/rbac/v3alpha/rbac.proto", fileDescriptor_0f8a7c98dce62dc2)
}

var fileDescriptor_0f8a7c98dce62dc2 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd0, 0xbf, 0x4b, 0x03, 0x31,
	0x14, 0x07, 0x70, 0x52, 0x4e, 0xd1, 0xb4, 0x83, 0x64, 0x51, 0xba, 0x28, 0x45, 0x41, 0x10, 0x5e,
	0xc4, 0xc3, 0x59, 0x8c, 0x9b, 0xd3, 0x11, 0x70, 0x96, 0xdc, 0x35, 0xf5, 0x02, 0xa1, 0x2f, 0xa4,
	0xe9, 0x69, 0x47, 0xff, 0x73, 0x49, 0x5e, 0x05, 0x3b, 0x08, 0xb7, 0xe5, 0xc7, 0xfb, 0x7c, 0x1f,
	0x7c, 0xf9, 0xbd, 0x5d, 0x0f, 0xb8, 0x93, 0x1d, 0xae, 0x57, 0xee, 0x43, 0xae, 0x9c, 0x4f, 0x36,
	0xca, 0x3e, 0xa5, 0x20, 0x63, 0x6b, 0x3a, 0x39, 0xd4, 0xc6, 0x87, 0xde, 0x94, 0x0b, 0x84, 0x88,
	0x09, 0xc5, 0x4d, 0x11, 0x40, 0x02, 0x48, 0x40, 0x16, 0x50, 0x86, 0xf6, 0x62, 0x7e, 0x7d, 0x10,
	0xfc, 0x4f, 0xd8, 0xfc, 0x7c, 0x30, 0xde, 0x2d, 0x4d, 0xb2, 0xf2, 0xf7, 0x40, 0x1f, 0x8b, 0x6f,
	0xc6, 0x2b, 0xad, 0x9e, 0x5f, 0xc4, 0x23, 0x3f, 0x8a, 0x5b, 0x6f, 0x37, 0x17, 0xec, 0x8a, 0xdd,
	0x4e, 0x1f, 0x2e, 0xe1, 0x60, 0xfd, 0xdf, 0x95, 0x90, 0xe7, 0x35, 0x4d, 0x0b, 0xc5, 0x67, 0x9b,
	0xde, 0x2c, 0xf1, 0xf3, 0x9d, 0xf4, 0x64, 0x9c, 0x9e, 0x12, 0xd2, 0xd9, 0x2c, 0xde, 0xf8, 0x2c,
	0x3f, 0x36, 0x36, 0x6a, 0xdc, 0x26, 0x2b, 0x9e, 0x78, 0x95, 0xc5, 0x3e, 0xeb, 0x0e, 0x46, 0x15,
	0x41, 0xb9, 0x05, 0xbe, 0x56, 0x27, 0xec, 0x6c, 0xa2, 0x14, 0xaf, 0x1d, 0x12, 0x0e, 0x11, 0xbf,
	0x76, 0xe3, 0x72, 0xd4, 0xa9, 0x6e, 0x4d, 0xd7, 0xe4, 0x72, 0x1a, 0xd6, 0x1e, 0x97, 0x96, 0xea,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x4b, 0xb7, 0xcd, 0xbf, 0x01, 0x00, 0x00,
}
