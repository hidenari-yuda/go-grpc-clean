// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: user.proto

// 名前衝突を避けるためにパッケージ名を指定

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User_UserType int32

const (
	User_NONE          User_UserType = 0
	User_NORMAL        User_UserType = 1
	User_ADMINISTRATOR User_UserType = 2
	User_GUEST         User_UserType = 3
	User_DISABLED      User_UserType = 4
)

// Enum value maps for User_UserType.
var (
	User_UserType_name = map[int32]string{
		0: "NONE",
		1: "NORMAL",
		2: "ADMINISTRATOR",
		3: "GUEST",
		4: "DISABLED",
	}
	User_UserType_value = map[string]int32{
		"NONE":          0,
		"NORMAL":        1,
		"ADMINISTRATOR": 2,
		"GUEST":         3,
		"DISABLED":      4,
	}
)

func (x User_UserType) Enum() *User_UserType {
	p := new(User_UserType)
	*p = x
	return p
}

func (x User_UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[0].Descriptor()
}

func (User_UserType) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[0]
}

func (x User_UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_UserType.Descriptor instead.
func (User_UserType) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0, 0}
}

type UserTypeRequest_UserType int32

const (
	UserTypeRequest_NONE          UserTypeRequest_UserType = 0
	UserTypeRequest_NORMAL        UserTypeRequest_UserType = 1
	UserTypeRequest_ADMINISTRATOR UserTypeRequest_UserType = 2
	UserTypeRequest_GUEST         UserTypeRequest_UserType = 3
	UserTypeRequest_DISABLED      UserTypeRequest_UserType = 4
)

// Enum value maps for UserTypeRequest_UserType.
var (
	UserTypeRequest_UserType_name = map[int32]string{
		0: "NONE",
		1: "NORMAL",
		2: "ADMINISTRATOR",
		3: "GUEST",
		4: "DISABLED",
	}
	UserTypeRequest_UserType_value = map[string]int32{
		"NONE":          0,
		"NORMAL":        1,
		"ADMINISTRATOR": 2,
		"GUEST":         3,
		"DISABLED":      4,
	}
)

func (x UserTypeRequest_UserType) Enum() *UserTypeRequest_UserType {
	p := new(UserTypeRequest_UserType)
	*p = x
	return p
}

func (x UserTypeRequest_UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserTypeRequest_UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[1].Descriptor()
}

func (UserTypeRequest_UserType) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[1]
}

func (x UserTypeRequest_UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserTypeRequest_UserType.Descriptor instead.
func (UserTypeRequest_UserType) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2, 0}
}

// ユーザー情報を表すメッセージ型
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid       string                 `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	FirebaseId string                 `protobuf:"bytes,3,opt,name=firebase_id,json=firebaseId,proto3" json:"firebase_id,omitempty"`
	Name       string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Email      string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Password   string                 `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	UserType   User_UserType          `protobuf:"varint,7,opt,name=user_type,json=userType,proto3,enum=user.User_UserType" json:"user_type,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
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
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *User) GetFirebaseId() string {
	if x != nil {
		return x.FirebaseId
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetUserType() User_UserType {
	if x != nil {
		return x.UserType
	}
	return User_NONE
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// ユーザーの情報リクエストに使用するメッセージ型
type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// ユーザーの情報リクエストに使用するメッセージ型
type UserTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserType UserTypeRequest_UserType `protobuf:"varint,1,opt,name=user_type,json=userType,proto3,enum=user.UserTypeRequest_UserType" json:"user_type,omitempty"`
}

func (x *UserTypeRequest) Reset() {
	*x = UserTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTypeRequest) ProtoMessage() {}

func (x *UserTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTypeRequest.ProtoReflect.Descriptor instead.
func (*UserTypeRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserTypeRequest) GetUserType() UserTypeRequest_UserType {
	if x != nil {
		return x.UserType
	}
	return UserTypeRequest_NONE
}

// ユーザー情報を返す際に使用するメッセージ型
type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	User    *User  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *UserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// ユーザー人数を返す際に使用するメッセージ型
type UserCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	UserCnt uint32 `protobuf:"varint,2,opt,name=user_cnt,json=userCnt,proto3" json:"user_cnt,omitempty"`
}

func (x *UserCountResponse) Reset() {
	*x = UserCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCountResponse) ProtoMessage() {}

func (x *UserCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCountResponse.ProtoReflect.Descriptor instead.
func (*UserCountResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserCountResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *UserCountResponse) GetUserCnt() uint32 {
	if x != nil {
		return x.UserCnt
	}
	return 0
}

// The request message containing the user's name.
type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *SignInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x87, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x30, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x4c, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10,
	0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54,
	0x4f, 0x52, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x55, 0x45, 0x53, 0x54, 0x10, 0x03, 0x12,
	0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x22, 0x1d, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9c, 0x01, 0x0a,
	0x0f, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4c, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x4f, 0x52,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x55, 0x45, 0x53, 0x54, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x22, 0x5e, 0x0a, 0x0c, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x44, 0x0a, 0x11, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6e,
	0x74, 0x22, 0x41, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x32, 0x9b, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x12,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x12,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x33, 0x0a,
	0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d,
	0x67, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_proto_goTypes = []interface{}{
	(User_UserType)(0),            // 0: user.User.UserType
	(UserTypeRequest_UserType)(0), // 1: user.UserTypeRequest.UserType
	(*User)(nil),                  // 2: user.User
	(*UserRequest)(nil),           // 3: user.UserRequest
	(*UserTypeRequest)(nil),       // 4: user.UserTypeRequest
	(*UserResponse)(nil),          // 5: user.UserResponse
	(*UserCountResponse)(nil),     // 6: user.UserCountResponse
	(*SignInRequest)(nil),         // 7: user.SignInRequest
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_user_proto_depIdxs = []int32{
	0,  // 0: user.User.user_type:type_name -> user.User.UserType
	8,  // 1: user.User.created_at:type_name -> google.protobuf.Timestamp
	8,  // 2: user.User.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 3: user.UserTypeRequest.user_type:type_name -> user.UserTypeRequest.UserType
	2,  // 4: user.UserResponse.user:type_name -> user.User
	2,  // 5: user.UserService.CreateUser:input_type -> user.User
	2,  // 6: user.UserService.UpdateUser:input_type -> user.User
	3,  // 7: user.UserService.GetUserById:input_type -> user.UserRequest
	4,  // 8: user.UserService.GetUsersByType:input_type -> user.UserTypeRequest
	7,  // 9: user.UserService.SignIn:input_type -> user.SignInRequest
	5,  // 10: user.UserService.CreateUser:output_type -> user.UserResponse
	5,  // 11: user.UserService.UpdateUser:output_type -> user.UserResponse
	5,  // 12: user.UserService.GetUserById:output_type -> user.UserResponse
	5,  // 13: user.UserService.GetUsersByType:output_type -> user.UserResponse
	5,  // 14: user.UserService.SignIn:output_type -> user.UserResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTypeRequest); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCountResponse); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRequest); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		EnumInfos:         file_user_proto_enumTypes,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
