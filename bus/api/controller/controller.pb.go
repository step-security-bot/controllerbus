// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/controllerbus/bus/api/controller/controller.proto

package bus_api_controller

import (
	fmt "fmt"
	math "math"

	api "github.com/aperturerobotics/controllerbus/bus/api"
	proto "github.com/golang/protobuf/proto"
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

// Config configures the API.
type Config struct {
	// ListenAddr is the address to listen on for connections.
	ListenAddr string `protobuf:"bytes,1,opt,name=listen_addr,json=listenAddr,proto3" json:"listen_addr,omitempty"`
	// BusApiConfig are options for controller bus api.
	BusApiConfig         *api.Config `protobuf:"bytes,2,opt,name=bus_api_config,json=busApiConfig,proto3" json:"bus_api_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_97d5967c7e353488, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetListenAddr() string {
	if m != nil {
		return m.ListenAddr
	}
	return ""
}

func (m *Config) GetBusApiConfig() *api.Config {
	if m != nil {
		return m.BusApiConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "bus.api.controller.Config")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/controllerbus/bus/api/controller/controller.proto", fileDescriptor_97d5967c7e353488)
}

var fileDescriptor_97d5967c7e353488 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8e, 0xbf, 0x0a, 0xc2, 0x40,
	0x0c, 0x87, 0xa9, 0x43, 0xc1, 0xab, 0x28, 0xdc, 0x54, 0x5c, 0x2c, 0x4e, 0x9d, 0xae, 0xa0, 0x38,
	0x39, 0x15, 0x67, 0x97, 0xbe, 0x40, 0xbd, 0x7f, 0xd6, 0x40, 0x6d, 0x8e, 0xdc, 0xdd, 0xfb, 0x4b,
	0x5b, 0xc4, 0xae, 0x0e, 0x81, 0xe4, 0x17, 0xf2, 0x7d, 0x61, 0xf7, 0x0e, 0xc2, 0x2b, 0x2a, 0xa1,
	0xf1, 0x5d, 0x49, 0x67, 0x29, 0x44, 0xb2, 0x84, 0x0a, 0x03, 0x68, 0x5f, 0x69, 0x1c, 0x02, 0x61,
	0xdf, 0x5b, 0x52, 0xd1, 0x57, 0x63, 0x49, 0x07, 0x8b, 0x74, 0xd1, 0x0a, 0x47, 0x18, 0x90, 0x73,
	0x15, 0xbd, 0x90, 0x0e, 0xc4, 0x6f, 0xb3, 0xbf, 0xfe, 0xaf, 0x18, 0x09, 0x13, 0xf0, 0xf8, 0x60,
	0xe9, 0x0d, 0x87, 0x27, 0x74, 0xfc, 0xc0, 0xb2, 0x1e, 0x7c, 0xb0, 0x43, 0x2b, 0x8d, 0xa1, 0x3c,
	0x29, 0x92, 0x72, 0xdd, 0xb0, 0x39, 0xaa, 0x8d, 0x21, 0x7e, 0x61, 0x5b, 0x15, 0x7d, 0x2b, 0x1d,
	0xb4, 0x7a, 0x3a, 0xc9, 0x57, 0x45, 0x52, 0x66, 0xa7, 0x9d, 0xf8, 0x3e, 0x35, 0x93, 0x9a, 0x8d,
	0x8a, 0xbe, 0x76, 0x30, 0x4f, 0x2a, 0x9d, 0x44, 0xe7, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c,
	0xe9, 0x37, 0xc1, 0x0a, 0x01, 0x00, 0x00,
}
