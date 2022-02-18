// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: mod.proto

package api

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_mod_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_mod_proto_rawDescGZIP(), []int{0}
}

func (x *PingReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IMsg string `protobuf:"bytes,1,opt,name=i_msg,proto3" json:"i_msg,omitempty"`
}

func (x *PingResp) Reset() {
	*x = PingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResp) ProtoMessage() {}

func (x *PingResp) ProtoReflect() protoreflect.Message {
	mi := &file_mod_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResp.ProtoReflect.Descriptor instead.
func (*PingResp) Descriptor() ([]byte, []int) {
	return file_mod_proto_rawDescGZIP(), []int{1}
}

func (x *PingResp) GetIMsg() string {
	if x != nil {
		return x.IMsg
	}
	return ""
}

type CreateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户名，最少4个字符， 字母数字下划线、横杠、点、@ example: richard@qq.com
	//
	// required
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// 是否是 root >1:yes, 0: no
	//
	// requied
	IsRoot int32 `protobuf:"varint,2,opt,name=is_root,json=isRoot,proto3" json:"is_root,omitempty"`
	// 密码
	//
	// required
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserReq) Reset() {
	*x = CreateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserReq) ProtoMessage() {}

func (x *CreateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_mod_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserReq.ProtoReflect.Descriptor instead.
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return file_mod_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserReq) GetIsRoot() int32 {
	if x != nil {
		return x.IsRoot
	}
	return 0
}

func (x *CreateUserReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *User `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateUserResp) Reset() {
	*x = CreateUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResp) ProtoMessage() {}

func (x *CreateUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_mod_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResp.ProtoReflect.Descriptor instead.
func (*CreateUserResp) Descriptor() ([]byte, []int) {
	return file_mod_proto_rawDescGZIP(), []int{3}
}

func (x *CreateUserResp) GetData() *User {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 数据库id
	Username string          `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	User     *UpdateUserBody `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUserReq) Reset() {
	*x = UpdateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserReq) ProtoMessage() {}

func (x *UpdateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_mod_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserReq.ProtoReflect.Descriptor instead.
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return file_mod_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateUserReq) GetUser() *UpdateUserBody {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateUserBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 密码
	//
	//
	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	// 新密码
	//
	NewPassword string `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	//是否跳过
	SkipPasscheck int32 `protobuf:"varint,3,opt,name=skip_passcheck,json=skipPasscheck,proto3" json:"skip_passcheck,omitempty"`
}

func (x *UpdateUserBody) Reset() {
	*x = UpdateUserBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserBody) ProtoMessage() {}

func (x *UpdateUserBody) ProtoReflect() protoreflect.Message {
	mi := &file_mod_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserBody.ProtoReflect.Descriptor instead.
func (*UpdateUserBody) Descriptor() ([]byte, []int) {
	return file_mod_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUserBody) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateUserBody) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *UpdateUserBody) GetSkipPasscheck() int32 {
	if x != nil {
		return x.SkipPasscheck
	}
	return 0
}

type UpdateUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *User `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateUserResp) Reset() {
	*x = UpdateUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResp) ProtoMessage() {}

func (x *UpdateUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_mod_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResp.ProtoReflect.Descriptor instead.
func (*UpdateUserResp) Descriptor() ([]byte, []int) {
	return file_mod_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserResp) GetData() *User {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 数据库id
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *DeleteUserReq) Reset() {
	*x = DeleteUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserReq) ProtoMessage() {}

func (x *DeleteUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_mod_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserReq.ProtoReflect.Descriptor instead.
func (*DeleteUserReq) Descriptor() ([]byte, []int) {
	return file_mod_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteUserReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户名，最少4个字符， 字母数字下划线、横杠、点、@ example: richard@qq.com
	//
	// required
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// 密码
	//
	// required
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_mod_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_mod_proto_rawDescGZIP(), []int{8}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *User `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_mod_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_mod_proto_rawDescGZIP(), []int{9}
}

func (x *LoginResp) GetData() *User {
	if x != nil {
		return x.Data
	}
	return nil
}

type LogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogoutReq) Reset() {
	*x = LogoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReq) ProtoMessage() {}

func (x *LogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_mod_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReq.ProtoReflect.Descriptor instead.
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return file_mod_proto_rawDescGZIP(), []int{10}
}

type ListUsersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUsersReq) Reset() {
	*x = ListUsersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersReq) ProtoMessage() {}

