// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: favorite.proto

package favorite

import (
	feed "GuTikTok/src/rpc/feed"
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

type FavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId    uint32 `protobuf:"varint,1,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`          // 用户id
	VideoId    uint32 `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`          // 视频id
	ActionType uint32 `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"` // 1-点赞，2-取消点赞
}

func (x *FavoriteRequest) Reset() {
	*x = FavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteRequest) ProtoMessage() {}

func (x *FavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteRequest.ProtoReflect.Descriptor instead.
func (*FavoriteRequest) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{0}
}

func (x *FavoriteRequest) GetActorId() uint32 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *FavoriteRequest) GetVideoId() uint32 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *FavoriteRequest) GetActionType() uint32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type FavoriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
}

func (x *FavoriteResponse) Reset() {
	*x = FavoriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteResponse) ProtoMessage() {}

func (x *FavoriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteResponse.ProtoReflect.Descriptor instead.
func (*FavoriteResponse) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{1}
}

func (x *FavoriteResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FavoriteResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type FavoriteListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId uint32 `protobuf:"varint,1,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"` // 发出请求的用户的id
	UserId  uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`    // 用户id
}

func (x *FavoriteListRequest) Reset() {
	*x = FavoriteListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteListRequest) ProtoMessage() {}

func (x *FavoriteListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteListRequest.ProtoReflect.Descriptor instead.
func (*FavoriteListRequest) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{2}
}

func (x *FavoriteListRequest) GetActorId() uint32 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *FavoriteListRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FavoriteListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32         `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string        `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	VideoList  []*feed.Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`     // 用户点赞视频列表
}

func (x *FavoriteListResponse) Reset() {
	*x = FavoriteListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteListResponse) ProtoMessage() {}

func (x *FavoriteListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteListResponse.ProtoReflect.Descriptor instead.
func (*FavoriteListResponse) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{3}
}

func (x *FavoriteListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FavoriteListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *FavoriteListResponse) GetVideoList() []*feed.Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type IsFavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId uint32 `protobuf:"varint,1,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"` // 用户id
	VideoId uint32 `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"` // 视频id
}

func (x *IsFavoriteRequest) Reset() {
	*x = IsFavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsFavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsFavoriteRequest) ProtoMessage() {}

func (x *IsFavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsFavoriteRequest.ProtoReflect.Descriptor instead.
func (*IsFavoriteRequest) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{4}
}

func (x *IsFavoriteRequest) GetActorId() uint32 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *IsFavoriteRequest) GetVideoId() uint32 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type IsFavoriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	Result     bool   `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`                           // 结果
}

func (x *IsFavoriteResponse) Reset() {
	*x = IsFavoriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsFavoriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsFavoriteResponse) ProtoMessage() {}

func (x *IsFavoriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsFavoriteResponse.ProtoReflect.Descriptor instead.
func (*IsFavoriteResponse) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{5}
}

func (x *IsFavoriteResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *IsFavoriteResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *IsFavoriteResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type CountFavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId uint32 `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"` // 视频id
}

func (x *CountFavoriteRequest) Reset() {
	*x = CountFavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountFavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountFavoriteRequest) ProtoMessage() {}

func (x *CountFavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountFavoriteRequest.ProtoReflect.Descriptor instead.
func (*CountFavoriteRequest) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{6}
}

func (x *CountFavoriteRequest) GetVideoId() uint32 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type CountFavoriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	Count      uint32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"` // 点赞数
}

func (x *CountFavoriteResponse) Reset() {
	*x = CountFavoriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountFavoriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountFavoriteResponse) ProtoMessage() {}

func (x *CountFavoriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountFavoriteResponse.ProtoReflect.Descriptor instead.
func (*CountFavoriteResponse) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{7}
}

func (x *CountFavoriteResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CountFavoriteResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *CountFavoriteResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CountUserFavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
}

func (x *CountUserFavoriteRequest) Reset() {
	*x = CountUserFavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountUserFavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountUserFavoriteRequest) ProtoMessage() {}

func (x *CountUserFavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountUserFavoriteRequest.ProtoReflect.Descriptor instead.
func (*CountUserFavoriteRequest) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{8}
}

func (x *CountUserFavoriteRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CountUserFavoriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	Count      uint32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"` // 点赞数
}

func (x *CountUserFavoriteResponse) Reset() {
	*x = CountUserFavoriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountUserFavoriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountUserFavoriteResponse) ProtoMessage() {}

func (x *CountUserFavoriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountUserFavoriteResponse.ProtoReflect.Descriptor instead.
func (*CountUserFavoriteResponse) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{9}
}

func (x *CountUserFavoriteResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CountUserFavoriteResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *CountUserFavoriteResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CountUserTotalFavoritedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId uint32 `protobuf:"varint,1,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`
	UserId  uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CountUserTotalFavoritedRequest) Reset() {
	*x = CountUserTotalFavoritedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountUserTotalFavoritedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountUserTotalFavoritedRequest) ProtoMessage() {}

func (x *CountUserTotalFavoritedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountUserTotalFavoritedRequest.ProtoReflect.Descriptor instead.
func (*CountUserTotalFavoritedRequest) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{10}
}

func (x *CountUserTotalFavoritedRequest) GetActorId() uint32 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *CountUserTotalFavoritedRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CountUserTotalFavoritedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	Count      uint32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"` // 点赞数
}

func (x *CountUserTotalFavoritedResponse) Reset() {
	*x = CountUserTotalFavoritedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountUserTotalFavoritedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountUserTotalFavoritedResponse) ProtoMessage() {}

func (x *CountUserTotalFavoritedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountUserTotalFavoritedResponse.ProtoReflect.Descriptor instead.
func (*CountUserTotalFavoritedResponse) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{11}
}

func (x *CountUserTotalFavoritedResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CountUserTotalFavoritedResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *CountUserTotalFavoritedResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_favorite_proto protoreflect.FileDescriptor

var file_favorite_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x1a, 0x0a,
	0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x0f, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x52, 0x0a, 0x10, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x49, 0x0a, 0x13, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x14, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x2e, 0x0a, 0x0a,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x11,
	0x49, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x12, 0x49, 0x73, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x31, 0x0a, 0x14, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x6d, 0x0a, 0x15, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x18, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x71, 0x0a, 0x19,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x54, 0x0a, 0x1e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x77, 0x0a, 0x1f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xc2,
	0x04, 0x0a, 0x0f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x49, 0x73,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x49, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x49, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x11, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x76, 0x0a, 0x17, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x47, 0x75, 0x47, 0x6f, 0x54, 0x69, 0x6b, 0x2f, 0x73,
	0x72, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_favorite_proto_rawDescOnce sync.Once
	file_favorite_proto_rawDescData = file_favorite_proto_rawDesc
)

func file_favorite_proto_rawDescGZIP() []byte {
	file_favorite_proto_rawDescOnce.Do(func() {
		file_favorite_proto_rawDescData = protoimpl.X.CompressGZIP(file_favorite_proto_rawDescData)
	})
	return file_favorite_proto_rawDescData
}

var file_favorite_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_favorite_proto_goTypes = []interface{}{
	(*FavoriteRequest)(nil),                 // 0: rpc.favorite.FavoriteRequest
	(*FavoriteResponse)(nil),                // 1: rpc.favorite.FavoriteResponse
	(*FavoriteListRequest)(nil),             // 2: rpc.favorite.FavoriteListRequest
	(*FavoriteListResponse)(nil),            // 3: rpc.favorite.FavoriteListResponse
	(*IsFavoriteRequest)(nil),               // 4: rpc.favorite.IsFavoriteRequest
	(*IsFavoriteResponse)(nil),              // 5: rpc.favorite.IsFavoriteResponse
	(*CountFavoriteRequest)(nil),            // 6: rpc.favorite.CountFavoriteRequest
	(*CountFavoriteResponse)(nil),           // 7: rpc.favorite.CountFavoriteResponse
	(*CountUserFavoriteRequest)(nil),        // 8: rpc.favorite.CountUserFavoriteRequest
	(*CountUserFavoriteResponse)(nil),       // 9: rpc.favorite.CountUserFavoriteResponse
	(*CountUserTotalFavoritedRequest)(nil),  // 10: rpc.favorite.CountUserTotalFavoritedRequest
	(*CountUserTotalFavoritedResponse)(nil), // 11: rpc.favorite.CountUserTotalFavoritedResponse
	(*feed.Video)(nil),                      // 12: rpc.feed.Video
}
var file_favorite_proto_depIdxs = []int32{
	12, // 0: rpc.favorite.FavoriteListResponse.video_list:type_name -> rpc.feed.Video
	0,  // 1: rpc.favorite.FavoriteService.FavoriteAction:input_type -> rpc.favorite.FavoriteRequest
	2,  // 2: rpc.favorite.FavoriteService.FavoriteList:input_type -> rpc.favorite.FavoriteListRequest
	4,  // 3: rpc.favorite.FavoriteService.IsFavorite:input_type -> rpc.favorite.IsFavoriteRequest
	6,  // 4: rpc.favorite.FavoriteService.CountFavorite:input_type -> rpc.favorite.CountFavoriteRequest
	8,  // 5: rpc.favorite.FavoriteService.CountUserFavorite:input_type -> rpc.favorite.CountUserFavoriteRequest
	10, // 6: rpc.favorite.FavoriteService.CountUserTotalFavorited:input_type -> rpc.favorite.CountUserTotalFavoritedRequest
	1,  // 7: rpc.favorite.FavoriteService.FavoriteAction:output_type -> rpc.favorite.FavoriteResponse
	3,  // 8: rpc.favorite.FavoriteService.FavoriteList:output_type -> rpc.favorite.FavoriteListResponse
	5,  // 9: rpc.favorite.FavoriteService.IsFavorite:output_type -> rpc.favorite.IsFavoriteResponse
	7,  // 10: rpc.favorite.FavoriteService.CountFavorite:output_type -> rpc.favorite.CountFavoriteResponse
	9,  // 11: rpc.favorite.FavoriteService.CountUserFavorite:output_type -> rpc.favorite.CountUserFavoriteResponse
	11, // 12: rpc.favorite.FavoriteService.CountUserTotalFavorited:output_type -> rpc.favorite.CountUserTotalFavoritedResponse
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_favorite_proto_init() }
func file_favorite_proto_init() {
	if File_favorite_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_favorite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteRequest); i {
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
		file_favorite_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteResponse); i {
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
		file_favorite_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteListRequest); i {
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
		file_favorite_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteListResponse); i {
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
		file_favorite_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsFavoriteRequest); i {
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
		file_favorite_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsFavoriteResponse); i {
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
		file_favorite_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountFavoriteRequest); i {
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
		file_favorite_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountFavoriteResponse); i {
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
		file_favorite_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountUserFavoriteRequest); i {
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
		file_favorite_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountUserFavoriteResponse); i {
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
		file_favorite_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountUserTotalFavoritedRequest); i {
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
		file_favorite_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountUserTotalFavoritedResponse); i {
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
			RawDescriptor: file_favorite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_favorite_proto_goTypes,
		DependencyIndexes: file_favorite_proto_depIdxs,
		MessageInfos:      file_favorite_proto_msgTypes,
	}.Build()
	File_favorite_proto = out.File
	file_favorite_proto_rawDesc = nil
	file_favorite_proto_goTypes = nil
	file_favorite_proto_depIdxs = nil
}
