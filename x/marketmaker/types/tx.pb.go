// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/marketmaker/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type MsgApplyMarketMaker struct {
	Address string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PairIds []uint64 `protobuf:"varint,2,rep,packed,name=pair_ids,json=pairIds,proto3" json:"pair_ids,omitempty" yaml:"pair_ids"`
}

func (m *MsgApplyMarketMaker) Reset()         { *m = MsgApplyMarketMaker{} }
func (m *MsgApplyMarketMaker) String() string { return proto.CompactTextString(m) }
func (*MsgApplyMarketMaker) ProtoMessage()    {}
func (*MsgApplyMarketMaker) Descriptor() ([]byte, []int) {
	return fileDescriptor_690b1f467487b678, []int{0}
}
func (m *MsgApplyMarketMaker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgApplyMarketMaker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgApplyMarketMaker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgApplyMarketMaker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgApplyMarketMaker.Merge(m, src)
}
func (m *MsgApplyMarketMaker) XXX_Size() int {
	return m.Size()
}
func (m *MsgApplyMarketMaker) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgApplyMarketMaker.DiscardUnknown(m)
}

var xxx_messageInfo_MsgApplyMarketMaker proto.InternalMessageInfo

type MsgApplyMarketMakerResponse struct {
}

func (m *MsgApplyMarketMakerResponse) Reset()         { *m = MsgApplyMarketMakerResponse{} }
func (m *MsgApplyMarketMakerResponse) String() string { return proto.CompactTextString(m) }
func (*MsgApplyMarketMakerResponse) ProtoMessage()    {}
func (*MsgApplyMarketMakerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_690b1f467487b678, []int{1}
}
func (m *MsgApplyMarketMakerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgApplyMarketMakerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgApplyMarketMakerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgApplyMarketMakerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgApplyMarketMakerResponse.Merge(m, src)
}
func (m *MsgApplyMarketMakerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgApplyMarketMakerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgApplyMarketMakerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgApplyMarketMakerResponse proto.InternalMessageInfo

type MsgClaimIncentives struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgClaimIncentives) Reset()         { *m = MsgClaimIncentives{} }
func (m *MsgClaimIncentives) String() string { return proto.CompactTextString(m) }
func (*MsgClaimIncentives) ProtoMessage()    {}
func (*MsgClaimIncentives) Descriptor() ([]byte, []int) {
	return fileDescriptor_690b1f467487b678, []int{2}
}
func (m *MsgClaimIncentives) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimIncentives) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimIncentives.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimIncentives) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimIncentives.Merge(m, src)
}
func (m *MsgClaimIncentives) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimIncentives) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimIncentives.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimIncentives proto.InternalMessageInfo

type MsgClaimIncentivesResponse struct {
}

func (m *MsgClaimIncentivesResponse) Reset()         { *m = MsgClaimIncentivesResponse{} }
func (m *MsgClaimIncentivesResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClaimIncentivesResponse) ProtoMessage()    {}
func (*MsgClaimIncentivesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_690b1f467487b678, []int{3}
}
func (m *MsgClaimIncentivesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimIncentivesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimIncentivesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimIncentivesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimIncentivesResponse.Merge(m, src)
}
func (m *MsgClaimIncentivesResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimIncentivesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimIncentivesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimIncentivesResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgApplyMarketMaker)(nil), "mantrachain.marketmaker.v1beta1.MsgApplyMarketMaker")
	proto.RegisterType((*MsgApplyMarketMakerResponse)(nil), "mantrachain.marketmaker.v1beta1.MsgApplyMarketMakerResponse")
	proto.RegisterType((*MsgClaimIncentives)(nil), "mantrachain.marketmaker.v1beta1.MsgClaimIncentives")
	proto.RegisterType((*MsgClaimIncentivesResponse)(nil), "mantrachain.marketmaker.v1beta1.MsgClaimIncentivesResponse")
}

func init() {
	proto.RegisterFile("mantrachain/marketmaker/v1beta1/tx.proto", fileDescriptor_690b1f467487b678)
}

