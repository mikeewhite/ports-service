// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: ports_service/v1/ports.proto

package ports_service

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

type AddPortsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ports []*Port `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
}

func (x *AddPortsRequest) Reset() {
	*x = AddPortsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_service_v1_ports_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPortsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPortsRequest) ProtoMessage() {}

func (x *AddPortsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ports_service_v1_ports_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPortsRequest.ProtoReflect.Descriptor instead.
func (*AddPortsRequest) Descriptor() ([]byte, []int) {
	return file_ports_service_v1_ports_proto_rawDescGZIP(), []int{0}
}

func (x *AddPortsRequest) GetPorts() []*Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

type AddPortsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddPortsResponse) Reset() {
	*x = AddPortsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_service_v1_ports_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPortsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPortsResponse) ProtoMessage() {}

func (x *AddPortsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ports_service_v1_ports_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPortsResponse.ProtoReflect.Descriptor instead.
func (*AddPortsResponse) Descriptor() ([]byte, []int) {
	return file_ports_service_v1_ports_proto_rawDescGZIP(), []int{1}
}

type ListPortsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPortsRequest) Reset() {
	*x = ListPortsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_service_v1_ports_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPortsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPortsRequest) ProtoMessage() {}

func (x *ListPortsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ports_service_v1_ports_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPortsRequest.ProtoReflect.Descriptor instead.
func (*ListPortsRequest) Descriptor() ([]byte, []int) {
	return file_ports_service_v1_ports_proto_rawDescGZIP(), []int{2}
}

type ListPortsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ports []*Port `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
}

func (x *ListPortsResponse) Reset() {
	*x = ListPortsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_service_v1_ports_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPortsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPortsResponse) ProtoMessage() {}

func (x *ListPortsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ports_service_v1_ports_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPortsResponse.ProtoReflect.Descriptor instead.
func (*ListPortsResponse) Descriptor() ([]byte, []int) {
	return file_ports_service_v1_ports_proto_rawDescGZIP(), []int{3}
}

func (x *ListPortsResponse) GetPorts() []*Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City        string    `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country     string    `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Alias       []string  `protobuf:"bytes,5,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions     []string  `protobuf:"bytes,6,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates []float32 `protobuf:"fixed32,7,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province    string    `protobuf:"bytes,8,opt,name=province,proto3" json:"province,omitempty"`
	Timezone    string    `protobuf:"bytes,9,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs      []string  `protobuf:"bytes,10,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code        string    `protobuf:"bytes,11,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_service_v1_ports_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_ports_service_v1_ports_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_ports_service_v1_ports_proto_rawDescGZIP(), []int{4}
}

func (x *Port) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Port) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Port) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Port) GetAlias() []string {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *Port) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *Port) GetCoordinates() []float32 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *Port) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Port) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Port) GetUnlocs() []string {
	if x != nil {
		return x.Unlocs
	}
	return nil
}

func (x *Port) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_ports_service_v1_ports_proto protoreflect.FileDescriptor

var file_ports_service_v1_ports_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x22, 0x3f, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x22, 0x12, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x8e, 0x02, 0x0a,
	0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xbb, 0x01,
	0x0a, 0x0c, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53,
	0x0a, 0x08, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x21, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x73,
	0x12, 0x22, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x6b, 0x65, 0x65, 0x77,
	0x68, 0x69, 0x74, 0x65, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ports_service_v1_ports_proto_rawDescOnce sync.Once
	file_ports_service_v1_ports_proto_rawDescData = file_ports_service_v1_ports_proto_rawDesc
)

func file_ports_service_v1_ports_proto_rawDescGZIP() []byte {
	file_ports_service_v1_ports_proto_rawDescOnce.Do(func() {
		file_ports_service_v1_ports_proto_rawDescData = protoimpl.X.CompressGZIP(file_ports_service_v1_ports_proto_rawDescData)
	})
	return file_ports_service_v1_ports_proto_rawDescData
}

var file_ports_service_v1_ports_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ports_service_v1_ports_proto_goTypes = []interface{}{
	(*AddPortsRequest)(nil),   // 0: ports_service.v1.AddPortsRequest
	(*AddPortsResponse)(nil),  // 1: ports_service.v1.AddPortsResponse
	(*ListPortsRequest)(nil),  // 2: ports_service.v1.ListPortsRequest
	(*ListPortsResponse)(nil), // 3: ports_service.v1.ListPortsResponse
	(*Port)(nil),              // 4: ports_service.v1.Port
}
var file_ports_service_v1_ports_proto_depIdxs = []int32{
	4, // 0: ports_service.v1.AddPortsRequest.ports:type_name -> ports_service.v1.Port
	4, // 1: ports_service.v1.ListPortsResponse.ports:type_name -> ports_service.v1.Port
	0, // 2: ports_service.v1.PortsService.AddPorts:input_type -> ports_service.v1.AddPortsRequest
	2, // 3: ports_service.v1.PortsService.ListPorts:input_type -> ports_service.v1.ListPortsRequest
	1, // 4: ports_service.v1.PortsService.AddPorts:output_type -> ports_service.v1.AddPortsResponse
	3, // 5: ports_service.v1.PortsService.ListPorts:output_type -> ports_service.v1.ListPortsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ports_service_v1_ports_proto_init() }
func file_ports_service_v1_ports_proto_init() {
	if File_ports_service_v1_ports_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ports_service_v1_ports_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPortsRequest); i {
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
		file_ports_service_v1_ports_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPortsResponse); i {
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
		file_ports_service_v1_ports_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPortsRequest); i {
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
		file_ports_service_v1_ports_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPortsResponse); i {
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
		file_ports_service_v1_ports_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
			RawDescriptor: file_ports_service_v1_ports_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ports_service_v1_ports_proto_goTypes,
		DependencyIndexes: file_ports_service_v1_ports_proto_depIdxs,
		MessageInfos:      file_ports_service_v1_ports_proto_msgTypes,
	}.Build()
	File_ports_service_v1_ports_proto = out.File
	file_ports_service_v1_ports_proto_rawDesc = nil
	file_ports_service_v1_ports_proto_goTypes = nil
	file_ports_service_v1_ports_proto_depIdxs = nil
}