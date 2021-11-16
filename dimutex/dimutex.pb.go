// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: DiMutex.proto

package dimutex

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Lamport int32  `protobuf:"varint,2,opt,name=Lamport,proto3" json:"Lamport,omitempty"`
	Id      int32  `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DiMutex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_DiMutex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_DiMutex_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Request) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

func (x *Request) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Lamport int32  `protobuf:"varint,2,opt,name=Lamport,proto3" json:"Lamport,omitempty"`
	Id      int32  `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DiMutex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_DiMutex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_DiMutex_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Reply) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

func (x *Reply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DiMutex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_DiMutex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_DiMutex_proto_rawDescGZIP(), []int{2}
}

var File_DiMutex_proto protoreflect.FileDescriptor

var file_DiMutex_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x44, 0x69, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x64, 0x69, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x22, 0x4d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x61,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4c, 0x61, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x49, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xfd, 0x01,
	0x0a, 0x07, 0x44, 0x69, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x12, 0x33, 0x0a, 0x0d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x2e, 0x64, 0x69, 0x6d,
	0x75, 0x74, 0x65, 0x78, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x64,
	0x69, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x33,
	0x0a, 0x0d, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x2e, 0x64, 0x69, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x64, 0x69, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0e, 0x48, 0x6f, 0x6c, 0x64, 0x41, 0x6e, 0x64, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x2e, 0x64, 0x69, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x64, 0x69, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x05, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x12, 0x0e, 0x2e, 0x64, 0x69, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x1a, 0x0e, 0x2e, 0x64, 0x69, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x29, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0e, 0x2e, 0x64, 0x69,
	0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x64, 0x69,
	0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x35, 0x5a,
	0x33, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x65, 0x70, 0x70, 0x65, 0x6e, 0x2f, 0x64, 0x69, 0x73, 0x79,
	0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x32, 0x3b, 0x64, 0x69, 0x6d,
	0x75, 0x74, 0x65, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DiMutex_proto_rawDescOnce sync.Once
	file_DiMutex_proto_rawDescData = file_DiMutex_proto_rawDesc
)

func file_DiMutex_proto_rawDescGZIP() []byte {
	file_DiMutex_proto_rawDescOnce.Do(func() {
		file_DiMutex_proto_rawDescData = protoimpl.X.CompressGZIP(file_DiMutex_proto_rawDescData)
	})
	return file_DiMutex_proto_rawDescData
}

var file_DiMutex_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_DiMutex_proto_goTypes = []interface{}{
	(*Request)(nil), // 0: dimutex.Request
	(*Reply)(nil),   // 1: dimutex.Reply
	(*Empty)(nil),   // 2: dimutex.Empty
}
var file_DiMutex_proto_depIdxs = []int32{
	0, // 0: dimutex.DiMutex.RequestAccess:input_type -> dimutex.Request
	0, // 1: dimutex.DiMutex.AnswerRequest:input_type -> dimutex.Request
	2, // 2: dimutex.DiMutex.HoldAndRelease:input_type -> dimutex.Empty
	1, // 3: dimutex.DiMutex.Grant:input_type -> dimutex.Reply
	2, // 4: dimutex.DiMutex.Hello:input_type -> dimutex.Empty
	2, // 5: dimutex.DiMutex.RequestAccess:output_type -> dimutex.Empty
	1, // 6: dimutex.DiMutex.AnswerRequest:output_type -> dimutex.Reply
	1, // 7: dimutex.DiMutex.HoldAndRelease:output_type -> dimutex.Reply
	2, // 8: dimutex.DiMutex.Grant:output_type -> dimutex.Empty
	2, // 9: dimutex.DiMutex.Hello:output_type -> dimutex.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DiMutex_proto_init() }
func file_DiMutex_proto_init() {
	if File_DiMutex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DiMutex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_DiMutex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_DiMutex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_DiMutex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_DiMutex_proto_goTypes,
		DependencyIndexes: file_DiMutex_proto_depIdxs,
		MessageInfos:      file_DiMutex_proto_msgTypes,
	}.Build()
	File_DiMutex_proto = out.File
	file_DiMutex_proto_rawDesc = nil
	file_DiMutex_proto_goTypes = nil
	file_DiMutex_proto_depIdxs = nil
}
