// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/vault/v1/common.proto

package types

import (
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

type VaultEarning struct {
	Type       string                                  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" yaml:"type"`
	Address    string                                  `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Percentage *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=percentage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"percentage,omitempty" yaml:"percentage"`
}

func (m *VaultEarning) Reset()         { *m = VaultEarning{} }
func (m *VaultEarning) String() string { return proto.CompactTextString(m) }
func (*VaultEarning) ProtoMessage()    {}
func (*VaultEarning) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7730c85ebb1b39a, []int{0}
}
func (m *VaultEarning) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultEarning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VaultEarning.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VaultEarning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultEarning.Merge(m, src)
}
func (m *VaultEarning) XXX_Size() int {
	return m.Size()
}
func (m *VaultEarning) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultEarning.DiscardUnknown(m)
}

var xxx_messageInfo_VaultEarning proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VaultEarning)(nil), "mantrachain.vault.v1.VaultEarning")
}

func init() { proto.RegisterFile("mantrachain/vault/v1/common.proto", fileDescriptor_e7730c85ebb1b39a) }

var fileDescriptor_e7730c85ebb1b39a = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcc, 0x4d, 0xcc, 0x2b,
	0x29, 0x4a, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x4b, 0x2c, 0xcd, 0x29, 0xd1, 0x2f, 0x33,
	0xd4, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x41,
	0x52, 0xa2, 0x07, 0x56, 0xa2, 0x57, 0x66, 0x28, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xa0,
	0x0f, 0x62, 0x41, 0xd4, 0x2a, 0x9d, 0x61, 0xe4, 0xe2, 0x09, 0x03, 0x29, 0x71, 0x4d, 0x2c, 0xca,
	0xcb, 0xcc, 0x4b, 0x17, 0x52, 0xe6, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x74, 0xe2, 0xff, 0x74, 0x4f, 0x9e, 0xbb, 0x32, 0x31, 0x37, 0xc7, 0x4a, 0x09, 0x24, 0xaa,
	0x14, 0x04, 0x96, 0x14, 0xd2, 0xe1, 0x62, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x96, 0x60,
	0x02, 0xab, 0x13, 0xfa, 0x74, 0x4f, 0x9e, 0x0f, 0xa2, 0x0e, 0x2a, 0xa1, 0x14, 0x04, 0x53, 0x22,
	0x94, 0xcc, 0xc5, 0x55, 0x90, 0x5a, 0x94, 0x9c, 0x9a, 0x57, 0x92, 0x98, 0x9e, 0x2a, 0xc1, 0x0c,
	0xd6, 0xe0, 0x7c, 0xe2, 0x9e, 0x3c, 0xe3, 0xad, 0x7b, 0xf2, 0x6a, 0xe9, 0x99, 0x25, 0x19, 0xa5,
	0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0x50, 0x4a, 0xb7, 0x38,
	0x25, 0x5b, 0x1f, 0x64, 0x59, 0xb1, 0x9e, 0x67, 0x5e, 0xc9, 0xa7, 0x7b, 0xf2, 0x82, 0x10, 0xe3,
	0x11, 0x26, 0x29, 0x05, 0x21, 0x19, 0x6b, 0xc5, 0xd2, 0xb1, 0x40, 0x9e, 0xc1, 0xc9, 0x67, 0xc5,
	0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71,
	0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x43, 0xb2,
	0xcc, 0x27, 0x33, 0x37, 0xd5, 0x19, 0x1c, 0x88, 0xc8, 0x01, 0x5a, 0x01, 0x0d, 0x52, 0xb0, 0xc5,
	0x49, 0x6c, 0xe0, 0x30, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xab, 0x38, 0x39, 0x25, 0x74,
	0x01, 0x00, 0x00,
}

func (this *VaultEarning) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VaultEarning)
	if !ok {
		that2, ok := that.(VaultEarning)
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
	if this.Type != that1.Type {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if that1.Percentage == nil {
		if this.Percentage != nil {
			return false
		}
	} else if !this.Percentage.Equal(*that1.Percentage) {
		return false
	}
	return true
}
func (m *VaultEarning) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultEarning) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultEarning) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Percentage != nil {
		{
			size := m.Percentage.Size()
			i -= size
			if _, err := m.Percentage.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintCommon(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VaultEarning) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.Percentage != nil {
		l = m.Percentage.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VaultEarning) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: VaultEarning: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultEarning: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Percentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.Percentage = &v
			if err := m.Percentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)
