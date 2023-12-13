// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aumega/guard/v1/common.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type MsgAccounts struct {
	Accounts []string `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (m *MsgAccounts) Reset()         { *m = MsgAccounts{} }
func (m *MsgAccounts) String() string { return proto.CompactTextString(m) }
func (*MsgAccounts) ProtoMessage()    {}
func (*MsgAccounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1fabd206b7c56d1, []int{0}
}
func (m *MsgAccounts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAccounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAccounts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAccounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAccounts.Merge(m, src)
}
func (m *MsgAccounts) XXX_Size() int {
	return m.Size()
}
func (m *MsgAccounts) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAccounts.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAccounts proto.InternalMessageInfo

func (m *MsgAccounts) GetAccounts() []string {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type MsgIndexes struct {
	Indexes [][]byte `protobuf:"bytes,1,rep,name=indexes,proto3" json:"indexes,omitempty"`
}

func (m *MsgIndexes) Reset()         { *m = MsgIndexes{} }
func (m *MsgIndexes) String() string { return proto.CompactTextString(m) }
func (*MsgIndexes) ProtoMessage()    {}
func (*MsgIndexes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1fabd206b7c56d1, []int{1}
}
func (m *MsgIndexes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIndexes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIndexes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIndexes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIndexes.Merge(m, src)
}
func (m *MsgIndexes) XXX_Size() int {
	return m.Size()
}
func (m *MsgIndexes) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIndexes.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIndexes proto.InternalMessageInfo

func (m *MsgIndexes) GetIndexes() [][]byte {
	if m != nil {
		return m.Indexes
	}
	return nil
}

type AccountPrivileges struct {
	Account    []byte `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Privileges []byte `protobuf:"bytes,2,opt,name=privileges,proto3" json:"privileges,omitempty"`
}

func (m *AccountPrivileges) Reset()         { *m = AccountPrivileges{} }
func (m *AccountPrivileges) String() string { return proto.CompactTextString(m) }
func (*AccountPrivileges) ProtoMessage()    {}
func (*AccountPrivileges) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1fabd206b7c56d1, []int{2}
}
func (m *AccountPrivileges) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountPrivileges) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountPrivileges.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountPrivileges) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountPrivileges.Merge(m, src)
}
func (m *AccountPrivileges) XXX_Size() int {
	return m.Size()
}
func (m *AccountPrivileges) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountPrivileges.DiscardUnknown(m)
}

var xxx_messageInfo_AccountPrivileges proto.InternalMessageInfo

func (m *AccountPrivileges) GetAccount() []byte {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *AccountPrivileges) GetPrivileges() []byte {
	if m != nil {
		return m.Privileges
	}
	return nil
}

type RequiredPrivileges struct {
	Index      []byte `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Privileges []byte `protobuf:"bytes,2,opt,name=privileges,proto3" json:"privileges,omitempty"`
	Kind       string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (m *RequiredPrivileges) Reset()         { *m = RequiredPrivileges{} }
func (m *RequiredPrivileges) String() string { return proto.CompactTextString(m) }
func (*RequiredPrivileges) ProtoMessage()    {}
func (*RequiredPrivileges) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1fabd206b7c56d1, []int{3}
}
func (m *RequiredPrivileges) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequiredPrivileges) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequiredPrivileges.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequiredPrivileges) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequiredPrivileges.Merge(m, src)
}
func (m *RequiredPrivileges) XXX_Size() int {
	return m.Size()
}
func (m *RequiredPrivileges) XXX_DiscardUnknown() {
	xxx_messageInfo_RequiredPrivileges.DiscardUnknown(m)
}

var xxx_messageInfo_RequiredPrivileges proto.InternalMessageInfo

func (m *RequiredPrivileges) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *RequiredPrivileges) GetPrivileges() []byte {
	if m != nil {
		return m.Privileges
	}
	return nil
}

func (m *RequiredPrivileges) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

type AuthzGrantRevokeMsgType struct {
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	Grant   bool   `protobuf:"varint,2,opt,name=grant,proto3" json:"grant,omitempty"`
}

func (m *AuthzGrantRevokeMsgType) Reset()         { *m = AuthzGrantRevokeMsgType{} }
func (m *AuthzGrantRevokeMsgType) String() string { return proto.CompactTextString(m) }
func (*AuthzGrantRevokeMsgType) ProtoMessage()    {}
func (*AuthzGrantRevokeMsgType) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1fabd206b7c56d1, []int{4}
}
func (m *AuthzGrantRevokeMsgType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthzGrantRevokeMsgType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthzGrantRevokeMsgType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthzGrantRevokeMsgType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthzGrantRevokeMsgType.Merge(m, src)
}
func (m *AuthzGrantRevokeMsgType) XXX_Size() int {
	return m.Size()
}
func (m *AuthzGrantRevokeMsgType) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthzGrantRevokeMsgType.DiscardUnknown(m)
}

var xxx_messageInfo_AuthzGrantRevokeMsgType proto.InternalMessageInfo

func (m *AuthzGrantRevokeMsgType) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *AuthzGrantRevokeMsgType) GetGrant() bool {
	if m != nil {
		return m.Grant
	}
	return false
}

type AuthzGrantRevokeMsgsTypes struct {
	Msgs []*AuthzGrantRevokeMsgType `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs,omitempty"`
}

