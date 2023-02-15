// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/marketplace/v1/marketplace.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type Marketplace struct {
	Index       []byte                                        `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty" yaml:"index"`
	Id          string                                        `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Name        string                                        `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
	Url         string                                        `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty" yaml:"url"`
	Description string                                        `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	Opened      bool                                          `protobuf:"varint,6,opt,name=opened,proto3" json:"opened,omitempty" yaml:"opened"`
	Options     []*MarketplaceOption                          `protobuf:"bytes,7,rep,name=options,proto3" json:"options,omitempty" yaml:"options"`
	Attributes  []*MarketplaceAttribute                       `protobuf:"bytes,8,rep,name=attributes,proto3" json:"attributes,omitempty" yaml:"attributes"`
	Images      []*MarketplaceImage                           `protobuf:"bytes,9,rep,name=images,proto3" json:"images,omitempty" yaml:"images"`
	Links       []*MarketplaceLink                            `protobuf:"bytes,10,rep,name=links,proto3" json:"links,omitempty" yaml:"links"`
	Data        *types.Any                                    `protobuf:"bytes,11,opt,name=data,proto3" json:"data,omitempty" yaml:"data"`
	Owner       github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,12,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner,omitempty" yaml:"owner"`
	Creator     github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,13,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty" yaml:"creator"`
}

func (m *Marketplace) Reset()         { *m = Marketplace{} }
func (m *Marketplace) String() string { return proto.CompactTextString(m) }
func (*Marketplace) ProtoMessage()    {}
func (*Marketplace) Descriptor() ([]byte, []int) {
	return fileDescriptor_80aa58710d9765e1, []int{0}
}
func (m *Marketplace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Marketplace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Marketplace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Marketplace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Marketplace.Merge(m, src)
}
func (m *Marketplace) XXX_Size() int {
	return m.Size()
}
func (m *Marketplace) XXX_DiscardUnknown() {
	xxx_messageInfo_Marketplace.DiscardUnknown(m)
}

var xxx_messageInfo_Marketplace proto.InternalMessageInfo

func (m *Marketplace) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *Marketplace) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Marketplace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Marketplace) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Marketplace) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Marketplace) GetOpened() bool {
	if m != nil {
		return m.Opened
	}
	return false
}

func (m *Marketplace) GetOptions() []*MarketplaceOption {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Marketplace) GetAttributes() []*MarketplaceAttribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Marketplace) GetImages() []*MarketplaceImage {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *Marketplace) GetLinks() []*MarketplaceLink {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Marketplace) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Marketplace) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Marketplace) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func init() {
	proto.RegisterType((*Marketplace)(nil), "mantrachain.marketplace.v1.Marketplace")
}

func init() {
	proto.RegisterFile("mantrachain/marketplace/v1/marketplace.proto", fileDescriptor_80aa58710d9765e1)
}

