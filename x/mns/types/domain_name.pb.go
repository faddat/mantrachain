// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/mns/v1/domain_name.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
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

type DomainName struct {
	Index      string                                        `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty" yaml:"index"`
	Domain     string                                        `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty" yaml:"domain"`
	DomainName string                                        `protobuf:"bytes,3,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty" yaml:"domain_name"`
	Did        string                                        `protobuf:"bytes,4,opt,name=did,proto3" json:"did,omitempty" yaml:"did"`
	Owner      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,5,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner,omitempty" yaml:"owner"`
	Creator    github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,6,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty" yaml:"creator"`
	ExpireAt   int64                                         `protobuf:"varint,7,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty" yaml:"expire_at"`
}

func (m *DomainName) Reset()         { *m = DomainName{} }
func (m *DomainName) String() string { return proto.CompactTextString(m) }
func (*DomainName) ProtoMessage()    {}
func (*DomainName) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ce39d959621ecc, []int{0}
}
func (m *DomainName) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DomainName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DomainName.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DomainName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainName.Merge(m, src)
}
func (m *DomainName) XXX_Size() int {
	return m.Size()
}
func (m *DomainName) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainName.DiscardUnknown(m)
}

var xxx_messageInfo_DomainName proto.InternalMessageInfo

func (m *DomainName) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *DomainName) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *DomainName) GetDomainName() string {
	if m != nil {
		return m.DomainName
	}
	return ""
}

func (m *DomainName) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *DomainName) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *DomainName) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *DomainName) GetExpireAt() int64 {
	if m != nil {
		return m.ExpireAt
	}
	return 0
}

func init() {
	proto.RegisterType((*DomainName)(nil), "mantrachain.mns.v1.DomainName")
}

func init() {
	proto.RegisterFile("mantrachain/mns/v1/domain_name.proto", fileDescriptor_07ce39d959621ecc)
}

var fileDescriptor_07ce39d959621ecc = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x3b, 0xc6, 0xb6, 0x76, 0xac, 0xa5, 0x0e, 0x45, 0x82, 0x8b, 0x24, 0x0c, 0x22, 0x15,
	0x6c, 0x42, 0x71, 0x21, 0xb8, 0x4b, 0xea, 0x4a, 0xc4, 0x45, 0x36, 0x82, 0x20, 0x65, 0x9a, 0x19,
	0xda, 0x41, 0x67, 0xa6, 0x64, 0x62, 0x6d, 0xdf, 0xc2, 0x97, 0x10, 0x7c, 0x14, 0x97, 0x5d, 0xde,
	0x55, 0xb8, 0xa4, 0x6f, 0x90, 0xe5, 0x5d, 0x5d, 0x3a, 0x93, 0x7b, 0x9b, 0xf5, 0x5d, 0xcd, 0xe1,
	0xfc, 0xdf, 0xf9, 0xcf, 0x0f, 0x73, 0xe0, 0x2b, 0x41, 0x64, 0x91, 0x93, 0x6c, 0x43, 0xb8, 0x8c,
	0x84, 0xd4, 0xd1, 0x6e, 0x1e, 0x51, 0x25, 0x08, 0x97, 0x4b, 0x49, 0x04, 0x0b, 0xb7, 0xb9, 0x2a,
	0x14, 0x42, 0x2d, 0x2a, 0x14, 0x52, 0x87, 0xbb, 0xf9, 0xcb, 0xc9, 0x5a, 0xad, 0x95, 0x91, 0xa3,
	0x73, 0x65, 0x49, 0xfc, 0xd7, 0x81, 0xf0, 0xa3, 0x99, 0xff, 0x42, 0x04, 0x43, 0xaf, 0x61, 0x97,
	0x4b, 0xca, 0xf6, 0x2e, 0x08, 0xc0, 0x74, 0x90, 0x8c, 0xeb, 0xd2, 0x1f, 0x1e, 0x88, 0xf8, 0xf9,
	0x01, 0x9b, 0x36, 0x4e, 0xad, 0x8c, 0xde, 0xc0, 0x9e, 0xdd, 0xea, 0x3e, 0x32, 0xe0, 0xf3, 0xba,
	0xf4, 0x9f, 0x59, 0xd0, 0xf6, 0x71, 0xda, 0x00, 0xe8, 0x3d, 0x7c, 0xda, 0x0a, 0xe8, 0x3a, 0x86,
	0x7f, 0x51, 0x97, 0x3e, 0x6a, 0xf3, 0x46, 0xc4, 0x29, 0xa4, 0x97, 0x2c, 0x01, 0x74, 0x28, 0xa7,
	0xee, 0x63, 0x33, 0x30, 0xaa, 0x4b, 0x1f, 0x36, 0x03, 0x9c, 0xe2, 0xf4, 0x2c, 0xa1, 0xaf, 0xb0,
	0xab, 0x7e, 0x4b, 0x96, 0xbb, 0xdd, 0x00, 0x4c, 0x87, 0x49, 0x7c, 0x49, 0x6b, 0xda, 0xf8, 0xa6,
	0xf4, 0x67, 0x6b, 0x5e, 0x6c, 0x7e, 0xad, 0xc2, 0x4c, 0x89, 0x28, 0x53, 0x5a, 0x28, 0xdd, 0x3c,
	0x33, 0x4d, 0x7f, 0x44, 0xc5, 0x61, 0xcb, 0x74, 0x18, 0x67, 0x59, 0x4c, 0x69, 0xce, 0xb4, 0x4e,
	0xad, 0x1f, 0xfa, 0x0e, 0xfb, 0x59, 0xce, 0x48, 0xa1, 0x72, 0xb7, 0x67, 0xac, 0x17, 0x75, 0xe9,
	0x8f, 0xac, 0x75, 0x23, 0x3c, 0xc0, 0xfc, 0xce, 0x13, 0xcd, 0xe1, 0x80, 0xed, 0xb7, 0x3c, 0x67,
	0x4b, 0x52, 0xb8, 0xfd, 0x00, 0x4c, 0x9d, 0x64, 0x52, 0x97, 0xfe, 0xd8, 0x2e, 0xb8, 0x97, 0x70,
	0xfa, 0xc4, 0xd6, 0x71, 0x91, 0x7c, 0xfa, 0x57, 0x79, 0xe0, 0x7f, 0xe5, 0x81, 0x63, 0xe5, 0x81,
	0xeb, 0xca, 0x03, 0x7f, 0x4e, 0x5e, 0xe7, 0x78, 0xf2, 0x3a, 0x57, 0x27, 0xaf, 0xf3, 0xed, 0x6d,
	0x2b, 0xc8, 0x67, 0x2e, 0xd8, 0xc2, 0x9e, 0x47, 0xeb, 0x54, 0xf6, 0xe6, 0x58, 0x4c, 0xa4, 0x55,
	0xcf, 0x7c, 0xfd, 0xbb, 0xdb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0x25, 0x2e, 0x60, 0x4c, 0x02,
	0x00, 0x00,
}

