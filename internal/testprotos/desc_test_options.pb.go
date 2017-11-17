// Code generated by protoc-gen-go. DO NOT EDIT.
// source: desc_test_options.proto

package testprotos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ReallySimpleEnum int32

const (
	ReallySimpleEnum_VALUE ReallySimpleEnum = 1
)

var ReallySimpleEnum_name = map[int32]string{
	1: "VALUE",
}
var ReallySimpleEnum_value = map[string]int32{
	"VALUE": 1,
}

func (x ReallySimpleEnum) Enum() *ReallySimpleEnum {
	p := new(ReallySimpleEnum)
	*p = x
	return p
}
func (x ReallySimpleEnum) String() string {
	return proto.EnumName(ReallySimpleEnum_name, int32(x))
}
func (x *ReallySimpleEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ReallySimpleEnum_value, data, "ReallySimpleEnum")
	if err != nil {
		return err
	}
	*x = ReallySimpleEnum(value)
	return nil
}
func (ReallySimpleEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

type ReallySimpleMessage struct {
	Id               *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReallySimpleMessage) Reset()                    { *m = ReallySimpleMessage{} }
func (m *ReallySimpleMessage) String() string            { return proto.CompactTextString(m) }
func (*ReallySimpleMessage) ProtoMessage()               {}
func (*ReallySimpleMessage) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *ReallySimpleMessage) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *ReallySimpleMessage) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

var E_Mfubar = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         10101,
	Name:          "testprotos.mfubar",
	Tag:           "varint,10101,opt,name=mfubar",
	Filename:      "desc_test_options.proto",
}

var E_Ffubar = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: ([]string)(nil),
	Field:         10101,
	Name:          "testprotos.ffubar",
	Tag:           "bytes,10101,rep,name=ffubar",
	Filename:      "desc_test_options.proto",
}

var E_Ffubarb = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: ([]byte)(nil),
	Field:         10102,
	Name:          "testprotos.ffubarb",
	Tag:           "bytes,10102,opt,name=ffubarb",
	Filename:      "desc_test_options.proto",
}

var E_Efubar = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         10101,
	Name:          "testprotos.efubar",
	Tag:           "varint,10101,opt,name=efubar",
	Filename:      "desc_test_options.proto",
}

var E_Efubars = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         10102,
	Name:          "testprotos.efubars",
	Tag:           "zigzag32,10102,opt,name=efubars",
	Filename:      "desc_test_options.proto",
}

var E_Efubarsf = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         10103,
	Name:          "testprotos.efubarsf",
	Tag:           "fixed32,10103,opt,name=efubarsf",
	Filename:      "desc_test_options.proto",
}

var E_Efubaru = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         10104,
	Name:          "testprotos.efubaru",
	Tag:           "varint,10104,opt,name=efubaru",
	Filename:      "desc_test_options.proto",
}

var E_Efubaruf = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         10105,
	Name:          "testprotos.efubaruf",
	Tag:           "fixed32,10105,opt,name=efubaruf",
	Filename:      "desc_test_options.proto",
}

var E_Evfubar = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         10101,
	Name:          "testprotos.evfubar",
	Tag:           "varint,10101,opt,name=evfubar",
	Filename:      "desc_test_options.proto",
}

var E_Evfubars = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         10102,
	Name:          "testprotos.evfubars",
	Tag:           "zigzag64,10102,opt,name=evfubars",
	Filename:      "desc_test_options.proto",
}

var E_Evfubarsf = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         10103,
	Name:          "testprotos.evfubarsf",
	Tag:           "fixed64,10103,opt,name=evfubarsf",
	Filename:      "desc_test_options.proto",
}

var E_Evfubaru = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         10104,
	Name:          "testprotos.evfubaru",
	Tag:           "varint,10104,opt,name=evfubaru",
	Filename:      "desc_test_options.proto",
}

var E_Evfubaruf = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         10105,
	Name:          "testprotos.evfubaruf",
	Tag:           "fixed64,10105,opt,name=evfubaruf",
	Filename:      "desc_test_options.proto",
}

var E_Sfubar = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*ReallySimpleMessage)(nil),
	Field:         10101,
	Name:          "testprotos.sfubar",
	Tag:           "bytes,10101,opt,name=sfubar",
	Filename:      "desc_test_options.proto",
}