var fileDescriptor_80aa58710d9765e1 = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x97, 0xad, 0x7f, 0x56, 0xa7, 0x1d, 0xcc, 0x1a, 0xc8, 0x54, 0x22, 0x89, 0x8c, 0x04,
	0x45, 0xac, 0x09, 0x1b, 0x1c, 0x80, 0x5b, 0xb3, 0x13, 0xd2, 0x10, 0x92, 0x39, 0x4c, 0x02, 0x71,
	0x70, 0x13, 0x93, 0x59, 0x4d, 0xec, 0x2a, 0x49, 0xc7, 0xfa, 0x2d, 0xf8, 0x18, 0x7c, 0x14, 0x8e,
	0x3b, 0x72, 0x8a, 0x50, 0x7b, 0xe0, 0x9e, 0x23, 0x27, 0x54, 0x3b, 0x55, 0xc3, 0x01, 0x54, 0x71,
	0x8a, 0xf3, 0x3e, 0xcf, 0xf3, 0x7b, 0x25, 0xfb, 0x7d, 0xc1, 0x71, 0x42, 0x45, 0x9e, 0xd2, 0xe0,
	0x92, 0x72, 0xe1, 0x25, 0x34, 0x9d, 0xb0, 0x7c, 0x1a, 0xd3, 0x80, 0x79, 0x57, 0x27, 0xf5, 0x5f,
	0x77, 0x9a, 0xca, 0x5c, 0xc2, 0x7e, 0xcd, 0xed, 0xd6, 0xe5, 0xab, 0x93, 0xfe, 0xbd, 0x48, 0xca,
	0x28, 0x66, 0x9e, 0x72, 0x8e, 0x67, 0x9f, 0x3c, 0x2a, 0xe6, 0x3a, 0xd6, 0x3f, 0x8a, 0x64, 0x24,
	0xd5, 0xd1, 0x5b, 0x9d, 0xaa, 0xea, 0xa3, 0x7f, 0xb4, 0x0e, 0x64, 0x92, 0x48, 0xa1, 0x8d, 0xf8,
	0x67, 0x0b, 0x98, 0x6f, 0x36, 0x3a, 0x7c, 0x08, 0x9a, 0x5c, 0x84, 0xec, 0x1a, 0x19, 0x8e, 0x31,
	0xe8, 0xfa, 0xb7, 0xcb, 0xc2, 0xee, 0xce, 0x69, 0x12, 0xbf, 0xc2, 0xaa, 0x8c, 0x89, 0x96, 0xe1,
	0x7d, 0xb0, 0xcb, 0x43, 0xb4, 0xeb, 0x18, 0x83, 0x8e, 0xdf, 0x2b, 0x0b, 0xbb, 0x53, 0x99, 0x42,
	0x4c, 0x76, 0x79, 0x08, 0x1f, 0x80, 0x86, 0xa0, 0x09, 0x43, 0x7b, 0xca, 0x70, 0xab, 0x2c, 0x6c,
	0x53, 0x1b, 0x56, 0x55, 0x4c, 0x94, 0x08, 0x1d, 0xb0, 0x37, 0x4b, 0x63, 0xd4, 0x50, 0x9e, 0x83,
	0xb2, 0xb0, 0x81, 0xf6, 0xcc, 0xd2, 0x18, 0x93, 0x95, 0x04, 0x5f, 0x00, 0x33, 0x64, 0x59, 0x90,
	0xf2, 0x69, 0xce, 0xa5, 0x40, 0x4d, 0xe5, 0xbc, 0x5b, 0x16, 0x36, 0xd4, 0xce, 0x9a, 0x88, 0x49,
	0xdd, 0x0a, 0x1f, 0x83, 0x96, 0x9c, 0x32, 0xc1, 0x42, 0xd4, 0x72, 0x8c, 0xc1, 0xbe, 0x7f, 0x58,
	0x16, 0x76, 0x4f, 0x87, 0x74, 0x1d, 0x93, 0xca, 0x00, 0x3f, 0x80, 0xb6, 0x54, 0xa1, 0x0c, 0xb5,
	0x9d, 0xbd, 0x81, 0x79, 0x3a, 0x74, 0xff, 0xfe, 0x14, 0x6e, 0xed, 0xb2, 0xde, 0xaa, 0x94, 0x0f,
	0xcb, 0xc2, 0x3e, 0x58, 0xa3, 0x15, 0x07, 0x93, 0x35, 0x11, 0x46, 0x00, 0xd0, 0x3c, 0x4f, 0xf9,
	0x78, 0x96, 0xb3, 0x0c, 0xed, 0x2b, 0xfe, 0xd3, 0x2d, 0xf9, 0xa3, 0x75, 0xd0, 0xbf, 0x53, 0x16,
	0xf6, 0xa1, 0x6e, 0xb1, 0xa1, 0x61, 0x52, 0x43, 0xc3, 0x0b, 0xd0, 0xe2, 0x09, 0x8d, 0x58, 0x86,
	0x3a, 0xaa, 0xc9, 0xf1, 0x96, 0x4d, 0x5e, 0xaf, 0x42, 0xf5, 0xeb, 0xd1, 0x14, 0x4c, 0x2a, 0x1c,
	0x7c, 0x07, 0x9a, 0x31, 0x17, 0x93, 0x0c, 0x01, 0xc5, 0x7d, 0xb2, 0x25, 0xf7, 0x9c, 0x8b, 0x49,
	0x7d, 0x7c, 0x14, 0x03, 0x13, 0xcd, 0x82, 0x2f, 0x41, 0x23, 0xa4, 0x39, 0x45, 0xa6, 0x63, 0x0c,
	0xcc, 0xd3, 0x23, 0x57, 0xcf, 0xb7, 0xbb, 0x9e, 0x6f, 0x77, 0x24, 0xe6, 0xf5, 0xa9, 0x59, 0x79,
	0x31, 0x51, 0x11, 0x78, 0x01, 0x9a, 0xf2, 0xb3, 0x60, 0x29, 0xea, 0xaa, 0x09, 0x1d, 0x6d, 0x5a,
	0xa8, 0x32, 0xfe, 0x55, 0xd8, 0xc3, 0x88, 0xe7, 0x97, 0xb3, 0xb1, 0x1b, 0xc8, 0xc4, 0x0b, 0x64,
	0x96, 0xc8, 0xac, 0xfa, 0x0c, 0xb3, 0x70, 0xe2, 0xe5, 0xf3, 0x29, 0xcb, 0xdc, 0x51, 0x10, 0x8c,
	0xc2, 0x30, 0x65, 0x59, 0x46, 0x34, 0x0f, 0x7e, 0x04, 0xed, 0x20, 0x65, 0x34, 0x97, 0x29, 0xea,
	0x29, 0xf4, 0xd9, 0xe6, 0x61, 0x2b, 0xe1, 0x3f, 0xe0, 0x6b, 0xa6, 0x4f, 0xbe, 0x2e, 0x2c, 0xe3,
	0xdb, 0xc2, 0x32, 0x6e, 0x16, 0x96, 0xf1, 0x63, 0x61, 0x19, 0x5f, 0x96, 0xd6, 0xce, 0xcd, 0xd2,
	0xda, 0xf9, 0xbe, 0xb4, 0x76, 0xde, 0x3f, 0xaf, 0x51, 0xcf, 0x79, 0xc2, 0xce, 0xaa, 0xcd, 0xdd,
	0x6c, 0xf1, 0xf5, 0x1f, 0x7b, 0xac, 0xfa, 0x8c, 0x5b, 0xea, 0xc2, 0x9e, 0xfd, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0xca, 0x4b, 0x53, 0x8c, 0x6a, 0x04, 0x00, 0x00,
}

