// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: apps/token/pb/rpc.proto

package token

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

type QueryType int32

const (
	QueryType_QUERY_TYPE_UNSPECIFIED QueryType = 0
	QueryType_QUERY_TYPE_USERNAME    QueryType = 1
)

// Enum value maps for QueryType.
var (
	QueryType_name = map[int32]string{
		0: "QUERY_TYPE_UNSPECIFIED",
		1: "QUERY_TYPE_USERNAME",
	}
	QueryType_value = map[string]int32{
		"QUERY_TYPE_UNSPECIFIED": 0,
		"QUERY_TYPE_USERNAME":    1,
	}
)

func (x QueryType) Enum() *QueryType {
	p := new(QueryType)
	*p = x
	return p
}

func (x QueryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryType) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_token_pb_rpc_proto_enumTypes[0].Descriptor()
}

func (QueryType) Type() protoreflect.EnumType {
	return &file_apps_token_pb_rpc_proto_enumTypes[0]
}

func (x QueryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryType.Descriptor instead.
func (QueryType) EnumDescriptor() ([]byte, []int) {
	return file_apps_token_pb_rpc_proto_rawDescGZIP(), []int{0}
}

type IssueTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty" validate:"required"`
	// @gotags: validate:"required"
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" validate:"required"`
}

func (x *IssueTokenRequest) Reset() {
	*x = IssueTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_token_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueTokenRequest) ProtoMessage() {}

func (x *IssueTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_token_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueTokenRequest.ProtoReflect.Descriptor instead.
func (*IssueTokenRequest) Descriptor() ([]byte, []int) {
	return file_apps_token_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *IssueTokenRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *IssueTokenRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DeleteTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty" validate:"required"` // 要被撤销的Token
	// @gotags: validate:"required"
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty" validate:"required"` // 需要refresh token 才能撤销
}

func (x *DeleteTokenRequest) Reset() {
	*x = DeleteTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_token_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTokenRequest) ProtoMessage() {}

func (x *DeleteTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_token_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTokenRequest.ProtoReflect.Descriptor instead.
func (*DeleteTokenRequest) Descriptor() ([]byte, []int) {
	return file_apps_token_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *DeleteTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty" validate:"required"` // 要被延长的Token
	// @gotags: validate:"required"
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty" validate:"required"` // 需要refresh token 才能延长
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_token_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_token_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_apps_token_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *RefreshTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RefreshTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type QueryTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	PageNum int64 `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty" validate:"required"`
	// @gotags: validate:"required"
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" validate:"required"`
	// @gotags: validate:"required"
	QueryType QueryType `protobuf:"varint,3,opt,name=query_type,json=queryType,proto3,enum=mcenter.token.QueryType" json:"query_type,omitempty" validate:"required"`
	Keyword   string    `protobuf:"bytes,4,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *QueryTokenRequest) Reset() {
	*x = QueryTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_token_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTokenRequest) ProtoMessage() {}

func (x *QueryTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_token_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTokenRequest.ProtoReflect.Descriptor instead.
func (*QueryTokenRequest) Descriptor() ([]byte, []int) {
	return file_apps_token_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *QueryTokenRequest) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *QueryTokenRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryTokenRequest) GetQueryType() QueryType {
	if x != nil {
		return x.QueryType
	}
	return QueryType_QUERY_TYPE_UNSPECIFIED
}

func (x *QueryTokenRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

type GetTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty" validate:"required"` // 要被获取的Token
}

func (x *GetTokenRequest) Reset() {
	*x = GetTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_token_pb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenRequest) ProtoMessage() {}

func (x *GetTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_token_pb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenRequest.ProtoReflect.Descriptor instead.
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return file_apps_token_pb_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *GetTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type ValidateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty" validate:"required"` // 要被验证的Token
}

func (x *ValidateTokenRequest) Reset() {
	*x = ValidateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_token_pb_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenRequest) ProtoMessage() {}

func (x *ValidateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_token_pb_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return file_apps_token_pb_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

var File_apps_token_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_token_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x11, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x5c, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5d,
	0x0a, 0x13, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9e, 0x01,
	0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x18, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x34,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x39, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2a,
	0x40, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16,
	0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x51, 0x55, 0x45, 0x52,
	0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x4e, 0x41, 0x4d, 0x45, 0x10,
	0x01, 0x32, 0xb4, 0x03, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12, 0x44, 0x0a, 0x0a, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x46, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21,
	0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x40, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x47, 0x0a, 0x0a, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53,
	0x65, 0x74, 0x12, 0x48, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x22, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x4a, 0x0a, 0x0d,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x61, 0x6f, 0x67, 0x79, 0x39, 0x31, 0x2f,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_token_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_token_pb_rpc_proto_rawDescData = file_apps_token_pb_rpc_proto_rawDesc
)

func file_apps_token_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_token_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_token_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_token_pb_rpc_proto_rawDescData)
	})
	return file_apps_token_pb_rpc_proto_rawDescData
}

