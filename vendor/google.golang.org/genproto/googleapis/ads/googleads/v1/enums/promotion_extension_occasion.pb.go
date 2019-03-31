// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/promotion_extension_occasion.proto

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

// A promotion extension occasion.
type PromotionExtensionOccasionEnum_PromotionExtensionOccasion int32

const (
	// Not specified.
	PromotionExtensionOccasionEnum_UNSPECIFIED PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 0
	// Used for return value only. Represents value unknown in this version.
	PromotionExtensionOccasionEnum_UNKNOWN PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 1
	// New Year's.
	PromotionExtensionOccasionEnum_NEW_YEARS PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 2
	// Chinese New Year.
	PromotionExtensionOccasionEnum_CHINESE_NEW_YEAR PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 3
	// Valentine's Day.
	PromotionExtensionOccasionEnum_VALENTINES_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 4
	// Easter.
	PromotionExtensionOccasionEnum_EASTER PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 5
	// Mother's Day.
	PromotionExtensionOccasionEnum_MOTHERS_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 6
	// Father's Day.
	PromotionExtensionOccasionEnum_FATHERS_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 7
	// Labor Day.
	PromotionExtensionOccasionEnum_LABOR_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 8
	// Back To School.
	PromotionExtensionOccasionEnum_BACK_TO_SCHOOL PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 9
	// Halloween.
	PromotionExtensionOccasionEnum_HALLOWEEN PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 10
	// Black Friday.
	PromotionExtensionOccasionEnum_BLACK_FRIDAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 11
	// Cyber Monday.
	PromotionExtensionOccasionEnum_CYBER_MONDAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 12
	// Christmas.
	PromotionExtensionOccasionEnum_CHRISTMAS PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 13
	// Boxing Day.
	PromotionExtensionOccasionEnum_BOXING_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 14
	// Independence Day in any country.
	PromotionExtensionOccasionEnum_INDEPENDENCE_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 15
	// National Day in any country.
	PromotionExtensionOccasionEnum_NATIONAL_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 16
	// End of any season.
	PromotionExtensionOccasionEnum_END_OF_SEASON PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 17
	// Winter Sale.
	PromotionExtensionOccasionEnum_WINTER_SALE PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 18
	// Summer sale.
	PromotionExtensionOccasionEnum_SUMMER_SALE PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 19
	// Fall Sale.
	PromotionExtensionOccasionEnum_FALL_SALE PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 20
	// Spring Sale.
	PromotionExtensionOccasionEnum_SPRING_SALE PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 21
	// Ramadan.
	PromotionExtensionOccasionEnum_RAMADAN PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 22
	// Eid al-Fitr.
	PromotionExtensionOccasionEnum_EID_AL_FITR PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 23
	// Eid al-Adha.
	PromotionExtensionOccasionEnum_EID_AL_ADHA PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 24
	// Singles Day.
	PromotionExtensionOccasionEnum_SINGLES_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 25
	// Women's Day.
	PromotionExtensionOccasionEnum_WOMENS_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 26
	// Holi.
	PromotionExtensionOccasionEnum_HOLI PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 27
	// Parent's Day.
	PromotionExtensionOccasionEnum_PARENTS_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 28
	// St. Nicholas Day.
	PromotionExtensionOccasionEnum_ST_NICHOLAS_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 29
	// Carnival.
	PromotionExtensionOccasionEnum_CARNIVAL PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 30
	// Epiphany, also known as Three Kings' Day.
	PromotionExtensionOccasionEnum_EPIPHANY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 31
	// Rosh Hashanah.
	PromotionExtensionOccasionEnum_ROSH_HASHANAH PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 32
	// Passover.
	PromotionExtensionOccasionEnum_PASSOVER PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 33
	// Hanukkah.
	PromotionExtensionOccasionEnum_HANUKKAH PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 34
	// Diwali.
	PromotionExtensionOccasionEnum_DIWALI PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 35
	// Navratri.
	PromotionExtensionOccasionEnum_NAVRATRI PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 36
	// Available in Thai: Songkran.
	PromotionExtensionOccasionEnum_SONGKRAN PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 37
	// Available in Japanese: Year-end Gift.
	PromotionExtensionOccasionEnum_YEAR_END_GIFT PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 38
)

