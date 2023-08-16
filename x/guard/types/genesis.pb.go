// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/guard/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the guard module's genesis state.
type GenesisState struct {
	Params                 Params                `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	AccountPrivilegesList  []*AccountPrivileges  `protobuf:"bytes,2,rep,name=account_privileges_list,json=accountPrivilegesList,proto3" json:"account_privileges_list,omitempty"`
	GuardTransferCoins     []byte                `protobuf:"bytes,3,opt,name=guard_transfer_coins,json=guardTransferCoins,proto3" json:"guard_transfer_coins,omitempty"`
	RequiredPrivilegesList []*RequiredPrivileges `protobuf:"bytes,4,rep,name=required_privileges_list,json=requiredPrivilegesList,proto3" json:"required_privileges_list,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ea5a8529cb82408, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetAccountPrivilegesList() []*AccountPrivileges {
	if m != nil {
		return m.AccountPrivilegesList
	}
	return nil
}

func (m *GenesisState) GetGuardTransferCoins() []byte {
	if m != nil {
		return m.GuardTransferCoins
	}
	return nil
}

func (m *GenesisState) GetRequiredPrivilegesList() []*RequiredPrivileges {
	if m != nil {
		return m.RequiredPrivilegesList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "mantrachain.guard.v1.GenesisState")
}

func init() {
	proto.RegisterFile("mantrachain/guard/v1/genesis.proto", fileDescriptor_0ea5a8529cb82408)
}

var fileDescriptor_0ea5a8529cb82408 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4f, 0xc2, 0x30,
	0x18, 0x86, 0x37, 0x20, 0x1c, 0x06, 0xa7, 0x05, 0x75, 0x21, 0x66, 0x22, 0x17, 0x77, 0xb1, 0x15,
	0xb8, 0x79, 0x03, 0x13, 0xbd, 0x88, 0x21, 0x93, 0x93, 0x97, 0xa5, 0x94, 0x5a, 0x9a, 0xb0, 0x76,
	0xb6, 0x1d, 0xd1, 0x7f, 0xe1, 0x5f, 0xf2, 0xc6, 0x91, 0xa3, 0x27, 0x63, 0xe0, 0x8f, 0x18, 0xba,
	0xc5, 0x2c, 0xb2, 0x78, 0x6b, 0xbe, 0xef, 0x79, 0xbf, 0x3e, 0xc9, 0xeb, 0x74, 0x63, 0xc4, 0xb5,
	0x44, 0x78, 0x81, 0x18, 0x87, 0x34, 0x45, 0x72, 0x0e, 0x57, 0x3d, 0x48, 0x09, 0x27, 0x8a, 0x29,
	0x90, 0x48, 0xa1, 0x85, 0xdb, 0x2a, 0x30, 0xc0, 0x30, 0x60, 0xd5, 0x6b, 0xb7, 0xa8, 0xa0, 0xc2,
	0x00, 0x70, 0xff, 0xca, 0xd8, 0xf6, 0x79, 0xe9, 0x3d, 0x2c, 0xe2, 0x58, 0xf0, 0x7f, 0x91, 0x04,
	0x49, 0x14, 0xe7, 0x3f, 0x76, 0x3f, 0x2a, 0x4e, 0xf3, 0x2e, 0x73, 0x78, 0xd4, 0x48, 0x13, 0xf7,
	0xda, 0xa9, 0x67, 0x80, 0x67, 0x77, 0xec, 0xa0, 0xd1, 0x3f, 0x05, 0x65, 0x4e, 0x60, 0x62, 0x98,
	0x51, 0x6d, 0xfd, 0x75, 0x66, 0x85, 0x79, 0xc2, 0x8d, 0x9c, 0x13, 0x84, 0xb1, 0x48, 0xb9, 0x8e,
	0x12, 0xc9, 0x56, 0x6c, 0x49, 0x28, 0x51, 0xd1, 0x92, 0x29, 0xed, 0x55, 0x3a, 0xd5, 0xa0, 0xd1,
	0xbf, 0x28, 0x3f, 0x36, 0xcc, 0x42, 0x93, 0xdf, 0x4c, 0x78, 0x84, 0xfe, 0x8e, 0xee, 0x99, 0xd2,
	0xee, 0x95, 0xd3, 0x32, 0xa1, 0x48, 0x4b, 0xc4, 0xd5, 0x33, 0x91, 0x11, 0x16, 0x8c, 0x2b, 0xaf,
	0xda, 0xb1, 0x83, 0x66, 0xe8, 0x9a, 0xdd, 0x34, 0x5f, 0xdd, 0xec, 0x37, 0xee, 0xcc, 0xf1, 0x24,
	0x79, 0x49, 0x99, 0x24, 0xf3, 0x03, 0xa7, 0x9a, 0x71, 0x0a, 0xca, 0x9d, 0xc2, 0x3c, 0x55, 0x90,
	0x3a, 0x96, 0x07, 0xb3, 0xbd, 0xd5, 0x68, 0xbc, 0xde, 0xfa, 0xf6, 0x66, 0xeb, 0xdb, 0xdf, 0x5b,
	0xdf, 0x7e, 0xdf, 0xf9, 0xd6, 0x66, 0xe7, 0x5b, 0x9f, 0x3b, 0xdf, 0x7a, 0x1a, 0x50, 0xa6, 0x17,
	0xe9, 0x0c, 0x60, 0x11, 0xc3, 0xf1, 0xf0, 0x61, 0x1a, 0x0e, 0x2f, 0x6f, 0x19, 0x47, 0x1c, 0x13,
	0x58, 0xac, 0xe6, 0x35, 0x2f, 0x47, 0xbf, 0x25, 0x44, 0xcd, 0xea, 0xa6, 0x99, 0xc1, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x4a, 0xe4, 0xfd, 0xd2, 0x31, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RequiredPrivilegesList) > 0 {
		for iNdEx := len(m.RequiredPrivilegesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RequiredPrivilegesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.GuardTransferCoins) > 0 {
		i -= len(m.GuardTransferCoins)
		copy(dAtA[i:], m.GuardTransferCoins)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.GuardTransferCoins)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AccountPrivilegesList) > 0 {
		for iNdEx := len(m.AccountPrivilegesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccountPrivilegesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.AccountPrivilegesList) > 0 {
		for _, e := range m.AccountPrivilegesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = len(m.GuardTransferCoins)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.RequiredPrivilegesList) > 0 {
		for _, e := range m.RequiredPrivilegesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountPrivilegesList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountPrivilegesList = append(m.AccountPrivilegesList, &AccountPrivileges{})
			if err := m.AccountPrivilegesList[len(m.AccountPrivilegesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuardTransferCoins", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GuardTransferCoins = append(m.GuardTransferCoins[:0], dAtA[iNdEx:postIndex]...)
			if m.GuardTransferCoins == nil {
				m.GuardTransferCoins = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequiredPrivilegesList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequiredPrivilegesList = append(m.RequiredPrivilegesList, &RequiredPrivileges{})
			if err := m.RequiredPrivilegesList[len(m.RequiredPrivilegesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
