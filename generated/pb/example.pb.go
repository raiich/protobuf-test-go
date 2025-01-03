// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: pb/example.proto

package pb

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

type MyU32 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	U32 uint32 `protobuf:"varint,1,opt,name=u32,proto3" json:"u32,omitempty"`
}

func (x *MyU32) Reset() {
	*x = MyU32{}
	mi := &file_pb_example_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MyU32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyU32) ProtoMessage() {}

func (x *MyU32) ProtoReflect() protoreflect.Message {
	mi := &file_pb_example_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyU32.ProtoReflect.Descriptor instead.
func (*MyU32) Descriptor() ([]byte, []int) {
	return file_pb_example_proto_rawDescGZIP(), []int{0}
}

func (x *MyU32) GetU32() uint32 {
	if x != nil {
		return x.U32
	}
	return 0
}

type MyStr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str string `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
}

func (x *MyStr) Reset() {
	*x = MyStr{}
	mi := &file_pb_example_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MyStr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyStr) ProtoMessage() {}

func (x *MyStr) ProtoReflect() protoreflect.Message {
	mi := &file_pb_example_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyStr.ProtoReflect.Descriptor instead.
func (*MyStr) Descriptor() ([]byte, []int) {
	return file_pb_example_proto_rawDescGZIP(), []int{1}
}

func (x *MyStr) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

type MySub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sub *MyStr `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
}

func (x *MySub) Reset() {
	*x = MySub{}
	mi := &file_pb_example_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MySub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MySub) ProtoMessage() {}

func (x *MySub) ProtoReflect() protoreflect.Message {
	mi := &file_pb_example_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MySub.ProtoReflect.Descriptor instead.
func (*MySub) Descriptor() ([]byte, []int) {
	return file_pb_example_proto_rawDescGZIP(), []int{2}
}

func (x *MySub) GetSub() *MyStr {
	if x != nil {
		return x.Sub
	}
	return nil
}

type MyRepeated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	U32 []uint32 `protobuf:"varint,1,rep,packed,name=u32,proto3" json:"u32,omitempty"`
	Str []string `protobuf:"bytes,2,rep,name=str,proto3" json:"str,omitempty"`
}

func (x *MyRepeated) Reset() {
	*x = MyRepeated{}
	mi := &file_pb_example_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MyRepeated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyRepeated) ProtoMessage() {}

func (x *MyRepeated) ProtoReflect() protoreflect.Message {
	mi := &file_pb_example_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyRepeated.ProtoReflect.Descriptor instead.
func (*MyRepeated) Descriptor() ([]byte, []int) {
	return file_pb_example_proto_rawDescGZIP(), []int{3}
}

func (x *MyRepeated) GetU32() []uint32 {
	if x != nil {
		return x.U32
	}
	return nil
}

func (x *MyRepeated) GetStr() []string {
	if x != nil {
		return x.Str
	}
	return nil
}

var File_pb_example_proto protoreflect.FileDescriptor

var file_pb_example_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x19, 0x0a, 0x05, 0x4d, 0x79, 0x55, 0x33, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x33, 0x32, 0x22, 0x19, 0x0a,
	0x05, 0x4d, 0x79, 0x53, 0x74, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x74, 0x72, 0x22, 0x21, 0x0a, 0x05, 0x4d, 0x79, 0x53, 0x75,
	0x62, 0x12, 0x18, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x4d, 0x79, 0x53, 0x74, 0x72, 0x52, 0x03, 0x73, 0x75, 0x62, 0x22, 0x30, 0x0a, 0x0a, 0x4d,
	0x79, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x33, 0x32,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x33, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x74, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x73, 0x74, 0x72, 0x42, 0x10, 0x5a,
	0x0e, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_example_proto_rawDescOnce sync.Once
	file_pb_example_proto_rawDescData = file_pb_example_proto_rawDesc
)

func file_pb_example_proto_rawDescGZIP() []byte {
	file_pb_example_proto_rawDescOnce.Do(func() {
		file_pb_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_example_proto_rawDescData)
	})
	return file_pb_example_proto_rawDescData
}

var file_pb_example_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_example_proto_goTypes = []any{
	(*MyU32)(nil),      // 0: MyU32
	(*MyStr)(nil),      // 1: MyStr
	(*MySub)(nil),      // 2: MySub
	(*MyRepeated)(nil), // 3: MyRepeated
}
var file_pb_example_proto_depIdxs = []int32{
	1, // 0: MySub.sub:type_name -> MyStr
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_example_proto_init() }
func file_pb_example_proto_init() {
	if File_pb_example_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_example_proto_goTypes,
		DependencyIndexes: file_pb_example_proto_depIdxs,
		MessageInfos:      file_pb_example_proto_msgTypes,
	}.Build()
	File_pb_example_proto = out.File
	file_pb_example_proto_rawDesc = nil
	file_pb_example_proto_goTypes = nil
	file_pb_example_proto_depIdxs = nil
}
