// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: pkg/pb/post/post.proto

package post

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

type ShowPostReportsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Count int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ShowPostReportsRequest) Reset() {
	*x = ShowPostReportsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_post_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowPostReportsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowPostReportsRequest) ProtoMessage() {}

func (x *ShowPostReportsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_post_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowPostReportsRequest.ProtoReflect.Descriptor instead.
func (*ShowPostReportsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_post_post_proto_rawDescGZIP(), []int{0}
}

func (x *ShowPostReportsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ShowPostReportsRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ReportPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepostedUserid int64  `protobuf:"varint,1,opt,name=RepostedUserid,proto3" json:"RepostedUserid,omitempty"`
	Postid         int64  `protobuf:"varint,2,opt,name=postid,proto3" json:"postid,omitempty"`
	Report         string `protobuf:"bytes,3,opt,name=Report,proto3" json:"Report,omitempty"`
}

func (x *ReportPostResponse) Reset() {
	*x = ReportPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_post_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportPostResponse) ProtoMessage() {}

func (x *ReportPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_post_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportPostResponse.ProtoReflect.Descriptor instead.
func (*ReportPostResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_post_post_proto_rawDescGZIP(), []int{1}
}

func (x *ReportPostResponse) GetRepostedUserid() int64 {
	if x != nil {
		return x.RepostedUserid
	}
	return 0
}

func (x *ReportPostResponse) GetPostid() int64 {
	if x != nil {
		return x.Postid
	}
	return 0
}

func (x *ReportPostResponse) GetReport() string {
	if x != nil {
		return x.Report
	}
	return ""
}

type ShowPostReportsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reports []*ReportPostResponse `protobuf:"bytes,1,rep,name=reports,proto3" json:"reports,omitempty"`
}

func (x *ShowPostReportsResponse) Reset() {
	*x = ShowPostReportsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_post_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowPostReportsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowPostReportsResponse) ProtoMessage() {}

func (x *ShowPostReportsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_post_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowPostReportsResponse.ProtoReflect.Descriptor instead.
func (*ShowPostReportsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_post_post_proto_rawDescGZIP(), []int{2}
}

func (x *ShowPostReportsResponse) GetReports() []*ReportPostResponse {
	if x != nil {
		return x.Reports
	}
	return nil
}

type GetAllpostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Count int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllpostsRequest) Reset() {
	*x = GetAllpostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_post_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllpostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllpostsRequest) ProtoMessage() {}

func (x *GetAllpostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_post_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllpostsRequest.ProtoReflect.Descriptor instead.
func (*GetAllpostsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_post_post_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllpostsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllpostsRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid   int64  `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Imageurl string `protobuf:"bytes,3,opt,name=imageurl,proto3" json:"imageurl,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_post_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_post_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_pkg_pb_post_post_proto_rawDescGZIP(), []int{4}
}

func (x *UserData) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *UserData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserData) GetImageurl() string {
	if x != nil {
		return x.Imageurl
	}
	return ""
}

type PostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User      *UserData              `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Caption   string                 `protobuf:"bytes,3,opt,name=caption,proto3" json:"caption,omitempty"`
	Tag       []string               `protobuf:"bytes,4,rep,name=tag,proto3" json:"tag,omitempty"`
	Url       string                 `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Like      int64                  `protobuf:"varint,6,opt,name=like,proto3" json:"like,omitempty"`
	Comment   int64                  `protobuf:"varint,7,opt,name=comment,proto3" json:"comment,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PostResponse) Reset() {
	*x = PostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_post_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostResponse) ProtoMessage() {}

func (x *PostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_post_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostResponse.ProtoReflect.Descriptor instead.
func (*PostResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_post_post_proto_rawDescGZIP(), []int{5}
}

func (x *PostResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PostResponse) GetUser() *UserData {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PostResponse) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *PostResponse) GetTag() []string {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *PostResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PostResponse) GetLike() int64 {
	if x != nil {
		return x.Like
	}
	return 0
}

func (x *PostResponse) GetComment() int64 {
	if x != nil {
		return x.Comment
	}
	return 0
}

func (x *PostResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GetAllpostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*PostResponse `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetAllpostsResponse) Reset() {
	*x = GetAllpostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_post_post_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllpostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllpostsResponse) ProtoMessage() {}

func (x *GetAllpostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_post_post_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllpostsResponse.ProtoReflect.Descriptor instead.
func (*GetAllpostsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_post_post_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllpostsResponse) GetPosts() []*PostResponse {
	if x != nil {
		return x.Posts
	}
	return nil
}

type CheckPostIDByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostID int64 `protobuf:"varint,1,opt,name=postID,proto3" json:"postID,omitempty"`
}

func (x *CheckPostIDByIDRequest) Reset() {
	*x = CheckPostIDByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_post_post_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPostIDByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPostIDByIDRequest) ProtoMessage() {}

