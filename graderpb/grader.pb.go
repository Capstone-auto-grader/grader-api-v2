// Code generated by protoc-gen-go. DO NOT EDIT.
// source: graderpb/grader.proto

package grader

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SubmitForGradingRequest struct {
	Tasks                []*Task  `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitForGradingRequest) Reset()         { *m = SubmitForGradingRequest{} }
func (m *SubmitForGradingRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitForGradingRequest) ProtoMessage()    {}
func (*SubmitForGradingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1cca22a1b70c39, []int{0}
}

func (m *SubmitForGradingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitForGradingRequest.Unmarshal(m, b)
}
func (m *SubmitForGradingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitForGradingRequest.Marshal(b, m, deterministic)
}
func (m *SubmitForGradingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitForGradingRequest.Merge(m, src)
}
func (m *SubmitForGradingRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitForGradingRequest.Size(m)
}
func (m *SubmitForGradingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitForGradingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitForGradingRequest proto.InternalMessageInfo

func (m *SubmitForGradingRequest) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type Task struct {
	ImageName   string `protobuf:"bytes,1,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	TestKey     string `protobuf:"bytes,2,opt,name=test_key,json=testKey,proto3" json:"test_key,omitempty"`
	ZipKey      string `protobuf:"bytes,3,opt,name=zip_key,json=zipKey,proto3" json:"zip_key,omitempty"`
	StudentName string `protobuf:"bytes,4,opt,name=student_name,json=studentName,proto3" json:"student_name,omitempty"`
	// Timeout is in seconds.
	Timeout              int32    `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
	CallbackUri          string   `protobuf:"bytes,6,opt,name=callback_uri,json=callbackUri,proto3" json:"callback_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1cca22a1b70c39, []int{1}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *Task) GetTestKey() string {
	if m != nil {
		return m.TestKey
	}
	return ""
}

func (m *Task) GetZipKey() string {
	if m != nil {
		return m.ZipKey
	}
	return ""
}

func (m *Task) GetStudentName() string {
	if m != nil {
		return m.StudentName
	}
	return ""
}

func (m *Task) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Task) GetCallbackUri() string {
	if m != nil {
		return m.CallbackUri
	}
	return ""
}

type SubmitForGradingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitForGradingResponse) Reset()         { *m = SubmitForGradingResponse{} }
func (m *SubmitForGradingResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitForGradingResponse) ProtoMessage()    {}
func (*SubmitForGradingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1cca22a1b70c39, []int{2}
}

func (m *SubmitForGradingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitForGradingResponse.Unmarshal(m, b)
}
func (m *SubmitForGradingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitForGradingResponse.Marshal(b, m, deterministic)
}
func (m *SubmitForGradingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitForGradingResponse.Merge(m, src)
}
func (m *SubmitForGradingResponse) XXX_Size() int {
	return xxx_messageInfo_SubmitForGradingResponse.Size(m)
}
func (m *SubmitForGradingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitForGradingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitForGradingResponse proto.InternalMessageInfo

type CreateImageRequest struct {
	ImageName            string   `protobuf:"bytes,1,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	ImageTar             []byte   `protobuf:"bytes,2,opt,name=image_tar,json=imageTar,proto3" json:"image_tar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateImageRequest) Reset()         { *m = CreateImageRequest{} }
func (m *CreateImageRequest) String() string { return proto.CompactTextString(m) }
func (*CreateImageRequest) ProtoMessage()    {}
func (*CreateImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1cca22a1b70c39, []int{3}
}

func (m *CreateImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImageRequest.Unmarshal(m, b)
}
func (m *CreateImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImageRequest.Marshal(b, m, deterministic)
}
func (m *CreateImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImageRequest.Merge(m, src)
}
func (m *CreateImageRequest) XXX_Size() int {
	return xxx_messageInfo_CreateImageRequest.Size(m)
}
func (m *CreateImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImageRequest proto.InternalMessageInfo

func (m *CreateImageRequest) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *CreateImageRequest) GetImageTar() []byte {
	if m != nil {
		return m.ImageTar
	}
	return nil
}

type CreateImageResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateImageResponse) Reset()         { *m = CreateImageResponse{} }
func (m *CreateImageResponse) String() string { return proto.CompactTextString(m) }
func (*CreateImageResponse) ProtoMessage()    {}
func (*CreateImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1cca22a1b70c39, []int{4}
}

func (m *CreateImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImageResponse.Unmarshal(m, b)
}
func (m *CreateImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImageResponse.Marshal(b, m, deterministic)
}
func (m *CreateImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImageResponse.Merge(m, src)
}
func (m *CreateImageResponse) XXX_Size() int {
	return xxx_messageInfo_CreateImageResponse.Size(m)
}
func (m *CreateImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImageResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SubmitForGradingRequest)(nil), "grader.SubmitForGradingRequest")
	proto.RegisterType((*Task)(nil), "grader.Task")
	proto.RegisterType((*SubmitForGradingResponse)(nil), "grader.SubmitForGradingResponse")
	proto.RegisterType((*CreateImageRequest)(nil), "grader.CreateImageRequest")
	proto.RegisterType((*CreateImageResponse)(nil), "grader.CreateImageResponse")
}

