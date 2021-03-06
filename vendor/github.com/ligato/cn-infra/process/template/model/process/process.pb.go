// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: process.proto

package process

/*
Package process provides a data model for process manager plugin template
*/

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Template struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cmd                  string            `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
	POptions             *TemplatePOptions `protobuf:"bytes,3,opt,name=p_options,json=pOptions" json:"p_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Template) Reset()         { *m = Template{} }
func (m *Template) String() string { return proto.CompactTextString(m) }
func (*Template) ProtoMessage()    {}
func (*Template) Descriptor() ([]byte, []int) {
	return fileDescriptor_process_a7127b1a63b78f8f, []int{0}
}
func (m *Template) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Template.Unmarshal(m, b)
}
func (m *Template) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Template.Marshal(b, m, deterministic)
}
func (dst *Template) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Template.Merge(dst, src)
}
func (m *Template) XXX_Size() int {
	return xxx_messageInfo_Template.Size(m)
}
func (m *Template) XXX_DiscardUnknown() {
	xxx_messageInfo_Template.DiscardUnknown(m)
}

var xxx_messageInfo_Template proto.InternalMessageInfo

func (m *Template) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Template) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (m *Template) GetPOptions() *TemplatePOptions {
	if m != nil {
		return m.POptions
	}
	return nil
}

type TemplatePOptions struct {
	Args                 []string `protobuf:"bytes,1,rep,name=args" json:"args,omitempty"`
	Restart              int32    `protobuf:"varint,2,opt,name=restart,proto3" json:"restart,omitempty"`
	Detach               bool     `protobuf:"varint,3,opt,name=detach,proto3" json:"detach,omitempty"`
	RunOnStartup         bool     `protobuf:"varint,4,opt,name=run_on_startup,json=runOnStartup,proto3" json:"run_on_startup,omitempty"`
	Notify               bool     `protobuf:"varint,5,opt,name=notify,proto3" json:"notify,omitempty"`
	AutoTerminate        bool     `protobuf:"varint,6,opt,name=auto_terminate,json=autoTerminate,proto3" json:"auto_terminate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemplatePOptions) Reset()         { *m = TemplatePOptions{} }
func (m *TemplatePOptions) String() string { return proto.CompactTextString(m) }
func (*TemplatePOptions) ProtoMessage()    {}
func (*TemplatePOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_process_a7127b1a63b78f8f, []int{0, 0}
}
func (m *TemplatePOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplatePOptions.Unmarshal(m, b)
}
func (m *TemplatePOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplatePOptions.Marshal(b, m, deterministic)
}
func (dst *TemplatePOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplatePOptions.Merge(dst, src)
}
func (m *TemplatePOptions) XXX_Size() int {
	return xxx_messageInfo_TemplatePOptions.Size(m)
}
func (m *TemplatePOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplatePOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TemplatePOptions proto.InternalMessageInfo

func (m *TemplatePOptions) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *TemplatePOptions) GetRestart() int32 {
	if m != nil {
		return m.Restart
	}
	return 0
}

func (m *TemplatePOptions) GetDetach() bool {
	if m != nil {
		return m.Detach
	}
	return false
}

func (m *TemplatePOptions) GetRunOnStartup() bool {
	if m != nil {
		return m.RunOnStartup
	}
	return false
}

func (m *TemplatePOptions) GetNotify() bool {
	if m != nil {
		return m.Notify
	}
	return false
}

func (m *TemplatePOptions) GetAutoTerminate() bool {
	if m != nil {
		return m.AutoTerminate
	}
	return false
}

func init() {
	proto.RegisterType((*Template)(nil), "process.Template")
	proto.RegisterType((*TemplatePOptions)(nil), "process.Template.pOptions")
}

func init() { proto.RegisterFile("process.proto", fileDescriptor_process_a7127b1a63b78f8f) }

var fileDescriptor_process_a7127b1a63b78f8f = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xc1, 0x4e, 0xc3, 0x30,
	0x0c, 0x86, 0x95, 0x75, 0xeb, 0x3a, 0xc3, 0x26, 0xe4, 0x03, 0x8a, 0x76, 0xaa, 0x10, 0x48, 0x3b,
	0xf5, 0x00, 0x07, 0x1e, 0x63, 0x52, 0xd9, 0xbd, 0x0a, 0x5d, 0x80, 0x4a, 0x34, 0x89, 0x1c, 0xf7,
	0xc0, 0xe3, 0xf0, 0x10, 0xbc, 0x1f, 0x8a, 0xbb, 0xec, 0xf6, 0x7f, 0xbf, 0xed, 0xff, 0x97, 0x0c,
	0xdb, 0x40, 0xbe, 0xb7, 0x31, 0x36, 0x81, 0x3c, 0x7b, 0x5c, 0x5f, 0xf0, 0xe1, 0x77, 0x01, 0xd5,
	0xc9, 0x8e, 0xe1, 0xdb, 0xb0, 0x45, 0x84, 0xa5, 0x33, 0xa3, 0xd5, 0xaa, 0x56, 0x87, 0x4d, 0x2b,
	0x1a, 0xef, 0xa0, 0xe8, 0xc7, 0xb3, 0x5e, 0x88, 0x95, 0x24, 0xbe, 0xc2, 0x26, 0x74, 0x3e, 0xf0,
	0xe0, 0x5d, 0xd4, 0x45, 0xad, 0x0e, 0x37, 0xcf, 0xfb, 0x26, 0xc7, 0xe7, 0xac, 0x26, 0x1c, 0xe7,
	0x8d, 0xb6, 0xca, 0x6a, 0xff, 0xa7, 0xe0, 0x0a, 0xa9, 0xcb, 0xd0, 0x67, 0xd4, 0xaa, 0x2e, 0x52,
	0x57, 0xd2, 0xa8, 0x61, 0x4d, 0x36, 0xb2, 0x21, 0x96, 0xbe, 0x55, 0x9b, 0x11, 0xef, 0xa1, 0x3c,
	0x5b, 0x36, 0xfd, 0x97, 0x14, 0x56, 0xed, 0x85, 0xf0, 0x11, 0x76, 0x34, 0xb9, 0xce, 0xbb, 0x4e,
	0xf6, 0xa6, 0xa0, 0x97, 0x32, 0xbf, 0xa5, 0xc9, 0x1d, 0xdd, 0xdb, 0xec, 0xa5, 0x6b, 0xe7, 0x79,
	0xf8, 0xf8, 0xd1, 0xab, 0xf9, 0x7a, 0x26, 0x7c, 0x82, 0x9d, 0x99, 0xd8, 0x77, 0x6c, 0x69, 0x1c,
	0x9c, 0x61, 0xab, 0x4b, 0x99, 0x6f, 0x93, 0x7b, 0xca, 0xe6, 0x7b, 0x29, 0x3f, 0x7b, 0xf9, 0x0f,
	0x00, 0x00, 0xff, 0xff, 0x86, 0x45, 0xd8, 0x07, 0x44, 0x01, 0x00, 0x00,
}
