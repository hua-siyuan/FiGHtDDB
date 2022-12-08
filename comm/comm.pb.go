// Copyright 2015 gRPC authors.
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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: comm.proto

package comm

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

// The request message containing sql string
type SqlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SqlStr string `protobuf:"bytes,1,opt,name=sqlStr,proto3" json:"sqlStr,omitempty"`
}

func (x *SqlRequest) Reset() {
	*x = SqlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqlRequest) ProtoMessage() {}

func (x *SqlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqlRequest.ProtoReflect.Descriptor instead.
func (*SqlRequest) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{0}
}

func (x *SqlRequest) GetSqlStr() string {
	if x != nil {
		return x.SqlStr
	}
	return ""
}

// The response message containing execution result
type SqlResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rc   int32  `protobuf:"varint,1,opt,name=rc,proto3" json:"rc,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SqlResult) Reset() {
	*x = SqlResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqlResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqlResult) ProtoMessage() {}

func (x *SqlResult) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqlResult.ProtoReflect.Descriptor instead.
func (*SqlResult) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{1}
}

func (x *SqlResult) GetRc() int32 {
	if x != nil {
		return x.Rc
	}
	return 0
}

func (x *SqlResult) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_comm_proto protoreflect.FileDescriptor

var file_comm_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f,
	0x6d, 0x6d, 0x22, 0x24, 0x0a, 0x0a, 0x53, 0x71, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x71, 0x6c, 0x53, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x71, 0x6c, 0x53, 0x74, 0x72, 0x22, 0x2f, 0x0a, 0x09, 0x53, 0x71, 0x6c, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x72, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x6a, 0x0a, 0x08, 0x44, 0x61, 0x74,
	0x61, 0x42, 0x61, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x71, 0x6c,
	0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x53, 0x71, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x53, 0x71, 0x6c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x53, 0x71, 0x6c,
	0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x53, 0x71, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x53, 0x71, 0x6c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x69, 0x47, 0x48, 0x74, 0x44, 0x44, 0x42, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comm_proto_rawDescOnce sync.Once
	file_comm_proto_rawDescData = file_comm_proto_rawDesc
)

func file_comm_proto_rawDescGZIP() []byte {
	file_comm_proto_rawDescOnce.Do(func() {
		file_comm_proto_rawDescData = protoimpl.X.CompressGZIP(file_comm_proto_rawDescData)
	})
	return file_comm_proto_rawDescData
}

var file_comm_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_comm_proto_goTypes = []interface{}{
	(*SqlRequest)(nil), // 0: comm.SqlRequest
	(*SqlResult)(nil),  // 1: comm.SqlResult
}
var file_comm_proto_depIdxs = []int32{
	0, // 0: comm.DataBase.SendSql:input_type -> comm.SqlRequest
	0, // 1: comm.DataBase.ExecSql:input_type -> comm.SqlRequest
	1, // 2: comm.DataBase.SendSql:output_type -> comm.SqlResult
	1, // 3: comm.DataBase.ExecSql:output_type -> comm.SqlResult
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_comm_proto_init() }
func file_comm_proto_init() {
	if File_comm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqlRequest); i {
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
		file_comm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqlResult); i {
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
			RawDescriptor: file_comm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comm_proto_goTypes,
		DependencyIndexes: file_comm_proto_depIdxs,
		MessageInfos:      file_comm_proto_msgTypes,
	}.Build()
	File_comm_proto = out.File
	file_comm_proto_rawDesc = nil
	file_comm_proto_goTypes = nil
	file_comm_proto_depIdxs = nil
}