func init() { proto.RegisterFile("graderpb/grader.proto", fileDescriptor_cf1cca22a1b70c39) }

var fileDescriptor_cf1cca22a1b70c39 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x8e, 0xd3, 0x30,
	0x14, 0x96, 0x67, 0xa6, 0xe9, 0xcc, 0x4b, 0x17, 0x23, 0xa3, 0x99, 0x09, 0x29, 0x88, 0x8c, 0x57,
	0x55, 0x17, 0x8d, 0x54, 0x76, 0x95, 0x58, 0x21, 0x51, 0x21, 0x24, 0x84, 0x42, 0x59, 0x47, 0x4e,
	0xfb, 0x14, 0x59, 0x69, 0xec, 0x60, 0x3b, 0x8b, 0x76, 0xc9, 0x15, 0x38, 0x0d, 0xe7, 0xe0, 0x06,
	0x88, 0x83, 0xa0, 0xd8, 0x89, 0x04, 0x94, 0x8a, 0x5d, 0xbe, 0x1f, 0x7f, 0xf1, 0xf7, 0xfc, 0xe0,
	0xae, 0xd4, 0x7c, 0x87, 0xba, 0x29, 0x52, 0xff, 0xb1, 0x68, 0xb4, 0xb2, 0x8a, 0x06, 0x1e, 0xc5,
	0xcf, 0x4a, 0xa5, 0xca, 0x3d, 0xa6, 0xbc, 0x11, 0x29, 0x97, 0x52, 0x59, 0x6e, 0x85, 0x92, 0xc6,
	0xbb, 0xd8, 0x2b, 0x78, 0xf8, 0xd8, 0x16, 0xb5, 0xb0, 0x6f, 0x94, 0x5e, 0x6b, 0xbe, 0x13, 0xb2,
	0xcc, 0xf0, 0x73, 0x8b, 0xc6, 0x52, 0x06, 0x23, 0xcb, 0x4d, 0x65, 0x22, 0x92, 0x5c, 0xce, 0xc2,
	0xe5, 0x64, 0xd1, 0xc7, 0x6f, 0xb8, 0xa9, 0x32, 0x2f, 0xb1, 0x6f, 0x04, 0xae, 0x3a, 0x4c, 0x9f,
	0x03, 0x88, 0x9a, 0x97, 0x98, 0x4b, 0x5e, 0x63, 0x44, 0x12, 0x32, 0xbb, 0xc9, 0x6e, 0x1c, 0xf3,
	0x9e, 0xd7, 0x48, 0x9f, 0xc2, 0xb5, 0x45, 0x63, 0xf3, 0x0a, 0x0f, 0xd1, 0x85, 0x13, 0xc7, 0x1d,
	0x7e, 0x87, 0x07, 0xfa, 0x00, 0xe3, 0xa3, 0x68, 0x9c, 0x72, 0xe9, 0x94, 0xe0, 0x28, 0x9a, 0x4e,
	0x78, 0x84, 0x89, 0xb1, 0xed, 0x0e, 0xa5, 0xf5, 0xa1, 0x57, 0x4e, 0x0d, 0x7b, 0xce, 0xc5, 0x46,
	0x30, 0xb6, 0xa2, 0x46, 0xd5, 0xda, 0x68, 0x94, 0x90, 0xd9, 0x28, 0x1b, 0x60, 0x77, 0x78, 0xcb,
	0xf7, 0xfb, 0x82, 0x6f, 0xab, 0xbc, 0xd5, 0x22, 0x0a, 0xfc, 0xe1, 0x81, 0xfb, 0xa4, 0x05, 0x8b,
	0x21, 0x3a, 0xad, 0x6e, 0x1a, 0x25, 0x0d, 0xb2, 0x0f, 0x40, 0x5f, 0x6b, 0xe4, 0x16, 0xdf, 0x76,
	0x15, 0x86, 0x89, 0xfc, 0xa7, 0xe4, 0x14, 0x3c, 0xc8, 0x2d, 0xd7, 0xae, 0xe5, 0x24, 0xbb, 0x76,
	0xc4, 0x86, 0x6b, 0x76, 0x07, 0x4f, 0xfe, 0x48, 0xf4, 0x3f, 0x5a, 0xfe, 0x20, 0x10, 0xac, 0xdd,
	0x5c, 0x69, 0x0d, 0xb7, 0x7f, 0xdf, 0x87, 0xbe, 0x18, 0x86, 0x7e, 0xe6, 0x91, 0xe2, 0xe4, 0xbc,
	0xa1, 0xaf, 0x72, 0xff, 0xe5, 0xfb, 0xcf, 0xaf, 0x17, 0xb7, 0x2c, 0x74, 0x1b, 0x60, 0x9c, 0x6d,
	0x45, 0xe6, 0x54, 0x40, 0xf8, 0xdb, 0x85, 0x68, 0x3c, 0x04, 0x9d, 0xf6, 0x8e, 0xa7, 0xff, 0xd4,
	0xfa, 0xfc, 0x47, 0x97, 0x3f, 0x65, 0xf7, 0x7e, 0xc3, 0x8c, 0x11, 0xa5, 0xac, 0x51, 0xda, 0x74,
	0xeb, 0xcc, 0x2b, 0x32, 0x2f, 0x02, 0xb7, 0x6b, 0x2f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x20,
	0x5b, 0xda, 0x99, 0xaa, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GraderClient is the client API for Grader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GraderClient interface {
	// SubmitForGrading provides an endpoint for allowing caller to submit a student's name
	// and assignment information for grading.
	//
	// This is a non-blocking call, it returns a "OK" as acknowledgement or "InvalidArgument"
	// for invalid parameters.
	//
	// Once the assignment has been graded, it will call a specific endpoint of the caller.
	SubmitForGrading(ctx context.Context, in *SubmitForGradingRequest, opts ...grpc.CallOption) (*SubmitForGradingResponse, error)
	// CreateImage creates an assignment with a given dockerfile and startup script.
	//
	// This endpoint creates an image on the docker host and returns an unique assignment ID.
	//
	// Calling this endpoint is REQUIRED before grading any assignments.
	CreateImage(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*CreateImageResponse, error)
}

