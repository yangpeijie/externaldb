// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: push.proto

package proto

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

// push v1, register request
type BlockSeqCB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	URL           string `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`
	Encode        string `protobuf:"bytes,3,opt,name=encode,proto3" json:"encode,omitempty"`
	IsHeader      bool   `protobuf:"varint,4,opt,name=isHeader,proto3" json:"isHeader,omitempty"`
	LastSequence  int64  `protobuf:"varint,5,opt,name=lastSequence,proto3" json:"lastSequence,omitempty"`
	LastHeight    int64  `protobuf:"varint,6,opt,name=lastHeight,proto3" json:"lastHeight,omitempty"`
	LastBlockHash string `protobuf:"bytes,7,opt,name=lastBlockHash,proto3" json:"lastBlockHash,omitempty"`
}

func (x *BlockSeqCB) Reset() {
	*x = BlockSeqCB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockSeqCB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockSeqCB) ProtoMessage() {}

func (x *BlockSeqCB) ProtoReflect() protoreflect.Message {
	mi := &file_push_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockSeqCB.ProtoReflect.Descriptor instead.
func (*BlockSeqCB) Descriptor() ([]byte, []int) {
	return file_push_proto_rawDescGZIP(), []int{0}
}

func (x *BlockSeqCB) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BlockSeqCB) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *BlockSeqCB) GetEncode() string {
	if x != nil {
		return x.Encode
	}
	return ""
}

func (x *BlockSeqCB) GetIsHeader() bool {
	if x != nil {
		return x.IsHeader
	}
	return false
}

func (x *BlockSeqCB) GetLastSequence() int64 {
	if x != nil {
		return x.LastSequence
	}
	return 0
}

func (x *BlockSeqCB) GetLastHeight() int64 {
	if x != nil {
		return x.LastHeight
	}
	return 0
}

func (x *BlockSeqCB) GetLastBlockHash() string {
	if x != nil {
		return x.LastBlockHash
	}
	return ""
}

// push v2, register responce
type ReplySubscribePushV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOk bool   `protobuf:"varint,1,opt,name=isOk,proto3" json:"isOk,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ReplySubscribePushV2) Reset() {
	*x = ReplySubscribePushV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplySubscribePushV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplySubscribePushV2) ProtoMessage() {}

func (x *ReplySubscribePushV2) ProtoReflect() protoreflect.Message {
	mi := &file_push_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplySubscribePushV2.ProtoReflect.Descriptor instead.
func (*ReplySubscribePushV2) Descriptor() ([]byte, []int) {
	return file_push_proto_rawDescGZIP(), []int{1}
}

func (x *ReplySubscribePushV2) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

func (x *ReplySubscribePushV2) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// push v2, register request
type PushSubscribeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	URL           string `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`
	Encode        string `protobuf:"bytes,3,opt,name=encode,proto3" json:"encode,omitempty"`
	LastSequence  int64  `protobuf:"varint,4,opt,name=lastSequence,proto3" json:"lastSequence,omitempty"`
	LastHeight    int64  `protobuf:"varint,5,opt,name=lastHeight,proto3" json:"lastHeight,omitempty"`
	LastBlockHash string `protobuf:"bytes,6,opt,name=lastBlockHash,proto3" json:"lastBlockHash,omitempty"`
	// 0:代表区块；1:代表区块头信息；2：代表交易回执
	Type int32 `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"`
	//允许订阅多个类型的交易回执
	Contract map[string]bool `protobuf:"bytes,8,rep,name=contract,proto3" json:"contract,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *PushSubscribeReq) Reset() {
	*x = PushSubscribeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushSubscribeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushSubscribeReq) ProtoMessage() {}

func (x *PushSubscribeReq) ProtoReflect() protoreflect.Message {
	mi := &file_push_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushSubscribeReq.ProtoReflect.Descriptor instead.
func (*PushSubscribeReq) Descriptor() ([]byte, []int) {
	return file_push_proto_rawDescGZIP(), []int{2}
}

func (x *PushSubscribeReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PushSubscribeReq) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *PushSubscribeReq) GetEncode() string {
	if x != nil {
		return x.Encode
	}
	return ""
}

func (x *PushSubscribeReq) GetLastSequence() int64 {
	if x != nil {
		return x.LastSequence
	}
	return 0
}

func (x *PushSubscribeReq) GetLastHeight() int64 {
	if x != nil {
		return x.LastHeight
	}
	return 0
}

func (x *PushSubscribeReq) GetLastBlockHash() string {
	if x != nil {
		return x.LastBlockHash
	}
	return ""
}

func (x *PushSubscribeReq) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *PushSubscribeReq) GetContract() map[string]bool {
	if x != nil {
		return x.Contract
	}
	return nil
}

var File_push_proto protoreflect.FileDescriptor

var file_push_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x71,
	0x43, 0x42, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x22, 0x3c, 0x0a, 0x14, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x75, 0x73, 0x68, 0x56, 0x32, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73,
	0x4f, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0xce, 0x02, 0x0a, 0x10, 0x50, 0x75, 0x73, 0x68, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6c,
	0x61, 0x73, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x6c, 0x61, 0x73, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6c,
	0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x75, 0x73, 0x68, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x1a, 0x3b, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_push_proto_rawDescOnce sync.Once
	file_push_proto_rawDescData = file_push_proto_rawDesc
)

func file_push_proto_rawDescGZIP() []byte {
	file_push_proto_rawDescOnce.Do(func() {
		file_push_proto_rawDescData = protoimpl.X.CompressGZIP(file_push_proto_rawDescData)
	})
	return file_push_proto_rawDescData
}

var file_push_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_push_proto_goTypes = []interface{}{
	(*BlockSeqCB)(nil),           // 0: proto.BlockSeqCB
	(*ReplySubscribePushV2)(nil), // 1: proto.ReplySubscribePushV2
	(*PushSubscribeReq)(nil),     // 2: proto.PushSubscribeReq
	nil,                          // 3: proto.PushSubscribeReq.ContractEntry
}
var file_push_proto_depIdxs = []int32{
	3, // 0: proto.PushSubscribeReq.contract:type_name -> proto.PushSubscribeReq.ContractEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_push_proto_init() }
func file_push_proto_init() {
	if File_push_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_push_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockSeqCB); i {
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
		file_push_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplySubscribePushV2); i {
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
		file_push_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushSubscribeReq); i {
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
			RawDescriptor: file_push_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_push_proto_goTypes,
		DependencyIndexes: file_push_proto_depIdxs,
		MessageInfos:      file_push_proto_msgTypes,
	}.Build()
	File_push_proto = out.File
	file_push_proto_rawDesc = nil
	file_push_proto_goTypes = nil
	file_push_proto_depIdxs = nil
}