var fileDescriptor_690b1f467487b678 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xbd, 0x8e, 0xda, 0x40,
	0x18, 0xb4, 0x01, 0x85, 0x64, 0x1b, 0x22, 0x93, 0x82, 0x38, 0x89, 0x8d, 0x5c, 0xd1, 0xc4, 0x2b,
	0x02, 0x4d, 0x48, 0x1a, 0x88, 0x14, 0x89, 0xc2, 0x29, 0x9c, 0x54, 0x69, 0xd0, 0xda, 0xde, 0x2c,
	0x2b, 0xbc, 0xbb, 0x96, 0x77, 0x41, 0x50, 0x47, 0x8a, 0x52, 0xa6, 0xc8, 0x03, 0xe4, 0x71, 0x52,
	0x52, 0x5e, 0x75, 0x3a, 0xc1, 0x1b, 0xdc, 0x13, 0x9c, 0xf0, 0xcf, 0x89, 0x03, 0x4e, 0x77, 0xa2,
	0xdb, 0xd1, 0xcc, 0x7c, 0xdf, 0x7c, 0xda, 0x01, 0x1d, 0x86, 0xb8, 0x4a, 0x51, 0x38, 0x45, 0x94,
	0x43, 0x86, 0xd2, 0x19, 0x56, 0x0c, 0xcd, 0x70, 0x0a, 0x17, 0xdd, 0x00, 0x2b, 0xd4, 0x85, 0x6a,
	0xe9, 0x26, 0xa9, 0x50, 0xc2, 0xb0, 0xf7, 0x94, 0xee, 0x9e, 0xd2, 0x2d, 0x94, 0xe6, 0x0b, 0x22,
	0x88, 0xc8, 0xb4, 0x70, 0xf7, 0xca, 0x6d, 0xe6, 0xcb, 0x50, 0x48, 0x26, 0xe4, 0x24, 0x27, 0x72,
	0x50, 0x50, 0x56, 0x8e, 0x60, 0x80, 0x24, 0xbe, 0xdd, 0x17, 0x0a, 0xca, 0x0b, 0xde, 0x26, 0x42,
	0x90, 0x18, 0xc3, 0x0c, 0x05, 0xf3, 0x1f, 0x50, 0x51, 0x86, 0xa5, 0x42, 0x2c, 0xc9, 0x05, 0x0e,
	0x06, 0x4d, 0x4f, 0x92, 0x61, 0x92, 0xc4, 0x2b, 0x2f, 0x0b, 0xe4, 0xed, 0x02, 0x19, 0x2d, 0x50,
	0x47, 0x51, 0x94, 0x62, 0x29, 0x5b, 0x7a, 0x5b, 0xef, 0x3c, 0xf3, 0x4b, 0x68, 0xb8, 0xe0, 0x69,
	0x82, 0x68, 0x3a, 0xa1, 0x91, 0x6c, 0x55, 0xda, 0xd5, 0x4e, 0x6d, 0xd4, 0xbc, 0xbe, 0xb4, 0x1b,
	0x2b, 0xc4, 0xe2, 0x81, 0x53, 0x32, 0x8e, 0x5f, 0xdf, 0x3d, 0xc7, 0x91, 0x1c, 0xd4, 0x7e, 0xff,
	0xb3, 0x35, 0xe7, 0x0d, 0x78, 0x75, 0x62, 0x8d, 0x8f, 0x65, 0x22, 0xb8, 0xc4, 0x4e, 0x1f, 0x18,
	0x9e, 0x24, 0x9f, 0x62, 0x44, 0xd9, 0x98, 0x87, 0x98, 0x2b, 0xba, 0xc0, 0xf2, 0xfe, 0x10, 0xc5,
	0xd0, 0xd7, 0xc0, 0x3c, 0x76, 0x95, 0x33, 0xdf, 0xfd, 0xad, 0x80, 0xaa, 0x27, 0x89, 0xf1, 0x4b,
	0x07, 0xcf, 0x8f, 0xee, 0xeb, 0xbb, 0x0f, 0x7c, 0x85, 0x7b, 0x22, 0xae, 0xf9, 0xf1, 0x1c, 0x57,
	0x19, 0xc8, 0xf8, 0xa9, 0x83, 0xc6, 0xe1, 0x89, 0xbd, 0xc7, 0x4c, 0x3c, 0x30, 0x99, 0x1f, 0xce,
	0x30, 0x95, 0x29, 0x46, 0x5f, 0xff, 0x6f, 0x2c, 0x7d, 0xbd, 0xb1, 0xf4, 0xab, 0x8d, 0xa5, 0xff,
	0xd9, 0x5a, 0xda, 0x7a, 0x6b, 0x69, 0x17, 0x5b, 0x4b, 0xfb, 0xfe, 0x9e, 0x50, 0x35, 0x9d, 0x07,
	0x6e, 0x28, 0x18, 0xf4, 0x86, 0x5f, 0xbe, 0xf9, 0xc3, 0xb7, 0x9f, 0x29, 0x47, 0x3c, 0xc4, 0x70,
	0xbf, 0xe1, 0xcb, 0x3b, 0x1d, 0x57, 0xab, 0x04, 0xcb, 0xe0, 0x49, 0x56, 0xa6, 0xde, 0x4d, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd7, 0xf5, 0xd7, 0x40, 0x0b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	ApplyMarketMaker(ctx context.Context, in *MsgApplyMarketMaker, opts ...grpc.CallOption) (*MsgApplyMarketMakerResponse, error)
	ClaimIncentives(ctx context.Context, in *MsgClaimIncentives, opts ...grpc.CallOption) (*MsgClaimIncentivesResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ApplyMarketMaker(ctx context.Context, in *MsgApplyMarketMaker, opts ...grpc.CallOption) (*MsgApplyMarketMakerResponse, error) {
	out := new(MsgApplyMarketMakerResponse)
	err := c.cc.Invoke(ctx, "/mantrachain.marketmaker.v1beta1.Msg/ApplyMarketMaker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClaimIncentives(ctx context.Context, in *MsgClaimIncentives, opts ...grpc.CallOption) (*MsgClaimIncentivesResponse, error) {
	out := new(MsgClaimIncentivesResponse)
	err := c.cc.Invoke(ctx, "/mantrachain.marketmaker.v1beta1.Msg/ClaimIncentives", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	ApplyMarketMaker(context.Context, *MsgApplyMarketMaker) (*MsgApplyMarketMakerResponse, error)
	ClaimIncentives(context.Context, *MsgClaimIncentives) (*MsgClaimIncentivesResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ApplyMarketMaker(ctx context.Context, req *MsgApplyMarketMaker) (*MsgApplyMarketMakerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyMarketMaker not implemented")
}
func (*UnimplementedMsgServer) ClaimIncentives(ctx context.Context, req *MsgClaimIncentives) (*MsgClaimIncentivesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimIncentives not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ApplyMarketMaker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgApplyMarketMaker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ApplyMarketMaker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mantrachain.marketmaker.v1beta1.Msg/ApplyMarketMaker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ApplyMarketMaker(ctx, req.(*MsgApplyMarketMaker))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClaimIncentives_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimIncentives)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimIncentives(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mantrachain.marketmaker.v1beta1.Msg/ClaimIncentives",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimIncentives(ctx, req.(*MsgClaimIncentives))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mantrachain.marketmaker.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyMarketMaker",
			Handler:    _Msg_ApplyMarketMaker_Handler,
		},
		{
			MethodName: "ClaimIncentives",
			Handler:    _Msg_ClaimIncentives_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mantrachain/marketmaker/v1beta1/tx.proto",
}

func (m *MsgApplyMarketMaker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgApplyMarketMaker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgApplyMarketMaker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PairIds) > 0 {
		dAtA2 := make([]byte, len(m.PairIds)*10)
		var j1 int
		for _, num := range m.PairIds {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTx(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgApplyMarketMakerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgApplyMarketMakerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgApplyMarketMakerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgClaimIncentives) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimIncentives) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimIncentives) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimIncentivesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimIncentivesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimIncentivesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgApplyMarketMaker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.PairIds) > 0 {
		l = 0
		for _, e := range m.PairIds {
			l += sovTx(uint64(e))
		}
		n += 1 + sovTx(uint64(l)) + l
	}
	return n
}

func (m *MsgApplyMarketMakerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgClaimIncentives) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgClaimIncentivesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgApplyMarketMaker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgApplyMarketMaker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgApplyMarketMaker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.PairIds = append(m.PairIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTx
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTx
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.PairIds) == 0 {
					m.PairIds = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTx
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.PairIds = append(m.PairIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field PairIds", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgApplyMarketMakerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgApplyMarketMakerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgApplyMarketMakerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgClaimIncentives) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClaimIncentives: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimIncentives: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgClaimIncentivesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClaimIncentivesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimIncentivesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
