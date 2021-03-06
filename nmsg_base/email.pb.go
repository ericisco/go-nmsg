// Code generated by protoc-gen-go.
// source: email.proto
// DO NOT EDIT!

package nmsg_base

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EmailType int32

const (
	EmailType_unknown     EmailType = 0
	EmailType_spamtrap    EmailType = 1
	EmailType_rej_network EmailType = 2
	EmailType_rej_content EmailType = 3
	EmailType_rej_user    EmailType = 4
)

var EmailType_name = map[int32]string{
	0: "unknown",
	1: "spamtrap",
	2: "rej_network",
	3: "rej_content",
	4: "rej_user",
}
var EmailType_value = map[string]int32{
	"unknown":     0,
	"spamtrap":    1,
	"rej_network": 2,
	"rej_content": 3,
	"rej_user":    4,
}

func (x EmailType) Enum() *EmailType {
	p := new(EmailType)
	*p = x
	return p
}
func (x EmailType) String() string {
	return proto.EnumName(EmailType_name, int32(x))
}
func (x *EmailType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EmailType_value, data, "EmailType")
	if err != nil {
		return err
	}
	*x = EmailType(value)
	return nil
}
func (EmailType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type Email struct {
	Type             *EmailType `protobuf:"varint,8,opt,name=type,enum=nmsg.base.EmailType" json:"type,omitempty"`
	Headers          []byte     `protobuf:"bytes,2,opt,name=headers" json:"headers,omitempty"`
	Srcip            []byte     `protobuf:"bytes,3,opt,name=srcip" json:"srcip,omitempty"`
	Srchost          []byte     `protobuf:"bytes,4,opt,name=srchost" json:"srchost,omitempty"`
	Helo             []byte     `protobuf:"bytes,5,opt,name=helo" json:"helo,omitempty"`
	From             []byte     `protobuf:"bytes,6,opt,name=from" json:"from,omitempty"`
	Rcpt             [][]byte   `protobuf:"bytes,7,rep,name=rcpt" json:"rcpt,omitempty"`
	Bodyurl          [][]byte   `protobuf:"bytes,9,rep,name=bodyurl" json:"bodyurl,omitempty"`
	Body             []byte     `protobuf:"bytes,10,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Email) Reset()                    { *m = Email{} }
func (m *Email) String() string            { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()               {}
func (*Email) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Email) GetType() EmailType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return EmailType_unknown
}

func (m *Email) GetHeaders() []byte {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Email) GetSrcip() []byte {
	if m != nil {
		return m.Srcip
	}
	return nil
}

func (m *Email) GetSrchost() []byte {
	if m != nil {
		return m.Srchost
	}
	return nil
}

func (m *Email) GetHelo() []byte {
	if m != nil {
		return m.Helo
	}
	return nil
}

func (m *Email) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Email) GetRcpt() [][]byte {
	if m != nil {
		return m.Rcpt
	}
	return nil
}

func (m *Email) GetBodyurl() [][]byte {
	if m != nil {
		return m.Bodyurl
	}
	return nil
}

func (m *Email) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Email)(nil), "nmsg.base.Email")
	proto.RegisterEnum("nmsg.base.EmailType", EmailType_name, EmailType_value)
}

func init() { proto.RegisterFile("email.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x8e, 0x41, 0x4e, 0xc3, 0x30,
	0x10, 0x45, 0x49, 0x93, 0x90, 0x66, 0x52, 0xc0, 0x8a, 0x58, 0xcc, 0xb2, 0xea, 0x0a, 0xb1, 0xc8,
	0x82, 0x3b, 0x70, 0x03, 0xc4, 0x16, 0xb9, 0xe9, 0x40, 0x4a, 0x13, 0x8f, 0x35, 0x76, 0x54, 0xf5,
	0x40, 0xdc, 0x13, 0xdb, 0x28, 0x74, 0xf9, 0xdf, 0x7b, 0xb6, 0x06, 0x1a, 0x9a, 0xf4, 0x71, 0xec,
	0xac, 0xb0, 0xe7, 0xb6, 0x36, 0x93, 0xfb, 0xea, 0xf6, 0xda, 0xd1, 0xee, 0x27, 0x83, 0xf2, 0x35,
	0xaa, 0x76, 0x07, 0x85, 0xbf, 0x58, 0xc2, 0xf5, 0x36, 0x7b, 0xba, 0x7f, 0x79, 0xec, 0xfe, 0x9b,
	0x2e, 0xf9, 0xb7, 0xe0, 0xda, 0x07, 0xa8, 0x06, 0xd2, 0x07, 0x12, 0x87, 0xab, 0x90, 0x6d, 0xda,
	0x3b, 0x28, 0x9d, 0xf4, 0x47, 0x8b, 0x79, 0x9a, 0xc1, 0x87, 0x39, 0xb0, 0xf3, 0x58, 0x24, 0xb0,
	0x81, 0x62, 0xa0, 0x91, 0xb1, 0x5c, 0xd6, 0xa7, 0xf0, 0x84, 0xb7, 0xcb, 0x92, 0xde, 0x7a, 0xac,
	0xb6, 0xf9, 0xdf, 0xd3, 0x3d, 0x1f, 0x2e, 0xb3, 0x8c, 0x58, 0x27, 0x10, 0x74, 0x04, 0x08, 0x31,
	0x7e, 0x7e, 0x87, 0xfa, 0x7a, 0x46, 0x03, 0xd5, 0x6c, 0x4e, 0x86, 0xcf, 0x46, 0xdd, 0x84, 0x6e,
	0xed, 0xac, 0x9e, 0xbc, 0x68, 0xab, 0xb2, 0xf0, 0x4d, 0x23, 0xf4, 0xfd, 0x61, 0xc8, 0x9f, 0x59,
	0x4e, 0x6a, 0xb5, 0x80, 0x9e, 0x8d, 0x27, 0xe3, 0x55, 0x1e, 0xfb, 0x08, 0x66, 0x47, 0xa2, 0x8a,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xe2, 0x02, 0x36, 0x18, 0x01, 0x00, 0x00,
}
