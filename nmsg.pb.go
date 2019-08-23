// Code generated by protoc-gen-go.
// source: nmsg.proto
// DO NOT EDIT!

/*
Package nmsg is a generated protocol buffer package.

It is generated from these files:
	nmsg.proto

It has these top-level messages:
	Nmsg
	NmsgFragment
	NmsgPayload
*/
package nmsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Nmsg struct {
	Payloads         []*NmsgPayload `protobuf:"bytes,1,rep,name=payloads" json:"payloads,omitempty"`
	PayloadCrcs      []uint32       `protobuf:"varint,2,rep,name=payload_crcs" json:"payload_crcs,omitempty"`
	Sequence         *uint32        `protobuf:"varint,3,opt,name=sequence" json:"sequence,omitempty"`
	SequenceId       *uint64        `protobuf:"varint,4,opt,name=sequence_id" json:"sequence_id,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Nmsg) Reset()                    { *m = Nmsg{} }
func (m *Nmsg) String() string            { return proto.CompactTextString(m) }
func (*Nmsg) ProtoMessage()               {}
func (*Nmsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Nmsg) GetPayloads() []*NmsgPayload {
	if m != nil {
		return m.Payloads
	}
	return nil
}

func (m *Nmsg) GetPayloadCrcs() []uint32 {
	if m != nil {
		return m.PayloadCrcs
	}
	return nil
}

func (m *Nmsg) GetSequence() uint32 {
	if m != nil && m.Sequence != nil {
		return *m.Sequence
	}
	return 0
}

func (m *Nmsg) GetSequenceId() uint64 {
	if m != nil && m.SequenceId != nil {
		return *m.SequenceId
	}
	return 0
}

type NmsgFragment struct {
	Id               *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Current          *uint32 `protobuf:"varint,2,req,name=current" json:"current,omitempty"`
	Last             *uint32 `protobuf:"varint,3,req,name=last" json:"last,omitempty"`
	Fragment         []byte  `protobuf:"bytes,4,req,name=fragment" json:"fragment,omitempty"`
	Crc              *uint32 `protobuf:"varint,5,opt,name=crc" json:"crc,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NmsgFragment) Reset()                    { *m = NmsgFragment{} }
func (m *NmsgFragment) String() string            { return proto.CompactTextString(m) }
func (*NmsgFragment) ProtoMessage()               {}
func (*NmsgFragment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NmsgFragment) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *NmsgFragment) GetCurrent() uint32 {
	if m != nil && m.Current != nil {
		return *m.Current
	}
	return 0
}

func (m *NmsgFragment) GetLast() uint32 {
	if m != nil && m.Last != nil {
		return *m.Last
	}
	return 0
}

func (m *NmsgFragment) GetFragment() []byte {
	if m != nil {
		return m.Fragment
	}
	return nil
}

func (m *NmsgFragment) GetCrc() uint32 {
	if m != nil && m.Crc != nil {
		return *m.Crc
	}
	return 0
}