var E_Sfubare = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*ReallySimpleEnum)(nil),
	Field:         10102,
	Name:          "testprotos.sfubare",
	Tag:           "varint,10102,opt,name=sfubare,enum=testprotos.ReallySimpleEnum",
	Filename:      "desc_test_options.proto",
}

var E_Mtfubar = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: ([]float32)(nil),
	Field:         10101,
	Name:          "testprotos.mtfubar",
	Tag:           "fixed32,10101,rep,name=mtfubar",
	Filename:      "desc_test_options.proto",
}

var E_Mtfubard = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*float64)(nil),
	Field:         10102,
	Name:          "testprotos.mtfubard",
	Tag:           "fixed64,10102,opt,name=mtfubard",
	Filename:      "desc_test_options.proto",
}

func init() {
	proto.RegisterType((*ReallySimpleMessage)(nil), "testprotos.ReallySimpleMessage")
	proto.RegisterEnum("testprotos.ReallySimpleEnum", ReallySimpleEnum_name, ReallySimpleEnum_value)
	proto.RegisterExtension(E_Mfubar)
	proto.RegisterExtension(E_Ffubar)
	proto.RegisterExtension(E_Ffubarb)
	proto.RegisterExtension(E_Efubar)
	proto.RegisterExtension(E_Efubars)
	proto.RegisterExtension(E_Efubarsf)
	proto.RegisterExtension(E_Efubaru)
	proto.RegisterExtension(E_Efubaruf)
	proto.RegisterExtension(E_Evfubar)
	proto.RegisterExtension(E_Evfubars)
	proto.RegisterExtension(E_Evfubarsf)
	proto.RegisterExtension(E_Evfubaru)
	proto.RegisterExtension(E_Evfubaruf)
	proto.RegisterExtension(E_Sfubar)
	proto.RegisterExtension(E_Sfubare)
	proto.RegisterExtension(E_Mtfubar)
	proto.RegisterExtension(E_Mtfubard)
}