func (m *AuthzGrantRevokeMsgsTypes) Reset()         { *m = AuthzGrantRevokeMsgsTypes{} }
func (m *AuthzGrantRevokeMsgsTypes) String() string { return proto.CompactTextString(m) }
func (*AuthzGrantRevokeMsgsTypes) ProtoMessage()    {}
func (*AuthzGrantRevokeMsgsTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1fabd206b7c56d1, []int{5}
}
func (m *AuthzGrantRevokeMsgsTypes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthzGrantRevokeMsgsTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthzGrantRevokeMsgsTypes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthzGrantRevokeMsgsTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthzGrantRevokeMsgsTypes.Merge(m, src)
}
func (m *AuthzGrantRevokeMsgsTypes) XXX_Size() int {
	return m.Size()
}
func (m *AuthzGrantRevokeMsgsTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthzGrantRevokeMsgsTypes.DiscardUnknown(m)
}

var xxx_messageInfo_AuthzGrantRevokeMsgsTypes proto.InternalMessageInfo

func (m *AuthzGrantRevokeMsgsTypes) GetMsgs() []*AuthzGrantRevokeMsgType {
	if m != nil {
		return m.Msgs
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgAccounts)(nil), "aumega.guard.v1.MsgAccounts")
	proto.RegisterType((*MsgIndexes)(nil), "aumega.guard.v1.MsgIndexes")
	proto.RegisterType((*AccountPrivileges)(nil), "aumega.guard.v1.AccountPrivileges")
	proto.RegisterType((*RequiredPrivileges)(nil), "aumega.guard.v1.RequiredPrivileges")
	proto.RegisterType((*AuthzGrantRevokeMsgType)(nil), "aumega.guard.v1.AuthzGrantRevokeMsgType")
	proto.RegisterType((*AuthzGrantRevokeMsgsTypes)(nil), "aumega.guard.v1.AuthzGrantRevokeMsgsTypes")
}

func init() { proto.RegisterFile("aumega/guard/v1/common.proto", fileDescriptor_a1fabd206b7c56d1) }