func (x *ListUsersReq) ProtoReflect() protoreflect.Message {
	mi := &file_mod_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersReq.ProtoReflect.Descriptor instead.
func (*ListUsersReq) Descriptor() ([]byte, []int) {
	return file_mod_proto_rawDescGZIP(), []int{11}
}

type ListUsersResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*User `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListUsersResp) Reset() {
	*x = ListUsersResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResp) ProtoMessage() {}

func (x *ListUsersResp) ProtoReflect() protoreflect.Message {
	mi := &file_mod_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResp.ProtoReflect.Descriptor instead.
func (*ListUsersResp) Descriptor() ([]byte, []int) {
	return file_mod_proto_rawDescGZIP(), []int{12}
}

func (x *ListUsersResp) GetData() []*User {
	if x != nil {
		return x.Data
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户名
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// 是否是管理员
	IsRoot int32 `protobuf:"varint,2,opt,name=is_root,json=isRoot,proto3" json:"is_root,omitempty"`
	// 最后登录时间
	LastLogin string `protobuf:"bytes,3,opt,name=last_login,json=lastLogin,proto3" json:"last_login,omitempty"`
	// 创建时间
	CreatedTime string `protobuf:"bytes,4,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// id
	Id int32 `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_mod_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_mod_proto_rawDescGZIP(), []int{13}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetIsRoot() int32 {
	if x != nil {
		return x.IsRoot
	}
	return 0
}

func (x *User) GetLastLogin() string {
	if x != nil {
		return x.LastLogin
	}
	return ""
}

func (x *User) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SuccessResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// success
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SuccessResp) Reset() {
	*x = SuccessResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResp) ProtoMessage() {}

func (x *SuccessResp) ProtoReflect() protoreflect.Message {
	mi := &file_mod_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResp.ProtoReflect.Descriptor instead.
func (*SuccessResp) Descriptor() ([]byte, []int) {
	return file_mod_proto_rawDescGZIP(), []int{14}
}

func (x *SuccessResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_mod_proto protoreflect.FileDescriptor

var file_mod_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2f, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xe2, 0xdf, 0x1f, 0x0e, 0x0a,
	0x0c, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x7b, 0x32, 0x2c, 0x36, 0x7d, 0x24, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x20, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x5f, 0x6d, 0x73, 0x67, 0x22, 0x94, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x46, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe2, 0xdf, 0x1f, 0x26, 0x0a, 0x24,
	0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x5b, 0x61, 0x2d, 0x7a,
	0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5c, 0x2d, 0x5f, 0x40, 0x5c, 0x2e, 0x5d, 0x7b, 0x33, 0x2c,
	0x32, 0x30, 0x7d, 0x24, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x69, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58,
	0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2f, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1d, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5c, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x27, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x7e, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x6b, 0x69,
	0x70, 0x50, 0x61, 0x73, 0x73, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x22, 0x2f, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x76, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x46, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a,
	0xe2, 0xdf, 0x1f, 0x26, 0x0a, 0x24, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d,
	0x39, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5c, 0x2d, 0x5f, 0x40,
	0x5c, 0x2e, 0x5d, 0x7b, 0x33, 0x2c, 0x32, 0x30, 0x7d, 0x24, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2a, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x0b, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x22, 0x0e, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x22, 0x2e, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x8d, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x1f, 0x0a, 0x0b, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x42, 0x05, 0x5a, 0x03, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mod_proto_rawDescOnce sync.Once
	file_mod_proto_rawDescData = file_mod_proto_rawDesc
)

func file_mod_proto_rawDescGZIP() []byte {
	file_mod_proto_rawDescOnce.Do(func() {
		file_mod_proto_rawDescData = protoimpl.X.CompressGZIP(file_mod_proto_rawDescData)
	})
	return file_mod_proto_rawDescData
}

var file_mod_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_mod_proto_goTypes = []interface{}{
	(*PingReq)(nil),        // 0: api.PingReq
	(*PingResp)(nil),       // 1: api.PingResp
	(*CreateUserReq)(nil),  // 2: api.CreateUserReq
	(*CreateUserResp)(nil), // 3: api.CreateUserResp
	(*UpdateUserReq)(nil),  // 4: api.UpdateUserReq
	(*UpdateUserBody)(nil), // 5: api.UpdateUserBody
	(*UpdateUserResp)(nil), // 6: api.UpdateUserResp
	(*DeleteUserReq)(nil),  // 7: api.DeleteUserReq
	(*LoginReq)(nil),       // 8: api.LoginReq
	(*LoginResp)(nil),      // 9: api.LoginResp
	(*LogoutReq)(nil),      // 10: api.LogoutReq
	(*ListUsersReq)(nil),   // 11: api.ListUsersReq
	(*ListUsersResp)(nil),  // 12: api.ListUsersResp
	(*User)(nil),           // 13: api.User
	(*SuccessResp)(nil),    // 14: api.SuccessResp
}
var file_mod_proto_depIdxs = []int32{
	13, // 0: api.CreateUserResp.data:type_name -> api.User
	5,  // 1: api.UpdateUserReq.user:type_name -> api.UpdateUserBody
	13, // 2: api.UpdateUserResp.data:type_name -> api.User
	13, // 3: api.LoginResp.data:type_name -> api.User
	13, // 4: api.ListUsersResp.data:type_name -> api.User
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_mod_proto_init() }
func file_mod_proto_init() {
	if File_mod_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mod_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mod_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mod_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mod_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mod_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mod_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserBody); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mod_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mod_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mod_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mod_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mod_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mod_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mod_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mod_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mod_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mod_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mod_proto_goTypes,
		DependencyIndexes: file_mod_proto_depIdxs,
		MessageInfos:      file_mod_proto_msgTypes,
	}.Build()
	File_mod_proto = out.File
	file_mod_proto_rawDesc = nil
	file_mod_proto_goTypes = nil
	file_mod_proto_depIdxs = nil
}