type NmsgPayload struct {
	Vid              *uint32 `protobuf:"varint,1,req,name=vid" json:"vid,omitempty"`
	Msgtype          *uint32 `protobuf:"varint,2,req,name=msgtype" json:"msgtype,omitempty"`
	TimeSec          *int64  `protobuf:"varint,3,req,name=time_sec" json:"time_sec,omitempty"`
	TimeNsec         *uint32 `protobuf:"fixed32,4,req,name=time_nsec" json:"time_nsec,omitempty"`
	Payload          []byte  `protobuf:"bytes,5,opt,name=payload" json:"payload,omitempty"`
	Source           *uint32 `protobuf:"varint,7,opt,name=source" json:"source,omitempty"`
	Operator         *uint32 `protobuf:"varint,8,opt,name=operator" json:"operator,omitempty"`
	Group            *uint32 `protobuf:"varint,9,opt,name=group" json:"group,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NmsgPayload) Reset()                    { *m = NmsgPayload{} }
func (m *NmsgPayload) String() string            { return proto.CompactTextString(m) }
func (*NmsgPayload) ProtoMessage()               {}
func (*NmsgPayload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NmsgPayload) GetVid() uint32 {
	if m != nil && m.Vid != nil {
		return *m.Vid
	}
	return 0
}

func (m *NmsgPayload) GetMsgtype() uint32 {
	if m != nil && m.Msgtype != nil {
		return *m.Msgtype
	}
	return 0
}

func (m *NmsgPayload) GetTimeSec() int64 {
	if m != nil && m.TimeSec != nil {
		return *m.TimeSec
	}
	return 0
}

func (m *NmsgPayload) GetTimeNsec() uint32 {
	if m != nil && m.TimeNsec != nil {
		return *m.TimeNsec
	}
	return 0
}

func (m *NmsgPayload) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *NmsgPayload) GetSource() uint32 {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return 0
}

func (m *NmsgPayload) GetOperator() uint32 {
	if m != nil && m.Operator != nil {
		return *m.Operator
	}
	return 0
}

func (m *NmsgPayload) GetGroup() uint32 {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return 0
}

func init() {
	proto.RegisterType((*Nmsg)(nil), "nmsg.Nmsg")
	proto.RegisterType((*NmsgFragment)(nil), "nmsg.NmsgFragment")
	proto.RegisterType((*NmsgPayload)(nil), "nmsg.NmsgPayload")
}

func init() { proto.RegisterFile("nmsg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0xc7, 0x25, 0xe9, 0xc5, 0x05, 0x6a, 0x18, 0x3c, 0x46, 0x61, 0xe9, 0xd4, 0x81,
	0x87, 0x60, 0x44, 0x8c, 0x6c, 0x91, 0xe5, 0x9a, 0x50, 0xa9, 0x89, 0x8d, 0xed, 0x20, 0xf5, 0x35,
	0x78, 0x62, 0xee, 0xdc, 0x54, 0xea, 0x94, 0xf8, 0xfb, 0xef, 0xbe, 0xbb, 0x03, 0x98, 0xc6, 0x38,
	0xec, 0x7d, 0x70, 0xc9, 0x49, 0x4e, 0xff, 0xdd, 0x37, 0xf0, 0x77, 0xfc, 0xca, 0x17, 0xa8, 0xbd,
	0x3e, 0x9f, 0x9c, 0x3e, 0x44, 0x55, 0xb4, 0xe5, 0xae, 0x79, 0xdd, 0xee, 0x73, 0x31, 0xa5, 0x1f,
	0x97, 0x44, 0x3e, 0x83, 0x58, 0x8a, 0x7a, 0x13, 0x4c, 0x54, 0x0c, 0x0b, 0x37, 0xf2, 0x11, 0xea,
	0x68, 0x7f, 0x66, 0x3b, 0x19, 0xab, 0xca, 0xb6, 0x40, 0xf2, 0x04, 0xcd, 0x95, 0xf4, 0xc7, 0x83,
	0xe2, 0x08, 0x79, 0xf7, 0x09, 0x82, 0x5c, 0x6f, 0x41, 0x0f, 0xa3, 0x9d, 0x92, 0x04, 0x60, 0x98,
	0x15, 0x2d, 0xc3, 0x86, 0x07, 0xa8, 0xcc, 0x1c, 0x02, 0x62, 0x74, 0x12, 0x10, 0xc0, 0x4f, 0x3a,
	0x26, 0xf4, 0xb1, 0xcb, 0x84, 0xaf, 0xa5, 0x0d, 0x65, 0x6c, 0x27, 0x64, 0x03, 0x25, 0x6e, 0xa0,
	0x56, 0x34, 0xae, 0xfb, 0x2b, 0xa0, 0xb9, 0x5d, 0x13, 0xc3, 0xdf, 0x5b, 0x35, 0x46, 0xe9, 0xec,
	0xed, 0xa2, 0x46, 0x59, 0x3a, 0x8e, 0xb6, 0x8f, 0xd6, 0x64, 0x7d, 0x29, 0xb7, 0xb0, 0xce, 0x64,
	0x22, 0x44, 0xfe, 0x8a, 0xba, 0x96, 0x4b, 0xf3, 0x0c, 0x21, 0xef, 0xe1, 0x2e, 0xba, 0x39, 0xe0,
	0x89, 0x55, 0x3e, 0x11, 0x2d, 0xce, 0xdb, 0xa0, 0x93, 0x0b, 0xaa, 0xce, 0x64, 0x03, 0xab, 0x21,
	0xb8, 0xd9, 0xab, 0x35, 0x3d, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x73, 0x07, 0x52, 0x34, 0x6b,
	0x01, 0x00, 0x00,
}