var fileDescriptor_a1fabd206b7c56d1 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x6b, 0x76, 0x61, 0xdb, 0x59, 0x24, 0x84, 0x55, 0x89, 0x74, 0x85, 0xac, 0x2a, 0x07,
	0xd4, 0xcb, 0x26, 0x5a, 0xe0, 0xc8, 0x25, 0x8b, 0x04, 0x02, 0x11, 0x84, 0xcc, 0x72, 0x80, 0x03,
	0xab, 0x6c, 0x62, 0xb9, 0x56, 0x1b, 0x3b, 0xd8, 0x49, 0xd4, 0xf2, 0x14, 0x3c, 0x06, 0x0f, 0xc0,
	0x43, 0x70, 0xac, 0x38, 0x71, 0x44, 0xe9, 0x8b, 0xa0, 0xd8, 0xa6, 0xaa, 0x10, 0x68, 0x6f, 0xf3,
	0xcd, 0x7c, 0xf3, 0xfb, 0xf2, 0x67, 0xe0, 0x7e, 0xd6, 0x94, 0x8c, 0x67, 0x31, 0x6f, 0x32, 0x5d,
	0xc4, 0xed, 0x59, 0x9c, 0xab, 0xb2, 0x54, 0x32, 0xaa, 0xb4, 0xaa, 0x15, 0xbe, 0xe3, 0xa6, 0x91,
	0x9d, 0x46, 0xed, 0xd9, 0xc9, 0x24, 0x57, 0xa6, 0x54, 0xe6, 0xd2, 0x8e, 0x63, 0x27, 0x9c, 0xf7,
	0x64, 0xcc, 0x15, 0x57, 0xae, 0xdf, 0x57, 0xae, 0x1b, 0x3e, 0x85, 0xe3, 0xd4, 0xf0, 0x24, 0xcf,
	0x55, 0x23, 0x6b, 0x83, 0x1f, 0xc3, 0x30, 0xf3, 0x75, 0x80, 0xa6, 0x07, 0xb3, 0xd1, 0x79, 0xf0,
	0xe3, 0xdb, 0xe9, 0xd8, 0x83, 0x92, 0xa2, 0xd0, 0xcc, 0x98, 0xb7, 0xb5, 0x16, 0x92, 0xd3, 0x9d,
	0x33, 0x7c, 0x00, 0x90, 0x1a, 0xfe, 0x42, 0x16, 0x6c, 0xc5, 0x0c, 0x0e, 0xe0, 0x48, 0xb8, 0xd2,
	0x22, 0x6e, 0xd3, 0x3f, 0x32, 0x4c, 0xe1, 0xae, 0x4f, 0x7a, 0xa3, 0x45, 0x2b, 0x96, 0x8c, 0x3b,
	0xbb, 0x07, 0x05, 0x68, 0x8a, 0x7a, 0xbb, 0x97, 0x98, 0x00, 0x54, 0x3b, 0x5f, 0x70, 0xc3, 0x0e,
	0xf7, 0x3a, 0xe1, 0x47, 0xc0, 0x94, 0x7d, 0x6a, 0x84, 0x66, 0xc5, 0x1e, 0x6f, 0x0c, 0x37, 0x6d,
	0x9e, 0xa7, 0x39, 0x71, 0x1d, 0x0b, 0x63, 0x38, 0x5c, 0x08, 0x59, 0x04, 0x07, 0x53, 0x34, 0x1b,
	0x51, 0x5b, 0x87, 0x2f, 0xe1, 0x5e, 0xd2, 0xd4, 0xf3, 0xcf, 0xcf, 0x75, 0x26, 0x6b, 0xca, 0x5a,
	0xb5, 0x60, 0xa9, 0xe1, 0x17, 0xeb, 0x8a, 0xe1, 0x09, 0x0c, 0xeb, 0x75, 0xc5, 0x2e, 0x1b, 0xbd,
	0xb4, 0x39, 0x23, 0x7a, 0xd4, 0xeb, 0x77, 0x7a, 0xd9, 0xe7, 0xf3, 0x7e, 0xc1, 0x86, 0x0c, 0xa9,
	0x13, 0xe1, 0x7b, 0x98, 0xfc, 0x83, 0x65, 0x7a, 0x98, 0xc1, 0x4f, 0xe0, 0xb0, 0x34, 0xdc, 0x7d,
	0xae, 0xe3, 0x87, 0xb3, 0xe8, 0xaf, 0xbf, 0x1a, 0xfd, 0xe7, 0x29, 0xa8, 0xdd, 0x3a, 0x7f, 0xf5,
	0xb5, 0x23, 0xe8, 0x7b, 0x47, 0xd0, 0xa6, 0x23, 0xe8, 0x57, 0x47, 0xd0, 0x97, 0x2d, 0x19, 0x6c,
	0xb6, 0x64, 0xf0, 0x73, 0x4b, 0x06, 0x1f, 0x22, 0x2e, 0xea, 0x79, 0x73, 0x15, 0xe5, 0xaa, 0x8c,
	0xd3, 0xe4, 0xf5, 0x05, 0x4d, 0x4e, 0x9f, 0x09, 0x99, 0xc9, 0x9c, 0xc5, 0xfe, 0xb4, 0x56, 0xfe,
	0xb8, 0xfa, 0x17, 0x30, 0x57, 0xb7, 0xec, 0x5d, 0x3c, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xa1,
	0x3c, 0xee, 0xb7, 0x79, 0x02, 0x00, 0x00,
}

