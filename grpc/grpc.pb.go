// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/controllerbus/grpc/grpc.proto

package controllerbus_grpc

import (
	context "context"
	fmt "fmt"
	controller "github.com/aperturerobotics/controllerbus/controller"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// GetBusInfoRequest is the request type for GetBusInfo.
type GetBusInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBusInfoRequest) Reset()         { *m = GetBusInfoRequest{} }
func (m *GetBusInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetBusInfoRequest) ProtoMessage()    {}
func (*GetBusInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d9069eab812c2c8, []int{0}
}

func (m *GetBusInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBusInfoRequest.Unmarshal(m, b)
}
func (m *GetBusInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBusInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetBusInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBusInfoRequest.Merge(m, src)
}
func (m *GetBusInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetBusInfoRequest.Size(m)
}
func (m *GetBusInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBusInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBusInfoRequest proto.InternalMessageInfo

// GetBusInfoResponse is the response type for GetBusInfo.
type GetBusInfoResponse struct {
	// RunningControllers is the list of running controllers.
	RunningControllers []*controller.Info `protobuf:"bytes,1,rep,name=running_controllers,json=runningControllers,proto3" json:"running_controllers,omitempty"`
	// RunningDirectives is the list of running directives.
	RunningDirectives    []*DirectiveInfo `protobuf:"bytes,2,rep,name=running_directives,json=runningDirectives,proto3" json:"running_directives,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetBusInfoResponse) Reset()         { *m = GetBusInfoResponse{} }
func (m *GetBusInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetBusInfoResponse) ProtoMessage()    {}
func (*GetBusInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d9069eab812c2c8, []int{1}
}

func (m *GetBusInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBusInfoResponse.Unmarshal(m, b)
}
func (m *GetBusInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBusInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetBusInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBusInfoResponse.Merge(m, src)
}
func (m *GetBusInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetBusInfoResponse.Size(m)
}
func (m *GetBusInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBusInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBusInfoResponse proto.InternalMessageInfo

func (m *GetBusInfoResponse) GetRunningControllers() []*controller.Info {
	if m != nil {
		return m.RunningControllers
	}
	return nil
}

func (m *GetBusInfoResponse) GetRunningDirectives() []*DirectiveInfo {
	if m != nil {
		return m.RunningDirectives
	}
	return nil
}

// DirectiveInfo contains directive information.
type DirectiveInfo struct {
	// Name is the directive name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// DebugVals contains the directive debug values.
	DebugVals            []*DebugValue `protobuf:"bytes,2,rep,name=debug_vals,json=debugVals,proto3" json:"debug_vals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DirectiveInfo) Reset()         { *m = DirectiveInfo{} }
func (m *DirectiveInfo) String() string { return proto.CompactTextString(m) }
func (*DirectiveInfo) ProtoMessage()    {}
func (*DirectiveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d9069eab812c2c8, []int{2}
}

func (m *DirectiveInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectiveInfo.Unmarshal(m, b)
}
func (m *DirectiveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectiveInfo.Marshal(b, m, deterministic)
}
func (m *DirectiveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectiveInfo.Merge(m, src)
}
func (m *DirectiveInfo) XXX_Size() int {
	return xxx_messageInfo_DirectiveInfo.Size(m)
}
func (m *DirectiveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectiveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DirectiveInfo proto.InternalMessageInfo

func (m *DirectiveInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DirectiveInfo) GetDebugVals() []*DebugValue {
	if m != nil {
		return m.DebugVals
	}
	return nil
}

// DebugValue is a debug value.
type DebugValue struct {
	// Key is the debug value key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Values are the debug value values.
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugValue) Reset()         { *m = DebugValue{} }
func (m *DebugValue) String() string { return proto.CompactTextString(m) }
func (*DebugValue) ProtoMessage()    {}
func (*DebugValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d9069eab812c2c8, []int{3}
}

func (m *DebugValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugValue.Unmarshal(m, b)
}
func (m *DebugValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugValue.Marshal(b, m, deterministic)
}
func (m *DebugValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugValue.Merge(m, src)
}
func (m *DebugValue) XXX_Size() int {
	return xxx_messageInfo_DebugValue.Size(m)
}
func (m *DebugValue) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugValue.DiscardUnknown(m)
}

var xxx_messageInfo_DebugValue proto.InternalMessageInfo

func (m *DebugValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DebugValue) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*GetBusInfoRequest)(nil), "controllerbus.grpc.GetBusInfoRequest")
	proto.RegisterType((*GetBusInfoResponse)(nil), "controllerbus.grpc.GetBusInfoResponse")
	proto.RegisterType((*DirectiveInfo)(nil), "controllerbus.grpc.DirectiveInfo")
	proto.RegisterType((*DebugValue)(nil), "controllerbus.grpc.DebugValue")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/controllerbus/grpc/grpc.proto", fileDescriptor_0d9069eab812c2c8)
}

