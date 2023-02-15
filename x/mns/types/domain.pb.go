// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/mns/v1/domain.proto

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

type Domain struct {
	Index      string                                        `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty" yaml:"index"`
	Domain     string                                        `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty" yaml:"domain"`
	DomainType DomainType                                    `protobuf:"bytes,3,opt,name=domain_type,json=domainType,proto3,casttype=DomainType" json:"domain_type,omitempty" yaml:"type"`
	Did        string                                        `protobuf:"bytes,4,opt,name=did,proto3" json:"did,omitempty" yaml:"did"`
	Owner      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,5,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner,omitempty" yaml:"owner"`
	Creator    github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,6,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty" yaml:"creator"`
	ExpireAt   int64                                         `protobuf:"varint,7,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty" yaml:"expire_at"`
}

func (m *Domain) Reset()         { *m = Domain{} }
func (m *Domain) String() string { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()    {}
func (*Domain) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0a76f1de68f8bc9, []int{0}
}
func (m *Domain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Domain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Domain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Domain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Domain.Merge(m, src)
}
func (m *Domain) XXX_Size() int {
	return m.Size()
}
func (m *Domain) XXX_DiscardUnknown() {
	xxx_messageInfo_Domain.DiscardUnknown(m)
}

var xxx_messageInfo_Domain proto.InternalMessageInfo

func (m *Domain) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Domain) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Domain) GetDomainType() DomainType {
	if m != nil {
		return m.DomainType
	}
	return ""
}

func (m *Domain) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *Domain) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Domain) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *Domain) GetExpireAt() int64 {
	if m != nil {
		return m.ExpireAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Domain)(nil), "mantrachain.mns.v1.Domain")
}

func init() { proto.RegisterFile("mantrachain/mns/v1/domain.proto", fileDescriptor_d0a76f1de68f8bc9) }

var fileDescriptor_d0a76f1de68f8bc9 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x31, 0x8b, 0xdb, 0x30,
	0x18, 0x8d, 0xea, 0xc6, 0x69, 0x94, 0x34, 0xa4, 0x22, 0x83, 0x29, 0xd4, 0x32, 0x1e, 0x4a, 0x0a,
	0x8d, 0x4d, 0xe8, 0xd6, 0xa1, 0x60, 0xa7, 0x53, 0xe9, 0x64, 0x0a, 0x85, 0x42, 0x09, 0x8e, 0x25,
	0x12, 0xd1, 0xca, 0x32, 0x96, 0x2f, 0x97, 0xfc, 0x8b, 0xfb, 0x19, 0xf9, 0x29, 0x37, 0x66, 0xbc,
	0xc9, 0x1c, 0xce, 0x3f, 0xf0, 0x78, 0xd3, 0x61, 0xc9, 0x97, 0xcb, 0x7c, 0x93, 0x1e, 0xef, 0xbd,
	0xef, 0x7d, 0x0f, 0xf4, 0x41, 0xcc, 0xe3, 0xb4, 0xc8, 0xe3, 0x64, 0x13, 0xb3, 0xd4, 0xe7, 0xa9,
	0xf4, 0xb7, 0x73, 0x9f, 0x08, 0x1e, 0xb3, 0xd4, 0xcb, 0x72, 0x51, 0x08, 0x84, 0x2e, 0x0c, 0x1e,
	0x4f, 0xa5, 0xb7, 0x9d, 0xbf, 0x9f, 0xac, 0xc5, 0x5a, 0x28, 0xd9, 0x6f, 0x90, 0x76, 0xba, 0x07,
	0x03, 0x9a, 0xdf, 0xd5, 0x28, 0xfa, 0x08, 0xbb, 0x2c, 0x25, 0x74, 0x67, 0x01, 0x07, 0x4c, 0xfb,
	0xe1, 0xb8, 0x2e, 0xf1, 0x70, 0x1f, 0xf3, 0xff, 0x5f, 0x5d, 0x45, 0xbb, 0x91, 0x96, 0xd1, 0x27,
	0x68, 0xea, 0x65, 0xd6, 0x2b, 0x65, 0x7c, 0x57, 0x97, 0xf8, 0xad, 0x36, 0x6a, 0xde, 0x8d, 0x5a,
	0x03, 0xfa, 0x06, 0x07, 0x1a, 0x2d, 0x8b, 0x7d, 0x46, 0x2d, 0x43, 0xf9, 0x3f, 0xd4, 0x25, 0x1e,
	0x68, 0x7f, 0xc3, 0xba, 0x0f, 0x25, 0x86, 0xba, 0xc2, 0xaf, 0x7d, 0x46, 0x23, 0x48, 0xce, 0x18,
	0x39, 0xd0, 0x20, 0x8c, 0x58, 0xaf, 0xd5, 0xdc, 0xa8, 0x2e, 0x31, 0x6c, 0xf7, 0x30, 0xe2, 0x46,
	0x8d, 0x84, 0x7e, 0xc3, 0xae, 0xb8, 0x4e, 0x69, 0x6e, 0x75, 0x1d, 0x30, 0x1d, 0x86, 0xc1, 0x73,
	0x69, 0x45, 0x37, 0xe1, 0xb3, 0x35, 0x2b, 0x36, 0x57, 0x2b, 0x2f, 0x11, 0xdc, 0x4f, 0x84, 0xe4,
	0x42, 0xb6, 0xcf, 0x4c, 0x92, 0x7f, 0x7e, 0x53, 0x41, 0x7a, 0x41, 0x92, 0x04, 0x84, 0xe4, 0x54,
	0xca, 0x48, 0xe7, 0xa1, 0xbf, 0xb0, 0x97, 0xe4, 0x34, 0x2e, 0x44, 0x6e, 0x99, 0x2a, 0x7a, 0x51,
	0x97, 0x78, 0xa4, 0xa3, 0x5b, 0xe1, 0x05, 0xe1, 0x4f, 0x99, 0x68, 0x0e, 0xfb, 0x74, 0x97, 0xb1,
	0x9c, 0x2e, 0xe3, 0xc2, 0xea, 0x39, 0x60, 0x6a, 0x84, 0x93, 0xba, 0xc4, 0x63, 0xbd, 0xe0, 0x2c,
	0xb9, 0xd1, 0x1b, 0x8d, 0x83, 0x22, 0xfc, 0x71, 0xa8, 0x6c, 0x70, 0x5b, 0xd9, 0xe0, 0x58, 0xd9,
	0xe0, 0xbe, 0xb2, 0xc1, 0xcd, 0xc9, 0xee, 0x1c, 0x4f, 0x76, 0xe7, 0xee, 0x64, 0x77, 0xfe, 0x7c,
	0xbe, 0x28, 0xf2, 0x93, 0x71, 0xba, 0xd0, 0xc7, 0x71, 0x71, 0x28, 0x3b, 0x75, 0x2a, 0xaa, 0xd2,
	0xca, 0x54, 0xbf, 0xff, 0xe5, 0x31, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x0a, 0xcd, 0x3b, 0x4a, 0x02,
	0x00, 0x00,
}

