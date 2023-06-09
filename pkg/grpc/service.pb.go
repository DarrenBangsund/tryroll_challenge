// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/grpc/service.proto

package grpc

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

type ERC20Tx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To     string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	From   string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Amount string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Hash   string `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *ERC20Tx) Reset() {
	*x = ERC20Tx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ERC20Tx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ERC20Tx) ProtoMessage() {}

func (x *ERC20Tx) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ERC20Tx.ProtoReflect.Descriptor instead.
func (*ERC20Tx) Descriptor() ([]byte, []int) {
	return file_api_grpc_service_proto_rawDescGZIP(), []int{0}
}

func (x *ERC20Tx) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *ERC20Tx) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *ERC20Tx) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *ERC20Tx) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type GetERC20TxReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	To     string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	From   string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	Limit  string `protobuf:"bytes,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset string `protobuf:"bytes,5,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetERC20TxReq) Reset() {
	*x = GetERC20TxReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetERC20TxReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetERC20TxReq) ProtoMessage() {}

func (x *GetERC20TxReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetERC20TxReq.ProtoReflect.Descriptor instead.
func (*GetERC20TxReq) Descriptor() ([]byte, []int) {
	return file_api_grpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetERC20TxReq) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *GetERC20TxReq) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *GetERC20TxReq) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *GetERC20TxReq) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *GetERC20TxReq) GetOffset() string {
	if x != nil {
		return x.Offset
	}
	return ""
}

type GetERC20TxRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txs []*ERC20Tx `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (x *GetERC20TxRes) Reset() {
	*x = GetERC20TxRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetERC20TxRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetERC20TxRes) ProtoMessage() {}

func (x *GetERC20TxRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetERC20TxRes.ProtoReflect.Descriptor instead.
func (*GetERC20TxRes) Descriptor() ([]byte, []int) {
	return file_api_grpc_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetERC20TxRes) GetTxs() []*ERC20Tx {
	if x != nil {
		return x.Txs
	}
	return nil
}

type GetERC20TokenTxWithToAndFromReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To   string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *GetERC20TokenTxWithToAndFromReq) Reset() {
	*x = GetERC20TokenTxWithToAndFromReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetERC20TokenTxWithToAndFromReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetERC20TokenTxWithToAndFromReq) ProtoMessage() {}

func (x *GetERC20TokenTxWithToAndFromReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetERC20TokenTxWithToAndFromReq.ProtoReflect.Descriptor instead.
func (*GetERC20TokenTxWithToAndFromReq) Descriptor() ([]byte, []int) {
	return file_api_grpc_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetERC20TokenTxWithToAndFromReq) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *GetERC20TokenTxWithToAndFromReq) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

type GetERC20TokenTxWithToAndFromRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txs []*ERC20Tx `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (x *GetERC20TokenTxWithToAndFromRes) Reset() {
	*x = GetERC20TokenTxWithToAndFromRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetERC20TokenTxWithToAndFromRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetERC20TokenTxWithToAndFromRes) ProtoMessage() {}

func (x *GetERC20TokenTxWithToAndFromRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetERC20TokenTxWithToAndFromRes.ProtoReflect.Descriptor instead.
func (*GetERC20TokenTxWithToAndFromRes) Descriptor() ([]byte, []int) {
	return file_api_grpc_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetERC20TokenTxWithToAndFromRes) GetTxs() []*ERC20Tx {
	if x != nil {
		return x.Txs
	}
	return nil
}

var File_api_grpc_service_proto protoreflect.FileDescriptor

var file_api_grpc_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x59,
	0x0a, 0x07, 0x45, 0x52, 0x43, 0x32, 0x30, 0x54, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x79, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x45, 0x52, 0x43, 0x32, 0x30, 0x54, 0x78, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x22, 0x30, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x45, 0x52, 0x43, 0x32, 0x30,
	0x54, 0x78, 0x52, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x52, 0x43, 0x32, 0x30, 0x54,
	0x78, 0x52, 0x03, 0x74, 0x78, 0x73, 0x22, 0x45, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x45, 0x52, 0x43,
	0x32, 0x30, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x78, 0x57, 0x69, 0x74, 0x68, 0x54, 0x6f, 0x41,
	0x6e, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x22, 0x42, 0x0a,
	0x1f, 0x47, 0x65, 0x74, 0x45, 0x52, 0x43, 0x32, 0x30, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x78,
	0x57, 0x69, 0x74, 0x68, 0x54, 0x6f, 0x41, 0x6e, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x12, 0x1f, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x52, 0x43, 0x32, 0x30, 0x54, 0x78, 0x52, 0x03, 0x74, 0x78,
	0x73, 0x32, 0x48, 0x0a, 0x0e, 0x54, 0x72, 0x79, 0x52, 0x6f, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x45, 0x52, 0x43, 0x32, 0x30, 0x54,
	0x78, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x52, 0x43, 0x32,
	0x30, 0x54, 0x78, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x52, 0x43, 0x32, 0x30, 0x54, 0x78, 0x52, 0x65, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x70,
	0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_service_proto_rawDescOnce sync.Once
	file_api_grpc_service_proto_rawDescData = file_api_grpc_service_proto_rawDesc
)

func file_api_grpc_service_proto_rawDescGZIP() []byte {
	file_api_grpc_service_proto_rawDescOnce.Do(func() {
		file_api_grpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_service_proto_rawDescData)
	})
	return file_api_grpc_service_proto_rawDescData
}

var file_api_grpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_grpc_service_proto_goTypes = []interface{}{
	(*ERC20Tx)(nil),                         // 0: grpc.ERC20Tx
	(*GetERC20TxReq)(nil),                   // 1: grpc.GetERC20TxReq
	(*GetERC20TxRes)(nil),                   // 2: grpc.GetERC20TxRes
	(*GetERC20TokenTxWithToAndFromReq)(nil), // 3: grpc.GetERC20TokenTxWithToAndFromReq
	(*GetERC20TokenTxWithToAndFromRes)(nil), // 4: grpc.GetERC20TokenTxWithToAndFromRes
}
var file_api_grpc_service_proto_depIdxs = []int32{
	0, // 0: grpc.GetERC20TxRes.txs:type_name -> grpc.ERC20Tx
	0, // 1: grpc.GetERC20TokenTxWithToAndFromRes.txs:type_name -> grpc.ERC20Tx
	1, // 2: grpc.TryRollService.GetERC20Tx:input_type -> grpc.GetERC20TxReq
	2, // 3: grpc.TryRollService.GetERC20Tx:output_type -> grpc.GetERC20TxRes
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_grpc_service_proto_init() }
func file_api_grpc_service_proto_init() {
	if File_api_grpc_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ERC20Tx); i {
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
		file_api_grpc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetERC20TxReq); i {
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
		file_api_grpc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetERC20TxRes); i {
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
		file_api_grpc_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetERC20TokenTxWithToAndFromReq); i {
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
		file_api_grpc_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetERC20TokenTxWithToAndFromRes); i {
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
			RawDescriptor: file_api_grpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_service_proto_goTypes,
		DependencyIndexes: file_api_grpc_service_proto_depIdxs,
		MessageInfos:      file_api_grpc_service_proto_msgTypes,
	}.Build()
	File_api_grpc_service_proto = out.File
	file_api_grpc_service_proto_rawDesc = nil
	file_api_grpc_service_proto_goTypes = nil
	file_api_grpc_service_proto_depIdxs = nil
}
