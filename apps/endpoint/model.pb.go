// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: apps/endpoint/pb/model.proto

package endpoint

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

type EndpointSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// @gotags: json:"items"
	Items []*Endpoint `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *EndpointSet) Reset() {
	*x = EndpointSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointSet) ProtoMessage() {}

func (x *EndpointSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointSet.ProtoReflect.Descriptor instead.
func (*EndpointSet) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_model_proto_rawDescGZIP(), []int{0}
}

func (x *EndpointSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *EndpointSet) GetItems() []*Endpoint {
	if x != nil {
		return x.Items
	}
	return nil
}

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: gorm:"embedded"
	Meta *Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty" gorm:"embedded"`
	// @gotags: gorm:"embedded"
	Spec *Spec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty" gorm:"embedded"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_model_proto_rawDescGZIP(), []int{1}
}

func (x *Endpoint) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Endpoint) GetSpec() *Spec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"created_at" gorm:"primaryKey"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"created_at" gorm:"primaryKey"`
	// @gotags: json:"created_at" gorm:"autoCreateTime"
	CreatedAt int64 `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at" gorm:"autoCreateTime"`
	// @gotags: json:"updated_at" gorm:"autoUpdateTime"
	UpdatedAt int64 `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at" gorm:"autoUpdateTime"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_model_proto_rawDescGZIP(), []int{2}
}

func (x *Meta) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Meta) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Meta) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

// 如果让Role 取关联Ep 表的普通ID，那么当服务重新上线后（注册时，会先删除、后创建ep，此时ID会变动）那么就会导致已经配置的权限失效
// 因此  identity = "svc.module.method.url" ，创建权限时使用 Role 表跟这个 identify 进行关联
// 核心是确保被引用的 identify 字段的值，在服务端重启后不能发生改变
type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务名称
	// @gotags: json:"service" validate:"required"
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service" validate:"required"`
	// 模块名称
	// @gotags: json:"resource" validate:"required"
	Resource string `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource" validate:"required"`
	// 操作行为
	// @gotags: json:"action" validate:"required"
	Action string `protobuf:"bytes,3,opt,name=action,proto3" json:"action" validate:"required"`
	// 端点唯一标识
	// identity = "service.model.path" = "cmdb.cmdb.post./cmdb/api/v1/secret"
	// @gotags: json:"identity" validate:"required"
	Identity string `protobuf:"bytes,5,opt,name=identity,proto3" json:"identity" validate:"required"`
	// 端点描述
	// @gotags: json:"desc" validate:"required"
	Desc string `protobuf:"bytes,6,opt,name=desc,proto3" json:"desc" validate:"required"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec.ProtoReflect.Descriptor instead.
func (*Spec) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_model_proto_rawDescGZIP(), []int{3}
}

func (x *Spec) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Spec) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Spec) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Spec) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *Spec) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

var File_apps_endpoint_pb_model_proto protoreflect.FileDescriptor

var file_apps_endpoint_pb_model_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f,
	0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x22, 0x55, 0x0a, 0x0b, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x62, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12,
	0x2a, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x54, 0x0a, 0x04, 0x4d,
	0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x84, 0x01, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x61, 0x6f, 0x67, 0x79, 0x39, 0x31, 0x2f,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_endpoint_pb_model_proto_rawDescOnce sync.Once
	file_apps_endpoint_pb_model_proto_rawDescData = file_apps_endpoint_pb_model_proto_rawDesc
)

func file_apps_endpoint_pb_model_proto_rawDescGZIP() []byte {
	file_apps_endpoint_pb_model_proto_rawDescOnce.Do(func() {
		file_apps_endpoint_pb_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_endpoint_pb_model_proto_rawDescData)
	})
	return file_apps_endpoint_pb_model_proto_rawDescData
}

var file_apps_endpoint_pb_model_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_endpoint_pb_model_proto_goTypes = []any{
	(*EndpointSet)(nil), // 0: mcenter.endpoint.EndpointSet
	(*Endpoint)(nil),    // 1: mcenter.endpoint.Endpoint
	(*Meta)(nil),        // 2: mcenter.endpoint.Meta
	(*Spec)(nil),        // 3: mcenter.endpoint.Spec
}
var file_apps_endpoint_pb_model_proto_depIdxs = []int32{
	1, // 0: mcenter.endpoint.EndpointSet.items:type_name -> mcenter.endpoint.Endpoint
	2, // 1: mcenter.endpoint.Endpoint.meta:type_name -> mcenter.endpoint.Meta
	3, // 2: mcenter.endpoint.Endpoint.spec:type_name -> mcenter.endpoint.Spec
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_endpoint_pb_model_proto_init() }
func file_apps_endpoint_pb_model_proto_init() {
	if File_apps_endpoint_pb_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_endpoint_pb_model_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EndpointSet); i {
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
		file_apps_endpoint_pb_model_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Endpoint); i {
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
		file_apps_endpoint_pb_model_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Meta); i {
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
		file_apps_endpoint_pb_model_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Spec); i {
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
			RawDescriptor: file_apps_endpoint_pb_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_endpoint_pb_model_proto_goTypes,
		DependencyIndexes: file_apps_endpoint_pb_model_proto_depIdxs,
		MessageInfos:      file_apps_endpoint_pb_model_proto_msgTypes,
	}.Build()
	File_apps_endpoint_pb_model_proto = out.File
	file_apps_endpoint_pb_model_proto_rawDesc = nil
	file_apps_endpoint_pb_model_proto_goTypes = nil
	file_apps_endpoint_pb_model_proto_depIdxs = nil
}