func (this *Domain) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Domain)
	if !ok {
		that2, ok := that.(Domain)
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
	if this.DomainType != that1.DomainType {
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
func (m *Domain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Domain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Domain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpireAt != 0 {
		i = encodeVarintDomain(dAtA, i, uint64(m.ExpireAt))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDomain(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintDomain(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintDomain(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DomainType) > 0 {
		i -= len(m.DomainType)
		copy(dAtA[i:], m.DomainType)
		i = encodeVarintDomain(dAtA, i, uint64(len(m.DomainType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintDomain(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintDomain(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDomain(dAtA []byte, offset int, v uint64) int {
	offset -= sovDomain(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Domain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovDomain(uint64(l))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovDomain(uint64(l))
	}
	l = len(m.DomainType)
	if l > 0 {
		n += 1 + l + sovDomain(uint64(l))
	}
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovDomain(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovDomain(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDomain(uint64(l))
	}
	if m.ExpireAt != 0 {
		n += 1 + sovDomain(uint64(m.ExpireAt))
	}
	return n
}

func sovDomain(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDomain(x uint64) (n int) {
	return sovDomain(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Domain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDomain
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
			return fmt.Errorf("proto: Domain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Domain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomain
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
				return ErrInvalidLengthDomain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomain
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
					return ErrIntOverflowDomain
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
				return ErrInvalidLengthDomain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DomainType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomain
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
				return ErrInvalidLengthDomain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DomainType = DomainType(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomain
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
				return ErrInvalidLengthDomain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomain
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
					return ErrIntOverflowDomain
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
				return ErrInvalidLengthDomain
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDomain
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
					return ErrIntOverflowDomain
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
				return ErrInvalidLengthDomain
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDomain
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
					return ErrIntOverflowDomain
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
			skippy, err := skipDomain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDomain
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
func skipDomain(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDomain
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
					return 0, ErrIntOverflowDomain
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
					return 0, ErrIntOverflowDomain
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
				return 0, ErrInvalidLengthDomain
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDomain
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDomain
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDomain        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDomain          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDomain = fmt.Errorf("proto: unexpected end of group")
)
