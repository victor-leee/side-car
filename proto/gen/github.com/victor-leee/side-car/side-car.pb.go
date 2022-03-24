// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.2
// source: side-car.proto

package side_car

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

type BaseResponse_Code int32

const (
	BaseResponse_CODE_SUCCESS BaseResponse_Code = 0
	BaseResponse_CODE_ERROR   BaseResponse_Code = 1
)

// Enum value maps for BaseResponse_Code.
var (
	BaseResponse_Code_name = map[int32]string{
		0: "CODE_SUCCESS",
		1: "CODE_ERROR",
	}
	BaseResponse_Code_value = map[string]int32{
		"CODE_SUCCESS": 0,
		"CODE_ERROR":   1,
	}
)

func (x BaseResponse_Code) Enum() *BaseResponse_Code {
	p := new(BaseResponse_Code)
	*p = x
	return p
}

func (x BaseResponse_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BaseResponse_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_side_car_proto_enumTypes[0].Descriptor()
}

func (BaseResponse_Code) Type() protoreflect.EnumType {
	return &file_side_car_proto_enumTypes[0]
}

func (x BaseResponse_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BaseResponse_Code.Descriptor instead.
func (BaseResponse_Code) EnumDescriptor() ([]byte, []int) {
	return file_side_car_proto_rawDescGZIP(), []int{0, 0}
}

type InitConnectionReq_ConnectionType int32

const (
	// APP_TO_CAR means the app sends request and block wait response from the connection
	InitConnectionReq_CONNECTION_TYPE_APP_TO_CAR InitConnectionReq_ConnectionType = 0
	// CAR_TO_APP means the side car sends request and block wait response from the apps
	InitConnectionReq_CONNECTION_TYPE_CAR_TO_APP InitConnectionReq_ConnectionType = 1
)

// Enum value maps for InitConnectionReq_ConnectionType.
var (
	InitConnectionReq_ConnectionType_name = map[int32]string{
		0: "CONNECTION_TYPE_APP_TO_CAR",
		1: "CONNECTION_TYPE_CAR_TO_APP",
	}
	InitConnectionReq_ConnectionType_value = map[string]int32{
		"CONNECTION_TYPE_APP_TO_CAR": 0,
		"CONNECTION_TYPE_CAR_TO_APP": 1,
	}
)

func (x InitConnectionReq_ConnectionType) Enum() *InitConnectionReq_ConnectionType {
	p := new(InitConnectionReq_ConnectionType)
	*p = x
	return p
}

func (x InitConnectionReq_ConnectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InitConnectionReq_ConnectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_side_car_proto_enumTypes[1].Descriptor()
}

func (InitConnectionReq_ConnectionType) Type() protoreflect.EnumType {
	return &file_side_car_proto_enumTypes[1]
}

func (x InitConnectionReq_ConnectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InitConnectionReq_ConnectionType.Descriptor instead.
func (InitConnectionReq_ConnectionType) EnumDescriptor() ([]byte, []int) {
	return file_side_car_proto_rawDescGZIP(), []int{2, 0}
}

type BaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    BaseResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=com.github.victor_leee.side_car.BaseResponse_Code" json:"code,omitempty"`
	Message string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_side_car_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_side_car_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_side_car_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResponse) GetCode() BaseResponse_Code {
	if x != nil {
		return x.Code
	}
	return BaseResponse_CODE_SUCCESS
}

func (x *BaseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// GetConfigReq is used to fetch config data from config center
type GetConfigReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *GetConfigReq) Reset() {
	*x = GetConfigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_side_car_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigReq) ProtoMessage() {}

func (x *GetConfigReq) ProtoReflect() protoreflect.Message {
	mi := &file_side_car_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigReq.ProtoReflect.Descriptor instead.
func (*GetConfigReq) Descriptor() ([]byte, []int) {
	return file_side_car_proto_rawDescGZIP(), []int{1}
}

func (x *GetConfigReq) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

// InitConnectionReq is sent from sdk when connections are established(for each connection)
// It is meant to make the side car capable to send requests from other side cars to internal apps
type InitConnectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionType InitConnectionReq_ConnectionType `protobuf:"varint,1,opt,name=connection_type,json=connectionType,proto3,enum=com.github.victor_leee.side_car.InitConnectionReq_ConnectionType" json:"connection_type,omitempty"`
}

func (x *InitConnectionReq) Reset() {
	*x = InitConnectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_side_car_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitConnectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitConnectionReq) ProtoMessage() {}

func (x *InitConnectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_side_car_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitConnectionReq.ProtoReflect.Descriptor instead.
func (*InitConnectionReq) Descriptor() ([]byte, []int) {
	return file_side_car_proto_rawDescGZIP(), []int{2}
}

func (x *InitConnectionReq) GetConnectionType() InitConnectionReq_ConnectionType {
	if x != nil {
		return x.ConnectionType
	}
	return InitConnectionReq_CONNECTION_TYPE_APP_TO_CAR
}

var File_side_car_proto protoreflect.FileDescriptor

var file_side_car_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x69, 0x64, 0x65, 0x2d, 0x63, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x69, 0x63,
	0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x65, 0x65, 0x65, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x63, 0x61,
	0x72, 0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x69,
	0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x65, 0x65, 0x65, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x63,
	0x61, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x0c,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x22, 0x31,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0xd1, 0x01, 0x0a, 0x11, 0x49, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x6a, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x69,
	0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x65, 0x65, 0x65, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x63,
	0x61, 0x72, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x50, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x54, 0x4f, 0x5f,
	0x43, 0x41, 0x52, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x41, 0x52, 0x5f, 0x54, 0x4f, 0x5f,
	0x41, 0x50, 0x50, 0x10, 0x01, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x6c, 0x65, 0x65, 0x65, 0x2f,
	0x73, 0x69, 0x64, 0x65, 0x2d, 0x63, 0x61, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_side_car_proto_rawDescOnce sync.Once
	file_side_car_proto_rawDescData = file_side_car_proto_rawDesc
)

func file_side_car_proto_rawDescGZIP() []byte {
	file_side_car_proto_rawDescOnce.Do(func() {
		file_side_car_proto_rawDescData = protoimpl.X.CompressGZIP(file_side_car_proto_rawDescData)
	})
	return file_side_car_proto_rawDescData
}

var file_side_car_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_side_car_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_side_car_proto_goTypes = []interface{}{
	(BaseResponse_Code)(0),                // 0: com.github.victor_leee.side_car.BaseResponse.Code
	(InitConnectionReq_ConnectionType)(0), // 1: com.github.victor_leee.side_car.InitConnectionReq.ConnectionType
	(*BaseResponse)(nil),                  // 2: com.github.victor_leee.side_car.BaseResponse
	(*GetConfigReq)(nil),                  // 3: com.github.victor_leee.side_car.GetConfigReq
	(*InitConnectionReq)(nil),             // 4: com.github.victor_leee.side_car.InitConnectionReq
}
var file_side_car_proto_depIdxs = []int32{
	0, // 0: com.github.victor_leee.side_car.BaseResponse.code:type_name -> com.github.victor_leee.side_car.BaseResponse.Code
	1, // 1: com.github.victor_leee.side_car.InitConnectionReq.connection_type:type_name -> com.github.victor_leee.side_car.InitConnectionReq.ConnectionType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_side_car_proto_init() }
func file_side_car_proto_init() {
	if File_side_car_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_side_car_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResponse); i {
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
		file_side_car_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigReq); i {
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
		file_side_car_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitConnectionReq); i {
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
			RawDescriptor: file_side_car_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_side_car_proto_goTypes,
		DependencyIndexes: file_side_car_proto_depIdxs,
		EnumInfos:         file_side_car_proto_enumTypes,
		MessageInfos:      file_side_car_proto_msgTypes,
	}.Build()
	File_side_car_proto = out.File
	file_side_car_proto_rawDesc = nil
	file_side_car_proto_goTypes = nil
	file_side_car_proto_depIdxs = nil
}
