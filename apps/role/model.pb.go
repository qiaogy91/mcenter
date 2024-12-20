// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: apps/role/pb/model.proto

package role

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

type RoleSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64   `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Role `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *RoleSet) Reset() {
	*x = RoleSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleSet) ProtoMessage() {}

func (x *RoleSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleSet.ProtoReflect.Descriptor instead.
func (*RoleSet) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_model_proto_rawDescGZIP(), []int{0}
}

func (x *RoleSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RoleSet) GetItems() []*Role {
	if x != nil {
		return x.Items
	}
	return nil
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: gorm:"embedded"
	Meta *Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty" gorm:"embedded"`
	// @gotags: gorm:"embedded"
	Spec *Spec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty" gorm:"embedded"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_model_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Role) GetSpec() *Spec {
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
		mi := &file_apps_role_pb_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_model_proto_msgTypes[2]
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
	return file_apps_role_pb_model_proto_rawDescGZIP(), []int{2}
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

type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"name" validate:"required"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" validate:"required"`
	// @gotags: json:"namespace" validate:"required"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" validate:"required"`
	// @gotags: json:"service" validate:"required" gorm:"json"
	Service []string `protobuf:"bytes,3,rep,name=service,proto3" json:"service" validate:"required" gorm:"json"`
	// @gotags: json:"resource" validate:"required" gorm:"json"
	Resource []string `protobuf:"bytes,4,rep,name=resource,proto3" json:"resource" validate:"required" gorm:"json"`
	// @gotags: json:"actions" validate:"required" gorm:"json"
	Actions []string `protobuf:"bytes,5,rep,name=actions,proto3" json:"actions" validate:"required" gorm:"json"`
	// @gotags: json:"description" validate:"required"
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description" validate:"required"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_model_proto_msgTypes[3]
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
	return file_apps_role_pb_model_proto_rawDescGZIP(), []int{3}
}

func (x *Spec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Spec) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Spec) GetService() []string {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *Spec) GetResource() []string {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Spec) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *Spec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_apps_role_pb_model_proto protoreflect.FileDescriptor

var file_apps_role_pb_model_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x49, 0x0a, 0x07, 0x52, 0x6f, 0x6c, 0x65,
	0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x56, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x54, 0x0a, 0x04, 0x4d,
	0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xaa, 0x01, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x27,
	0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x61,
	0x6f, 0x67, 0x79, 0x39, 0x31, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_role_pb_model_proto_rawDescOnce sync.Once
	file_apps_role_pb_model_proto_rawDescData = file_apps_role_pb_model_proto_rawDesc
)

func file_apps_role_pb_model_proto_rawDescGZIP() []byte {
	file_apps_role_pb_model_proto_rawDescOnce.Do(func() {
		file_apps_role_pb_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_role_pb_model_proto_rawDescData)
	})
	return file_apps_role_pb_model_proto_rawDescData
}

var file_apps_role_pb_model_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_role_pb_model_proto_goTypes = []any{
	(*RoleSet)(nil), // 0: mcenter.role.RoleSet
	(*Role)(nil),    // 1: mcenter.role.Role
	(*Meta)(nil),    // 2: mcenter.role.Meta
	(*Spec)(nil),    // 3: mcenter.role.Spec
}
var file_apps_role_pb_model_proto_depIdxs = []int32{
	1, // 0: mcenter.role.RoleSet.items:type_name -> mcenter.role.Role
	2, // 1: mcenter.role.Role.meta:type_name -> mcenter.role.Meta
	3, // 2: mcenter.role.Role.spec:type_name -> mcenter.role.Spec
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_role_pb_model_proto_init() }
func file_apps_role_pb_model_proto_init() {
	if File_apps_role_pb_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_role_pb_model_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RoleSet); i {
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
		file_apps_role_pb_model_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Role); i {
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
		file_apps_role_pb_model_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_apps_role_pb_model_proto_msgTypes[3].Exporter = func(v any, i int) any {
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
			RawDescriptor: file_apps_role_pb_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_role_pb_model_proto_goTypes,
		DependencyIndexes: file_apps_role_pb_model_proto_depIdxs,
		MessageInfos:      file_apps_role_pb_model_proto_msgTypes,
	}.Build()
	File_apps_role_pb_model_proto = out.File
	file_apps_role_pb_model_proto_rawDesc = nil
	file_apps_role_pb_model_proto_goTypes = nil
	file_apps_role_pb_model_proto_depIdxs = nil
}