func init() { proto.RegisterFile("desc_test_options.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcd, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xe5, 0x34, 0xb5, 0x93, 0x01, 0x4a, 0x30, 0x07, 0x22, 0xd4, 0xd2, 0xc0, 0x29, 0xe2,
	0x60, 0x8b, 0x22, 0xa0, 0x59, 0x10, 0x02, 0xa4, 0x72, 0xe2, 0x43, 0x9a, 0x8a, 0x0a, 0xb8, 0x54,
	0x4e, 0xbc, 0x4e, 0x16, 0xd9, 0x5e, 0xcb, 0xbb, 0x5b, 0x89, 0x3f, 0x9c, 0xaf, 0x1b, 0xf2, 0xae,
	0x3f, 0x12, 0x35, 0x91, 0x7d, 0x7b, 0x89, 0xde, 0x6f, 0xde, 0xdb, 0x89, 0x26, 0x70, 0x2f, 0xa4,
	0x62, 0x71, 0x29, 0xa9, 0x90, 0x97, 0x3c, 0x93, 0x8c, 0xa7, 0xc2, 0xcb, 0x72, 0x2e, 0xb9, 0x0b,
	0xc5, 0x77, 0x5a, 0x8a, 0xfb, 0x93, 0x25, 0xe7, 0xcb, 0x98, 0xfa, 0xfa, 0xe3, 0x5c, 0x45, 0x7e,
	0x01, 0xe5, 0x2c, 0x93, 0x3c, 0x37, 0xee, 0x47, 0x33, 0xb8, 0x8b, 0x34, 0x88, 0xe3, 0x9f, 0xe7,
	0x2c, 0xc9, 0x62, 0xfa, 0x91, 0x0a, 0x11, 0x2c, 0xa9, 0x7b, 0x00, 0x3d, 0x16, 0x8e, 0xad, 0x89,
	0x35, 0xed, 0x63, 0x8f, 0x85, 0xae, 0x0b, 0xfd, 0x34, 0x48, 0xe8, 0xb8, 0x37, 0xb1, 0xa6, 0x43,
	0xd4, 0xfa, 0xf1, 0x11, 0x8c, 0xd6, 0xd1, 0xb3, 0x54, 0x25, 0xee, 0x10, 0xf6, 0x2f, 0xde, 0x7e,
	0xf8, 0x72, 0x36, 0xb2, 0xc8, 0x29, 0xd8, 0x49, 0xa4, 0xe6, 0x41, 0xee, 0x1e, 0x7b, 0xa6, 0x86,
	0x57, 0xd5, 0xf0, 0xca, 0x98, 0xcf, 0xa6, 0xf8, 0xf8, 0xd7, 0xa7, 0x89, 0x35, 0x1d, 0x60, 0xe9,
	0x27, 0xcf, 0xc1, 0x8e, 0x0c, 0x79, 0x74, 0x8d, 0x7c, 0xcf, 0x68, 0x1c, 0xae, 0x71, 0x7b, 0xd3,
	0x21, 0x96, 0x6e, 0x72, 0x0a, 0x8e, 0x51, 0xf3, 0x36, 0xf0, 0x77, 0x11, 0x78, 0x13, 0x2b, 0x3b,
	0x79, 0x06, 0x36, 0x35, 0x89, 0x87, 0xd7, 0xc0, 0xe2, 0x5d, 0x1b, 0x45, 0xf7, 0xb1, 0x34, 0x93,
	0x17, 0xe0, 0x18, 0x25, 0x5a, 0x38, 0x9d, 0x77, 0x07, 0x2b, 0x37, 0x99, 0xc1, 0xa0, 0x94, 0x51,
	0x0b, 0xf9, 0xa7, 0x20, 0x6f, 0x63, 0x6d, 0x6f, 0x32, 0x55, 0x0b, 0xf9, 0xb7, 0x20, 0x6f, 0x55,
	0x99, 0xaa, 0xc9, 0x54, 0x6d, 0x99, 0xff, 0x0a, 0xd2, 0xc1, 0xda, 0x4e, 0x5e, 0x81, 0x43, 0xaf,
	0xcc, 0x7e, 0x1e, 0x6e, 0x25, 0x2f, 0x82, 0x58, 0x6d, 0xfe, 0x9a, 0x7b, 0x58, 0x21, 0xe4, 0x35,
	0x0c, 0x4a, 0x29, 0xba, 0xe0, 0x7a, 0x57, 0x2e, 0xd6, 0x0c, 0x79, 0x03, 0xc3, 0x4a, 0x47, 0x5d,
	0x06, 0xe8, 0x95, 0x8d, 0xb0, 0x81, 0xd6, 0x1a, 0xa8, 0x2e, 0x03, 0xf4, 0xe6, 0xfa, 0x75, 0x03,
	0xb5, 0xd6, 0x40, 0x75, 0x6a, 0xa0, 0x17, 0x68, 0x63, 0x03, 0x91, 0x6f, 0x60, 0x8b, 0x5d, 0xc7,
	0x70, 0x4e, 0xf3, 0x2b, 0xb6, 0xd8, 0x5c, 0xdf, 0x8d, 0x93, 0x63, 0xaf, 0xb9, 0x63, 0x6f, 0xcb,
	0x89, 0x62, 0x39, 0x90, 0x7c, 0x05, 0xc7, 0x28, 0xda, 0x3e, 0x5b, 0xef, 0xf6, 0xe0, 0xe4, 0x70,
	0xd7, 0xec, 0xe2, 0x25, 0x58, 0x8d, 0x23, 0x33, 0x70, 0x12, 0x69, 0x5a, 0x3f, 0xd8, 0x72, 0xc2,
	0x72, 0xc5, 0x37, 0x2f, 0xb1, 0x87, 0x95, 0x9f, 0xbc, 0x84, 0x41, 0x29, 0xc3, 0x56, 0x56, 0x97,
	0xb2, 0xb0, 0x06, 0xde, 0x3d, 0xfd, 0xfe, 0x64, 0xc9, 0xe4, 0x4a, 0xcd, 0xbd, 0x05, 0x4f, 0xfc,
	0x1f, 0x2b, 0x95, 0x64, 0xe6, 0x1f, 0x2c, 0xa7, 0x51, 0x4c, 0x17, 0xd2, 0x67, 0xa9, 0xa4, 0x79,
	0x1a, 0xc4, 0x7e, 0xf3, 0x8c, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xf8, 0x98, 0xc6, 0x10,
	0x05, 0x00, 0x00,
}