func (this *DomainName) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DomainName)
	if !ok {
		that2, ok := that.(DomainName)
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
	if this.Index != that1.Index {
		return false
	}
	if this.Domain != that1.Domain {
		return false
	}
	if this.DomainName != that1.DomainName {
		return false
	}
	if this.Did != that1.Did {
		return false
	}
	if !bytes.Equal(this.Owner, that1.Owner) {
		return false
	}
	if !bytes.Equal(this.Creator, that1.Creator) {
		return false
	}
	if this.ExpireAt != that1.ExpireAt {
		return false
	}
	return true
}
func (m *DomainName) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DomainName) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DomainName) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpireAt != 0 {
		i = encodeVarintDomainName(dAtA, i, uint64(m.ExpireAt))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDomainName(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintDomainName(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintDomainName(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DomainName) > 0 {
		i -= len(m.DomainName)
		copy(dAtA[i:], m.DomainName)
		i = encodeVarintDomainName(dAtA, i, uint64(len(m.DomainName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintDomainName(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintDomainName(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDomainName(dAtA []byte, offset int, v uint64) int {
	offset -= sovDomainName(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DomainName) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovDomainName(uint64(l))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovDomainName(uint64(l))
	}
	l = len(m.DomainName)
	if l > 0 {
		n += 1 + l + sovDomainName(uint64(l))
	}
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovDomainName(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovDomainName(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDomainName(uint64(l))
	}
	if m.ExpireAt != 0 {
		n += 1 + sovDomainName(uint64(m.ExpireAt))
	}
	return n
}

func sovDomainName(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDomainName(x uint64) (n int) {
	return sovDomainName(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DomainName) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDomainName
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
			return fmt.Errorf("proto: DomainName: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DomainName: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomainName
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
				return ErrInvalidLengthDomainName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomainName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomainName
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
				return ErrInvalidLengthDomainName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomainName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DomainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomainName
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
				return ErrInvalidLengthDomainName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomainName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DomainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomainName
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
				return ErrInvalidLengthDomainName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomainName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomainName
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
				return ErrInvalidLengthDomainName
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDomainName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomainName
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
				return ErrInvalidLengthDomainName
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDomainName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireAt", wireType)
			}
			m.ExpireAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomainName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpireAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDomainName(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDomainName
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
func skipDomainName(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDomainName
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
					return 0, ErrIntOverflowDomainName
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
					return 0, ErrIntOverflowDomainName
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
				return 0, ErrInvalidLengthDomainName
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDomainName
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDomainName
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDomainName        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDomainName          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDomainName = fmt.Errorf("proto: unexpected end of group")
)