var PromotionExtensionOccasionEnum_PromotionExtensionOccasion_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "NEW_YEARS",
	3:  "CHINESE_NEW_YEAR",
	4:  "VALENTINES_DAY",
	5:  "EASTER",
	6:  "MOTHERS_DAY",
	7:  "FATHERS_DAY",
	8:  "LABOR_DAY",
	9:  "BACK_TO_SCHOOL",
	10: "HALLOWEEN",
	11: "BLACK_FRIDAY",
	12: "CYBER_MONDAY",
	13: "CHRISTMAS",
	14: "BOXING_DAY",
	15: "INDEPENDENCE_DAY",
	16: "NATIONAL_DAY",
	17: "END_OF_SEASON",
	18: "WINTER_SALE",
	19: "SUMMER_SALE",
	20: "FALL_SALE",
	21: "SPRING_SALE",
	22: "RAMADAN",
	23: "EID_AL_FITR",
	24: "EID_AL_ADHA",
	25: "SINGLES_DAY",
	26: "WOMENS_DAY",
	27: "HOLI",
	28: "PARENTS_DAY",
	29: "ST_NICHOLAS_DAY",
	30: "CARNIVAL",
	31: "EPIPHANY",
	32: "ROSH_HASHANAH",
	33: "PASSOVER",
	34: "HANUKKAH",
	35: "DIWALI",
	36: "NAVRATRI",
	37: "SONGKRAN",
	38: "YEAR_END_GIFT",
}
var PromotionExtensionOccasionEnum_PromotionExtensionOccasion_value = map[string]int32{
	"UNSPECIFIED":      0,
	"UNKNOWN":          1,
	"NEW_YEARS":        2,
	"CHINESE_NEW_YEAR": 3,
	"VALENTINES_DAY":   4,
	"EASTER":           5,
	"MOTHERS_DAY":      6,
	"FATHERS_DAY":      7,
	"LABOR_DAY":        8,
	"BACK_TO_SCHOOL":   9,
	"HALLOWEEN":        10,
	"BLACK_FRIDAY":     11,
	"CYBER_MONDAY":     12,
	"CHRISTMAS":        13,
	"BOXING_DAY":       14,
	"INDEPENDENCE_DAY": 15,
	"NATIONAL_DAY":     16,
	"END_OF_SEASON":    17,
	"WINTER_SALE":      18,
	"SUMMER_SALE":      19,
	"FALL_SALE":        20,
	"SPRING_SALE":      21,
	"RAMADAN":          22,
	"EID_AL_FITR":      23,
	"EID_AL_ADHA":      24,
	"SINGLES_DAY":      25,
	"WOMENS_DAY":       26,
	"HOLI":             27,
	"PARENTS_DAY":      28,
	"ST_NICHOLAS_DAY":  29,
	"CARNIVAL":         30,
	"EPIPHANY":         31,
	"ROSH_HASHANAH":    32,
	"PASSOVER":         33,
	"HANUKKAH":         34,
	"DIWALI":           35,
	"NAVRATRI":         36,
	"SONGKRAN":         37,
	"YEAR_END_GIFT":    38,
}

func (x PromotionExtensionOccasionEnum_PromotionExtensionOccasion) String() string {
	return proto.EnumName(PromotionExtensionOccasionEnum_PromotionExtensionOccasion_name, int32(x))
}
func (PromotionExtensionOccasionEnum_PromotionExtensionOccasion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_promotion_extension_occasion_e447c55310caddd1, []int{0, 0}
}

// Container for enum describing a promotion extension occasion.
// For more information about the occasions please check:
// https://support.google.com/google-ads/answer/7367521
type PromotionExtensionOccasionEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PromotionExtensionOccasionEnum) Reset()         { *m = PromotionExtensionOccasionEnum{} }
func (m *PromotionExtensionOccasionEnum) String() string { return proto.CompactTextString(m) }
func (*PromotionExtensionOccasionEnum) ProtoMessage()    {}
func (*PromotionExtensionOccasionEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_promotion_extension_occasion_e447c55310caddd1, []int{0}
}
func (m *PromotionExtensionOccasionEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PromotionExtensionOccasionEnum.Unmarshal(m, b)
}
func (m *PromotionExtensionOccasionEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PromotionExtensionOccasionEnum.Marshal(b, m, deterministic)
}
func (dst *PromotionExtensionOccasionEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PromotionExtensionOccasionEnum.Merge(dst, src)
}
func (m *PromotionExtensionOccasionEnum) XXX_Size() int {
	return xxx_messageInfo_PromotionExtensionOccasionEnum.Size(m)
}
func (m *PromotionExtensionOccasionEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PromotionExtensionOccasionEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PromotionExtensionOccasionEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PromotionExtensionOccasionEnum)(nil), "google.ads.googleads.v1.enums.PromotionExtensionOccasionEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.PromotionExtensionOccasionEnum_PromotionExtensionOccasion", PromotionExtensionOccasionEnum_PromotionExtensionOccasion_name, PromotionExtensionOccasionEnum_PromotionExtensionOccasion_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/promotion_extension_occasion.proto", fileDescriptor_promotion_extension_occasion_e447c55310caddd1)
}

