// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: dimutex.proto

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

type AccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Lamport int32  `protobuf:"varint,2,opt,name=Lamport,proto3" json:"Lamport,omitempty"`
	Id      int32  `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"` //overvej denne
}

func (x *AccessRequest) Reset() {
	*x = AccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dimutex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequest) ProtoMessage() {}

func (x *AccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dimutex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequest.ProtoReflect.Descriptor instead.
func (*AccessRequest) Descriptor() ([]byte, []int) {
	return file_dimutex_proto_rawDescGZIP(), []int{0}
}

func (x *AccessRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AccessRequest) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

func (x *AccessRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

//to indicate whether a node was given access?
type AccessGrant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	AccessGranted bool  `protobuf:"varint,2,opt,name=AccessGranted,proto3" json:"AccessGranted,omitempty"`
}

func (x *AccessGrant) Reset() {
	*x = AccessGrant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dimutex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessGrant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessGrant) ProtoMessage() {}

func (x *AccessGrant) ProtoReflect() protoreflect.Message {
	mi := &file_dimutex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessGrant.ProtoReflect.Descriptor instead.
func (*AccessGrant) Descriptor() ([]byte, []int) {
	return file_dimutex_proto_rawDescGZIP(), []int{1}
}

func (x *AccessGrant) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccessGrant) GetAccessGranted() bool {
	if x != nil {
		return x.AccessGranted
	}
	return false
}

type ReleaseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`  //ikke nødvendigt, men lidt nice
	Lamport int32  `protobuf:"varint,2,opt,name=Lamport,proto3" json:"Lamport,omitempty"` //hvornår blev der anmodet om release/rent faktisk released
}

func (x *ReleaseMessage) Reset() {
	*x = ReleaseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dimutex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseMessage) ProtoMessage() {}

func (x *ReleaseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_dimutex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseMessage.ProtoReflect.Descriptor instead.
func (*ReleaseMessage) Descriptor() ([]byte, []int) {
	return file_dimutex_proto_rawDescGZIP(), []int{2}
}

func (x *ReleaseMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ReleaseMessage) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

type RequestAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestAnswer) Reset() {
	*x = RequestAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dimutex_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAnswer) ProtoMessage() {}

func (x *RequestAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_dimutex_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAnswer.ProtoReflect.Descriptor instead.
func (*RequestAnswer) Descriptor() ([]byte, []int) {
	return file_dimutex_proto_rawDescGZIP(), []int{3}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dimutex_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_dimutex_proto_msgTypes[4]
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
	return file_dimutex_proto_rawDescGZIP(), []int{4}
}

var File_dimutex_proto protoreflect.FileDescriptor

var file_dimutex_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x69, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x64, 0x69, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x22, 0x53, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x43, 0x0a,
	0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x65, 0x64, 0x22, 0x44, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x32, 0x86, 0x02, 0x0a, 0x07, 0x44, 0x69, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x12, 0x3f,
	0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x16, 0x2e, 0x64, 0x69, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x64, 0x69, 0x6d, 0x75, 0x74, 0x65,
	0x78, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x0d, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x2e, 0x64, 0x69, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x64, 0x69, 0x6d, 0x75, 0x74,
	0x65, 0x78, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x43, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x17, 0x2e, 0x64, 0x69, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x2e, 0x64,
	0x69, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0e, 0x48, 0x6f, 0x6c, 0x64, 0x41,
	0x6e, 0x64, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x2e, 0x64, 0x69, 0x6d, 0x75,
	0x74, 0x65, 0x78, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x64, 0x69, 0x6d, 0x75,
	0x74, 0x65, 0x78, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x61, 0x65, 0x70, 0x70, 0x65, 0x6e, 0x2f, 0x64, 0x69, 0x73, 0x79, 0x73, 0x2d,
	0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x32, 0x3b, 0x64, 0x69, 0x6d, 0x75, 0x74,
	0x65, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dimutex_proto_rawDescOnce sync.Once
	file_dimutex_proto_rawDescData = file_dimutex_proto_rawDesc
)

func file_dimutex_proto_rawDescGZIP() []byte {
	file_dimutex_proto_rawDescOnce.Do(func() {
		file_dimutex_proto_rawDescData = protoimpl.X.CompressGZIP(file_dimutex_proto_rawDescData)
	})
	return file_dimutex_proto_rawDescData
}

var file_dimutex_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_dimutex_proto_goTypes = []interface{}{
	(*AccessRequest)(nil),  // 0: dimutex.AccessRequest
	(*AccessGrant)(nil),    // 1: dimutex.AccessGrant
	(*ReleaseMessage)(nil), // 2: dimutex.ReleaseMessage
	(*RequestAnswer)(nil),  // 3: dimutex.RequestAnswer
	(*Empty)(nil),          // 4: dimutex.Empty
}
var file_dimutex_proto_depIdxs = []int32{
	0, // 0: dimutex.DiMutex.RequestAccess:input_type -> dimutex.AccessRequest
	0, // 1: dimutex.DiMutex.AnswerRequest:input_type -> dimutex.AccessRequest
	2, // 2: dimutex.DiMutex.ReleaseAccess:input_type -> dimutex.ReleaseMessage
	4, // 3: dimutex.DiMutex.HoldAndRelease:input_type -> dimutex.Empty
	1, // 4: dimutex.DiMutex.RequestAccess:output_type -> dimutex.AccessGrant
	3, // 5: dimutex.DiMutex.AnswerRequest:output_type -> dimutex.RequestAnswer
	2, // 6: dimutex.DiMutex.ReleaseAccess:output_type -> dimutex.ReleaseMessage
	4, // 7: dimutex.DiMutex.HoldAndRelease:output_type -> dimutex.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dimutex_proto_init() }
func file_dimutex_proto_init() {
	if File_dimutex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dimutex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessRequest); i {
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
		file_dimutex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessGrant); i {
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
		file_dimutex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseMessage); i {
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
		file_dimutex_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAnswer); i {
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
		file_dimutex_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_dimutex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dimutex_proto_goTypes,
		DependencyIndexes: file_dimutex_proto_depIdxs,
		MessageInfos:      file_dimutex_proto_msgTypes,
	}.Build()
	File_dimutex_proto = out.File
	file_dimutex_proto_rawDesc = nil
	file_dimutex_proto_goTypes = nil
	file_dimutex_proto_depIdxs = nil
}
