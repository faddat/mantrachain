// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/bridge/v1/params.proto

package types

import (
	fmt "fmt"
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

// Params defines the parameters for the module.
type Params struct {
	AdminAccount                           string `protobuf:"bytes,1,opt,name=admin_account,json=adminAccount,proto3" json:"admin_account,omitempty"`
	ValidBridgeId                          string `protobuf:"bytes,2,opt,name=valid_bridge_id,json=validBridgeId,proto3" json:"valid_bridge_id,omitempty"`
	ValidBridgeCw20ContractNameMinLength   int32  `protobuf:"varint,3,opt,name=valid_bridge_cw20_contract_name_min_length,json=validBridgeCw20ContractNameMinLength,proto3" json:"valid_bridge_cw20_contract_name_min_length,omitempty"`
	ValidBridgeCw20ContractNameMaxLength   int32  `protobuf:"varint,4,opt,name=valid_bridge_cw20_contract_name_max_length,json=validBridgeCw20ContractNameMaxLength,proto3" json:"valid_bridge_cw20_contract_name_max_length,omitempty"`
	ValidBridgeCw20ContractSymbolMinLength int32  `protobuf:"varint,5,opt,name=valid_bridge_cw20_contract_symbol_min_length,json=validBridgeCw20ContractSymbolMinLength,proto3" json:"valid_bridge_cw20_contract_symbol_min_length,omitempty"`
	ValidBridgeCw20ContractSymbolMaxLength int32  `protobuf:"varint,6,opt,name=valid_bridge_cw20_contract_symbol_max_length,json=validBridgeCw20ContractSymbolMaxLength,proto3" json:"valid_bridge_cw20_contract_symbol_max_length,omitempty"`
	ValidMintListMetadataMaxCount          int32  `protobuf:"varint,7,opt,name=valid_mint_list_metadata_max_count,json=validMintListMetadataMaxCount,proto3" json:"valid_mint_list_metadata_max_count,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7ca7e14e86ab95c, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAdminAccount() string {
	if m != nil {
		return m.AdminAccount
	}
	return ""
}

func (m *Params) GetValidBridgeId() string {
	if m != nil {
		return m.ValidBridgeId
	}
	return ""
}

func (m *Params) GetValidBridgeCw20ContractNameMinLength() int32 {
	if m != nil {
		return m.ValidBridgeCw20ContractNameMinLength
	}
	return 0
}

func (m *Params) GetValidBridgeCw20ContractNameMaxLength() int32 {
	if m != nil {
		return m.ValidBridgeCw20ContractNameMaxLength
	}
	return 0
}

func (m *Params) GetValidBridgeCw20ContractSymbolMinLength() int32 {
	if m != nil {
		return m.ValidBridgeCw20ContractSymbolMinLength
	}
	return 0
}

func (m *Params) GetValidBridgeCw20ContractSymbolMaxLength() int32 {
	if m != nil {
		return m.ValidBridgeCw20ContractSymbolMaxLength
	}
	return 0
}

func (m *Params) GetValidMintListMetadataMaxCount() int32 {
	if m != nil {
		return m.ValidMintListMetadataMaxCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "mantrachain.bridge.v1.Params")
}

func init() {
	proto.RegisterFile("mantrachain/bridge/v1/params.proto", fileDescriptor_e7ca7e14e86ab95c)
}

var fileDescriptor_e7ca7e14e86ab95c = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6a, 0xe2, 0x50,
	0x14, 0x86, 0xcd, 0x8c, 0x3a, 0xcc, 0x65, 0x64, 0x20, 0x50, 0x70, 0xd3, 0x20, 0xb6, 0x88, 0x94,
	0x92, 0xa8, 0x7d, 0x82, 0x9a, 0x95, 0x60, 0x4a, 0xb1, 0x9b, 0x52, 0x0a, 0x97, 0x93, 0xe4, 0xa2,
	0x17, 0x72, 0x6f, 0x24, 0x39, 0x6a, 0x7c, 0x0b, 0x1f, 0xab, 0x4b, 0x97, 0x5d, 0x16, 0x7d, 0x91,
	0xe2, 0x89, 0xd6, 0x74, 0xd1, 0xd6, 0xed, 0xcf, 0xc7, 0x77, 0x7e, 0x38, 0x3f, 0x6b, 0x2a, 0xd0,
	0x98, 0x40, 0x30, 0x01, 0xa9, 0x1d, 0x3f, 0x91, 0xe1, 0x58, 0x38, 0xf3, 0xae, 0x33, 0x85, 0x04,
	0x54, 0x6a, 0x4f, 0x93, 0x18, 0x63, 0xf3, 0xac, 0xc0, 0xd8, 0x39, 0x63, 0xcf, 0xbb, 0xcd, 0x55,
	0x99, 0x55, 0xef, 0x89, 0x33, 0x2f, 0x58, 0x0d, 0x42, 0x25, 0x35, 0x87, 0x20, 0x88, 0x67, 0x1a,
	0xeb, 0x46, 0xc3, 0x68, 0xff, 0x1d, 0xfd, 0xa3, 0xf0, 0x36, 0xcf, 0xcc, 0x16, 0xfb, 0x3f, 0x87,
	0x48, 0x86, 0x3c, 0x57, 0x70, 0x19, 0xd6, 0x7f, 0x11, 0x56, 0xa3, 0xb8, 0x4f, 0xe9, 0x20, 0x34,
	0x1f, 0xd9, 0xd5, 0x27, 0x2e, 0x58, 0xf4, 0x3a, 0x3c, 0x88, 0xa9, 0x02, 0x72, 0x0d, 0x4a, 0xf0,
	0xdd, 0xa9, 0x48, 0xe8, 0x31, 0x4e, 0xea, 0xbf, 0x1b, 0x46, 0xbb, 0x32, 0xba, 0x2c, 0x28, 0xdc,
	0x45, 0xaf, 0xe3, 0xee, 0xf1, 0x3b, 0x50, 0xc2, 0x93, 0x7a, 0x48, 0xec, 0x49, 0x66, 0xc8, 0x0e,
	0xe6, 0xf2, 0xcf, 0x66, 0xc8, 0xf6, 0xe6, 0x67, 0x76, 0xfd, 0x8d, 0x39, 0x5d, 0x2a, 0x3f, 0x8e,
	0x8a, 0xad, 0x2b, 0xe4, 0x6e, 0x7d, 0xe1, 0x7e, 0x20, 0xfe, 0xd8, 0xfb, 0x34, 0xfb, 0xb1, 0x79,
	0xf5, 0x14, 0xfb, 0x47, 0xf7, 0x01, 0x6b, 0xe6, 0x76, 0x25, 0x35, 0xf2, 0x48, 0xa6, 0xc8, 0x95,
	0x40, 0x08, 0x01, 0x81, 0xa4, 0xf9, 0x47, 0xff, 0x90, 0xf3, 0x9c, 0x48, 0x4f, 0x6a, 0x1c, 0xca,
	0x14, 0xbd, 0x3d, 0xe6, 0x41, 0xe6, 0xee, 0xa0, 0xfe, 0xe0, 0x65, 0x63, 0x19, 0xeb, 0x8d, 0x65,
	0xbc, 0x6d, 0x2c, 0x63, 0xb5, 0xb5, 0x4a, 0xeb, 0xad, 0x55, 0x7a, 0xdd, 0x5a, 0xa5, 0x27, 0x67,
	0x2c, 0x71, 0x32, 0xf3, 0xed, 0x20, 0x56, 0xce, 0x50, 0x2a, 0xe1, 0xd2, 0xe0, 0x8a, 0xe3, 0xcb,
	0x0e, 0xf3, 0xc3, 0xe5, 0x54, 0xa4, 0x7e, 0x95, 0xb6, 0x77, 0xf3, 0x1e, 0x00, 0x00, 0xff, 0xff,
	0x85, 0x08, 0xe2, 0xde, 0xa1, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ValidMintListMetadataMaxCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ValidMintListMetadataMaxCount))
		i--
		dAtA[i] = 0x38
	}
	if m.ValidBridgeCw20ContractSymbolMaxLength != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ValidBridgeCw20ContractSymbolMaxLength))
		i--
		dAtA[i] = 0x30
	}
	if m.ValidBridgeCw20ContractSymbolMinLength != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ValidBridgeCw20ContractSymbolMinLength))
		i--
		dAtA[i] = 0x28
	}
	if m.ValidBridgeCw20ContractNameMaxLength != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ValidBridgeCw20ContractNameMaxLength))
		i--
		dAtA[i] = 0x20
	}
	if m.ValidBridgeCw20ContractNameMinLength != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ValidBridgeCw20ContractNameMinLength))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ValidBridgeId) > 0 {
		i -= len(m.ValidBridgeId)
		copy(dAtA[i:], m.ValidBridgeId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ValidBridgeId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AdminAccount) > 0 {
		i -= len(m.AdminAccount)
		copy(dAtA[i:], m.AdminAccount)
		i = encodeVarintParams(dAtA, i, uint64(len(m.AdminAccount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AdminAccount)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ValidBridgeId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.ValidBridgeCw20ContractNameMinLength != 0 {
		n += 1 + sovParams(uint64(m.ValidBridgeCw20ContractNameMinLength))
	}
	if m.ValidBridgeCw20ContractNameMaxLength != 0 {
		n += 1 + sovParams(uint64(m.ValidBridgeCw20ContractNameMaxLength))
	}
	if m.ValidBridgeCw20ContractSymbolMinLength != 0 {
		n += 1 + sovParams(uint64(m.ValidBridgeCw20ContractSymbolMinLength))
	}
	if m.ValidBridgeCw20ContractSymbolMaxLength != 0 {
		n += 1 + sovParams(uint64(m.ValidBridgeCw20ContractSymbolMaxLength))
	}
	if m.ValidMintListMetadataMaxCount != 0 {
		n += 1 + sovParams(uint64(m.ValidMintListMetadataMaxCount))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdminAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidBridgeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidBridgeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidBridgeCw20ContractNameMinLength", wireType)
			}
			m.ValidBridgeCw20ContractNameMinLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidBridgeCw20ContractNameMinLength |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidBridgeCw20ContractNameMaxLength", wireType)
			}
			m.ValidBridgeCw20ContractNameMaxLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidBridgeCw20ContractNameMaxLength |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidBridgeCw20ContractSymbolMinLength", wireType)
			}
			m.ValidBridgeCw20ContractSymbolMinLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidBridgeCw20ContractSymbolMinLength |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidBridgeCw20ContractSymbolMaxLength", wireType)
			}
			m.ValidBridgeCw20ContractSymbolMaxLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidBridgeCw20ContractSymbolMaxLength |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidMintListMetadataMaxCount", wireType)
			}
			m.ValidMintListMetadataMaxCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidMintListMetadataMaxCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