var fileDescriptor_promotion_extension_occasion_e447c55310caddd1 = []byte{
	// 669 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcb, 0x6e, 0xdb, 0x3a,
	0x10, 0xbd, 0x71, 0xde, 0xcc, 0x8b, 0x51, 0x72, 0x5f, 0xb9, 0x79, 0xdc, 0xe4, 0x3e, 0x96, 0x32,
	0x8c, 0xee, 0xd4, 0x4d, 0x47, 0x12, 0x6d, 0x11, 0xa6, 0x87, 0x02, 0x29, 0xdb, 0x75, 0x61, 0x40,
	0x70, 0x63, 0xc3, 0x08, 0x10, 0x4b, 0x46, 0xe4, 0x04, 0xfd, 0x95, 0x6e, 0xbb, 0xec, 0x37, 0xf4,
	0x0b, 0xfa, 0x29, 0x5d, 0xf7, 0x03, 0x8a, 0x91, 0xe2, 0xb4, 0x9b, 0x74, 0x63, 0xcc, 0x9c, 0x73,
	0xe6, 0x1c, 0x93, 0xe2, 0xb0, 0x57, 0xd3, 0x3c, 0x9f, 0xde, 0x4e, 0xea, 0xa3, 0x71, 0x51, 0xaf,
	0x4a, 0xaa, 0x1e, 0x1a, 0xf5, 0x49, 0x76, 0x3f, 0x2b, 0xea, 0xf3, 0xbb, 0x7c, 0x96, 0x2f, 0x6e,
	0xf2, 0x2c, 0x9d, 0xbc, 0x5b, 0x4c, 0xb2, 0x82, 0xaa, 0xfc, 0xfa, 0x7a, 0x44, 0x85, 0x3b, 0xbf,
	0xcb, 0x17, 0xb9, 0x73, 0x56, 0x8d, 0xb9, 0xa3, 0x71, 0xe1, 0x3e, 0x39, 0xb8, 0x0f, 0x0d, 0xb7,
	0x74, 0x38, 0x39, 0x5d, 0x06, 0xcc, 0x6f, 0xea, 0xa3, 0x2c, 0xcb, 0x17, 0x23, 0xb2, 0x2b, 0xaa,
	0xe1, 0xab, 0x4f, 0xeb, 0xec, 0x3c, 0x5e, 0x66, 0x88, 0x65, 0x84, 0x7e, 0x4c, 0x10, 0xd9, 0xfd,
	0xec, 0xea, 0xfd, 0x3a, 0x3b, 0x79, 0x5e, 0xe2, 0x1c, 0xb0, 0x9d, 0x2e, 0xda, 0x58, 0x04, 0xb2,
	0x29, 0x45, 0xc8, 0x7f, 0x71, 0x76, 0xd8, 0x66, 0x17, 0xdb, 0xa8, 0xfb, 0xc8, 0x57, 0x9c, 0x3d,
	0xb6, 0x8d, 0xa2, 0x9f, 0x0e, 0x04, 0x18, 0xcb, 0x6b, 0xce, 0x31, 0xe3, 0x41, 0x24, 0x51, 0x58,
	0x91, 0x2e, 0x61, 0xbe, 0xea, 0x38, 0x6c, 0xbf, 0x07, 0x4a, 0x60, 0x42, 0x4c, 0x1a, 0xc2, 0x80,
	0xaf, 0x39, 0x8c, 0x6d, 0x08, 0xb0, 0x89, 0x30, 0x7c, 0x9d, 0x22, 0x3a, 0x3a, 0x89, 0x84, 0xa9,
	0xc8, 0x0d, 0x02, 0x9a, 0xf0, 0x1d, 0xd8, 0xa4, 0x18, 0x05, 0xbe, 0x36, 0x65, 0xbb, 0x45, 0x86,
	0x3e, 0x04, 0xed, 0x34, 0xd1, 0xa9, 0x0d, 0x22, 0xad, 0x15, 0xdf, 0x26, 0x49, 0x04, 0x4a, 0xe9,
	0xbe, 0x10, 0xc8, 0x99, 0xc3, 0xd9, 0xae, 0xaf, 0x48, 0xd3, 0x34, 0x92, 0x86, 0x76, 0x08, 0x09,
	0x06, 0xbe, 0x30, 0x69, 0x47, 0x23, 0x21, 0xbb, 0x34, 0x12, 0x44, 0x46, 0xda, 0xa4, 0x03, 0x96,
	0xef, 0x39, 0xfb, 0x8c, 0xf9, 0xfa, 0xb5, 0xc4, 0x56, 0x99, 0xb2, 0x4f, 0x87, 0x91, 0x18, 0x8a,
	0x58, 0x60, 0x28, 0x30, 0x10, 0x25, 0x7a, 0x40, 0x36, 0x08, 0x89, 0xd4, 0x08, 0xaa, 0x44, 0xb8,
	0x73, 0xc8, 0xf6, 0x04, 0x86, 0xa9, 0x6e, 0xa6, 0x56, 0x80, 0xd5, 0xc8, 0x0f, 0xe9, 0x00, 0x7d,
	0x89, 0x89, 0x30, 0xa9, 0x05, 0x25, 0xb8, 0x43, 0x80, 0xed, 0x76, 0x3a, 0x4b, 0xe0, 0x88, 0xb2,
	0x9b, 0xa0, 0x54, 0xd5, 0x1e, 0x97, 0x7c, 0x6c, 0x28, 0xbb, 0x04, 0x7e, 0xa5, 0x5b, 0x36, 0xd0,
	0x81, 0x10, 0x90, 0xff, 0x46, 0xac, 0x90, 0x61, 0x0a, 0x2a, 0x6d, 0xca, 0xc4, 0xf0, 0xdf, 0x7f,
	0x00, 0x20, 0x8c, 0x80, 0xff, 0x51, 0xce, 0x4b, 0x6c, 0xa9, 0xc7, 0xfb, 0xfd, 0x93, 0x0e, 0xd3,
	0xd7, 0x1d, 0x81, 0x55, 0x7f, 0xe2, 0x6c, 0xb1, 0xb5, 0x48, 0x2b, 0xc9, 0xff, 0x22, 0x69, 0x0c,
	0x46, 0x60, 0x52, 0x51, 0xa7, 0xce, 0x11, 0x3b, 0xb0, 0x49, 0x8a, 0x32, 0x88, 0xb4, 0x82, 0x0a,
	0x3c, 0x73, 0x76, 0xd9, 0x56, 0x00, 0x06, 0x65, 0x0f, 0x14, 0x3f, 0xa7, 0x4e, 0xc4, 0x32, 0x8e,
	0x00, 0x07, 0xfc, 0x82, 0x0e, 0x6c, 0xb4, 0x8d, 0xd2, 0x08, 0x6c, 0x04, 0x08, 0x11, 0xff, 0x9b,
	0x04, 0x31, 0x58, 0xab, 0x7b, 0xc2, 0xf0, 0x4b, 0xea, 0x22, 0xc0, 0x6e, 0xbb, 0x0d, 0x11, 0xbf,
	0xa2, 0x4f, 0x1d, 0xca, 0x3e, 0x28, 0xc9, 0xff, 0x21, 0x06, 0xa1, 0x67, 0x20, 0x31, 0x92, 0xff,
	0x4b, 0x9d, 0xd5, 0xd8, 0x6a, 0x1b, 0x40, 0xfe, 0x1f, 0xd9, 0xd2, 0x83, 0x49, 0xe9, 0x32, 0x5b,
	0xb2, 0x99, 0xf0, 0xff, 0xfd, 0xaf, 0x2b, 0xec, 0xf2, 0x3a, 0x9f, 0xb9, 0x3f, 0x5d, 0x01, 0xff,
	0xe2, 0xf9, 0xe7, 0x1b, 0xd3, 0x16, 0xc4, 0x2b, 0x6f, 0xfc, 0x47, 0x87, 0x69, 0x7e, 0x3b, 0xca,
	0xa6, 0x6e, 0x7e, 0x37, 0xad, 0x4f, 0x27, 0x59, 0xb9, 0x23, 0xcb, 0xb5, 0x9c, 0xdf, 0x14, 0xcf,
	0x6c, 0xe9, 0xcb, 0xf2, 0xf7, 0x43, 0x6d, 0xb5, 0x05, 0xf0, 0xb1, 0x76, 0xd6, 0xaa, 0xac, 0x60,
	0x5c, 0xb8, 0x55, 0x49, 0x55, 0xaf, 0xe1, 0xd2, 0x36, 0x15, 0x9f, 0x97, 0xfc, 0x10, 0xc6, 0xc5,
	0xf0, 0x89, 0x1f, 0xf6, 0x1a, 0xc3, 0x92, 0xff, 0x52, 0xbb, 0xac, 0x40, 0xcf, 0x83, 0x71, 0xe1,
	0x79, 0x4f, 0x0a, 0xcf, 0xeb, 0x35, 0x3c, 0xaf, 0xd4, 0xbc, 0xdd, 0x28, 0xff, 0xd8, 0x8b, 0x6f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xb8, 0x9b, 0x12, 0x3d, 0x04, 0x00, 0x00,
}
