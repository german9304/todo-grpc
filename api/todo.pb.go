// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: api/todo.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content   string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Completed bool   `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_api_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_api_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Todo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Todo) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

type TodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TodosRequest) Reset() {
	*x = TodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodosRequest) ProtoMessage() {}

func (x *TodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodosRequest.ProtoReflect.Descriptor instead.
func (*TodosRequest) Descriptor() ([]byte, []int) {
	return file_api_todo_proto_rawDescGZIP(), []int{1}
}

type TodosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *TodosResponse) Reset() {
	*x = TodosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodosResponse) ProtoMessage() {}

func (x *TodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodosResponse.ProtoReflect.Descriptor instead.
func (*TodosResponse) Descriptor() ([]byte, []int) {
	return file_api_todo_proto_rawDescGZIP(), []int{2}
}

func (x *TodosResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type TodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TodoRequest) Reset() {
	*x = TodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoRequest) ProtoMessage() {}

func (x *TodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoRequest.ProtoReflect.Descriptor instead.
func (*TodoRequest) Descriptor() ([]byte, []int) {
	return file_api_todo_proto_rawDescGZIP(), []int{3}
}

func (x *TodoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *TodoResponse) Reset() {
	*x = TodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoResponse) ProtoMessage() {}

func (x *TodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoResponse.ProtoReflect.Descriptor instead.
func (*TodoResponse) Descriptor() ([]byte, []int) {
	return file_api_todo_proto_rawDescGZIP(), []int{4}
}

func (x *TodoResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

var File_api_todo_proto protoreflect.FileDescriptor

var file_api_todo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x4e, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x0d, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52,
	0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x1d, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x0c, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74,
	0x6f, 0x64, 0x6f, 0x32, 0x70, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x11, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2d, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x10,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_todo_proto_rawDescOnce sync.Once
	file_api_todo_proto_rawDescData = file_api_todo_proto_rawDesc
)

func file_api_todo_proto_rawDescGZIP() []byte {
	file_api_todo_proto_rawDescOnce.Do(func() {
		file_api_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_todo_proto_rawDescData)
	})
	return file_api_todo_proto_rawDescData
}

var file_api_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_todo_proto_goTypes = []interface{}{
	(*Todo)(nil),          // 0: api.Todo
	(*TodosRequest)(nil),  // 1: api.TodosRequest
	(*TodosResponse)(nil), // 2: api.TodosResponse
	(*TodoRequest)(nil),   // 3: api.TodoRequest
	(*TodoResponse)(nil),  // 4: api.TodoResponse
}
var file_api_todo_proto_depIdxs = []int32{
	0, // 0: api.TodosResponse.todo:type_name -> api.Todo
	0, // 1: api.TodoResponse.todo:type_name -> api.Todo
	1, // 2: api.TodoService.Todos:input_type -> api.TodosRequest
	3, // 3: api.TodoService.Todo:input_type -> api.TodoRequest
	2, // 4: api.TodoService.Todos:output_type -> api.TodosResponse
	4, // 5: api.TodoService.Todo:output_type -> api.TodoResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_todo_proto_init() }
func file_api_todo_proto_init() {
	if File_api_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_api_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodosRequest); i {
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
		file_api_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodosResponse); i {
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
		file_api_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoRequest); i {
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
		file_api_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoResponse); i {
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
			RawDescriptor: file_api_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_todo_proto_goTypes,
		DependencyIndexes: file_api_todo_proto_depIdxs,
		MessageInfos:      file_api_todo_proto_msgTypes,
	}.Build()
	File_api_todo_proto = out.File
	file_api_todo_proto_rawDesc = nil
	file_api_todo_proto_goTypes = nil
	file_api_todo_proto_depIdxs = nil
}
