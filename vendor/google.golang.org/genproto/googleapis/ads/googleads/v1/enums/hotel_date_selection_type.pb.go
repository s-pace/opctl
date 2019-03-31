// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/hotel_date_selection_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible hotel date selection types.
type HotelDateSelectionTypeEnum_HotelDateSelectionType int32

const (
	// Not specified.
	HotelDateSelectionTypeEnum_UNSPECIFIED HotelDateSelectionTypeEnum_HotelDateSelectionType = 0
	// Used for return value only. Represents value unknown in this version.
	HotelDateSelectionTypeEnum_UNKNOWN HotelDateSelectionTypeEnum_HotelDateSelectionType = 1
	// Dates selected by default.
	HotelDateSelectionTypeEnum_DEFAULT_SELECTION HotelDateSelectionTypeEnum_HotelDateSelectionType = 50
	// Dates selected by the user.
	HotelDateSelectionTypeEnum_USER_SELECTED HotelDateSelectionTypeEnum_HotelDateSelectionType = 51
)

var HotelDateSelectionTypeEnum_HotelDateSelectionType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	50: "DEFAULT_SELECTION",
	51: "USER_SELECTED",
}
var HotelDateSelectionTypeEnum_HotelDateSelectionType_value = map[string]int32{
	"UNSPECIFIED":       0,
	"UNKNOWN":           1,
	"DEFAULT_SELECTION": 50,
	"USER_SELECTED":     51,
}

func (x HotelDateSelectionTypeEnum_HotelDateSelectionType) String() string {
	return proto.EnumName(HotelDateSelectionTypeEnum_HotelDateSelectionType_name, int32(x))
}
func (HotelDateSelectionTypeEnum_HotelDateSelectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_hotel_date_selection_type_6f727549b636f879, []int{0, 0}
}

// Container for enum describing possible hotel date selection types
type HotelDateSelectionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotelDateSelectionTypeEnum) Reset()         { *m = HotelDateSelectionTypeEnum{} }
func (m *HotelDateSelectionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*HotelDateSelectionTypeEnum) ProtoMessage()    {}
func (*HotelDateSelectionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_hotel_date_selection_type_6f727549b636f879, []int{0}
}
func (m *HotelDateSelectionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotelDateSelectionTypeEnum.Unmarshal(m, b)
}
func (m *HotelDateSelectionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotelDateSelectionTypeEnum.Marshal(b, m, deterministic)
}
func (dst *HotelDateSelectionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotelDateSelectionTypeEnum.Merge(dst, src)
}
func (m *HotelDateSelectionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_HotelDateSelectionTypeEnum.Size(m)
}
func (m *HotelDateSelectionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_HotelDateSelectionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_HotelDateSelectionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HotelDateSelectionTypeEnum)(nil), "google.ads.googleads.v1.enums.HotelDateSelectionTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.HotelDateSelectionTypeEnum_HotelDateSelectionType", HotelDateSelectionTypeEnum_HotelDateSelectionType_name, HotelDateSelectionTypeEnum_HotelDateSelectionType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/hotel_date_selection_type.proto", fileDescriptor_hotel_date_selection_type_6f727549b636f879)
}

var fileDescriptor_hotel_date_selection_type_6f727549b636f879 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xc1, 0x4e, 0xc2, 0x40,
	0x14, 0x14, 0x4c, 0x34, 0x59, 0x62, 0x2c, 0x4d, 0xf4, 0x80, 0x72, 0x80, 0x0f, 0xd8, 0xa6, 0x72,
	0x5b, 0xe3, 0xa1, 0xd0, 0x05, 0x89, 0xa4, 0x10, 0x4b, 0x31, 0x31, 0x4d, 0xea, 0x4a, 0x37, 0x95,
	0xa4, 0xec, 0x36, 0xec, 0x42, 0xc2, 0xc5, 0x8f, 0xf1, 0xe8, 0xa7, 0xf8, 0x29, 0x9e, 0xfc, 0x04,
	0xb3, 0xbb, 0x94, 0x13, 0x7a, 0x79, 0x99, 0xbc, 0x79, 0x33, 0x79, 0x33, 0xe0, 0x2e, 0xe3, 0x3c,
	0xcb, 0xa9, 0x43, 0x52, 0xe1, 0x18, 0xa8, 0xd0, 0xc6, 0x75, 0x28, 0x5b, 0x2f, 0x85, 0xf3, 0xc6,
	0x25, 0xcd, 0x93, 0x94, 0x48, 0x9a, 0x08, 0x9a, 0xd3, 0xb9, 0x5c, 0x70, 0x96, 0xc8, 0x6d, 0x41,
	0x61, 0xb1, 0xe2, 0x92, 0xdb, 0x4d, 0xa3, 0x81, 0x24, 0x15, 0x70, 0x2f, 0x87, 0x1b, 0x17, 0x6a,
	0x79, 0xe3, 0xba, 0x74, 0x2f, 0x16, 0x0e, 0x61, 0x8c, 0x4b, 0xa2, 0x0c, 0x84, 0x11, 0xb7, 0xdf,
	0x41, 0xe3, 0x5e, 0xf9, 0xfb, 0x44, 0xd2, 0xb0, 0x74, 0x9f, 0x6e, 0x0b, 0x8a, 0xd9, 0x7a, 0xd9,
	0x7e, 0x01, 0x97, 0x87, 0x59, 0xfb, 0x1c, 0xd4, 0xa2, 0x20, 0x9c, 0xe0, 0xde, 0xb0, 0x3f, 0xc4,
	0xbe, 0x75, 0x64, 0xd7, 0xc0, 0x69, 0x14, 0x3c, 0x04, 0xe3, 0xa7, 0xc0, 0xaa, 0xd8, 0x17, 0xa0,
	0xee, 0xe3, 0xbe, 0x17, 0x8d, 0xa6, 0x49, 0x88, 0x47, 0xb8, 0x37, 0x1d, 0x8e, 0x03, 0xeb, 0xc6,
	0xae, 0x83, 0xb3, 0x28, 0xc4, 0x8f, 0xbb, 0x1d, 0xf6, 0xad, 0x4e, 0xf7, 0xa7, 0x02, 0x5a, 0x73,
	0xbe, 0x84, 0xff, 0x66, 0xe8, 0x5e, 0x1d, 0xfe, 0x62, 0xa2, 0x22, 0x4c, 0x2a, 0xcf, 0xdd, 0x9d,
	0x3a, 0xe3, 0x39, 0x61, 0x19, 0xe4, 0xab, 0xcc, 0xc9, 0x28, 0xd3, 0x01, 0xcb, 0x42, 0x8b, 0x85,
	0xf8, 0xa3, 0xdf, 0x5b, 0x3d, 0x3f, 0xaa, 0xc7, 0x03, 0xcf, 0xfb, 0xac, 0x36, 0x07, 0xc6, 0xca,
	0x4b, 0x05, 0x34, 0x50, 0xa1, 0x99, 0x0b, 0x55, 0x1f, 0xe2, 0xab, 0xe4, 0x63, 0x2f, 0x15, 0xf1,
	0x9e, 0x8f, 0x67, 0x6e, 0xac, 0xf9, 0xef, 0x6a, 0xcb, 0x2c, 0x11, 0xf2, 0x52, 0x81, 0xd0, 0xfe,
	0x02, 0xa1, 0x99, 0x8b, 0x90, 0xbe, 0x79, 0x3d, 0xd1, 0x8f, 0x75, 0x7e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x46, 0xaf, 0xfd, 0xab, 0xf7, 0x01, 0x00, 0x00,
}
