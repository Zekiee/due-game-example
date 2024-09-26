// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.20.3
// source: pb_rpsgame/pb_rpsgame.proto

package pb_rpsgame

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

type BotStatusReqStatusType int32

const (
	BotStatusReq_waiting    BotStatusReqStatusType = 0
	BotStatusReq_connecting BotStatusReqStatusType = 1
	BotStatusReq_close      BotStatusReqStatusType = 2
)

// Enum value maps for BotStatusReqStatusType.
var (
	BotStatusReqStatusType_name = map[int32]string{
		0: "waiting",
		1: "connecting",
		2: "close",
	}
	BotStatusReqStatusType_value = map[string]int32{
		"waiting":    0,
		"connecting": 1,
		"close":      2,
	}
)

func (x BotStatusReqStatusType) Enum() *BotStatusReqStatusType {
	p := new(BotStatusReqStatusType)
	*p = x
	return p
}

func (x BotStatusReqStatusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BotStatusReqStatusType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_rpsgame_pb_rpsgame_proto_enumTypes[0].Descriptor()
}

func (BotStatusReqStatusType) Type() protoreflect.EnumType {
	return &file_pb_rpsgame_pb_rpsgame_proto_enumTypes[0]
}

func (x BotStatusReqStatusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BotStatusReqStatusType.Descriptor instead.
func (BotStatusReqStatusType) EnumDescriptor() ([]byte, []int) {
	return file_pb_rpsgame_pb_rpsgame_proto_rawDescGZIP(), []int{1, 0}
}

type BotStatusReqGenderType int32

const (
	BotStatusReq_man   BotStatusReqGenderType = 0
	BotStatusReq_woman BotStatusReqGenderType = 1
)

// Enum value maps for BotStatusReqGenderType.
var (
	BotStatusReqGenderType_name = map[int32]string{
		0: "man",
		1: "woman",
	}
	BotStatusReqGenderType_value = map[string]int32{
		"man":   0,
		"woman": 1,
	}
)

func (x BotStatusReqGenderType) Enum() *BotStatusReqGenderType {
	p := new(BotStatusReqGenderType)
	*p = x
	return p
}

func (x BotStatusReqGenderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BotStatusReqGenderType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_rpsgame_pb_rpsgame_proto_enumTypes[1].Descriptor()
}

func (BotStatusReqGenderType) Type() protoreflect.EnumType {
	return &file_pb_rpsgame_pb_rpsgame_proto_enumTypes[1]
}

func (x BotStatusReqGenderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BotStatusReqGenderType.Descriptor instead.
func (BotStatusReqGenderType) EnumDescriptor() ([]byte, []int) {
	return file_pb_rpsgame_pb_rpsgame_proto_rawDescGZIP(), []int{1, 1}
}

type PInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityId   int32  `protobuf:"varint,1,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	Country  string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Region   string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Province string `protobuf:"bytes,4,opt,name=province,proto3" json:"province,omitempty"`
	City     string `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	Isp      string `protobuf:"bytes,6,opt,name=isp,proto3" json:"isp,omitempty"`
}

func (x *PInfo) Reset() {
	*x = PInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rpsgame_pb_rpsgame_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PInfo) ProtoMessage() {}

func (x *PInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rpsgame_pb_rpsgame_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PInfo.ProtoReflect.Descriptor instead.
func (*PInfo) Descriptor() ([]byte, []int) {
	return file_pb_rpsgame_pb_rpsgame_proto_rawDescGZIP(), []int{0}
}

func (x *PInfo) GetCityId() int32 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *PInfo) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *PInfo) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *PInfo) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *PInfo) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *PInfo) GetIsp() string {
	if x != nil {
		return x.Isp
	}
	return ""
}

type BotStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotId   string                 `protobuf:"bytes,1,opt,name=bot_id,json=botId,proto3" json:"bot_id,omitempty"`
	X       float32                `protobuf:"fixed32,2,opt,name=x,proto3" json:"x,omitempty"`
	Y       float32                `protobuf:"fixed32,3,opt,name=y,proto3" json:"y,omitempty"`
	EyeX    float32                `protobuf:"fixed32,4,opt,name=eye_x,json=eyeX,proto3" json:"eye_x,omitempty"`
	EyeY    float32                `protobuf:"fixed32,5,opt,name=eye_y,json=eyeY,proto3" json:"eye_y,omitempty"`
	Msg     string                 `protobuf:"bytes,6,opt,name=msg,proto3" json:"msg,omitempty"`
	RealX   float32                `protobuf:"fixed32,7,opt,name=real_x,json=realX,proto3" json:"real_x,omitempty"`
	RealY   float32                `protobuf:"fixed32,8,opt,name=real_y,json=realY,proto3" json:"real_y,omitempty"`
	Status  BotStatusReqStatusType `protobuf:"varint,9,opt,name=status,proto3,enum=pb_rpsgame.BotStatusReqStatusType" json:"status,omitempty"`
	Name    string                 `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Gender  BotStatusReqGenderType `protobuf:"varint,11,opt,name=gender,proto3,enum=pb_rpsgame.BotStatusReqGenderType" json:"gender,omitempty"`
	PosInfo *PInfo                 `protobuf:"bytes,12,opt,name=pos_info,json=posInfo,proto3" json:"pos_info,omitempty"`
}

func (x *BotStatusReq) Reset() {
	*x = BotStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rpsgame_pb_rpsgame_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotStatusReq) ProtoMessage() {}

func (x *BotStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rpsgame_pb_rpsgame_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotStatusReq.ProtoReflect.Descriptor instead.
func (*BotStatusReq) Descriptor() ([]byte, []int) {
	return file_pb_rpsgame_pb_rpsgame_proto_rawDescGZIP(), []int{1}
}

func (x *BotStatusReq) GetBotId() string {
	if x != nil {
		return x.BotId
	}
	return ""
}

func (x *BotStatusReq) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *BotStatusReq) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *BotStatusReq) GetEyeX() float32 {
	if x != nil {
		return x.EyeX
	}
	return 0
}

func (x *BotStatusReq) GetEyeY() float32 {
	if x != nil {
		return x.EyeY
	}
	return 0
}

func (x *BotStatusReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *BotStatusReq) GetRealX() float32 {
	if x != nil {
		return x.RealX
	}
	return 0
}

func (x *BotStatusReq) GetRealY() float32 {
	if x != nil {
		return x.RealY
	}
	return 0
}

func (x *BotStatusReq) GetStatus() BotStatusReqStatusType {
	if x != nil {
		return x.Status
	}
	return BotStatusReq_waiting
}

func (x *BotStatusReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BotStatusReq) GetGender() BotStatusReqGenderType {
	if x != nil {
		return x.Gender
	}
	return BotStatusReq_man
}

func (x *BotStatusReq) GetPosInfo() *PInfo {
	if x != nil {
		return x.PosInfo
	}
	return nil
}

type BotStatusRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotStatus []*BotStatusReq `protobuf:"bytes,1,rep,name=bot_status,json=botStatus,proto3" json:"bot_status,omitempty"`
}

func (x *BotStatusRes) Reset() {
	*x = BotStatusRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rpsgame_pb_rpsgame_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotStatusRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotStatusRes) ProtoMessage() {}

func (x *BotStatusRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rpsgame_pb_rpsgame_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotStatusRes.ProtoReflect.Descriptor instead.
func (*BotStatusRes) Descriptor() ([]byte, []int) {
	return file_pb_rpsgame_pb_rpsgame_proto_rawDescGZIP(), []int{2}
}

func (x *BotStatusRes) GetBotStatus() []*BotStatusReq {
	if x != nil {
		return x.BotStatus
	}
	return nil
}

var File_pb_rpsgame_pb_rpsgame_proto protoreflect.FileDescriptor

var file_pb_rpsgame_pb_rpsgame_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x62, 0x5f, 0x72, 0x70, 0x73, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x62, 0x5f,
	0x72, 0x70, 0x73, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70,
	0x62, 0x5f, 0x72, 0x70, 0x73, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x05, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x73, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x73, 0x70,
	0x22, 0xc3, 0x03, 0x0a, 0x0c, 0x42, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x01, 0x79, 0x12, 0x13, 0x0a, 0x05, 0x65, 0x79, 0x65, 0x5f, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x04, 0x65, 0x79, 0x65, 0x58, 0x12, 0x13, 0x0a, 0x05, 0x65, 0x79, 0x65,
	0x5f, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x65, 0x79, 0x65, 0x59, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x58, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x6c, 0x5f,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x59, 0x12, 0x3c,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24,
	0x2e, 0x70, 0x62, 0x5f, 0x72, 0x70, 0x73, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x42, 0x6f, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x3c, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x24, 0x2e, 0x70, 0x62, 0x5f, 0x72, 0x70, 0x73, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x42, 0x6f,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x08, 0x70, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x62, 0x5f, 0x72, 0x70, 0x73, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x35, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x77,
	0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x10, 0x02, 0x22, 0x21, 0x0a, 0x0b, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x6d, 0x61, 0x6e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x77,
	0x6f, 0x6d, 0x61, 0x6e, 0x10, 0x01, 0x22, 0x47, 0x0a, 0x0c, 0x42, 0x6f, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x62, 0x6f, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x5f,
	0x72, 0x70, 0x73, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x42, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x52, 0x09, 0x62, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x27, 0x5a, 0x25, 0x64, 0x75, 0x65, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x62,
	0x5f, 0x72, 0x70, 0x73, 0x67, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_rpsgame_pb_rpsgame_proto_rawDescOnce sync.Once
	file_pb_rpsgame_pb_rpsgame_proto_rawDescData = file_pb_rpsgame_pb_rpsgame_proto_rawDesc
)

func file_pb_rpsgame_pb_rpsgame_proto_rawDescGZIP() []byte {
	file_pb_rpsgame_pb_rpsgame_proto_rawDescOnce.Do(func() {
		file_pb_rpsgame_pb_rpsgame_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_rpsgame_pb_rpsgame_proto_rawDescData)
	})
	return file_pb_rpsgame_pb_rpsgame_proto_rawDescData
}

var file_pb_rpsgame_pb_rpsgame_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pb_rpsgame_pb_rpsgame_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_rpsgame_pb_rpsgame_proto_goTypes = []interface{}{
	(BotStatusReqStatusType)(0), // 0: pb_rpsgame.BotStatusReq.status_type
	(BotStatusReqGenderType)(0), // 1: pb_rpsgame.BotStatusReq.gender_type
	(*PInfo)(nil),               // 2: pb_rpsgame.pInfo
	(*BotStatusReq)(nil),        // 3: pb_rpsgame.BotStatusReq
	(*BotStatusRes)(nil),        // 4: pb_rpsgame.BotStatusRes
}
var file_pb_rpsgame_pb_rpsgame_proto_depIdxs = []int32{
	0, // 0: pb_rpsgame.BotStatusReq.status:type_name -> pb_rpsgame.BotStatusReq.status_type
	1, // 1: pb_rpsgame.BotStatusReq.gender:type_name -> pb_rpsgame.BotStatusReq.gender_type
	2, // 2: pb_rpsgame.BotStatusReq.pos_info:type_name -> pb_rpsgame.pInfo
	3, // 3: pb_rpsgame.BotStatusRes.bot_status:type_name -> pb_rpsgame.BotStatusReq
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pb_rpsgame_pb_rpsgame_proto_init() }
func file_pb_rpsgame_pb_rpsgame_proto_init() {
	if File_pb_rpsgame_pb_rpsgame_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_rpsgame_pb_rpsgame_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PInfo); i {
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
		file_pb_rpsgame_pb_rpsgame_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotStatusReq); i {
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
		file_pb_rpsgame_pb_rpsgame_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotStatusRes); i {
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
			RawDescriptor: file_pb_rpsgame_pb_rpsgame_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_rpsgame_pb_rpsgame_proto_goTypes,
		DependencyIndexes: file_pb_rpsgame_pb_rpsgame_proto_depIdxs,
		EnumInfos:         file_pb_rpsgame_pb_rpsgame_proto_enumTypes,
		MessageInfos:      file_pb_rpsgame_pb_rpsgame_proto_msgTypes,
	}.Build()
	File_pb_rpsgame_pb_rpsgame_proto = out.File
	file_pb_rpsgame_pb_rpsgame_proto_rawDesc = nil
	file_pb_rpsgame_pb_rpsgame_proto_goTypes = nil
	file_pb_rpsgame_pb_rpsgame_proto_depIdxs = nil
}