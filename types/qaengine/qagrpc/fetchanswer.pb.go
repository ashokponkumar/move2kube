/*
 *  Copyright IBM Corporation 2020, 2021
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

//
//Copyright IBM Corporation 2021
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative fetchanswer.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: fetchanswer.proto

package qagrpc

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

type Problem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type        string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Hints       []string `protobuf:"bytes,4,rep,name=hints,proto3" json:"hints,omitempty"`
	Options     []string `protobuf:"bytes,5,rep,name=options,proto3" json:"options,omitempty"`
	Default     []string `protobuf:"bytes,6,rep,name=default,proto3" json:"default,omitempty"`
}

func (x *Problem) Reset() {
	*x = Problem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fetchanswer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Problem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem) ProtoMessage() {}

func (x *Problem) ProtoReflect() protoreflect.Message {
	mi := &file_fetchanswer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Problem.ProtoReflect.Descriptor instead.
func (*Problem) Descriptor() ([]byte, []int) {
	return file_fetchanswer_proto_rawDescGZIP(), []int{0}
}

func (x *Problem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Problem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Problem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Problem) GetHints() []string {
	if x != nil {
		return x.Hints
	}
	return nil
}

func (x *Problem) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *Problem) GetDefault() []string {
	if x != nil {
		return x.Default
	}
	return nil
}

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer []string `protobuf:"bytes,1,rep,name=answer,proto3" json:"answer,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fetchanswer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_fetchanswer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_fetchanswer_proto_rawDescGZIP(), []int{1}
}

func (x *Answer) GetAnswer() []string {
	if x != nil {
		return x.Answer
	}
	return nil
}

var File_fetchanswer_proto protoreflect.FileDescriptor

var file_fetchanswer_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x65, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x71, 0x61, 0x67, 0x72, 0x70, 0x63, 0x22, 0x99, 0x01, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x68, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x69,
	0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0x20, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x32, 0x3c, 0x0a, 0x08, 0x51, 0x41, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x71, 0x61, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x1a, 0x0e, 0x2e, 0x71, 0x61, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6e, 0x76, 0x65, 0x79, 0x6f, 0x72, 0x2f, 0x6d,
	0x6f, 0x76, 0x65, 0x32, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x71,
	0x61, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x71, 0x61, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fetchanswer_proto_rawDescOnce sync.Once
	file_fetchanswer_proto_rawDescData = file_fetchanswer_proto_rawDesc
)

func file_fetchanswer_proto_rawDescGZIP() []byte {
	file_fetchanswer_proto_rawDescOnce.Do(func() {
		file_fetchanswer_proto_rawDescData = protoimpl.X.CompressGZIP(file_fetchanswer_proto_rawDescData)
	})
	return file_fetchanswer_proto_rawDescData
}

var file_fetchanswer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fetchanswer_proto_goTypes = []interface{}{
	(*Problem)(nil), // 0: qagrpc.Problem
	(*Answer)(nil),  // 1: qagrpc.Answer
}
var file_fetchanswer_proto_depIdxs = []int32{
	0, // 0: qagrpc.QAEngine.FetchAnswer:input_type -> qagrpc.Problem
	1, // 1: qagrpc.QAEngine.FetchAnswer:output_type -> qagrpc.Answer
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fetchanswer_proto_init() }
func file_fetchanswer_proto_init() {
	if File_fetchanswer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fetchanswer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Problem); i {
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
		file_fetchanswer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
			RawDescriptor: file_fetchanswer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fetchanswer_proto_goTypes,
		DependencyIndexes: file_fetchanswer_proto_depIdxs,
		MessageInfos:      file_fetchanswer_proto_msgTypes,
	}.Build()
	File_fetchanswer_proto = out.File
	file_fetchanswer_proto_rawDesc = nil
	file_fetchanswer_proto_goTypes = nil
	file_fetchanswer_proto_depIdxs = nil
}