var file_apps_token_pb_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_token_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apps_token_pb_rpc_proto_goTypes = []any{
	(QueryType)(0),               // 0: mcenter.token.QueryType
	(*IssueTokenRequest)(nil),    // 1: mcenter.token.IssueTokenRequest
	(*DeleteTokenRequest)(nil),   // 2: mcenter.token.DeleteTokenRequest
	(*RefreshTokenRequest)(nil),  // 3: mcenter.token.RefreshTokenRequest
	(*QueryTokenRequest)(nil),    // 4: mcenter.token.QueryTokenRequest
	(*GetTokenRequest)(nil),      // 5: mcenter.token.GetTokenRequest
	(*ValidateTokenRequest)(nil), // 6: mcenter.token.ValidateTokenRequest
	(*Token)(nil),                // 7: mcenter.token.Token
	(*TokenSet)(nil),             // 8: mcenter.token.TokenSet
}
var file_apps_token_pb_rpc_proto_depIdxs = []int32{
	0, // 0: mcenter.token.QueryTokenRequest.query_type:type_name -> mcenter.token.QueryType
	1, // 1: mcenter.token.Rpc.IssueToken:input_type -> mcenter.token.IssueTokenRequest
	2, // 2: mcenter.token.Rpc.DeleteToken:input_type -> mcenter.token.DeleteTokenRequest
	5, // 3: mcenter.token.Rpc.GetToken:input_type -> mcenter.token.GetTokenRequest
	4, // 4: mcenter.token.Rpc.QueryToken:input_type -> mcenter.token.QueryTokenRequest
	3, // 5: mcenter.token.Rpc.RefreshToken:input_type -> mcenter.token.RefreshTokenRequest
	6, // 6: mcenter.token.Rpc.ValidateToken:input_type -> mcenter.token.ValidateTokenRequest
	7, // 7: mcenter.token.Rpc.IssueToken:output_type -> mcenter.token.Token
	7, // 8: mcenter.token.Rpc.DeleteToken:output_type -> mcenter.token.Token
	7, // 9: mcenter.token.Rpc.GetToken:output_type -> mcenter.token.Token
	8, // 10: mcenter.token.Rpc.QueryToken:output_type -> mcenter.token.TokenSet
	7, // 11: mcenter.token.Rpc.RefreshToken:output_type -> mcenter.token.Token
	7, // 12: mcenter.token.Rpc.ValidateToken:output_type -> mcenter.token.Token
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_token_pb_rpc_proto_init() }
func file_apps_token_pb_rpc_proto_init() {
	if File_apps_token_pb_rpc_proto != nil {
		return
	}
	file_apps_token_pb_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_token_pb_rpc_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IssueTokenRequest); i {
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
		file_apps_token_pb_rpc_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTokenRequest); i {
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
		file_apps_token_pb_rpc_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RefreshTokenRequest); i {
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
		file_apps_token_pb_rpc_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*QueryTokenRequest); i {
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
		file_apps_token_pb_rpc_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetTokenRequest); i {
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
		file_apps_token_pb_rpc_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ValidateTokenRequest); i {
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
			RawDescriptor: file_apps_token_pb_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_token_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_token_pb_rpc_proto_depIdxs,
		EnumInfos:         file_apps_token_pb_rpc_proto_enumTypes,
		MessageInfos:      file_apps_token_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_token_pb_rpc_proto = out.File
	file_apps_token_pb_rpc_proto_rawDesc = nil
	file_apps_token_pb_rpc_proto_goTypes = nil
	file_apps_token_pb_rpc_proto_depIdxs = nil
}