type graderClient struct {
	cc *grpc.ClientConn
}

func NewGraderClient(cc *grpc.ClientConn) GraderClient {
	return &graderClient{cc}
}

func (c *graderClient) SubmitForGrading(ctx context.Context, in *SubmitForGradingRequest, opts ...grpc.CallOption) (*SubmitForGradingResponse, error) {
	out := new(SubmitForGradingResponse)
	err := c.cc.Invoke(ctx, "/grader.Grader/SubmitForGrading", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graderClient) CreateImage(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*CreateImageResponse, error) {
	out := new(CreateImageResponse)
	err := c.cc.Invoke(ctx, "/grader.Grader/CreateImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GraderServer is the server API for Grader service.
type GraderServer interface {
	// SubmitForGrading provides an endpoint for allowing caller to submit a student's name
	// and assignment information for grading.
	//
	// This is a non-blocking call, it returns a "OK" as acknowledgement or "InvalidArgument"
	// for invalid parameters.
	//
	// Once the assignment has been graded, it will call a specific endpoint of the caller.
	SubmitForGrading(context.Context, *SubmitForGradingRequest) (*SubmitForGradingResponse, error)
	// CreateImage creates an assignment with a given dockerfile and startup script.
	//
	// This endpoint creates an image on the docker host and returns an unique assignment ID.
	//
	// Calling this endpoint is REQUIRED before grading any assignments.
	CreateImage(context.Context, *CreateImageRequest) (*CreateImageResponse, error)
}

// UnimplementedGraderServer can be embedded to have forward compatible implementations.
type UnimplementedGraderServer struct {
}

func (*UnimplementedGraderServer) SubmitForGrading(ctx context.Context, req *SubmitForGradingRequest) (*SubmitForGradingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitForGrading not implemented")
}
func (*UnimplementedGraderServer) CreateImage(ctx context.Context, req *CreateImageRequest) (*CreateImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateImage not implemented")
}

func RegisterGraderServer(s *grpc.Server, srv GraderServer) {
	s.RegisterService(&_Grader_serviceDesc, srv)
}

func _Grader_SubmitForGrading_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitForGradingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraderServer).SubmitForGrading(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grader.Grader/SubmitForGrading",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraderServer).SubmitForGrading(ctx, req.(*SubmitForGradingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grader_CreateImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraderServer).CreateImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grader.Grader/CreateImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraderServer).CreateImage(ctx, req.(*CreateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Grader_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grader.Grader",
	HandlerType: (*GraderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitForGrading",
			Handler:    _Grader_SubmitForGrading_Handler,
		},
		{
			MethodName: "CreateImage",
			Handler:    _Grader_CreateImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "graderpb/grader.proto",
}