func (x *CheckPostIDByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_post_post_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPostIDByIDRequest.ProtoReflect.Descriptor instead.
func (*CheckPostIDByIDRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_post_post_proto_rawDescGZIP(), []int{7}
}

func (x *CheckPostIDByIDRequest) GetPostID() int64 {
	if x != nil {
		return x.PostID
	}
	return 0
}

type CheckPostIDByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist bool `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
}

func (x *CheckPostIDByIDResponse) Reset() {
	*x = CheckPostIDByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_post_post_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPostIDByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPostIDByIDResponse) ProtoMessage() {}

func (x *CheckPostIDByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_post_post_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPostIDByIDResponse.ProtoReflect.Descriptor instead.
func (*CheckPostIDByIDResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_post_post_proto_rawDescGZIP(), []int{8}
}

func (x *CheckPostIDByIDResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

type RemovePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostID int64 `protobuf:"varint,1,opt,name=postID,proto3" json:"postID,omitempty"`
}

func (x *RemovePostRequest) Reset() {
	*x = RemovePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_post_post_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePostRequest) ProtoMessage() {}

func (x *RemovePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_post_post_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePostRequest.ProtoReflect.Descriptor instead.
func (*RemovePostRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_post_post_proto_rawDescGZIP(), []int{9}
}

func (x *RemovePostRequest) GetPostID() int64 {
	if x != nil {
		return x.PostID
	}
	return 0
}

type RemovePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemovePostResponse) Reset() {
	*x = RemovePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_post_post_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePostResponse) ProtoMessage() {}

func (x *RemovePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_post_post_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePostResponse.ProtoReflect.Descriptor instead.
func (*RemovePostResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_post_post_proto_rawDescGZIP(), []int{10}
}

var File_pkg_pb_post_post_proto protoreflect.FileDescriptor

var file_pkg_pb_post_post_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x42, 0x0a, 0x16, 0x53, 0x68, 0x6f, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x6c, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x22, 0x4d, 0x0a, 0x17, 0x53, 0x68, 0x6f, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x22, 0x3e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x5a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x75, 0x72, 0x6c, 0x22, 0xe9, 0x01, 0x0a,
	0x0c, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x61, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c,
	0x69, 0x6b, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x30, 0x0a, 0x16, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x22, 0x2f, 0x0a, 0x17, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x11,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x22, 0x14, 0x0a, 0x12, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xba, 0x02, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x50, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x73,
	0x12, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d,
	0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_post_post_proto_rawDescOnce sync.Once
	file_pkg_pb_post_post_proto_rawDescData = file_pkg_pb_post_post_proto_rawDesc
)

func file_pkg_pb_post_post_proto_rawDescGZIP() []byte {
	file_pkg_pb_post_post_proto_rawDescOnce.Do(func() {
		file_pkg_pb_post_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_post_post_proto_rawDescData)
	})
	return file_pkg_pb_post_post_proto_rawDescData
}

var file_pkg_pb_post_post_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pkg_pb_post_post_proto_goTypes = []interface{}{
	(*ShowPostReportsRequest)(nil),  // 0: post.ShowPostReportsRequest
	(*ReportPostResponse)(nil),      // 1: post.ReportPostResponse
	(*ShowPostReportsResponse)(nil), // 2: post.ShowPostReportsResponse
	(*GetAllpostsRequest)(nil),      // 3: post.GetAllpostsRequest
	(*UserData)(nil),                // 4: post.UserData
	(*PostResponse)(nil),            // 5: post.PostResponse
	(*GetAllpostsResponse)(nil),     // 6: post.GetAllpostsResponse
	(*CheckPostIDByIDRequest)(nil),  // 7: post.CheckPostIDByIDRequest
	(*CheckPostIDByIDResponse)(nil), // 8: post.CheckPostIDByIDResponse
	(*RemovePostRequest)(nil),       // 9: post.RemovePostRequest
	(*RemovePostResponse)(nil),      // 10: post.RemovePostResponse
	(*timestamppb.Timestamp)(nil),   // 11: google.protobuf.Timestamp
}
var file_pkg_pb_post_post_proto_depIdxs = []int32{
	1,  // 0: post.ShowPostReportsResponse.reports:type_name -> post.ReportPostResponse
	4,  // 1: post.PostResponse.user:type_name -> post.UserData
	11, // 2: post.PostResponse.created_at:type_name -> google.protobuf.Timestamp
	5,  // 3: post.GetAllpostsResponse.posts:type_name -> post.PostResponse
	0,  // 4: post.PostService.ShowPostReports:input_type -> post.ShowPostReportsRequest
	3,  // 5: post.PostService.GetAllPosts:input_type -> post.GetAllpostsRequest
	7,  // 6: post.PostService.CheckPostIDByID:input_type -> post.CheckPostIDByIDRequest
	9,  // 7: post.PostService.RemovePost:input_type -> post.RemovePostRequest
	2,  // 8: post.PostService.ShowPostReports:output_type -> post.ShowPostReportsResponse
	6,  // 9: post.PostService.GetAllPosts:output_type -> post.GetAllpostsResponse
	8,  // 10: post.PostService.CheckPostIDByID:output_type -> post.CheckPostIDByIDResponse
	10, // 11: post.PostService.RemovePost:output_type -> post.RemovePostResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_pb_post_post_proto_init() }
func file_pkg_pb_post_post_proto_init() {
	if File_pkg_pb_post_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_post_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowPostReportsRequest); i {
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
		file_pkg_pb_post_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportPostResponse); i {
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
		file_pkg_pb_post_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowPostReportsResponse); i {
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
		file_pkg_pb_post_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllpostsRequest); i {
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
		file_pkg_pb_post_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_pkg_pb_post_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostResponse); i {
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
		file_pkg_pb_post_post_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllpostsResponse); i {
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
		file_pkg_pb_post_post_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPostIDByIDRequest); i {
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
		file_pkg_pb_post_post_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPostIDByIDResponse); i {
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
		file_pkg_pb_post_post_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePostRequest); i {
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
		file_pkg_pb_post_post_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePostResponse); i {
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
			RawDescriptor: file_pkg_pb_post_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_post_post_proto_goTypes,
		DependencyIndexes: file_pkg_pb_post_post_proto_depIdxs,
		MessageInfos:      file_pkg_pb_post_post_proto_msgTypes,
	}.Build()
	File_pkg_pb_post_post_proto = out.File
	file_pkg_pb_post_post_proto_rawDesc = nil
	file_pkg_pb_post_post_proto_goTypes = nil
	file_pkg_pb_post_post_proto_depIdxs = nil
}
