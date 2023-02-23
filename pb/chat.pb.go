// Specify the version of the protocol

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: chat.proto

// Specify package name to avoid name collision

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ユーザー情報を表すメッセージ型
type ChatGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid         string      `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserId       int64       `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name         string      `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Desc         string      `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	CreatedAt    string      `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string      `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IsDeleted    bool        `protobuf:"varint,8,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	ChatList     []*Chat     `protobuf:"bytes,9,rep,name=chat_list,json=chatList,proto3" json:"chat_list,omitempty"`
	ChatUserList []*ChatUser `protobuf:"bytes,10,rep,name=chat_user_list,json=chatUserList,proto3" json:"chat_user_list,omitempty"`
}

func (x *ChatGroup) Reset() {
	*x = ChatGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatGroup) ProtoMessage() {}

func (x *ChatGroup) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatGroup.ProtoReflect.Descriptor instead.
func (*ChatGroup) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatGroup) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChatGroup) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ChatGroup) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ChatGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChatGroup) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ChatGroup) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ChatGroup) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ChatGroup) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *ChatGroup) GetChatList() []*Chat {
	if x != nil {
		return x.ChatList
	}
	return nil
}

func (x *ChatGroup) GetChatUserList() []*ChatUser {
	if x != nil {
		return x.ChatUserList
	}
	return nil
}

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid      string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	GroupId   int64  `protobuf:"varint,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	From      int64  `protobuf:"varint,5,opt,name=from,proto3" json:"from,omitempty"`
	Content   string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IsDeleted bool   `protobuf:"varint,9,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Chat) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Chat) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Chat) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *Chat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Chat) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *Chat) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Chat) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Chat) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Chat) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

type ChatUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid      string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	GroupId   int64  `protobuf:"varint,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	UserId    int64  `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IsDeleted bool   `protobuf:"varint,7,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
}

func (x *ChatUser) Reset() {
	*x = ChatUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatUser) ProtoMessage() {}

func (x *ChatUser) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatUser.ProtoReflect.Descriptor instead.
func (*ChatUser) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ChatUser) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChatUser) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ChatUser) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *ChatUser) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ChatUser) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ChatUser) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ChatUser) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

// req
type ChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat *Chat `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *ChatRequest) Reset() {
	*x = ChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRequest) ProtoMessage() {}

func (x *ChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRequest.ProtoReflect.Descriptor instead.
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ChatRequest) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type ChatIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ChatIdRequest) Reset() {
	*x = ChatIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatIdRequest) ProtoMessage() {}

func (x *ChatIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatIdRequest.ProtoReflect.Descriptor instead.
func (*ChatIdRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *ChatIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// res
type ChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Chat  *Chat  `protobuf:"bytes,2,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *ChatResponse) Reset() {
	*x = ChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResponse) ProtoMessage() {}

func (x *ChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResponse.ProtoReflect.Descriptor instead.
func (*ChatResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{5}
}

func (x *ChatResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ChatResponse) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type ChatBoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Result bool   `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ChatBoolResponse) Reset() {
	*x = ChatBoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatBoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatBoolResponse) ProtoMessage() {}

func (x *ChatBoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatBoolResponse.ProtoReflect.Descriptor instead.
func (*ChatBoolResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{6}
}

func (x *ChatBoolResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ChatBoolResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type GetChatStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId int64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GetChatStreamRequest) Reset() {
	*x = GetChatStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatStreamRequest) ProtoMessage() {}

func (x *GetChatStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatStreamRequest.ProtoReflect.Descriptor instead.
func (*GetChatStreamRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{7}
}

func (x *GetChatStreamRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type ChatGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatGroup *ChatGroup `protobuf:"bytes,1,opt,name=chat_group,json=chatGroup,proto3" json:"chat_group,omitempty"`
}

func (x *ChatGroupRequest) Reset() {
	*x = ChatGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatGroupRequest) ProtoMessage() {}

func (x *ChatGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatGroupRequest.ProtoReflect.Descriptor instead.
func (*ChatGroupRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{8}
}

func (x *ChatGroupRequest) GetChatGroup() *ChatGroup {
	if x != nil {
		return x.ChatGroup
	}
	return nil
}

type ChatGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error     string     `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	ChatGroup *ChatGroup `protobuf:"bytes,2,opt,name=chat_group,json=chatGroup,proto3" json:"chat_group,omitempty"`
}

func (x *ChatGroupResponse) Reset() {
	*x = ChatGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatGroupResponse) ProtoMessage() {}

