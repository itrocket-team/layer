// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: layer/daemons/median_values.proto

package types

import (
	_ "github.com/cosmos/gogoproto/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// MedianValues is the median value for a market
type MedianValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// marketId is the market id for a pair
	MarketId uint32 `protobuf:"varint,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	Price    uint64 `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	Exponent int32  `protobuf:"varint,3,opt,name=exponent,proto3" json:"exponent,omitempty"`
}

func (x *MedianValues) Reset() {
	*x = MedianValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer_daemons_median_values_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MedianValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MedianValues) ProtoMessage() {}

func (x *MedianValues) ProtoReflect() protoreflect.Message {
	mi := &file_layer_daemons_median_values_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MedianValues.ProtoReflect.Descriptor instead.
func (*MedianValues) Descriptor() ([]byte, []int) {
	return file_layer_daemons_median_values_proto_rawDescGZIP(), []int{0}
}

func (x *MedianValues) GetMarketId() uint32 {
	if x != nil {
		return x.MarketId
	}
	return 0
}

func (x *MedianValues) GetPrice() uint64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *MedianValues) GetExponent() int32 {
	if x != nil {
		return x.Exponent
	}
	return 0
}

// GetAllMedianValuesRequest is the request for the GetAllMedianValues rpc
type GetAllMedianValuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllMedianValuesRequest) Reset() {
	*x = GetAllMedianValuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer_daemons_median_values_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMedianValuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMedianValuesRequest) ProtoMessage() {}

func (x *GetAllMedianValuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_layer_daemons_median_values_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMedianValuesRequest.ProtoReflect.Descriptor instead.
func (*GetAllMedianValuesRequest) Descriptor() ([]byte, []int) {
	return file_layer_daemons_median_values_proto_rawDescGZIP(), []int{1}
}

// GetAllMedianValuesResponse is the response for the GetAllMedianValues rpc
type GetAllMedianValuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MedianValues []*MedianValues `protobuf:"bytes,1,rep,name=median_values,json=medianValues,proto3" json:"median_values,omitempty"`
}

func (x *GetAllMedianValuesResponse) Reset() {
	*x = GetAllMedianValuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer_daemons_median_values_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMedianValuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMedianValuesResponse) ProtoMessage() {}

func (x *GetAllMedianValuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_layer_daemons_median_values_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMedianValuesResponse.ProtoReflect.Descriptor instead.
func (*GetAllMedianValuesResponse) Descriptor() ([]byte, []int) {
	return file_layer_daemons_median_values_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllMedianValuesResponse) GetMedianValues() []*MedianValues {
	if x != nil {
		return x.MedianValues
	}
	return nil
}

// GetMedianValuesRequest is the request for the GetMedianValues rpc
type GetMedianValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// query data to fetch prices for
	QueryData []byte `protobuf:"bytes,1,opt,name=query_data,json=queryData,proto3" json:"query_data,omitempty"`
}

func (x *GetMedianValueRequest) Reset() {
	*x = GetMedianValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer_daemons_median_values_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMedianValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMedianValueRequest) ProtoMessage() {}

func (x *GetMedianValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_layer_daemons_median_values_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMedianValueRequest.ProtoReflect.Descriptor instead.
func (*GetMedianValueRequest) Descriptor() ([]byte, []int) {
	return file_layer_daemons_median_values_proto_rawDescGZIP(), []int{3}
}

func (x *GetMedianValueRequest) GetQueryData() []byte {
	if x != nil {
		return x.QueryData
	}
	return nil
}

// GetMedianValueResponse is the response for the GetMedianValue rpc
type GetMedianValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MedianValues *MedianValues `protobuf:"bytes,1,opt,name=median_values,json=medianValues,proto3" json:"median_values,omitempty"`
}

func (x *GetMedianValueResponse) Reset() {
	*x = GetMedianValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer_daemons_median_values_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMedianValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMedianValueResponse) ProtoMessage() {}

