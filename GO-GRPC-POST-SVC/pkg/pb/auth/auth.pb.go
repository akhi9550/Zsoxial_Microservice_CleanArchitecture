// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: pkg/pb/auth/auth.proto

package auth

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

type CheckUserAvalilabilityWithUserIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CheckUserAvalilabilityWithUserIDRequest) Reset() {
	*x = CheckUserAvalilabilityWithUserIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserAvalilabilityWithUserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserAvalilabilityWithUserIDRequest) ProtoMessage() {}

func (x *CheckUserAvalilabilityWithUserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserAvalilabilityWithUserIDRequest.ProtoReflect.Descriptor instead.
func (*CheckUserAvalilabilityWithUserIDRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *CheckUserAvalilabilityWithUserIDRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CheckUserAvalilabilityWithUserIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *CheckUserAvalilabilityWithUserIDResponse) Reset() {
	*x = CheckUserAvalilabilityWithUserIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserAvalilabilityWithUserIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserAvalilabilityWithUserIDResponse) ProtoMessage() {}

func (x *CheckUserAvalilabilityWithUserIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserAvalilabilityWithUserIDResponse.ProtoReflect.Descriptor instead.
func (*CheckUserAvalilabilityWithUserIDResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *CheckUserAvalilabilityWithUserIDResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type UserDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserDataRequest) Reset() {
	*x = UserDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDataRequest) ProtoMessage() {}

func (x *UserDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDataRequest.ProtoReflect.Descriptor instead.
func (*UserDataRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *UserDataRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username     string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	ProfilePhoto string `protobuf:"bytes,3,opt,name=profile_photo,json=profilePhoto,proto3" json:"profile_photo,omitempty"`
}

func (x *UserDataResponse) Reset() {
	*x = UserDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDataResponse) ProtoMessage() {}

func (x *UserDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDataResponse.ProtoReflect.Descriptor instead.
func (*UserDataResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *UserDataResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserDataResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserDataResponse) GetProfilePhoto() string {
	if x != nil {
		return x.ProfilePhoto
	}
	return ""
}

type CheckUserAvalilabilityWithTagUserIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag *Tag `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *CheckUserAvalilabilityWithTagUserIDRequest) Reset() {
	*x = CheckUserAvalilabilityWithTagUserIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserAvalilabilityWithTagUserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserAvalilabilityWithTagUserIDRequest) ProtoMessage() {}

func (x *CheckUserAvalilabilityWithTagUserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserAvalilabilityWithTagUserIDRequest.ProtoReflect.Descriptor instead.
func (*CheckUserAvalilabilityWithTagUserIDRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *CheckUserAvalilabilityWithTagUserIDRequest) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User []string `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_auth_proto_rawDescGZIP(), []int{5}
}

func (x *Tag) GetUser() []string {
	if x != nil {
		return x.User
	}
	return nil
}

type CheckUserAvalilabilityWithTagUserIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *CheckUserAvalilabilityWithTagUserIDResponse) Reset() {
	*x = CheckUserAvalilabilityWithTagUserIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserAvalilabilityWithTagUserIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserAvalilabilityWithTagUserIDResponse) ProtoMessage() {}

func (x *CheckUserAvalilabilityWithTagUserIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserAvalilabilityWithTagUserIDResponse.ProtoReflect.Descriptor instead.
func (*CheckUserAvalilabilityWithTagUserIDResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_auth_proto_rawDescGZIP(), []int{6}
}

func (x *CheckUserAvalilabilityWithTagUserIDResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type GetUserNameWithTagUserIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag *Tag `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *GetUserNameWithTagUserIDRequest) Reset() {
	*x = GetUserNameWithTagUserIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserNameWithTagUserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserNameWithTagUserIDRequest) ProtoMessage() {}

func (x *GetUserNameWithTagUserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserNameWithTagUserIDRequest.ProtoReflect.Descriptor instead.
func (*GetUserNameWithTagUserIDRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_auth_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserNameWithTagUserIDRequest) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

type TagUsernames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *TagUsernames) Reset() {
	*x = TagUsernames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagUsernames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagUsernames) ProtoMessage() {}

func (x *TagUsernames) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagUsernames.ProtoReflect.Descriptor instead.
func (*TagUsernames) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_auth_proto_rawDescGZIP(), []int{8}
}

func (x *TagUsernames) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetUserNameWithTagUserIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name []*TagUsernames `protobuf:"bytes,1,rep,name=name,proto3" json:"name,omitempty"`
}

func (x *GetUserNameWithTagUserIDResponse) Reset() {
	*x = GetUserNameWithTagUserIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserNameWithTagUserIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserNameWithTagUserIDResponse) ProtoMessage() {}

func (x *GetUserNameWithTagUserIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserNameWithTagUserIDResponse.ProtoReflect.Descriptor instead.
func (*GetUserNameWithTagUserIDResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_auth_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserNameWithTagUserIDResponse) GetName() []*TagUsernames {
	if x != nil {
		return x.Name
	}
	return nil
}

type GetFollowingUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetFollowingUsersRequest) Reset() {
	*x = GetFollowingUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_auth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowingUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowingUsersRequest) ProtoMessage() {}

func (x *GetFollowingUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_auth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowingUsersRequest.ProtoReflect.Descriptor instead.
func (*GetFollowingUsersRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_auth_proto_rawDescGZIP(), []int{10}
}

func (x *GetFollowingUsersRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type Followuser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followinguser int64 `protobuf:"varint,1,opt,name=followinguser,proto3" json:"followinguser,omitempty"`
}

func (x *Followuser) Reset() {
	*x = Followuser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_auth_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Followuser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Followuser) ProtoMessage() {}

func (x *Followuser) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_auth_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Followuser.ProtoReflect.Descriptor instead.
func (*Followuser) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_auth_proto_rawDescGZIP(), []int{11}
}

func (x *Followuser) GetFollowinguser() int64 {
	if x != nil {
		return x.Followinguser
	}
	return 0
}

type GetFollowingUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User []*Followuser `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
}

func (x *GetFollowingUsersResponse) Reset() {
	*x = GetFollowingUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_auth_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowingUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowingUsersResponse) ProtoMessage() {}

func (x *GetFollowingUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_auth_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowingUsersResponse.ProtoReflect.Descriptor instead.
func (*GetFollowingUsersResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_auth_proto_rawDescGZIP(), []int{12}
}

func (x *GetFollowingUsersResponse) GetUser() []*Followuser {
	if x != nil {
		return x.User
	}
	return nil
}

var File_pkg_pb_auth_auth_proto protoreflect.FileDescriptor

var file_pkg_pb_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x39,
	0x0a, 0x27, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x6c, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x28, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x63,
	0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x68,
	0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x2a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x76, 0x61, 0x6c, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x57, 0x69, 0x74,
	0x68, 0x54, 0x61, 0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x19,
	0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x43, 0x0a, 0x2b, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x57, 0x69, 0x74, 0x68, 0x54, 0x61, 0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x3e,
	0x0a, 0x1f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x54, 0x61, 0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x2a,
	0x0a, 0x0c, 0x54, 0x61, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x20, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x57, 0x69, 0x74, 0x68, 0x54, 0x61, 0x67,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x32, 0x0a, 0x0a, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x75, 0x73, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x75, 0x73, 0x65, 0x72, 0x22, 0x41,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x75, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x32, 0xa4, 0x04, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x83, 0x01, 0x0a, 0x20, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x76, 0x61, 0x6c, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x57, 0x69, 0x74, 0x68,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x8c, 0x01, 0x0a, 0x23, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x57,
	0x69, 0x74, 0x68, 0x54, 0x61, 0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x30, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61,
	0x6c, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x57, 0x69, 0x74, 0x68, 0x54, 0x61,
	0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x76, 0x61, 0x6c, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x57, 0x69, 0x74, 0x68,
	0x54, 0x61, 0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x57, 0x69, 0x74, 0x68, 0x54, 0x61, 0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x25, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x57, 0x69, 0x74, 0x68, 0x54, 0x61, 0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x57, 0x69, 0x74, 0x68, 0x54, 0x61, 0x67,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x56, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_pb_auth_auth_proto_rawDescOnce sync.Once
	file_pkg_pb_auth_auth_proto_rawDescData = file_pkg_pb_auth_auth_proto_rawDesc
)

func file_pkg_pb_auth_auth_proto_rawDescGZIP() []byte {
	file_pkg_pb_auth_auth_proto_rawDescOnce.Do(func() {
		file_pkg_pb_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_auth_auth_proto_rawDescData)
	})
	return file_pkg_pb_auth_auth_proto_rawDescData
}

var file_pkg_pb_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_pkg_pb_auth_auth_proto_goTypes = []interface{}{
	(*CheckUserAvalilabilityWithUserIDRequest)(nil),    // 0: user.CheckUserAvalilabilityWithUserIDRequest
	(*CheckUserAvalilabilityWithUserIDResponse)(nil),   // 1: user.CheckUserAvalilabilityWithUserIDResponse
	(*UserDataRequest)(nil),                            // 2: user.UserDataRequest
	(*UserDataResponse)(nil),                           // 3: user.UserDataResponse
	(*CheckUserAvalilabilityWithTagUserIDRequest)(nil), // 4: user.CheckUserAvalilabilityWithTagUserIDRequest
	(*Tag)(nil), // 5: user.Tag
	(*CheckUserAvalilabilityWithTagUserIDResponse)(nil), // 6: user.CheckUserAvalilabilityWithTagUserIDResponse
	(*GetUserNameWithTagUserIDRequest)(nil),             // 7: user.GetUserNameWithTagUserIDRequest
	(*TagUsernames)(nil),                                // 8: user.TagUsernames
	(*GetUserNameWithTagUserIDResponse)(nil),            // 9: user.GetUserNameWithTagUserIDResponse
	(*GetFollowingUsersRequest)(nil),                    // 10: user.GetFollowingUsersRequest
	(*Followuser)(nil),                                  // 11: user.Followuser
	(*GetFollowingUsersResponse)(nil),                   // 12: user.GetFollowingUsersResponse
}
var file_pkg_pb_auth_auth_proto_depIdxs = []int32{
	5,  // 0: user.CheckUserAvalilabilityWithTagUserIDRequest.tag:type_name -> user.Tag
	5,  // 1: user.GetUserNameWithTagUserIDRequest.tag:type_name -> user.Tag
	8,  // 2: user.GetUserNameWithTagUserIDResponse.name:type_name -> user.TagUsernames
	11, // 3: user.GetFollowingUsersResponse.user:type_name -> user.Followuser
	0,  // 4: user.AuthService.CheckUserAvalilabilityWithUserID:input_type -> user.CheckUserAvalilabilityWithUserIDRequest
	2,  // 5: user.AuthService.UserData:input_type -> user.UserDataRequest
	4,  // 6: user.AuthService.CheckUserAvalilabilityWithTagUserID:input_type -> user.CheckUserAvalilabilityWithTagUserIDRequest
	7,  // 7: user.AuthService.GetUserNameWithTagUserID:input_type -> user.GetUserNameWithTagUserIDRequest
	10, // 8: user.AuthService.GetFollowingUsers:input_type -> user.GetFollowingUsersRequest
	1,  // 9: user.AuthService.CheckUserAvalilabilityWithUserID:output_type -> user.CheckUserAvalilabilityWithUserIDResponse
	3,  // 10: user.AuthService.UserData:output_type -> user.UserDataResponse
	6,  // 11: user.AuthService.CheckUserAvalilabilityWithTagUserID:output_type -> user.CheckUserAvalilabilityWithTagUserIDResponse
	9,  // 12: user.AuthService.GetUserNameWithTagUserID:output_type -> user.GetUserNameWithTagUserIDResponse
	12, // 13: user.AuthService.GetFollowingUsers:output_type -> user.GetFollowingUsersResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_pb_auth_auth_proto_init() }
func file_pkg_pb_auth_auth_proto_init() {
	if File_pkg_pb_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserAvalilabilityWithUserIDRequest); i {
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
		file_pkg_pb_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserAvalilabilityWithUserIDResponse); i {
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
		file_pkg_pb_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDataRequest); i {
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
		file_pkg_pb_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDataResponse); i {
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
		file_pkg_pb_auth_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserAvalilabilityWithTagUserIDRequest); i {
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
		file_pkg_pb_auth_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_pkg_pb_auth_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserAvalilabilityWithTagUserIDResponse); i {
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
		file_pkg_pb_auth_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserNameWithTagUserIDRequest); i {
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
		file_pkg_pb_auth_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagUsernames); i {
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
		file_pkg_pb_auth_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserNameWithTagUserIDResponse); i {
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
		file_pkg_pb_auth_auth_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowingUsersRequest); i {
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
		file_pkg_pb_auth_auth_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Followuser); i {
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
		file_pkg_pb_auth_auth_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowingUsersResponse); i {
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
			RawDescriptor: file_pkg_pb_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_auth_auth_proto_goTypes,
		DependencyIndexes: file_pkg_pb_auth_auth_proto_depIdxs,
		MessageInfos:      file_pkg_pb_auth_auth_proto_msgTypes,
	}.Build()
	File_pkg_pb_auth_auth_proto = out.File
	file_pkg_pb_auth_auth_proto_rawDesc = nil
	file_pkg_pb_auth_auth_proto_goTypes = nil
	file_pkg_pb_auth_auth_proto_depIdxs = nil
}