func (this *MsgAccounts) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgAccounts)
	if !ok {
		that2, ok := that.(MsgAccounts)
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
	if len(this.Accounts) != len(that1.Accounts) {
		return false
	}
	for i := range this.Accounts {
		if this.Accounts[i] != that1.Accounts[i] {
			return false
		}
	}
	return true
}
func (this *MsgIndexes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgIndexes)
	if !ok {
		that2, ok := that.(MsgIndexes)
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
	if len(this.Indexes) != len(that1.Indexes) {
		return false
	}
	for i := range this.Indexes {
		if !bytes.Equal(this.Indexes[i], that1.Indexes[i]) {
			return false
		}
	}
	return true
}
func (this *AccountPrivileges) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccountPrivileges)
	if !ok {
		that2, ok := that.(AccountPrivileges)
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
	if !bytes.Equal(this.Account, that1.Account) {
		return false
	}
	if !bytes.Equal(this.Privileges, that1.Privileges) {
		return false
	}
	return true
}
func (this *RequiredPrivileges) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequiredPrivileges)
	if !ok {
		that2, ok := that.(RequiredPrivileges)
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
	if !bytes.Equal(this.Index, that1.Index) {
		return false
	}
	if !bytes.Equal(this.Privileges, that1.Privileges) {
		return false
	}
	if this.Kind != that1.Kind {
		return false
	}
	return true
}
func (this *AuthzGrantRevokeMsgType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AuthzGrantRevokeMsgType)
	if !ok {
		that2, ok := that.(AuthzGrantRevokeMsgType)
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
	if this.TypeUrl != that1.TypeUrl {
		return false
	}
	if this.Grant != that1.Grant {
		return false
	}
	return true
}
func (this *AuthzGrantRevokeMsgsTypes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AuthzGrantRevokeMsgsTypes)
	if !ok {
		that2, ok := that.(AuthzGrantRevokeMsgsTypes)
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
	if len(this.Msgs) != len(that1.Msgs) {
		return false
	}
	for i := range this.Msgs {
		if !this.Msgs[i].Equal(that1.Msgs[i]) {
			return false
		}
	}
	return true
}
func (m *MsgAccounts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAccounts) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAccounts) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for iNdEx := len(m.Accounts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Accounts[iNdEx])
			copy(dAtA[i:], m.Accounts[iNdEx])
			i = encodeVarintCommon(dAtA, i, uint64(len(m.Accounts[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgIndexes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIndexes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIndexes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Indexes) > 0 {
		for iNdEx := len(m.Indexes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Indexes[iNdEx])
			copy(dAtA[i:], m.Indexes[iNdEx])
			i = encodeVarintCommon(dAtA, i, uint64(len(m.Indexes[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AccountPrivileges) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountPrivileges) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountPrivileges) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Privileges) > 0 {
		i -= len(m.Privileges)
		copy(dAtA[i:], m.Privileges)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Privileges)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequiredPrivileges) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequiredPrivileges) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequiredPrivileges) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Kind) > 0 {
		i -= len(m.Kind)
		copy(dAtA[i:], m.Kind)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Kind)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Privileges) > 0 {
		i -= len(m.Privileges)
		copy(dAtA[i:], m.Privileges)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Privileges)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthzGrantRevokeMsgType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthzGrantRevokeMsgType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthzGrantRevokeMsgType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Grant {
		i--
		if m.Grant {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.TypeUrl) > 0 {
		i -= len(m.TypeUrl)
		copy(dAtA[i:], m.TypeUrl)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.TypeUrl)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthzGrantRevokeMsgsTypes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthzGrantRevokeMsgsTypes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthzGrantRevokeMsgsTypes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for iNdEx := len(m.Msgs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Msgs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommon(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *MsgAccounts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for _, s := range m.Accounts {
			l = len(s)
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	return n
}

func (m *MsgIndexes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Indexes) > 0 {
		for _, b := range m.Indexes {
			l = len(b)
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	return n
}

func (m *AccountPrivileges) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Privileges)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *RequiredPrivileges) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Privileges)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Kind)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *AuthzGrantRevokeMsgType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TypeUrl)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.Grant {
		n += 2
	}
	return n
}

func (m *AuthzGrantRevokeMsgsTypes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAccounts) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAccounts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAccounts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accounts", wireType)
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
			m.Accounts = append(m.Accounts, string(dAtA[iNdEx:postIndex]))
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
func (m *MsgIndexes) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgIndexes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIndexes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Indexes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Indexes = append(m.Indexes, make([]byte, postIndex-iNdEx))
			copy(m.Indexes[len(m.Indexes)-1], dAtA[iNdEx:postIndex])
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
func (m *AccountPrivileges) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AccountPrivileges: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountPrivileges: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = append(m.Account[:0], dAtA[iNdEx:postIndex]...)
			if m.Account == nil {
				m.Account = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Privileges", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Privileges = append(m.Privileges[:0], dAtA[iNdEx:postIndex]...)
			if m.Privileges == nil {
				m.Privileges = []byte{}
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
func (m *RequiredPrivileges) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RequiredPrivileges: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequiredPrivileges: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = append(m.Index[:0], dAtA[iNdEx:postIndex]...)
			if m.Index == nil {
				m.Index = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Privileges", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Privileges = append(m.Privileges[:0], dAtA[iNdEx:postIndex]...)
			if m.Privileges == nil {
				m.Privileges = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
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
			m.Kind = string(dAtA[iNdEx:postIndex])
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
func (m *AuthzGrantRevokeMsgType) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AuthzGrantRevokeMsgType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthzGrantRevokeMsgType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeUrl", wireType)
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
			m.TypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grant", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Grant = bool(v != 0)
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
func (m *AuthzGrantRevokeMsgsTypes) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AuthzGrantRevokeMsgsTypes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthzGrantRevokeMsgsTypes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msgs = append(m.Msgs, &AuthzGrantRevokeMsgType{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
