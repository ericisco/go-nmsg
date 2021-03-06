// Code generated by protoc-gen-go.
// source: xml.proto
// DO NOT EDIT!

package nmsg_base

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Xml struct {
	Xmltype          []byte `protobuf:"bytes,1,req,name=xmltype" json:"xmltype,omitempty"`
	Xmlpayload       []byte `protobuf:"bytes,2,req,name=xmlpayload" json:"xmlpayload,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Xml) Reset()                    { *m = Xml{} }
func (m *Xml) String() string            { return proto.CompactTextString(m) }
func (*Xml) ProtoMessage()               {}
func (*Xml) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *Xml) GetXmltype() []byte {
	if m != nil {
		return m.Xmltype
	}
	return nil
}

func (m *Xml) GetXmlpayload() []byte {
	if m != nil {
		return m.Xmlpayload
	}
	return nil
}

func init() {
	proto.RegisterType((*Xml)(nil), "nmsg.base.Xml")
}

func init() { proto.RegisterFile("xml.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 86 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xac, 0xc8, 0xcd, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0xcb, 0x2d, 0x4e, 0xd7, 0x4b, 0x4a, 0x2c, 0x4e,
	0x55, 0xd2, 0xe2, 0x62, 0x8e, 0xc8, 0xcd, 0x11, 0xe2, 0xe7, 0x62, 0x07, 0x4a, 0x97, 0x54, 0x16,
	0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x69, 0xf0, 0x08, 0x09, 0x71, 0x71, 0x01, 0x05, 0x0a, 0x12, 0x2b,
	0x73, 0xf2, 0x13, 0x53, 0x24, 0x98, 0x40, 0x62, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x92,
	0x10, 0x81, 0x42, 0x00, 0x00, 0x00,
}