var fileDescriptor_0d9069eab812c2c8 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xdf, 0xbc, 0x95, 0x42, 0x47, 0x84, 0x76, 0x2a, 0x52, 0x7a, 0x90, 0x1a, 0x50, 0x7a,
	0x4a, 0xa1, 0x82, 0xe0, 0xc1, 0x83, 0x55, 0x11, 0x6f, 0x12, 0xc1, 0x9b, 0x94, 0x64, 0x3b, 0xc6,
	0xc5, 0x74, 0x37, 0xee, 0x9f, 0x80, 0x1f, 0xc9, 0x6f, 0x29, 0xd9, 0xa6, 0x6e, 0x4a, 0x0b, 0x7a,
	0x09, 0xcf, 0xce, 0xf3, 0xec, 0x6f, 0x32, 0x3b, 0x70, 0x99, 0x71, 0xf3, 0x66, 0xd3, 0x88, 0xc9,
	0xe5, 0x24, 0x29, 0x48, 0x19, 0xab, 0x48, 0xc9, 0x54, 0x1a, 0xce, 0xf4, 0x84, 0x49, 0x61, 0x94,
	0xcc, 0x73, 0x52, 0xa9, 0xd5, 0x93, 0x4c, 0x15, 0xcc, 0x7d, 0xa2, 0x42, 0x49, 0x23, 0x11, 0x37,
	0xec, 0xa8, 0x72, 0x86, 0x77, 0x7f, 0xc7, 0xf9, 0x53, 0x43, 0xae, 0xd0, 0x61, 0x1f, 0x7a, 0xf7,
	0x64, 0x66, 0x56, 0x3f, 0x88, 0x57, 0x19, 0xd3, 0x87, 0x25, 0x6d, 0xc2, 0xaf, 0x00, 0xb0, 0x59,
	0xd5, 0x85, 0x14, 0x9a, 0xf0, 0x1a, 0xfa, 0xca, 0x0a, 0xc1, 0x45, 0x36, 0xf7, 0x1c, 0x3d, 0x08,
	0x46, 0xad, 0xf1, 0xfe, 0xb4, 0x1b, 0x35, 0xd8, 0xee, 0x1a, 0xd6, 0xe1, 0x1b, 0x9f, 0xc5, 0x47,
	0x58, 0x57, 0xe7, 0x0b, 0xae, 0x88, 0x19, 0x5e, 0x92, 0x1e, 0xfc, 0x77, 0x84, 0x93, 0x68, 0x7b,
	0xcc, 0xe8, 0x76, 0x9d, 0x72, 0xc8, 0x5e, 0x7d, 0xf9, 0xa7, 0xaa, 0xc3, 0x14, 0x0e, 0x36, 0x32,
	0x88, 0xb0, 0x27, 0x92, 0x25, 0x0d, 0x82, 0x51, 0x30, 0xee, 0xc4, 0x4e, 0xe3, 0x15, 0xc0, 0x82,
	0x52, 0x9b, 0xcd, 0xcb, 0x24, 0x5f, 0xb7, 0x3b, 0xde, 0xd9, 0xae, 0x4a, 0x3d, 0x27, 0xb9, 0xa5,
	0xb8, 0xb3, 0xa8, 0xb5, 0x0e, 0x2f, 0x00, 0xbc, 0x81, 0x5d, 0x68, 0xbd, 0xd3, 0x67, 0xcd, 0xaf,
	0x24, 0x1e, 0x41, 0xbb, 0xac, 0xac, 0x15, 0xba, 0x13, 0xd7, 0xa7, 0xa9, 0x85, 0x43, 0x3f, 0xfc,
	0xcc, 0xea, 0x27, 0x52, 0x25, 0x67, 0x84, 0x2f, 0x00, 0xfe, 0x79, 0xf1, 0x74, 0xd7, 0x8f, 0x6c,
	0x2d, 0x65, 0x78, 0xf6, 0x5b, 0x6c, 0xb5, 0xa5, 0xf0, 0x5f, 0xda, 0x76, 0xab, 0x3d, 0xff, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xb9, 0x96, 0x20, 0x56, 0x72, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ControllerBusServiceClient is the client API for ControllerBusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControllerBusServiceClient interface {
	// GetBusInfo requests information about the controller bus.
	GetBusInfo(ctx context.Context, in *GetBusInfoRequest, opts ...grpc.CallOption) (*GetBusInfoResponse, error)
}

type controllerBusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewControllerBusServiceClient(cc grpc.ClientConnInterface) ControllerBusServiceClient {
	return &controllerBusServiceClient{cc}
}

func (c *controllerBusServiceClient) GetBusInfo(ctx context.Context, in *GetBusInfoRequest, opts ...grpc.CallOption) (*GetBusInfoResponse, error) {
	out := new(GetBusInfoResponse)
	err := c.cc.Invoke(ctx, "/controllerbus.grpc.ControllerBusService/GetBusInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControllerBusServiceServer is the server API for ControllerBusService service.
type ControllerBusServiceServer interface {
	// GetBusInfo requests information about the controller bus.
	GetBusInfo(context.Context, *GetBusInfoRequest) (*GetBusInfoResponse, error)
}

// UnimplementedControllerBusServiceServer can be embedded to have forward compatible implementations.
type UnimplementedControllerBusServiceServer struct {
}

func (*UnimplementedControllerBusServiceServer) GetBusInfo(ctx context.Context, req *GetBusInfoRequest) (*GetBusInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusInfo not implemented")
}

func RegisterControllerBusServiceServer(s *grpc.Server, srv ControllerBusServiceServer) {
	s.RegisterService(&_ControllerBusService_serviceDesc, srv)
}

func _ControllerBusService_GetBusInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerBusServiceServer).GetBusInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controllerbus.grpc.ControllerBusService/GetBusInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerBusServiceServer).GetBusInfo(ctx, req.(*GetBusInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ControllerBusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "controllerbus.grpc.ControllerBusService",
	HandlerType: (*ControllerBusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBusInfo",
			Handler:    _ControllerBusService_GetBusInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/aperturerobotics/controllerbus/grpc/grpc.proto",
}
