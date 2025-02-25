// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/tokenfactory/v1beta1/genesis.proto

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

// GenesisState defines the tokenfactory module's genesis state.
type GenesisState struct {
	// params defines the parameters of the module.
	Params        Params         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	FactoryDenoms []GenesisDenom `protobuf:"bytes,2,rep,name=factory_denoms,json=factoryDenoms,proto3" json:"factory_denoms" yaml:"factory_denoms"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5749c3f71850298b, []int{0}
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

func (m *GenesisState) GetFactoryDenoms() []GenesisDenom {
	if m != nil {
		return m.FactoryDenoms
	}
	return nil
}

// GenesisDenom defines a tokenfactory denom that is defined within genesis
// state. The structure contains DenomAuthorityMetadata which defines the
// denom's admin.
type GenesisDenom struct {
	Denom               string                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
	AuthorityMetadata   DenomAuthorityMetadata `protobuf:"bytes,2,opt,name=authority_metadata,json=authorityMetadata,proto3" json:"authority_metadata" yaml:"authority_metadata"`
	HookContractAddress string                 `protobuf:"bytes,3,opt,name=hook_contract_address,json=hookContractAddress,proto3" json:"hook_contract_address,omitempty"`
}

func (m *GenesisDenom) Reset()         { *m = GenesisDenom{} }
func (m *GenesisDenom) String() string { return proto.CompactTextString(m) }
func (*GenesisDenom) ProtoMessage()    {}
func (*GenesisDenom) Descriptor() ([]byte, []int) {
	return fileDescriptor_5749c3f71850298b, []int{1}
}
func (m *GenesisDenom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisDenom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisDenom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisDenom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisDenom.Merge(m, src)
}
func (m *GenesisDenom) XXX_Size() int {
	return m.Size()
}
func (m *GenesisDenom) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisDenom.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisDenom proto.InternalMessageInfo

func (m *GenesisDenom) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *GenesisDenom) GetAuthorityMetadata() DenomAuthorityMetadata {
	if m != nil {
		return m.AuthorityMetadata
	}
	return DenomAuthorityMetadata{}
}

func (m *GenesisDenom) GetHookContractAddress() string {
	if m != nil {
		return m.HookContractAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "osmosis.tokenfactory.v1beta1.GenesisState")
	proto.RegisterType((*GenesisDenom)(nil), "osmosis.tokenfactory.v1beta1.GenesisDenom")
}

func init() {
	proto.RegisterFile("osmosis/tokenfactory/v1beta1/genesis.proto", fileDescriptor_5749c3f71850298b)
}

var fileDescriptor_5749c3f71850298b = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0xae, 0xd2, 0x40,
	0x14, 0x86, 0x3b, 0x80, 0x24, 0x16, 0x34, 0x5a, 0x25, 0xa9, 0x04, 0x5b, 0xe8, 0xc2, 0x10, 0x12,
	0xdb, 0x80, 0x2c, 0x08, 0xbb, 0x16, 0x12, 0x57, 0x18, 0x53, 0x5d, 0xb9, 0x69, 0x86, 0x76, 0x6c,
	0x1b, 0x6c, 0xa7, 0xe9, 0x0c, 0xc6, 0xbe, 0x80, 0x6b, 0x1f, 0xc1, 0x07, 0xf1, 0x01, 0x58, 0xb2,
	0x74, 0x45, 0x0c, 0x6c, 0x5c, 0xf3, 0x02, 0x9a, 0xce, 0x4c, 0xee, 0xbd, 0x5c, 0x08, 0xbb, 0xe9,
	0x99, 0xef, 0xfc, 0xe7, 0x3f, 0x7f, 0x47, 0x1e, 0x60, 0x92, 0x60, 0x12, 0x13, 0x8b, 0xe2, 0x15,
	0x4a, 0x3f, 0x43, 0x9f, 0xe2, 0xbc, 0xb0, 0xbe, 0x0e, 0x97, 0x88, 0xc2, 0xa1, 0x15, 0xa2, 0x14,
	0x91, 0x98, 0x98, 0x59, 0x8e, 0x29, 0x56, 0x3a, 0x82, 0x35, 0xef, 0xb2, 0xa6, 0x60, 0xdb, 0xcf,
	0x43, 0x1c, 0x62, 0x06, 0x5a, 0xe5, 0x89, 0xf7, 0xb4, 0x7b, 0x17, 0xf5, 0x33, 0x98, 0xc3, 0x44,
	0xc8, 0xb6, 0xc7, 0x57, 0x2d, 0xc0, 0x35, 0x8d, 0x70, 0x1e, 0xd3, 0x62, 0x81, 0x28, 0x0c, 0x20,
	0x85, 0xbc, 0xcb, 0xf8, 0x05, 0xe4, 0xe6, 0x5b, 0x6e, 0xef, 0x03, 0x85, 0x14, 0x29, 0x53, 0xb9,
	0xce, 0x65, 0x55, 0xd0, 0x05, 0xfd, 0xc6, 0xa8, 0x63, 0x5e, 0xb4, 0xfb, 0x9e, 0x31, 0x4e, 0x6d,
	0xb3, 0xd3, 0x25, 0x57, 0x74, 0x28, 0x99, 0xfc, 0x58, 0xdc, 0x7b, 0x01, 0x4a, 0x71, 0x42, 0xd4,
	0x4a, 0xb7, 0xda, 0x6f, 0x8c, 0x06, 0xe6, 0xb5, 0x95, 0x4d, 0x31, 0x7f, 0x5e, 0xb6, 0x38, 0x2f,
	0x4b, 0xc5, 0xe3, 0x4e, 0x6f, 0x15, 0x30, 0xf9, 0x32, 0x35, 0x4e, 0xf5, 0x0c, 0xf7, 0x91, 0x28,
	0xcc, 0xf9, 0xf7, 0xbf, 0x5b, 0xfb, 0xac, 0xa2, 0xbc, 0x92, 0x1f, 0x30, 0x94, 0xb9, 0x7f, 0xe8,
	0x3c, 0x39, 0xee, 0xf4, 0x26, 0x57, 0x62, 0x65, 0xc3, 0xe5, 0xd7, 0xca, 0x77, 0x20, 0x2b, 0x37,
	0x99, 0x78, 0x89, 0x08, 0x45, 0xad, 0xb0, 0x9d, 0xc7, 0xd7, 0xfd, 0xb2, 0x49, 0xf6, 0xfd, 0x40,
	0x9d, 0x9e, 0x70, 0xfe, 0x82, 0xcf, 0x3b, 0x57, 0x37, 0xdc, 0xa7, 0x67, 0xbf, 0x41, 0x99, 0xc8,
	0xad, 0x08, 0xe3, 0x95, 0xe7, 0xe3, 0x94, 0xe6, 0xd0, 0xa7, 0x1e, 0x0c, 0x82, 0x1c, 0x11, 0xa2,
	0x56, 0xd9, 0x02, 0x65, 0xc0, 0xc0, 0x7d, 0x56, 0x22, 0x33, 0x41, 0xd8, 0x1c, 0x98, 0xd6, 0xfe,
	0xfe, 0xd4, 0x81, 0xe3, 0x6e, 0xf6, 0x1a, 0xd8, 0xee, 0x35, 0xf0, 0x67, 0xaf, 0x81, 0x1f, 0x07,
	0x4d, 0xda, 0x1e, 0x34, 0xe9, 0xf7, 0x41, 0x93, 0x3e, 0x4d, 0xc2, 0x98, 0x46, 0xeb, 0xa5, 0xe9,
	0xe3, 0xc4, 0x5a, 0xd8, 0xef, 0x3e, 0xba, 0xf6, 0xeb, 0x59, 0x04, 0xe3, 0xd4, 0x4a, 0x20, 0x13,
	0x62, 0xe7, 0x6f, 0xa7, 0xcf, 0x85, 0x16, 0x19, 0x22, 0xcb, 0x3a, 0x7b, 0x1b, 0x6f, 0xfe, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x0f, 0xa5, 0xe3, 0xdb, 0xd6, 0x02, 0x00, 0x00,
}

func (this *GenesisDenom) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenesisDenom)
	if !ok {
		that2, ok := that.(GenesisDenom)
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
	if this.Denom != that1.Denom {
		return false
	}
	if !this.AuthorityMetadata.Equal(&that1.AuthorityMetadata) {
		return false
	}
	if this.HookContractAddress != that1.HookContractAddress {
		return false
	}
	return true
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
	if len(m.FactoryDenoms) > 0 {
		for iNdEx := len(m.FactoryDenoms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FactoryDenoms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *GenesisDenom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisDenom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisDenom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.HookContractAddress) > 0 {
		i -= len(m.HookContractAddress)
		copy(dAtA[i:], m.HookContractAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.HookContractAddress)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.AuthorityMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
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
	if len(m.FactoryDenoms) > 0 {
		for _, e := range m.FactoryDenoms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisDenom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.AuthorityMetadata.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.HookContractAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field FactoryDenoms", wireType)
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
			m.FactoryDenoms = append(m.FactoryDenoms, GenesisDenom{})
			if err := m.FactoryDenoms[len(m.FactoryDenoms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GenesisDenom) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisDenom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisDenom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorityMetadata", wireType)
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
			if err := m.AuthorityMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HookContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HookContractAddress = string(dAtA[iNdEx:postIndex])
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
