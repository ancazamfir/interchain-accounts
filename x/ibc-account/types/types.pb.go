// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/account/types.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Type int32

const (
	Type_REGISTER Type = 0
	Type_RUNTX    Type = 1
)

var Type_name = map[int32]string{
	0: "REGISTER",
	1: "RUNTX",
}

var Type_value = map[string]int32{
	"REGISTER": 0,
	"RUNTX":    1,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0902a1528b05ff0d, []int{0}
}

type IBCTxRaw struct {
	BodyBytes []byte `protobuf:"bytes,1,opt,name=body_bytes,json=bodyBytes,proto3" json:"body_bytes,omitempty"`
}

func (m *IBCTxRaw) Reset()         { *m = IBCTxRaw{} }
func (m *IBCTxRaw) String() string { return proto.CompactTextString(m) }
func (*IBCTxRaw) ProtoMessage()    {}
func (*IBCTxRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_0902a1528b05ff0d, []int{0}
}
func (m *IBCTxRaw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCTxRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCTxRaw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCTxRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCTxRaw.Merge(m, src)
}
func (m *IBCTxRaw) XXX_Size() int {
	return m.Size()
}
func (m *IBCTxRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCTxRaw.DiscardUnknown(m)
}

var xxx_messageInfo_IBCTxRaw proto.InternalMessageInfo

func (m *IBCTxRaw) GetBodyBytes() []byte {
	if m != nil {
		return m.BodyBytes
	}
	return nil
}

type IBCTxBody struct {
	Messages []*types.Any `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (m *IBCTxBody) Reset()         { *m = IBCTxBody{} }
func (m *IBCTxBody) String() string { return proto.CompactTextString(m) }
func (*IBCTxBody) ProtoMessage()    {}
func (*IBCTxBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_0902a1528b05ff0d, []int{1}
}
func (m *IBCTxBody) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCTxBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCTxBody.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCTxBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCTxBody.Merge(m, src)
}
func (m *IBCTxBody) XXX_Size() int {
	return m.Size()
}
func (m *IBCTxBody) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCTxBody.DiscardUnknown(m)
}

var xxx_messageInfo_IBCTxBody proto.InternalMessageInfo

func (m *IBCTxBody) GetMessages() []*types.Any {
	if m != nil {
		return m.Messages
	}
	return nil
}

type IBCAccountPacketData struct {
	Type Type   `protobuf:"varint,1,opt,name=type,proto3,enum=ibc.account.Type" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *IBCAccountPacketData) Reset()         { *m = IBCAccountPacketData{} }
func (m *IBCAccountPacketData) String() string { return proto.CompactTextString(m) }
func (*IBCAccountPacketData) ProtoMessage()    {}
func (*IBCAccountPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0902a1528b05ff0d, []int{2}
}
func (m *IBCAccountPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCAccountPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCAccountPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCAccountPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCAccountPacketData.Merge(m, src)
}
func (m *IBCAccountPacketData) XXX_Size() int {
	return m.Size()
}
func (m *IBCAccountPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCAccountPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_IBCAccountPacketData proto.InternalMessageInfo

func (m *IBCAccountPacketData) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_REGISTER
}

func (m *IBCAccountPacketData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type IBCAccountPacketAcknowledgement struct {
	Type    Type   `protobuf:"varint,1,opt,name=type,proto3,enum=ibc.account.Type" json:"type,omitempty"`
	ChainID string `protobuf:"bytes,2,opt,name=chainID,proto3" json:"chainID,omitempty"`
	Code    uint32 `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Data    []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Error   string `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *IBCAccountPacketAcknowledgement) Reset()         { *m = IBCAccountPacketAcknowledgement{} }
func (m *IBCAccountPacketAcknowledgement) String() string { return proto.CompactTextString(m) }
func (*IBCAccountPacketAcknowledgement) ProtoMessage()    {}
func (*IBCAccountPacketAcknowledgement) Descriptor() ([]byte, []int) {
	return fileDescriptor_0902a1528b05ff0d, []int{3}
}
func (m *IBCAccountPacketAcknowledgement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCAccountPacketAcknowledgement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCAccountPacketAcknowledgement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCAccountPacketAcknowledgement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCAccountPacketAcknowledgement.Merge(m, src)
}
func (m *IBCAccountPacketAcknowledgement) XXX_Size() int {
	return m.Size()
}
func (m *IBCAccountPacketAcknowledgement) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCAccountPacketAcknowledgement.DiscardUnknown(m)
}

