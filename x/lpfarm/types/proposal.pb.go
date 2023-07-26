// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/lpfarm/v1beta1/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type FarmingPlanProposal struct {
	Title                 string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description           string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CreatePlanRequests    []CreatePlanRequest    `protobuf:"bytes,3,rep,name=create_plan_requests,json=createPlanRequests,proto3" json:"create_plan_requests"`
	TerminatePlanRequests []TerminatePlanRequest `protobuf:"bytes,4,rep,name=terminate_plan_requests,json=terminatePlanRequests,proto3" json:"terminate_plan_requests"`
}

func (m *FarmingPlanProposal) Reset()      { *m = FarmingPlanProposal{} }
func (*FarmingPlanProposal) ProtoMessage() {}
func (*FarmingPlanProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_54159baf10c48731, []int{0}
}
func (m *FarmingPlanProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FarmingPlanProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FarmingPlanProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FarmingPlanProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FarmingPlanProposal.Merge(m, src)
}
func (m *FarmingPlanProposal) XXX_Size() int {
	return m.Size()
}
func (m *FarmingPlanProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_FarmingPlanProposal.DiscardUnknown(m)
}

var xxx_messageInfo_FarmingPlanProposal proto.InternalMessageInfo

type CreatePlanRequest struct {
	Description        string             `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	FarmingPoolAddress string             `protobuf:"bytes,2,opt,name=farming_pool_address,json=farmingPoolAddress,proto3" json:"farming_pool_address,omitempty"`
	RewardAllocations  []RewardAllocation `protobuf:"bytes,3,rep,name=reward_allocations,json=rewardAllocations,proto3" json:"reward_allocations"`
	StartTime          time.Time          `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	EndTime            time.Time          `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time"`
}

