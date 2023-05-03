// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: google/monitoring/v3/group.proto

package monitoringpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The description of a dynamic collection of monitored resources. Each group
// has a filter that is matched against monitored resources and their associated
// metadata. If a group's filter matches an available monitored resource, then
// that resource is a member of that group.  Groups can contain any number of
// monitored resources, and each monitored resource can be a member of any
// number of groups.
//
// Groups can be nested in parent-child hierarchies. The `parentName` field
// identifies an optional parent for each group.  If a group has a parent, then
// the only monitored resources available to be matched by the group's filter
// are the resources contained in the parent group.  In other words, a group
// contains the monitored resources that match its filter and the filters of all
// the group's ancestors.  A group without a parent can contain any monitored
// resource.
//
// For example, consider an infrastructure running a set of instances with two
// user-defined tags: `"environment"` and `"role"`. A parent group has a filter,
// `environment="production"`.  A child of that parent group has a filter,
// `role="transcoder"`.  The parent group contains all instances in the
// production environment, regardless of their roles.  The child group contains
// instances that have the transcoder role *and* are in the production
// environment.
//
// The monitored resources contained in a group can change at any moment,
// depending on what resources exist and what filters are associated with the
// group and its ancestors.
type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The name of this group. The format is:
	//
	//	projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID]
	//
	// When creating a group, this field is ignored and a new name is created
	// consisting of the project specified in the call to `CreateGroup`
	// and a unique `[GROUP_ID]` that is generated automatically.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A user-assigned name for this group, used only for display purposes.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The name of the group's parent, if it has one. The format is:
	//
	//	projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID]
	//
	// For groups with no parent, `parent_name` is the empty string, `""`.
	ParentName string `protobuf:"bytes,3,opt,name=parent_name,json=parentName,proto3" json:"parent_name,omitempty"`
	// The filter used to determine which monitored resources belong to this
	// group.
	Filter string `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
	// If true, the members of this group are considered to be a cluster.
	// The system can perform additional analysis on groups that are clusters.
	IsCluster bool `protobuf:"varint,6,opt,name=is_cluster,json=isCluster,proto3" json:"is_cluster,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_v3_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_v3_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_google_monitoring_v3_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Group) GetParentName() string {
	if x != nil {
		return x.ParentName
	}
	return ""
}

func (x *Group) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *Group) GetIsCluster() bool {
	if x != nil {
		return x.IsCluster
	}
	return false
}

var File_google_monitoring_v3_group_proto protoreflect.FileDescriptor

var file_google_monitoring_v3_group_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x02, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x99, 0x01, 0xea,
	0x41, 0x95, 0x01, 0x0a, 0x1f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x21, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f,
	0x7b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x7d, 0x12, 0x2b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x7d, 0x12, 0x1f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x7d, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x7d, 0x12, 0x01, 0x2a, 0x42, 0xc2, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x33, 0x42, 0x0a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x3b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x33,
	0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x33, 0xea, 0x02, 0x1d,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_monitoring_v3_group_proto_rawDescOnce sync.Once
	file_google_monitoring_v3_group_proto_rawDescData = file_google_monitoring_v3_group_proto_rawDesc
)

func file_google_monitoring_v3_group_proto_rawDescGZIP() []byte {
	file_google_monitoring_v3_group_proto_rawDescOnce.Do(func() {
		file_google_monitoring_v3_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_monitoring_v3_group_proto_rawDescData)
	})
	return file_google_monitoring_v3_group_proto_rawDescData
}

var file_google_monitoring_v3_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_monitoring_v3_group_proto_goTypes = []interface{}{
	(*Group)(nil), // 0: google.monitoring.v3.Group
}
var file_google_monitoring_v3_group_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_monitoring_v3_group_proto_init() }
func file_google_monitoring_v3_group_proto_init() {
	if File_google_monitoring_v3_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_monitoring_v3_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
			RawDescriptor: file_google_monitoring_v3_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_monitoring_v3_group_proto_goTypes,
		DependencyIndexes: file_google_monitoring_v3_group_proto_depIdxs,
		MessageInfos:      file_google_monitoring_v3_group_proto_msgTypes,
	}.Build()
	File_google_monitoring_v3_group_proto = out.File
	file_google_monitoring_v3_group_proto_rawDesc = nil
	file_google_monitoring_v3_group_proto_goTypes = nil
	file_google_monitoring_v3_group_proto_depIdxs = nil
}