var xxx_messageInfo_IBCAccountPacketAcknowledgement proto.InternalMessageInfo

func (m *IBCAccountPacketAcknowledgement) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_REGISTER
}

func (m *IBCAccountPacketAcknowledgement) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *IBCAccountPacketAcknowledgement) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *IBCAccountPacketAcknowledgement) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *IBCAccountPacketAcknowledgement) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterEnum("ibc.account.Type", Type_name, Type_value)
	proto.RegisterType((*IBCTxRaw)(nil), "ibc.account.IBCTxRaw")
	proto.RegisterType((*IBCTxBody)(nil), "ibc.account.IBCTxBody")
	proto.RegisterType((*IBCAccountPacketData)(nil), "ibc.account.IBCAccountPacketData")
	proto.RegisterType((*IBCAccountPacketAcknowledgement)(nil), "ibc.account.IBCAccountPacketAcknowledgement")
}

func init() { proto.RegisterFile("ibc/account/types.proto", fileDescriptor_0902a1528b05ff0d) }

var fileDescriptor_0902a1528b05ff0d = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0xc6, 0x3b, 0xde, 0xa2, 0x74, 0xee, 0xd5, 0xe0, 0x84, 0xc4, 0x6a, 0x62, 0x21, 0x4d, 0x4c,
	0xd0, 0x84, 0x19, 0x83, 0x2b, 0x17, 0x2e, 0x28, 0x10, 0xd3, 0x8d, 0xd1, 0xa1, 0x26, 0xc6, 0x8d,
	0x99, 0x99, 0x8e, 0xa5, 0x81, 0x76, 0x48, 0x3b, 0x04, 0xfa, 0x16, 0x3e, 0x82, 0x8f, 0xe3, 0x92,
	0xa5, 0x4b, 0x03, 0x2f, 0x62, 0x3a, 0x05, 0xfc, 0xb3, 0xba, 0xbb, 0x73, 0xce, 0xf7, 0xcd, 0xf9,
	0x7e, 0x93, 0x19, 0xf8, 0x28, 0xe5, 0x82, 0x30, 0x21, 0xd4, 0x26, 0xd7, 0x44, 0x57, 0x6b, 0x59,
	0xe2, 0x75, 0xa1, 0xb4, 0x42, 0xd7, 0x29, 0x17, 0xf8, 0x24, 0x3c, 0x79, 0x9c, 0x28, 0x95, 0xac,
	0x24, 0x31, 0x12, 0xdf, 0x7c, 0x25, 0x2c, 0xaf, 0x1a, 0x9f, 0xff, 0x1c, 0xb6, 0xc3, 0x60, 0x12,
	0xed, 0x28, 0xdb, 0xa2, 0xa7, 0x10, 0x72, 0x15, 0x57, 0x5f, 0x78, 0xa5, 0x65, 0xe9, 0x82, 0x3e,
	0x18, 0xdc, 0x50, 0xa7, 0x9e, 0x04, 0xf5, 0xc0, 0x7f, 0x03, 0x1d, 0x63, 0x0d, 0x54, 0x5c, 0xa1,
	0x97, 0xb0, 0x9d, 0xc9, 0xb2, 0x64, 0x89, 0x71, 0x5e, 0x0d, 0xae, 0x47, 0x5d, 0xdc, 0xa4, 0xe0,
	0x73, 0x0a, 0x1e, 0xe7, 0x15, 0xbd, 0xb8, 0xfc, 0x0f, 0xb0, 0x1b, 0x06, 0x93, 0x71, 0x83, 0xf4,
	0x9e, 0x89, 0xa5, 0xd4, 0x53, 0xa6, 0x19, 0x7a, 0x06, 0xed, 0x1a, 0xdc, 0xe4, 0x3d, 0x18, 0x3d,
	0xc4, 0x7f, 0x81, 0xe3, 0xa8, 0x5a, 0x4b, 0x6a, 0x64, 0x84, 0xa0, 0x1d, 0x33, 0xcd, 0xdc, 0x3b,
	0x06, 0xcb, 0xd4, 0xfe, 0x77, 0x00, 0x7b, 0xff, 0xef, 0x1c, 0x8b, 0x65, 0xae, 0xb6, 0x2b, 0x19,
	0x27, 0x32, 0x93, 0xb9, 0xbe, 0xed, 0x7a, 0x17, 0xde, 0x13, 0x0b, 0x96, 0xe6, 0xe1, 0xd4, 0x24,
	0x38, 0xf4, 0xdc, 0xd6, 0xc1, 0x42, 0xc5, 0xd2, 0xbd, 0xea, 0x83, 0xc1, 0x7d, 0x6a, 0xea, 0x0b,
	0x8c, 0xfd, 0x07, 0x06, 0x75, 0x61, 0x4b, 0x16, 0x85, 0x2a, 0xdc, 0x96, 0x39, 0xdf, 0x34, 0x2f,
	0x7a, 0xd0, 0xae, 0x53, 0xd0, 0x0d, 0x6c, 0xd3, 0xd9, 0xdb, 0x70, 0x1e, 0xcd, 0x68, 0xc7, 0x42,
	0x0e, 0x6c, 0xd1, 0x8f, 0xef, 0xa2, 0x4f, 0x1d, 0x10, 0xcc, 0x7f, 0x1c, 0x3c, 0xb0, 0x3f, 0x78,
	0xe0, 0xd7, 0xc1, 0x03, 0xdf, 0x8e, 0x9e, 0xb5, 0x3f, 0x7a, 0xd6, 0xcf, 0xa3, 0x67, 0x7d, 0x7e,
	0x9d, 0xa4, 0x7a, 0xb1, 0xe1, 0x58, 0xa8, 0x8c, 0x08, 0x55, 0x66, 0xaa, 0x24, 0x69, 0xae, 0x65,
	0x61, 0xc0, 0x86, 0xa7, 0x3b, 0x94, 0x64, 0x47, 0x52, 0x2e, 0x86, 0xff, 0xfc, 0x01, 0x7e, 0xd7,
	0xbc, 0xc1, 0xab, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x31, 0x83, 0x53, 0xe2, 0x1f, 0x02, 0x00,
	0x00,
}

func (m *IBCTxRaw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCTxRaw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCTxRaw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BodyBytes) > 0 {
		i -= len(m.BodyBytes)
		copy(dAtA[i:], m.BodyBytes)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BodyBytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IBCTxBody) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCTxBody) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCTxBody) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Messages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *IBCAccountPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCAccountPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCAccountPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *IBCAccountPacketAcknowledgement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCAccountPacketAcknowledgement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCAccountPacketAcknowledgement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	if m.Code != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IBCTxRaw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BodyBytes)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *IBCTxBody) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *IBCAccountPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovTypes(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *IBCAccountPacketAcknowledgement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovTypes(uint64(m.Type))
	}
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Code != 0 {
		n += 1 + sovTypes(uint64(m.Code))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IBCTxRaw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IBCTxRaw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCTxRaw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BodyBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BodyBytes = append(m.BodyBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.BodyBytes == nil {
				m.BodyBytes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IBCTxBody) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IBCTxBody: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCTxBody: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, &types.Any{})
			if err := m.Messages[len(m.Messages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IBCAccountPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IBCAccountPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCAccountPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IBCAccountPacketAcknowledgement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IBCAccountPacketAcknowledgement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCAccountPacketAcknowledgement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