func (x *GetMedianValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_layer_daemons_median_values_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMedianValueResponse.ProtoReflect.Descriptor instead.
func (*GetMedianValueResponse) Descriptor() ([]byte, []int) {
	return file_layer_daemons_median_values_proto_rawDescGZIP(), []int{4}
}

func (x *GetMedianValueResponse) GetMedianValues() *MedianValues {
	if x != nil {
		return x.MedianValues
	}
	return nil
}

var File_layer_daemons_median_values_proto protoreflect.FileDescriptor

var file_layer_daemons_median_values_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x73, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x5e, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x40, 0x0a, 0x0d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x52, 0x0c, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x22, 0x36, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x22, 0x5a, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x0c, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x6e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x32, 0xbd, 0x02, 0x0a, 0x13, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x93,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x64, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x22, 0x12, 0x20, 0x2f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x65, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x12, 0x8f, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x24, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28, 0x2f, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2f, 0x7b, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x7d, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6c, 0x6c, 0x6f, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_layer_daemons_median_values_proto_rawDescOnce sync.Once
	file_layer_daemons_median_values_proto_rawDescData = file_layer_daemons_median_values_proto_rawDesc
)

func file_layer_daemons_median_values_proto_rawDescGZIP() []byte {
	file_layer_daemons_median_values_proto_rawDescOnce.Do(func() {
		file_layer_daemons_median_values_proto_rawDescData = protoimpl.X.CompressGZIP(file_layer_daemons_median_values_proto_rawDescData)
	})
	return file_layer_daemons_median_values_proto_rawDescData
}

var file_layer_daemons_median_values_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_layer_daemons_median_values_proto_goTypes = []interface{}{
	(*MedianValues)(nil),               // 0: layer.daemons.MedianValues
	(*GetAllMedianValuesRequest)(nil),  // 1: layer.daemons.GetAllMedianValuesRequest
	(*GetAllMedianValuesResponse)(nil), // 2: layer.daemons.GetAllMedianValuesResponse
	(*GetMedianValueRequest)(nil),      // 3: layer.daemons.GetMedianValueRequest
	(*GetMedianValueResponse)(nil),     // 4: layer.daemons.GetMedianValueResponse
}
var file_layer_daemons_median_values_proto_depIdxs = []int32{
	0, // 0: layer.daemons.GetAllMedianValuesResponse.median_values:type_name -> layer.daemons.MedianValues
	0, // 1: layer.daemons.GetMedianValueResponse.median_values:type_name -> layer.daemons.MedianValues
	1, // 2: layer.daemons.MedianValuesService.GetAllMedianValues:input_type -> layer.daemons.GetAllMedianValuesRequest
	3, // 3: layer.daemons.MedianValuesService.GetMedianValue:input_type -> layer.daemons.GetMedianValueRequest
	2, // 4: layer.daemons.MedianValuesService.GetAllMedianValues:output_type -> layer.daemons.GetAllMedianValuesResponse
	4, // 5: layer.daemons.MedianValuesService.GetMedianValue:output_type -> layer.daemons.GetMedianValueResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_layer_daemons_median_values_proto_init() }
func file_layer_daemons_median_values_proto_init() {
	if File_layer_daemons_median_values_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_layer_daemons_median_values_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MedianValues); i {
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
		file_layer_daemons_median_values_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMedianValuesRequest); i {
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
		file_layer_daemons_median_values_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMedianValuesResponse); i {
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
		file_layer_daemons_median_values_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMedianValueRequest); i {
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
		file_layer_daemons_median_values_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMedianValueResponse); i {
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
			RawDescriptor: file_layer_daemons_median_values_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_layer_daemons_median_values_proto_goTypes,
		DependencyIndexes: file_layer_daemons_median_values_proto_depIdxs,
		MessageInfos:      file_layer_daemons_median_values_proto_msgTypes,
	}.Build()
	File_layer_daemons_median_values_proto = out.File
	file_layer_daemons_median_values_proto_rawDesc = nil
	file_layer_daemons_median_values_proto_goTypes = nil
	file_layer_daemons_median_values_proto_depIdxs = nil
}
