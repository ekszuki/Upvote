// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: pkg/protos/coins.proto

package coins

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateCoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=Description,proto3" json:"Description,omitempty"`
	Short       string `protobuf:"bytes,2,opt,name=Short,proto3" json:"Short,omitempty"`
}

func (x *CreateCoinRequest) Reset() {
	*x = CreateCoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_coins_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoinRequest) ProtoMessage() {}

func (x *CreateCoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_coins_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoinRequest.ProtoReflect.Descriptor instead.
func (*CreateCoinRequest) Descriptor() ([]byte, []int) {
	return file_pkg_protos_coins_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCoinRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateCoinRequest) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

type CreateCoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	Short       string `protobuf:"bytes,3,opt,name=Short,proto3" json:"Short,omitempty"`
}

func (x *CreateCoinResponse) Reset() {
	*x = CreateCoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_coins_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoinResponse) ProtoMessage() {}

func (x *CreateCoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_coins_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoinResponse.ProtoReflect.Descriptor instead.
func (*CreateCoinResponse) Descriptor() ([]byte, []int) {
	return file_pkg_protos_coins_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCoinResponse) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CreateCoinResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateCoinResponse) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

type DeleteCoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteCoinRequest) Reset() {
	*x = DeleteCoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_coins_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCoinRequest) ProtoMessage() {}

func (x *DeleteCoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_coins_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCoinRequest.ProtoReflect.Descriptor instead.
func (*DeleteCoinRequest) Descriptor() ([]byte, []int) {
	return file_pkg_protos_coins_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteCoinRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ActiveCoinsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActiveCoins []*CreateCoinResponse `protobuf:"bytes,1,rep,name=ActiveCoins,proto3" json:"ActiveCoins,omitempty"`
}

func (x *ActiveCoinsResponse) Reset() {
	*x = ActiveCoinsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_coins_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveCoinsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveCoinsResponse) ProtoMessage() {}

func (x *ActiveCoinsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_coins_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveCoinsResponse.ProtoReflect.Descriptor instead.
func (*ActiveCoinsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_protos_coins_proto_rawDescGZIP(), []int{3}
}

func (x *ActiveCoinsResponse) GetActiveCoins() []*CreateCoinResponse {
	if x != nil {
		return x.ActiveCoins
	}
	return nil
}

type VoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinID int32 `protobuf:"varint,1,opt,name=CoinID,proto3" json:"CoinID,omitempty"`
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_coins_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_coins_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_pkg_protos_coins_proto_rawDescGZIP(), []int{4}
}

func (x *VoteRequest) GetCoinID() int32 {
	if x != nil {
		return x.CoinID
	}
	return 0
}

var File_pkg_protos_coins_proto protoreflect.FileDescriptor

var file_pkg_protos_coins_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x69,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x22, 0x5c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x49, 0x44, 0x22, 0x4c, 0x0a, 0x13, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43,
	0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f,
	0x69, 0x6e, 0x73, 0x22, 0x25, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x44, 0x32, 0xf7, 0x01, 0x0a, 0x0b, 0x43,
	0x6f, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x69,
	0x6e, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x69,
	0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x56, 0x6f, 0x74, 0x65, 0x55, 0x50, 0x12, 0x0c, 0x2e, 0x56,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_protos_coins_proto_rawDescOnce sync.Once
	file_pkg_protos_coins_proto_rawDescData = file_pkg_protos_coins_proto_rawDesc
)

func file_pkg_protos_coins_proto_rawDescGZIP() []byte {
	file_pkg_protos_coins_proto_rawDescOnce.Do(func() {
		file_pkg_protos_coins_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protos_coins_proto_rawDescData)
	})
	return file_pkg_protos_coins_proto_rawDescData
}

var file_pkg_protos_coins_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_protos_coins_proto_goTypes = []interface{}{
	(*CreateCoinRequest)(nil),   // 0: CreateCoinRequest
	(*CreateCoinResponse)(nil),  // 1: CreateCoinResponse
	(*DeleteCoinRequest)(nil),   // 2: DeleteCoinRequest
	(*ActiveCoinsResponse)(nil), // 3: ActiveCoinsResponse
	(*VoteRequest)(nil),         // 4: VoteRequest
	(*emptypb.Empty)(nil),       // 5: google.protobuf.Empty
}
var file_pkg_protos_coins_proto_depIdxs = []int32{
	1, // 0: ActiveCoinsResponse.ActiveCoins:type_name -> CreateCoinResponse
	0, // 1: CoinService.CreateCoin:input_type -> CreateCoinRequest
	2, // 2: CoinService.DeleteCoin:input_type -> DeleteCoinRequest
	5, // 3: CoinService.ListActiveCoins:input_type -> google.protobuf.Empty
	4, // 4: CoinService.VoteUP:input_type -> VoteRequest
	1, // 5: CoinService.CreateCoin:output_type -> CreateCoinResponse
	5, // 6: CoinService.DeleteCoin:output_type -> google.protobuf.Empty
	3, // 7: CoinService.ListActiveCoins:output_type -> ActiveCoinsResponse
	5, // 8: CoinService.VoteUP:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_protos_coins_proto_init() }
func file_pkg_protos_coins_proto_init() {
	if File_pkg_protos_coins_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protos_coins_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCoinRequest); i {
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
		file_pkg_protos_coins_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCoinResponse); i {
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
		file_pkg_protos_coins_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCoinRequest); i {
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
		file_pkg_protos_coins_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveCoinsResponse); i {
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
		file_pkg_protos_coins_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteRequest); i {
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
			RawDescriptor: file_pkg_protos_coins_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_protos_coins_proto_goTypes,
		DependencyIndexes: file_pkg_protos_coins_proto_depIdxs,
		MessageInfos:      file_pkg_protos_coins_proto_msgTypes,
	}.Build()
	File_pkg_protos_coins_proto = out.File
	file_pkg_protos_coins_proto_rawDesc = nil
	file_pkg_protos_coins_proto_goTypes = nil
	file_pkg_protos_coins_proto_depIdxs = nil
}
