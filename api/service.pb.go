// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: service.proto

package api

import (
	context "context"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x09, 0x6d, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xeb, 0x05,
	0x0a, 0x08, 0x48, 0x49, 0x54, 0x4c, 0x42, 0x61, 0x63, 0x6b, 0x12, 0x4d, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x69, 0x6e, 0x67, 0x92, 0x41, 0x11, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x09,
	0x70, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x65, 0x73, 0x74, 0x12, 0x80, 0x01, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x49, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x2f, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x0c, 0xe5, 0x88, 0x9b, 0xe5, 0xbb, 0xba, 0xe7, 0x94, 0xa8, 0xe6, 0x88,
	0xb7, 0x1a, 0x19, 0xe9, 0x9c, 0x80, 0xe8, 0xa6, 0x81, 0xe6, 0x8c, 0x87, 0xe5, 0xae, 0x9a, 0xe7,
	0x94, 0xa8, 0xe6, 0x88, 0xb7, 0x72, 0x6f, 0x6c, 0x65, 0x20, 0x69, 0x64, 0x12, 0x85, 0x01, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x4e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x1a, 0x17, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x92, 0x41, 0x26, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x1e, 0xe6, 0x9b, 0xb4, 0xe6, 0x96, 0xb0, 0xe7, 0x94, 0xa8, 0xe6,
	0x88, 0xb7, 0xe4, 0xbf, 0xa1, 0xe6, 0x81, 0xaf, 0xef, 0xbc, 0x8c, 0xe5, 0xaf, 0x86, 0xe7, 0xa0,
	0x81, 0xe7, 0xad, 0x89, 0x12, 0x6a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19,
	0x2a, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x92, 0x41, 0x14, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x0c, 0xe5, 0x88, 0xa0, 0xe9, 0x99, 0xa4, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7,
	0x12, 0x57, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a,
	0x01, 0x2a, 0x92, 0x41, 0x14, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x0c, 0xe7, 0x94, 0xa8,
	0xe6, 0x88, 0xb7, 0xe7, 0x99, 0xbb, 0xe5, 0xbd, 0x95, 0x12, 0x59, 0x0a, 0x06, 0x4c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x92, 0x41, 0x14,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x0c, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7, 0x99,
	0xbb, 0xe5, 0x87, 0xba, 0x12, 0x65, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e,
	0x12, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x92, 0x41,
	0x1a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0xe5, 0x88, 0x97, 0xe5, 0x87, 0xba, 0xe6,
	0x89, 0x80, 0xe6, 0x9c, 0x89, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0x42, 0x54, 0x5a, 0x03, 0x61,
	0x70, 0x69, 0x92, 0x41, 0x4c, 0x12, 0x33, 0x0a, 0x0c, 0x74, 0x61, 0x74, 0x73, 0x2d, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x1e, 0x74, 0x61, 0x74, 0x73, 0x2d, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x20, 0x68, 0x74, 0x74, 0x70, 0x20, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0xe6,
	0x96, 0x87, 0xe6, 0xa1, 0xa3, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x1a, 0x12, 0x31, 0x30, 0x2e, 0x34,
	0x2e, 0x31, 0x39, 0x36, 0x2e, 0x31, 0x33, 0x33, 0x3a, 0x33, 0x30, 0x30, 0x39, 0x30, 0x2a, 0x01,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*PingReq)(nil),        // 0: api.PingReq
	(*CreateUserReq)(nil),  // 1: api.CreateUserReq
	(*UpdateUserReq)(nil),  // 2: api.UpdateUserReq
	(*DeleteUserReq)(nil),  // 3: api.DeleteUserReq
	(*LoginReq)(nil),       // 4: api.LoginReq
	(*LogoutReq)(nil),      // 5: api.LogoutReq
	(*ListUsersReq)(nil),   // 6: api.ListUsersReq
	(*PingResp)(nil),       // 7: api.PingResp
	(*CreateUserResp)(nil), // 8: api.CreateUserResp
	(*UpdateUserResp)(nil), // 9: api.UpdateUserResp
	(*SuccessResp)(nil),    // 10: api.SuccessResp
	(*LoginResp)(nil),      // 11: api.LoginResp
	(*ListUsersResp)(nil),  // 12: api.ListUsersResp
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: api.HITLBack.Ping:input_type -> api.PingReq
	1,  // 1: api.HITLBack.CreateUser:input_type -> api.CreateUserReq
	2,  // 2: api.HITLBack.UpdateUser:input_type -> api.UpdateUserReq
	3,  // 3: api.HITLBack.DeleteUser:input_type -> api.DeleteUserReq
	4,  // 4: api.HITLBack.Login:input_type -> api.LoginReq
	5,  // 5: api.HITLBack.Logout:input_type -> api.LogoutReq
	6,  // 6: api.HITLBack.ListUsers:input_type -> api.ListUsersReq
	7,  // 7: api.HITLBack.Ping:output_type -> api.PingResp
	8,  // 8: api.HITLBack.CreateUser:output_type -> api.CreateUserResp
	9,  // 9: api.HITLBack.UpdateUser:output_type -> api.UpdateUserResp
	10, // 10: api.HITLBack.DeleteUser:output_type -> api.SuccessResp
	11, // 11: api.HITLBack.Login:output_type -> api.LoginResp
	10, // 12: api.HITLBack.Logout:output_type -> api.SuccessResp
	12, // 13: api.HITLBack.ListUsers:output_type -> api.ListUsersResp
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_mod_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HITLBackClient is the client API for HITLBack service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HITLBackClient interface {
	// ping
	//
	// 简单测试
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error)
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
	DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*SuccessResp, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*SuccessResp, error)
	ListUsers(ctx context.Context, in *ListUsersReq, opts ...grpc.CallOption) (*ListUsersResp, error)
}