func (x *ChatGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatGroupResponse.ProtoReflect.Descriptor instead.
func (*ChatGroupResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{9}
}

func (x *ChatGroupResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ChatGroupResponse) GetChatGroup() *ChatGroup {
	if x != nil {
		return x.ChatGroup
	}
	return nil
}

type ChatGroupListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error         string       `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	ChatGroupList []*ChatGroup `protobuf:"bytes,2,rep,name=chat_group_list,json=chatGroupList,proto3" json:"chat_group_list,omitempty"`
}

func (x *ChatGroupListResponse) Reset() {
	*x = ChatGroupListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatGroupListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatGroupListResponse) ProtoMessage() {}

func (x *ChatGroupListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatGroupListResponse.ProtoReflect.Descriptor instead.
func (*ChatGroupListResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{10}
}

func (x *ChatGroupListResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ChatGroupListResponse) GetChatGroupList() []*ChatGroup {
	if x != nil {
		return x.ChatGroupList
	}
	return nil
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x22, 0xac, 0x02, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x12, 0x27, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0e, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0xe4, 0x01, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0xbf, 0x01, 0x0a, 0x08, 0x43, 0x68, 0x61,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x2d, 0x0a, 0x0b, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x63, 0x68, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x43, 0x68, 0x61,
	0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x44, 0x0a, 0x0c, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x1e, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74,
	0x22, 0x40, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x31, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x09,
	0x63, 0x68, 0x61, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x59, 0x0a, 0x11, 0x43, 0x68, 0x61,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x09, 0x63, 0x68, 0x61, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x22, 0x66, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0d, 0x63,
	0x68, 0x61, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xab, 0x02, 0x0a,
	0x10, 0x43, 0x68, 0x61, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x17, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x1a, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x13, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8b, 0x02, 0x0a, 0x0b, 0x43,
	0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x0a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x1a, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0a,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_chat_proto_goTypes = []interface{}{
	(*ChatGroup)(nil),             // 0: chat.ChatGroup
	(*Chat)(nil),                  // 1: chat.Chat
	(*ChatUser)(nil),              // 2: chat.ChatUser
	(*ChatRequest)(nil),           // 3: chat.ChatRequest
	(*ChatIdRequest)(nil),         // 4: chat.ChatIdRequest
	(*ChatResponse)(nil),          // 5: chat.ChatResponse
	(*ChatBoolResponse)(nil),      // 6: chat.ChatBoolResponse
	(*GetChatStreamRequest)(nil),  // 7: chat.GetChatStreamRequest
	(*ChatGroupRequest)(nil),      // 8: chat.ChatGroupRequest
	(*ChatGroupResponse)(nil),     // 9: chat.ChatGroupResponse
	(*ChatGroupListResponse)(nil), // 10: chat.ChatGroupListResponse
}
var file_chat_proto_depIdxs = []int32{
	1,  // 0: chat.ChatGroup.chat_list:type_name -> chat.Chat
	2,  // 1: chat.ChatGroup.chat_user_list:type_name -> chat.ChatUser
	1,  // 2: chat.ChatRequest.chat:type_name -> chat.Chat
	1,  // 3: chat.ChatResponse.chat:type_name -> chat.Chat
	0,  // 4: chat.ChatGroupRequest.chat_group:type_name -> chat.ChatGroup
	0,  // 5: chat.ChatGroupResponse.chat_group:type_name -> chat.ChatGroup
	0,  // 6: chat.ChatGroupListResponse.chat_group_list:type_name -> chat.ChatGroup
	0,  // 7: chat.ChatGroupService.Create:input_type -> chat.ChatGroup
	0,  // 8: chat.ChatGroupService.Update:input_type -> chat.ChatGroup
	4,  // 9: chat.ChatGroupService.Delete:input_type -> chat.ChatIdRequest
	4,  // 10: chat.ChatGroupService.GetById:input_type -> chat.ChatIdRequest
	4,  // 11: chat.ChatGroupService.GetByUserId:input_type -> chat.ChatIdRequest
	1,  // 12: chat.ChatService.Create:input_type -> chat.Chat
	1,  // 13: chat.ChatService.Update:input_type -> chat.Chat
	4,  // 14: chat.ChatService.Delete:input_type -> chat.ChatIdRequest
	4,  // 15: chat.ChatService.GetById:input_type -> chat.ChatIdRequest
	7,  // 16: chat.ChatService.GetStream:input_type -> chat.GetChatStreamRequest
	9,  // 17: chat.ChatGroupService.Create:output_type -> chat.ChatGroupResponse
	9,  // 18: chat.ChatGroupService.Update:output_type -> chat.ChatGroupResponse
	6,  // 19: chat.ChatGroupService.Delete:output_type -> chat.ChatBoolResponse
	9,  // 20: chat.ChatGroupService.GetById:output_type -> chat.ChatGroupResponse
	10, // 21: chat.ChatGroupService.GetByUserId:output_type -> chat.ChatGroupListResponse
	5,  // 22: chat.ChatService.Create:output_type -> chat.ChatResponse
	5,  // 23: chat.ChatService.Update:output_type -> chat.ChatResponse
	6,  // 24: chat.ChatService.Delete:output_type -> chat.ChatBoolResponse
	5,  // 25: chat.ChatService.GetById:output_type -> chat.ChatResponse
	5,  // 26: chat.ChatService.GetStream:output_type -> chat.ChatResponse
	17, // [17:27] is the sub-list for method output_type
	7,  // [7:17] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatGroup); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chat); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatUser); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRequest); i {
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
		file_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatIdRequest); i {
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
		file_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatResponse); i {
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
		file_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatBoolResponse); i {
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
		file_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatStreamRequest); i {
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
		file_chat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatGroupRequest); i {
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
		file_chat_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatGroupResponse); i {
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
		file_chat_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatGroupListResponse); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
