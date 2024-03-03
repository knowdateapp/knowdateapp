// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: api/v1/common/knowledge_base.proto

package common

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

type KnowledgeBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId     string                 `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Topic       string                 `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Collections []*Collection          `protobuf:"bytes,5,rep,name=collections,proto3" json:"collections,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *KnowledgeBase) Reset() {
	*x = KnowledgeBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_common_knowledge_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KnowledgeBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KnowledgeBase) ProtoMessage() {}

func (x *KnowledgeBase) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_common_knowledge_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KnowledgeBase.ProtoReflect.Descriptor instead.
func (*KnowledgeBase) Descriptor() ([]byte, []int) {
	return file_api_v1_common_knowledge_base_proto_rawDescGZIP(), []int{0}
}

func (x *KnowledgeBase) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *KnowledgeBase) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *KnowledgeBase) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *KnowledgeBase) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *KnowledgeBase) GetCollections() []*Collection {
	if x != nil {
		return x.Collections
	}
	return nil
}

func (x *KnowledgeBase) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_api_v1_common_knowledge_base_proto protoreflect.FileDescriptor

var file_api_v1_common_knowledge_base_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x6b, 0x6e, 0x6f, 0x77, 0x64, 0x61, 0x74, 0x65, 0x61, 0x70,
	0x70, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x0d, 0x4b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x0b, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x64, 0x61, 0x74, 0x65, 0x61, 0x70, 0x70, 0x2e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x58, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6e, 0x6f, 0x77, 0x64, 0x61, 0x74, 0x65, 0x61, 0x70, 0x70, 0x2f,
	0x6b, 0x6e, 0x6f, 0x77, 0x64, 0x61, 0x74, 0x65, 0x61, 0x70, 0x70, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_common_knowledge_base_proto_rawDescOnce sync.Once
	file_api_v1_common_knowledge_base_proto_rawDescData = file_api_v1_common_knowledge_base_proto_rawDesc
)

func file_api_v1_common_knowledge_base_proto_rawDescGZIP() []byte {
	file_api_v1_common_knowledge_base_proto_rawDescOnce.Do(func() {
		file_api_v1_common_knowledge_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_common_knowledge_base_proto_rawDescData)
	})
	return file_api_v1_common_knowledge_base_proto_rawDescData
}

var file_api_v1_common_knowledge_base_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1_common_knowledge_base_proto_goTypes = []interface{}{
	(*KnowledgeBase)(nil),         // 0: knowdateapp.knowledge.v1.common.KnowledgeBase
	(*Collection)(nil),            // 1: knowdateapp.knowledge.v1.common.Collection
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_api_v1_common_knowledge_base_proto_depIdxs = []int32{
	1, // 0: knowdateapp.knowledge.v1.common.KnowledgeBase.collections:type_name -> knowdateapp.knowledge.v1.common.Collection
	2, // 1: knowdateapp.knowledge.v1.common.KnowledgeBase.created_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_common_knowledge_base_proto_init() }
func file_api_v1_common_knowledge_base_proto_init() {
	if File_api_v1_common_knowledge_base_proto != nil {
		return
	}
	file_api_v1_common_collection_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_common_knowledge_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KnowledgeBase); i {
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
			RawDescriptor: file_api_v1_common_knowledge_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_common_knowledge_base_proto_goTypes,
		DependencyIndexes: file_api_v1_common_knowledge_base_proto_depIdxs,
		MessageInfos:      file_api_v1_common_knowledge_base_proto_msgTypes,
	}.Build()
	File_api_v1_common_knowledge_base_proto = out.File
	file_api_v1_common_knowledge_base_proto_rawDesc = nil
	file_api_v1_common_knowledge_base_proto_goTypes = nil
	file_api_v1_common_knowledge_base_proto_depIdxs = nil
}