type hITLBackClient struct {
	cc grpc.ClientConnInterface
}

func NewHITLBackClient(cc grpc.ClientConnInterface) HITLBackClient {
	return &hITLBackClient{cc}
}

func (c *hITLBackClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error) {
	out := new(PingResp)
	err := c.cc.Invoke(ctx, "/api.HITLBack/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hITLBackClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error) {
	out := new(CreateUserResp)
	err := c.cc.Invoke(ctx, "/api.HITLBack/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hITLBackClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	out := new(UpdateUserResp)
	err := c.cc.Invoke(ctx, "/api.HITLBack/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hITLBackClient) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	out := new(SuccessResp)
	err := c.cc.Invoke(ctx, "/api.HITLBack/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hITLBackClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/api.HITLBack/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hITLBackClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	out := new(SuccessResp)
	err := c.cc.Invoke(ctx, "/api.HITLBack/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hITLBackClient) ListUsers(ctx context.Context, in *ListUsersReq, opts ...grpc.CallOption) (*ListUsersResp, error) {
	out := new(ListUsersResp)
	err := c.cc.Invoke(ctx, "/api.HITLBack/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HITLBackServer is the server API for HITLBack service.
type HITLBackServer interface {
	// ping
	//
	// 简单测试
	Ping(context.Context, *PingReq) (*PingResp, error)
	CreateUser(context.Context, *CreateUserReq) (*CreateUserResp, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error)
	DeleteUser(context.Context, *DeleteUserReq) (*SuccessResp, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Logout(context.Context, *LogoutReq) (*SuccessResp, error)
	ListUsers(context.Context, *ListUsersReq) (*ListUsersResp, error)
}

// UnimplementedHITLBackServer can be embedded to have forward compatible implementations.
type UnimplementedHITLBackServer struct {
}

func (*UnimplementedHITLBackServer) Ping(context.Context, *PingReq) (*PingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedHITLBackServer) CreateUser(context.Context, *CreateUserReq) (*CreateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedHITLBackServer) UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedHITLBackServer) DeleteUser(context.Context, *DeleteUserReq) (*SuccessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (*UnimplementedHITLBackServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedHITLBackServer) Logout(context.Context, *LogoutReq) (*SuccessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (*UnimplementedHITLBackServer) ListUsers(context.Context, *ListUsersReq) (*ListUsersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}

func RegisterHITLBackServer(s *grpc.Server, srv HITLBackServer) {
	s.RegisterService(&_HITLBack_serviceDesc, srv)
}

func _HITLBack_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HITLBackServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.HITLBack/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HITLBackServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HITLBack_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HITLBackServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.HITLBack/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HITLBackServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HITLBack_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HITLBackServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.HITLBack/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HITLBackServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HITLBack_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HITLBackServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.HITLBack/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HITLBackServer).DeleteUser(ctx, req.(*DeleteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HITLBack_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HITLBackServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.HITLBack/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HITLBackServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HITLBack_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HITLBackServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.HITLBack/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HITLBackServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HITLBack_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HITLBackServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.HITLBack/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HITLBackServer).ListUsers(ctx, req.(*ListUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _HITLBack_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.HITLBack",
	HandlerType: (*HITLBackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _HITLBack_Ping_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _HITLBack_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _HITLBack_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _HITLBack_DeleteUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _HITLBack_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _HITLBack_Logout_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _HITLBack_ListUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}