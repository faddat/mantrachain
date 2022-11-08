// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vault/v1/chain_validator_bridge.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type ChainValidatorBridge struct {
	BridgeId string                                 `protobuf:"bytes,1,opt,name=bridge_id,json=bridgeId,proto3" json:"bridge_id,omitempty" yaml:"bridge_id"`
	Creator  string                                 `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
	Staked   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=staked,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"staked" yaml:"staked"`
}

func (m *ChainValidatorBridge) Reset()         { *m = ChainValidatorBridge{} }
func (m *ChainValidatorBridge) String() string { return proto.CompactTextString(m) }
func (*ChainValidatorBridge) ProtoMessage()    {}
func (*ChainValidatorBridge) Descriptor() ([]byte, []int) {
	return fileDescriptor_911e9b9d4d33c790, []int{0}
}
func (m *ChainValidatorBridge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainValidatorBridge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainValidatorBridge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainValidatorBridge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainValidatorBridge.Merge(m, src)
}
func (m *ChainValidatorBridge) XXX_Size() int {
	return m.Size()
}
func (m *ChainValidatorBridge) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainValidatorBridge.DiscardUnknown(m)
}

var xxx_messageInfo_ChainValidatorBridge proto.InternalMessageInfo

func (m *ChainValidatorBridge) GetBridgeId() string {
	if m != nil {
		return m.BridgeId
	}
	return ""
}

func (m *ChainValidatorBridge) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*ChainValidatorBridge)(nil), "LimeChain.mantrachain.vault.v1.ChainValidatorBridge")
}

func init() {
	proto.RegisterFile("vault/v1/chain_validator_bridge.proto", fileDescriptor_911e9b9d4d33c790)
}

var fileDescriptor_911e9b9d4d33c790 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4a, 0xc4, 0x30,
	0x1c, 0xc6, 0x1b, 0x85, 0xd3, 0x2b, 0x28, 0x52, 0x3a, 0x14, 0x87, 0x54, 0x0a, 0x8a, 0x83, 0x26,
	0x14, 0x37, 0x17, 0xa1, 0x3a, 0x28, 0x38, 0x75, 0x50, 0x70, 0x29, 0x69, 0x13, 0x7a, 0xe1, 0xda,
	0xcb, 0xd1, 0xe6, 0x8a, 0xf7, 0x16, 0x3e, 0xd6, 0x8d, 0x1d, 0xc5, 0xa1, 0x48, 0xfb, 0x06, 0x7d,
	0x82, 0xa3, 0x49, 0xef, 0xb8, 0x29, 0x7f, 0xf8, 0xbe, 0xdf, 0x2f, 0xf0, 0x99, 0xd7, 0x15, 0x59,
	0x65, 0x12, 0x57, 0x3e, 0x4e, 0x66, 0x84, 0x2f, 0xa2, 0x8a, 0x64, 0x9c, 0x12, 0x29, 0x8a, 0x28,
	0x2e, 0x38, 0x4d, 0x19, 0x5a, 0x16, 0x42, 0x0a, 0x0b, 0xbe, 0xf3, 0x9c, 0x3d, 0x0f, 0x0d, 0x94,
	0x93, 0x85, 0x2c, 0x88, 0x6a, 0x23, 0x05, 0xa3, 0xca, 0xbf, 0xb4, 0x53, 0x91, 0x0a, 0x55, 0xc5,
	0xc3, 0xa5, 0x29, 0xaf, 0x06, 0xa6, 0xad, 0xa0, 0x8f, 0x9d, 0x35, 0x50, 0x52, 0xcb, 0x37, 0xa7,
	0x5a, 0x1f, 0x71, 0xea, 0x80, 0x2b, 0x70, 0x3b, 0x0d, 0xec, 0xbe, 0x71, 0x2f, 0xd6, 0x24, 0xcf,
	0x1e, 0xbd, 0x7d, 0xe4, 0x85, 0xa7, 0xfa, 0x7e, 0xa3, 0xd6, 0x9d, 0x79, 0x92, 0x14, 0x6c, 0x70,
	0x38, 0x47, 0x0a, 0xb0, 0xfa, 0xc6, 0x3d, 0xd7, 0xc0, 0x18, 0x78, 0xe1, 0xae, 0x62, 0x7d, 0x9a,
	0x93, 0x52, 0x92, 0x39, 0xa3, 0xce, 0xb1, 0x2a, 0x3f, 0x6d, 0x1a, 0xd7, 0xf8, 0x6b, 0xdc, 0x9b,
	0x94, 0xcb, 0xd9, 0x2a, 0x46, 0x89, 0xc8, 0x71, 0x22, 0xca, 0x5c, 0x94, 0xe3, 0x73, 0x5f, 0xd2,
	0x39, 0x96, 0xeb, 0x25, 0x2b, 0xd1, 0x0b, 0x4b, 0xfa, 0xc6, 0x3d, 0xd3, 0x6a, 0x6d, 0xf1, 0xc2,
	0x51, 0x17, 0xbc, 0x6e, 0x5a, 0x08, 0xea, 0x16, 0x82, 0xff, 0x16, 0x82, 0x9f, 0x0e, 0x1a, 0x75,
	0x07, 0x8d, 0xdf, 0x0e, 0x1a, 0x5f, 0xe8, 0x40, 0xbd, 0x5f, 0x0b, 0x1f, 0xac, 0x85, 0xbf, 0xb1,
	0x1e, 0x5b, 0x7d, 0x13, 0x4f, 0xd4, 0x46, 0x0f, 0xdb, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x99,
	0x41, 0x4f, 0x82, 0x01, 0x00, 0x00,
}

func (m *ChainValidatorBridge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainValidatorBridge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainValidatorBridge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Staked.Size()
		i -= size
		if _, err := m.Staked.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintChainValidatorBridge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintChainValidatorBridge(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BridgeId) > 0 {
		i -= len(m.BridgeId)
		copy(dAtA[i:], m.BridgeId)
		i = encodeVarintChainValidatorBridge(dAtA, i, uint64(len(m.BridgeId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintChainValidatorBridge(dAtA []byte, offset int, v uint64) int {
	offset -= sovChainValidatorBridge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainValidatorBridge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BridgeId)
	if l > 0 {
		n += 1 + l + sovChainValidatorBridge(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovChainValidatorBridge(uint64(l))
	}
	l = m.Staked.Size()
	n += 1 + l + sovChainValidatorBridge(uint64(l))
	return n
}

func sovChainValidatorBridge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChainValidatorBridge(x uint64) (n int) {
	return sovChainValidatorBridge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainValidatorBridge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChainValidatorBridge
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
			return fmt.Errorf("proto: ChainValidatorBridge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainValidatorBridge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainValidatorBridge
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
				return ErrInvalidLengthChainValidatorBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainValidatorBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BridgeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainValidatorBridge
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
				return ErrInvalidLengthChainValidatorBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainValidatorBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staked", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainValidatorBridge
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
				return ErrInvalidLengthChainValidatorBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainValidatorBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Staked.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChainValidatorBridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChainValidatorBridge
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
func skipChainValidatorBridge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChainValidatorBridge
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
					return 0, ErrIntOverflowChainValidatorBridge
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
					return 0, ErrIntOverflowChainValidatorBridge
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
				return 0, ErrInvalidLengthChainValidatorBridge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChainValidatorBridge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChainValidatorBridge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChainValidatorBridge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChainValidatorBridge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChainValidatorBridge = fmt.Errorf("proto: unexpected end of group")
)
