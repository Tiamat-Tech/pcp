// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: p2p.proto

package proto

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

// A message object that is shared among all requests.
type MessageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unix time in milliseconds.
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The ID of the node that created the message (not the peer that may have sent it)
	NodeId string `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Authoring node Secp256k1 public key (32bytes) - protobufs serialized
	NodePubKey []byte `protobuf:"bytes,3,opt,name=node_pub_key,json=nodePubKey,proto3" json:"node_pub_key,omitempty"`
	// The signature of the message data.
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	// The actual message content.
	//
	// Types that are assignable to Payload:
	//	*MessageData_SendRequest
	//	*MessageData_SendResponse
	Payload isMessageData_Payload `protobuf_oneof:"payload"`
}

func (x *MessageData) Reset() {
	*x = MessageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageData) ProtoMessage() {}

func (x *MessageData) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageData.ProtoReflect.Descriptor instead.
func (*MessageData) Descriptor() ([]byte, []int) {
	return file_p2p_proto_rawDescGZIP(), []int{0}
}

func (x *MessageData) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *MessageData) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *MessageData) GetNodePubKey() []byte {
	if x != nil {
		return x.NodePubKey
	}
	return nil
}

func (x *MessageData) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (m *MessageData) GetPayload() isMessageData_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *MessageData) GetSendRequest() *SendRequest {
	if x, ok := x.GetPayload().(*MessageData_SendRequest); ok {
		return x.SendRequest
	}
	return nil
}

func (x *MessageData) GetSendResponse() *SendResponse {
	if x, ok := x.GetPayload().(*MessageData_SendResponse); ok {
		return x.SendResponse
	}
	return nil
}

type isMessageData_Payload interface {
	isMessageData_Payload()
}

type MessageData_SendRequest struct {
	SendRequest *SendRequest `protobuf:"bytes,9,opt,name=send_request,json=sendRequest,proto3,oneof"`
}

type MessageData_SendResponse struct {
	SendResponse *SendResponse `protobuf:"bytes,10,opt,name=send_response,json=sendResponse,proto3,oneof"`
}

func (*MessageData_SendRequest) isMessageData_Payload() {}

func (*MessageData_SendResponse) isMessageData_Payload() {}

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileSize int64  `protobuf:"varint,2,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_p2p_proto_rawDescGZIP(), []int{1}
}

func (x *SendRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *SendRequest) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accepted bool `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_p2p_proto_rawDescGZIP(), []int{2}
}

func (x *SendResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

var File_p2p_proto protoreflect.FileDescriptor

var file_p2p_proto_rawDesc = []byte{
	0x0a, 0x09, 0x70, 0x32, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x0b,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x75, 0x62, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x75,
	0x62, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x31, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x73,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x47, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0x2a, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x42, 0x25, 0x5a, 0x23, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6e, 0x6e, 0x69, 0x73,
	0x2d, 0x74, 0x72, 0x61, 0x2f, 0x70, 0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_p2p_proto_rawDescOnce sync.Once
	file_p2p_proto_rawDescData = file_p2p_proto_rawDesc
)

func file_p2p_proto_rawDescGZIP() []byte {
	file_p2p_proto_rawDescOnce.Do(func() {
		file_p2p_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2p_proto_rawDescData)
	})
	return file_p2p_proto_rawDescData
}

var file_p2p_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_p2p_proto_goTypes = []interface{}{
	(*MessageData)(nil),  // 0: MessageData
	(*SendRequest)(nil),  // 1: SendRequest
	(*SendResponse)(nil), // 2: SendResponse
}
var file_p2p_proto_depIdxs = []int32{
	1, // 0: MessageData.send_request:type_name -> SendRequest
	2, // 1: MessageData.send_response:type_name -> SendResponse
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_p2p_proto_init() }
func file_p2p_proto_init() {
	if File_p2p_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_p2p_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageData); i {
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
		file_p2p_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_p2p_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
	file_p2p_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MessageData_SendRequest)(nil),
		(*MessageData_SendResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_p2p_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_p2p_proto_goTypes,
		DependencyIndexes: file_p2p_proto_depIdxs,
		MessageInfos:      file_p2p_proto_msgTypes,
	}.Build()
	File_p2p_proto = out.File
	file_p2p_proto_rawDesc = nil
	file_p2p_proto_goTypes = nil
	file_p2p_proto_depIdxs = nil
}