func (this *Marketplace) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Marketplace)
	if !ok {
		that2, ok := that.(Marketplace)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Index, that1.Index) {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Url != that1.Url {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Opened != that1.Opened {
		return false
	}
	if len(this.Options) != len(that1.Options) {
		return false
	}
	for i := range this.Options {
		if !this.Options[i].Equal(that1.Options[i]) {
			return false
		}
	}
	if len(this.Attributes) != len(that1.Attributes) {
		return false
	}
	for i := range this.Attributes {
		if !this.Attributes[i].Equal(that1.Attributes[i]) {
			return false
		}
	}
	if len(this.Images) != len(that1.Images) {
		return false
	}
	for i := range this.Images {
		if !this.Images[i].Equal(that1.Images[i]) {
			return false
		}
	}
	if len(this.Links) != len(that1.Links) {
		return false
	}
	for i := range this.Links {
		if !this.Links[i].Equal(that1.Links[i]) {
			return false
		}
	}
	if !this.Data.Equal(that1.Data) {
		return false
	}
	if !bytes.Equal(this.Owner, that1.Owner) {
		return false
	}
	if !bytes.Equal(this.Creator, that1.Creator) {
		return false
	}
	return true
}
func (m *Marketplace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Marketplace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Marketplace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMarketplace(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintMarketplace(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x62
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMarketplace(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Links) > 0 {
		for iNdEx := len(m.Links) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Links[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMarketplace(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.Images) > 0 {
		for iNdEx := len(m.Images) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Images[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMarketplace(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMarketplace(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Options) > 0 {
		for iNdEx := len(m.Options) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Options[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMarketplace(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.Opened {
		i--
		if m.Opened {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintMarketplace(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintMarketplace(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMarketplace(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintMarketplace(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintMarketplace(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMarketplace(dAtA []byte, offset int, v uint64) int {
	offset -= sovMarketplace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Marketplace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovMarketplace(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovMarketplace(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMarketplace(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovMarketplace(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMarketplace(uint64(l))
	}
	if m.Opened {
		n += 2
	}
	if len(m.Options) > 0 {
		for _, e := range m.Options {
			l = e.Size()
			n += 1 + l + sovMarketplace(uint64(l))
		}
	}
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovMarketplace(uint64(l))
		}
	}
	if len(m.Images) > 0 {
		for _, e := range m.Images {
			l = e.Size()
			n += 1 + l + sovMarketplace(uint64(l))
		}
	}
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovMarketplace(uint64(l))
		}
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovMarketplace(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovMarketplace(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMarketplace(uint64(l))
	}
	return n
}

func sovMarketplace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMarketplace(x uint64) (n int) {
	return sovMarketplace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Marketplace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarketplace
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
			return fmt.Errorf("proto: Marketplace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Marketplace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketplace
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
				return ErrInvalidLengthMarketplace
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketplace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = append(m.Index[:0], dAtA[iNdEx:postIndex]...)
			if m.Index == nil {
				m.Index = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketplace
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
				return ErrInvalidLengthMarketplace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketplace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketplace
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
				return ErrInvalidLengthMarketplace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketplace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketplace
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
				return ErrInvalidLengthMarketplace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketplace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketplace
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
				return ErrInvalidLengthMarketplace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketplace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Opened", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketplace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Opened = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketplace
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
				return ErrInvalidLengthMarketplace
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarketplace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = append(m.Options, &MarketplaceOption{})
			if err := m.Options[len(m.Options)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketplace
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
				return ErrInvalidLengthMarketplace
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarketplace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, &MarketplaceAttribute{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Images", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketplace
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
				return ErrInvalidLengthMarketplace
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarketplace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Images = append(m.Images, &MarketplaceImage{})
			if err := m.Images[len(m.Images)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketplace
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
				return ErrInvalidLengthMarketplace
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarketplace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, &MarketplaceLink{})
			if err := m.Links[len(m.Links)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketplace
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
				return ErrInvalidLengthMarketplace
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarketplace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketplace
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
				return ErrInvalidLengthMarketplace
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketplace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketplace
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
				return ErrInvalidLengthMarketplace
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketplace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarketplace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarketplace
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
func skipMarketplace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMarketplace
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
					return 0, ErrIntOverflowMarketplace
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
					return 0, ErrIntOverflowMarketplace
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
				return 0, ErrInvalidLengthMarketplace
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMarketplace
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMarketplace
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMarketplace        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMarketplace          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMarketplace = fmt.Errorf("proto: unexpected end of group")
)