func (m *CreatePlanRequest) Reset()         { *m = CreatePlanRequest{} }
func (m *CreatePlanRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePlanRequest) ProtoMessage()    {}
func (*CreatePlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_54159baf10c48731, []int{1}
}
func (m *CreatePlanRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePlanRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreatePlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlanRequest.Merge(m, src)
}
func (m *CreatePlanRequest) XXX_Size() int {
	return m.Size()
}
func (m *CreatePlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlanRequest proto.InternalMessageInfo

type TerminatePlanRequest struct {
	PlanId uint64 `protobuf:"varint,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
}

func (m *TerminatePlanRequest) Reset()         { *m = TerminatePlanRequest{} }
func (m *TerminatePlanRequest) String() string { return proto.CompactTextString(m) }
func (*TerminatePlanRequest) ProtoMessage()    {}
func (*TerminatePlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_54159baf10c48731, []int{2}
}
func (m *TerminatePlanRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TerminatePlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TerminatePlanRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TerminatePlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminatePlanRequest.Merge(m, src)
}
func (m *TerminatePlanRequest) XXX_Size() int {
	return m.Size()
}
func (m *TerminatePlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminatePlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TerminatePlanRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FarmingPlanProposal)(nil), "mantrachain.lpfarm.v1beta1.FarmingPlanProposal")
	proto.RegisterType((*CreatePlanRequest)(nil), "mantrachain.lpfarm.v1beta1.CreatePlanRequest")
	proto.RegisterType((*TerminatePlanRequest)(nil), "mantrachain.lpfarm.v1beta1.TerminatePlanRequest")
}

func init() {
	proto.RegisterFile("mantrachain/lpfarm/v1beta1/proposal.proto", fileDescriptor_54159baf10c48731)
}

var fileDescriptor_54159baf10c48731 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xbb, 0x6e, 0xdb, 0x30,
	0x14, 0x95, 0x1c, 0xe7, 0x45, 0x4f, 0x61, 0x55, 0xc4, 0xf0, 0x20, 0x1b, 0x5e, 0xea, 0x02, 0x8d,
	0x94, 0xa4, 0x9d, 0xba, 0x14, 0x76, 0x80, 0x00, 0x1d, 0x5a, 0x18, 0x82, 0xa7, 0x2e, 0xc2, 0xb5,
	0x44, 0x2b, 0x02, 0x28, 0x92, 0x25, 0xe9, 0x3e, 0xfe, 0x22, 0x63, 0xc7, 0xce, 0xfd, 0x12, 0x8f,
	0x1e, 0x3b, 0xf5, 0x61, 0xff, 0x48, 0x41, 0x4a, 0x2a, 0x0c, 0x3b, 0x35, 0xd0, 0x4d, 0xe4, 0x39,
	0xe7, 0x9e, 0x7b, 0x8f, 0x2e, 0xd1, 0xd3, 0x02, 0x98, 0x96, 0x90, 0xdc, 0x41, 0xce, 0x42, 0x2a,
	0x66, 0x20, 0x8b, 0xf0, 0xc3, 0xd5, 0x94, 0x68, 0xb8, 0x0a, 0x85, 0xe4, 0x82, 0x2b, 0xa0, 0x81,
	0x90, 0x5c, 0x73, 0xdc, 0xd9, 0xa0, 0x06, 0x25, 0x35, 0xa8, 0xa8, 0x1d, 0x2f, 0xe3, 0x19, 0xb7,
	0xb4, 0xd0, 0x7c, 0x95, 0x8a, 0x4e, 0x37, 0xe3, 0x3c, 0xa3, 0x24, 0xb4, 0xa7, 0xe9, 0x7c, 0x16,
	0xea, 0xbc, 0x20, 0x4a, 0x43, 0x21, 0x2a, 0xc2, 0x93, 0x3d, 0xee, 0x95, 0x83, 0x25, 0xf6, 0xbf,
	0x35, 0xd0, 0xa3, 0x5b, 0x90, 0x45, 0xce, 0xb2, 0x31, 0x05, 0x36, 0xae, 0x3a, 0xc3, 0x1e, 0x3a,
	0xd4, 0xb9, 0xa6, 0xa4, 0xed, 0xf6, 0xdc, 0xc1, 0x69, 0x54, 0x1e, 0x70, 0x0f, 0xb5, 0x52, 0xa2,
	0x12, 0x99, 0x0b, 0x9d, 0x73, 0xd6, 0x6e, 0x58, 0x6c, 0xf3, 0x0a, 0x13, 0xe4, 0x25, 0x92, 0x80,
	0x26, 0xb1, 0xa0, 0xc0, 0x62, 0x49, 0xde, 0xcf, 0x89, 0xd2, 0xaa, 0x7d, 0xd0, 0x3b, 0x18, 0xb4,
	0xae, 0x2f, 0x82, 0x7f, 0x8f, 0x1a, 0xdc, 0x58, 0x9d, 0xe9, 0x22, 0x2a, 0x55, 0xa3, 0xe6, 0xe2,
	0x47, 0xd7, 0x89, 0x70, 0xb2, 0x0d, 0x28, 0xcc, 0xd0, 0xb9, 0x26, 0xa6, 0xeb, 0x5d, 0xa7, 0xa6,
	0x75, 0xba, 0xdc, 0xe7, 0x34, 0xa9, 0xa5, 0xbb, 0x66, 0x8f, 0xf5, 0x03, 0x98, 0x7a, 0xd9, 0xfc,
	0xf2, 0xb5, 0xeb, 0xf4, 0x97, 0x0d, 0x74, 0xb6, 0xd3, 0xe5, 0x76, 0x28, 0xee, 0x6e, 0x28, 0x97,
	0xc8, 0x9b, 0x95, 0x19, 0xc7, 0x82, 0x73, 0x1a, 0x43, 0x9a, 0x4a, 0xa2, 0x54, 0x95, 0x1f, 0xae,
	0xb0, 0x31, 0xe7, 0x74, 0x58, 0x22, 0x18, 0x10, 0x96, 0xe4, 0x23, 0xc8, 0x34, 0x06, 0x4a, 0x79,
	0x02, 0xa6, 0x4c, 0x1d, 0xe2, 0xb3, 0x7d, 0xa3, 0x45, 0x56, 0x35, 0xfc, 0x2b, 0xaa, 0xc6, 0x3a,
	0x93, 0x5b, 0xf7, 0x0a, 0xdf, 0x20, 0xa4, 0x34, 0x48, 0x1d, 0x9b, 0xdd, 0x69, 0x37, 0x7b, 0xee,
	0xa0, 0x75, 0xdd, 0x09, 0xca, 0xc5, 0x0a, 0xea, 0xc5, 0x0a, 0x26, 0xf5, 0x62, 0x8d, 0x4e, 0x4c,
	0xa1, 0xfb, 0x9f, 0x5d, 0x37, 0x3a, 0xb5, 0x3a, 0x83, 0xe0, 0x57, 0xe8, 0x84, 0xb0, 0xb4, 0x2c,
	0x71, 0xf8, 0x1f, 0x25, 0x8e, 0x09, 0x4b, 0xcd, 0x7d, 0x3f, 0x44, 0xde, 0x43, 0x7f, 0x03, 0x9f,
	0xa3, 0x63, 0xfb, 0x5b, 0xf3, 0xd4, 0x06, 0xda, 0x8c, 0x8e, 0xcc, 0xf1, 0x75, 0x3a, 0x8a, 0x16,
	0xbf, 0x7d, 0x67, 0xb1, 0xf2, 0xdd, 0xe5, 0xca, 0x77, 0x7f, 0xad, 0x7c, 0xf7, 0x7e, 0xed, 0x3b,
	0xcb, 0xb5, 0xef, 0x7c, 0x5f, 0xfb, 0xce, 0xbb, 0x17, 0x59, 0xae, 0xef, 0xe6, 0xd3, 0x20, 0xe1,
	0x45, 0xf8, 0x66, 0xf8, 0x76, 0x12, 0x0d, 0x2f, 0x6e, 0x73, 0x06, 0x2c, 0x21, 0xe1, 0xe6, 0x8b,
	0xf8, 0x54, 0xbf, 0x09, 0xfd, 0x59, 0x10, 0x35, 0x3d, 0xb2, 0xbd, 0x3e, 0xff, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x3c, 0xce, 0x97, 0x2e, 0xb4, 0x03, 0x00, 0x00,
}

func (m *FarmingPlanProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FarmingPlanProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FarmingPlanProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TerminatePlanRequests) > 0 {
		for iNdEx := len(m.TerminatePlanRequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TerminatePlanRequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.CreatePlanRequests) > 0 {
		for iNdEx := len(m.CreatePlanRequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CreatePlanRequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CreatePlanRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePlanRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreatePlanRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintProposal(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintProposal(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if len(m.RewardAllocations) > 0 {
		for iNdEx := len(m.RewardAllocations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RewardAllocations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.FarmingPoolAddress) > 0 {
		i -= len(m.FarmingPoolAddress)
		copy(dAtA[i:], m.FarmingPoolAddress)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.FarmingPoolAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TerminatePlanRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TerminatePlanRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TerminatePlanRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PlanId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.PlanId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FarmingPlanProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.CreatePlanRequests) > 0 {
		for _, e := range m.CreatePlanRequests {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if len(m.TerminatePlanRequests) > 0 {
		for _, e := range m.TerminatePlanRequests {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func (m *CreatePlanRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.FarmingPoolAddress)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.RewardAllocations) > 0 {
		for _, e := range m.RewardAllocations {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovProposal(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovProposal(uint64(l))
	return n
}

func (m *TerminatePlanRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PlanId != 0 {
		n += 1 + sovProposal(uint64(m.PlanId))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FarmingPlanProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: FarmingPlanProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FarmingPlanProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatePlanRequests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatePlanRequests = append(m.CreatePlanRequests, CreatePlanRequest{})
			if err := m.CreatePlanRequests[len(m.CreatePlanRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TerminatePlanRequests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TerminatePlanRequests = append(m.TerminatePlanRequests, TerminatePlanRequest{})
			if err := m.TerminatePlanRequests[len(m.TerminatePlanRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *CreatePlanRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: CreatePlanRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePlanRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FarmingPoolAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FarmingPoolAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAllocations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardAllocations = append(m.RewardAllocations, RewardAllocation{})
			if err := m.RewardAllocations[len(m.RewardAllocations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *TerminatePlanRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: TerminatePlanRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TerminatePlanRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanId", wireType)
			}
			m.PlanId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
