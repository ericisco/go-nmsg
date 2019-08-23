// Code generated by protoc-gen-go.
// source: logline.proto
// DO NOT EDIT!

package nmsg_base

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LogLine struct {
	Category         []byte `protobuf:"bytes,1,opt,name=category" json:"category,omitempty"`
	Message          []byte `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *LogLine) Reset()                    { *m = LogLine{} }
func (m *LogLine) String() string            { return proto.CompactTextString(m) }
func (*LogLine) ProtoMessage()               {}
func (*LogLine) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *LogLine) GetCategory() []byte {
	if m != nil {
		return m.Category
	}
	return nil
}

func (m *LogLine) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*LogLine)(nil), "nmsg.base.LogLine")
}

func init() { proto.RegisterFile("logline.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 93 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xc9, 0x4f, 0xcf,
	0xc9, 0xcc, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0xcb, 0x2d, 0x4e, 0xd7,
	0x4b, 0x4a, 0x2c, 0x4e, 0x55, 0xd2, 0xe1, 0x62, 0xf7, 0xc9, 0x4f, 0xf7, 0x01, 0xca, 0x09, 0x09,
	0x70, 0x71, 0x24, 0x27, 0x96, 0xa4, 0xa6, 0xe7, 0x17, 0x55, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0,
	0x08, 0xf1, 0x73, 0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x81, 0x04, 0x00,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x4e, 0xe7, 0x48, 0x48, 0x00, 0x00, 0x00,
}