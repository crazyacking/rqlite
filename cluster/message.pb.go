// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: message.proto

package cluster

import (
	command "github.com/rqlite/rqlite/command"
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

type Command_Type int32

const (
	Command_COMMAND_TYPE_UNKNOWN          Command_Type = 0
	Command_COMMAND_TYPE_GET_NODE_API_URL Command_Type = 1
	Command_COMMAND_TYPE_EXECUTE          Command_Type = 2
	Command_COMMAND_TYPE_QUERY            Command_Type = 3
)

// Enum value maps for Command_Type.
var (
	Command_Type_name = map[int32]string{
		0: "COMMAND_TYPE_UNKNOWN",
		1: "COMMAND_TYPE_GET_NODE_API_URL",
		2: "COMMAND_TYPE_EXECUTE",
		3: "COMMAND_TYPE_QUERY",
	}
	Command_Type_value = map[string]int32{
		"COMMAND_TYPE_UNKNOWN":          0,
		"COMMAND_TYPE_GET_NODE_API_URL": 1,
		"COMMAND_TYPE_EXECUTE":          2,
		"COMMAND_TYPE_QUERY":            3,
	}
)

func (x Command_Type) Enum() *Command_Type {
	p := new(Command_Type)
	*p = x
	return p
}

func (x Command_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Command_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (Command_Type) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x Command_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Command_Type.Descriptor instead.
func (Command_Type) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1, 0}
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Command_Type `protobuf:"varint,1,opt,name=type,proto3,enum=cluster.Command_Type" json:"type,omitempty"`
	// Types that are assignable to Request:
	//	*Command_ExecuteRequest
	//	*Command_QueryRequest
	Request isCommand_Request `protobuf_oneof:"request"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *Command) GetType() Command_Type {
	if x != nil {
		return x.Type
	}
	return Command_COMMAND_TYPE_UNKNOWN
}

func (m *Command) GetRequest() isCommand_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *Command) GetExecuteRequest() *command.ExecuteRequest {
	if x, ok := x.GetRequest().(*Command_ExecuteRequest); ok {
		return x.ExecuteRequest
	}
	return nil
}

func (x *Command) GetQueryRequest() *command.QueryRequest {
	if x, ok := x.GetRequest().(*Command_QueryRequest); ok {
		return x.QueryRequest
	}
	return nil
}

type isCommand_Request interface {
	isCommand_Request()
}

type Command_ExecuteRequest struct {
	ExecuteRequest *command.ExecuteRequest `protobuf:"bytes,2,opt,name=execute_request,json=executeRequest,proto3,oneof"`
}

type Command_QueryRequest struct {
	QueryRequest *command.QueryRequest `protobuf:"bytes,3,opt,name=query_request,json=queryRequest,proto3,oneof"`
}

func (*Command_ExecuteRequest) isCommand_Request() {}

func (*Command_QueryRequest) isCommand_Request() {}

type CommandExecuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   string                   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Results []*command.ExecuteResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *CommandExecuteResponse) Reset() {
	*x = CommandExecuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandExecuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandExecuteResponse) ProtoMessage() {}

func (x *CommandExecuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandExecuteResponse.ProtoReflect.Descriptor instead.
func (*CommandExecuteResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *CommandExecuteResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CommandExecuteResponse) GetResults() []*command.ExecuteResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type CommandQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string               `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Rows  []*command.QueryRows `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *CommandQueryResponse) Reset() {
	*x = CommandQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandQueryResponse) ProtoMessage() {}

func (x *CommandQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandQueryResponse.ProtoReflect.Descriptor instead.
func (*CommandQueryResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *CommandQueryResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CommandQueryResponse) GetRows() []*command.QueryRows {
	if x != nil {
		return x.Rows
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1b, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0xb8, 0x02, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x75, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x14, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x4d, 0x4d, 0x41,
	0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x4e, 0x4f, 0x44, 0x45,
	0x5f, 0x41, 0x50, 0x49, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f,
	0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55,
	0x54, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x03, 0x42, 0x09, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x60, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x30, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x54, 0x0a, 0x14, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x42,
	0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x71,
	0x6c, 0x69, 0x74, 0x65, 0x2f, 0x72, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_message_proto_goTypes = []interface{}{
	(Command_Type)(0),              // 0: cluster.Command.Type
	(*Address)(nil),                // 1: cluster.Address
	(*Command)(nil),                // 2: cluster.Command
	(*CommandExecuteResponse)(nil), // 3: cluster.CommandExecuteResponse
	(*CommandQueryResponse)(nil),   // 4: cluster.CommandQueryResponse
	(*command.ExecuteRequest)(nil), // 5: command.ExecuteRequest
	(*command.QueryRequest)(nil),   // 6: command.QueryRequest
	(*command.ExecuteResult)(nil),  // 7: command.ExecuteResult
	(*command.QueryRows)(nil),      // 8: command.QueryRows
}
var file_message_proto_depIdxs = []int32{
	0, // 0: cluster.Command.type:type_name -> cluster.Command.Type
	5, // 1: cluster.Command.execute_request:type_name -> command.ExecuteRequest
	6, // 2: cluster.Command.query_request:type_name -> command.QueryRequest
	7, // 3: cluster.CommandExecuteResponse.results:type_name -> command.ExecuteResult
	8, // 4: cluster.CommandQueryResponse.rows:type_name -> command.QueryRows
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandExecuteResponse); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandQueryResponse); i {
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
	file_message_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Command_ExecuteRequest)(nil),
		(*Command_QueryRequest)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
