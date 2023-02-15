// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/bridge/v1/cw_20_contract.proto

package types

import (
	fmt "fmt"
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

type Cw20Contract struct {
	CodeId  uint64 `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	Ver     string `protobuf:"bytes,2,opt,name=ver,proto3" json:"ver,omitempty"`
	Path    string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Cw20Contract) Reset()         { *m = Cw20Contract{} }
func (m *Cw20Contract) String() string { return proto.CompactTextString(m) }
func (*Cw20Contract) ProtoMessage()    {}
func (*Cw20Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_02ba575d1b8b8424, []int{0}
}
func (m *Cw20Contract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cw20Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Cw20Contract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Cw20Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cw20Contract.Merge(m, src)
}
func (m *Cw20Contract) XXX_Size() int {
	return m.Size()
}
func (m *Cw20Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Cw20Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Cw20Contract proto.InternalMessageInfo

func (m *Cw20Contract) GetCodeId() uint64 {
	if m != nil {
		return m.CodeId
	}
	return 0
}

func (m *Cw20Contract) GetVer() string {
	if m != nil {
		return m.Ver
	}
	return ""
}

func (m *Cw20Contract) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Cw20Contract) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Cw20Contract)(nil), "mantrachain.bridge.v1.Cw20Contract")
}

func init() {
	proto.RegisterFile("mantrachain/bridge/v1/cw_20_contract.proto", fileDescriptor_02ba575d1b8b8424)
}

var fileDescriptor_02ba575d1b8b8424 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xca, 0x4d, 0xcc, 0x2b,
	0x29, 0x4a, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x2a, 0xca, 0x4c, 0x49, 0x4f, 0xd5, 0x2f,
	0x33, 0xd4, 0x4f, 0x2e, 0x8f, 0x37, 0x32, 0x88, 0x4f, 0xce, 0x07, 0xcb, 0x95, 0xe8, 0x15, 0x14,
	0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x22, 0xa9, 0xd5, 0x83, 0xa8, 0xd5, 0x2b, 0x33, 0x94, 0x12, 0x49,
	0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd0, 0x07, 0xb1, 0x20, 0x8a, 0x95, 0xd2, 0xb9, 0x78, 0x9c, 0xcb,
	0x8d, 0x0c, 0x9c, 0xa1, 0x46, 0x08, 0x89, 0x73, 0xb1, 0x27, 0xe7, 0xa7, 0xa4, 0xc6, 0x67, 0xa6,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0xb1, 0x81, 0xb8, 0x9e, 0x29, 0x42, 0x02, 0x5c, 0xcc,
	0x65, 0xa9, 0x45, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x90, 0x10, 0x17, 0x4b,
	0x41, 0x62, 0x49, 0x86, 0x04, 0x33, 0x58, 0x08, 0xcc, 0x16, 0x92, 0xe0, 0x62, 0x4f, 0x2e, 0x4a,
	0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x01, 0x0b, 0xc3, 0xb8, 0x4e, 0xbe, 0x2b, 0x1e, 0xc9, 0x31,
	0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb,
	0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x7e, 0x7a, 0x66, 0x49, 0x46,
	0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x4f, 0x66, 0x6e, 0xaa, 0x33, 0xd8, 0xa3, 0xc8, 0x9e,
	0xae, 0x80, 0x79, 0xbb, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x7c, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbf, 0xea, 0x9f, 0x20, 0x19, 0x01, 0x00, 0x00,
}

func (this *Cw20Contract) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Cw20Contract)
	if !ok {
		that2, ok := that.(Cw20Contract)
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
	if this.CodeId != that1.CodeId {
		return false
	}
	if this.Ver != that1.Ver {
		return false
	}
	if this.Path != that1.Path {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	return true
}
func (m *Cw20Contract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cw20Contract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Cw20Contract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintCw_20Contract(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintCw_20Contract(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Ver) > 0 {
		i -= len(m.Ver)
		copy(dAtA[i:], m.Ver)
		i = encodeVarintCw_20Contract(dAtA, i, uint64(len(m.Ver)))
		i--
		dAtA[i] = 0x12
	}
	if m.CodeId != 0 {
		i = encodeVarintCw_20Contract(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCw_20Contract(dAtA []byte, offset int, v uint64) int {
	offset -= sovCw_20Contract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Cw20Contract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeId != 0 {
		n += 1 + sovCw_20Contract(uint64(m.CodeId))
	}
	l = len(m.Ver)
	if l > 0 {
		n += 1 + l + sovCw_20Contract(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovCw_20Contract(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovCw_20Contract(uint64(l))
	}
	return n
}

func sovCw_20Contract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCw_20Contract(x uint64) (n int) {
	return sovCw_20Contract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Cw20Contract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCw_20Contract
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
			return fmt.Errorf("proto: Cw20Contract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cw20Contract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCw_20Contract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCw_20Contract
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
				return ErrInvalidLengthCw_20Contract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCw_20Contract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCw_20Contract
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
				return ErrInvalidLengthCw_20Contract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCw_20Contract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCw_20Contract
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
				return ErrInvalidLengthCw_20Contract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCw_20Contract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCw_20Contract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCw_20Contract
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
func skipCw_20Contract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCw_20Contract
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
					return 0, ErrIntOverflowCw_20Contract
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
					return 0, ErrIntOverflowCw_20Contract
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
				return 0, ErrInvalidLengthCw_20Contract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCw_20Contract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCw_20Contract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCw_20Contract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCw_20Contract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCw_20Contract = fmt.Errorf("proto: unexpected end of group")
)
