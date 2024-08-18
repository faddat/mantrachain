// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: mantrachain/liquidity/v1beta1/tx.proto

package liquidityv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Msg_UpdateParams_FullMethodName      = "/mantrachain.liquidity.v1beta1.Msg/UpdateParams"
	Msg_CreatePair_FullMethodName        = "/mantrachain.liquidity.v1beta1.Msg/CreatePair"
	Msg_UpdatePairSwapFee_FullMethodName = "/mantrachain.liquidity.v1beta1.Msg/UpdatePairSwapFee"
	Msg_CreatePool_FullMethodName        = "/mantrachain.liquidity.v1beta1.Msg/CreatePool"
	Msg_CreateRangedPool_FullMethodName  = "/mantrachain.liquidity.v1beta1.Msg/CreateRangedPool"
	Msg_Deposit_FullMethodName           = "/mantrachain.liquidity.v1beta1.Msg/Deposit"
	Msg_Withdraw_FullMethodName          = "/mantrachain.liquidity.v1beta1.Msg/Withdraw"
	Msg_LimitOrder_FullMethodName        = "/mantrachain.liquidity.v1beta1.Msg/LimitOrder"
	Msg_MarketOrder_FullMethodName       = "/mantrachain.liquidity.v1beta1.Msg/MarketOrder"
	Msg_MMOrder_FullMethodName           = "/mantrachain.liquidity.v1beta1.Msg/MMOrder"
	Msg_CancelOrder_FullMethodName       = "/mantrachain.liquidity.v1beta1.Msg/CancelOrder"
	Msg_CancelAllOrders_FullMethodName   = "/mantrachain.liquidity.v1beta1.Msg/CancelAllOrders"
	Msg_CancelMMOrder_FullMethodName     = "/mantrachain.liquidity.v1beta1.Msg/CancelMMOrder"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Msg defines the Msg service.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	// CreatePair defines a method for creating a pair
	CreatePair(ctx context.Context, in *MsgCreatePair, opts ...grpc.CallOption) (*MsgCreatePairResponse, error)
	// UpdatePairSwapFee defines a method for creating a pair
	UpdatePairSwapFee(ctx context.Context, in *MsgUpdatePairSwapFee, opts ...grpc.CallOption) (*MsgUpdatePairSwapFeeResponse, error)
	// CreatePool defines a method for creating a pool
	CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error)
	// CreateRangePool defines a method for creating a ranged pool
	CreateRangedPool(ctx context.Context, in *MsgCreateRangedPool, opts ...grpc.CallOption) (*MsgCreateRangedPoolResponse, error)
	// Deposit defines a method for depositing coins to the pool
	Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error)
	// Withdraw defines a method for withdrawing pool coin from the pool
	Withdraw(ctx context.Context, in *MsgWithdraw, opts ...grpc.CallOption) (*MsgWithdrawResponse, error)
	// LimitOrder defines a method for making a limit order
	LimitOrder(ctx context.Context, in *MsgLimitOrder, opts ...grpc.CallOption) (*MsgLimitOrderResponse, error)
	// MarketOrder defines a method for making a market order
	MarketOrder(ctx context.Context, in *MsgMarketOrder, opts ...grpc.CallOption) (*MsgMarketOrderResponse, error)
	// MsgMMOrder defines a method for making a MM(market making) order
	MMOrder(ctx context.Context, in *MsgMMOrder, opts ...grpc.CallOption) (*MsgMMOrderResponse, error)
	// CancelOrder defines a method for cancelling an order
	CancelOrder(ctx context.Context, in *MsgCancelOrder, opts ...grpc.CallOption) (*MsgCancelOrderResponse, error)
	// CancelAllOrders defines a method for cancelling all orders
	CancelAllOrders(ctx context.Context, in *MsgCancelAllOrders, opts ...grpc.CallOption) (*MsgCancelAllOrdersResponse, error)
	// CancelMMOrder defines a method for cancelling previously placed market making orders
	CancelMMOrder(ctx context.Context, in *MsgCancelMMOrder, opts ...grpc.CallOption) (*MsgCancelMMOrderResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePair(ctx context.Context, in *MsgCreatePair, opts ...grpc.CallOption) (*MsgCreatePairResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgCreatePairResponse)
	err := c.cc.Invoke(ctx, Msg_CreatePair_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdatePairSwapFee(ctx context.Context, in *MsgUpdatePairSwapFee, opts ...grpc.CallOption) (*MsgUpdatePairSwapFeeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgUpdatePairSwapFeeResponse)
	err := c.cc.Invoke(ctx, Msg_UpdatePairSwapFee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgCreatePoolResponse)
	err := c.cc.Invoke(ctx, Msg_CreatePool_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateRangedPool(ctx context.Context, in *MsgCreateRangedPool, opts ...grpc.CallOption) (*MsgCreateRangedPoolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgCreateRangedPoolResponse)
	err := c.cc.Invoke(ctx, Msg_CreateRangedPool_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgDepositResponse)
	err := c.cc.Invoke(ctx, Msg_Deposit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Withdraw(ctx context.Context, in *MsgWithdraw, opts ...grpc.CallOption) (*MsgWithdrawResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgWithdrawResponse)
	err := c.cc.Invoke(ctx, Msg_Withdraw_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) LimitOrder(ctx context.Context, in *MsgLimitOrder, opts ...grpc.CallOption) (*MsgLimitOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgLimitOrderResponse)
	err := c.cc.Invoke(ctx, Msg_LimitOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MarketOrder(ctx context.Context, in *MsgMarketOrder, opts ...grpc.CallOption) (*MsgMarketOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgMarketOrderResponse)
	err := c.cc.Invoke(ctx, Msg_MarketOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MMOrder(ctx context.Context, in *MsgMMOrder, opts ...grpc.CallOption) (*MsgMMOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgMMOrderResponse)
	err := c.cc.Invoke(ctx, Msg_MMOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelOrder(ctx context.Context, in *MsgCancelOrder, opts ...grpc.CallOption) (*MsgCancelOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgCancelOrderResponse)
	err := c.cc.Invoke(ctx, Msg_CancelOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelAllOrders(ctx context.Context, in *MsgCancelAllOrders, opts ...grpc.CallOption) (*MsgCancelAllOrdersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgCancelAllOrdersResponse)
	err := c.cc.Invoke(ctx, Msg_CancelAllOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelMMOrder(ctx context.Context, in *MsgCancelMMOrder, opts ...grpc.CallOption) (*MsgCancelMMOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgCancelMMOrderResponse)
	err := c.cc.Invoke(ctx, Msg_CancelMMOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility.
//
// Msg defines the Msg service.
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	// CreatePair defines a method for creating a pair
	CreatePair(context.Context, *MsgCreatePair) (*MsgCreatePairResponse, error)
	// UpdatePairSwapFee defines a method for creating a pair
	UpdatePairSwapFee(context.Context, *MsgUpdatePairSwapFee) (*MsgUpdatePairSwapFeeResponse, error)
	// CreatePool defines a method for creating a pool
	CreatePool(context.Context, *MsgCreatePool) (*MsgCreatePoolResponse, error)
	// CreateRangePool defines a method for creating a ranged pool
	CreateRangedPool(context.Context, *MsgCreateRangedPool) (*MsgCreateRangedPoolResponse, error)
	// Deposit defines a method for depositing coins to the pool
	Deposit(context.Context, *MsgDeposit) (*MsgDepositResponse, error)
	// Withdraw defines a method for withdrawing pool coin from the pool
	Withdraw(context.Context, *MsgWithdraw) (*MsgWithdrawResponse, error)
	// LimitOrder defines a method for making a limit order
	LimitOrder(context.Context, *MsgLimitOrder) (*MsgLimitOrderResponse, error)
	// MarketOrder defines a method for making a market order
	MarketOrder(context.Context, *MsgMarketOrder) (*MsgMarketOrderResponse, error)
	// MsgMMOrder defines a method for making a MM(market making) order
	MMOrder(context.Context, *MsgMMOrder) (*MsgMMOrderResponse, error)
	// CancelOrder defines a method for cancelling an order
	CancelOrder(context.Context, *MsgCancelOrder) (*MsgCancelOrderResponse, error)
	// CancelAllOrders defines a method for cancelling all orders
	CancelAllOrders(context.Context, *MsgCancelAllOrders) (*MsgCancelAllOrdersResponse, error)
	// CancelMMOrder defines a method for cancelling previously placed market making orders
	CancelMMOrder(context.Context, *MsgCancelMMOrder) (*MsgCancelMMOrderResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMsgServer struct{}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreatePair(context.Context, *MsgCreatePair) (*MsgCreatePairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePair not implemented")
}
func (UnimplementedMsgServer) UpdatePairSwapFee(context.Context, *MsgUpdatePairSwapFee) (*MsgUpdatePairSwapFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePairSwapFee not implemented")
}
func (UnimplementedMsgServer) CreatePool(context.Context, *MsgCreatePool) (*MsgCreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}
func (UnimplementedMsgServer) CreateRangedPool(context.Context, *MsgCreateRangedPool) (*MsgCreateRangedPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRangedPool not implemented")
}
func (UnimplementedMsgServer) Deposit(context.Context, *MsgDeposit) (*MsgDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedMsgServer) Withdraw(context.Context, *MsgWithdraw) (*MsgWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedMsgServer) LimitOrder(context.Context, *MsgLimitOrder) (*MsgLimitOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LimitOrder not implemented")
}
func (UnimplementedMsgServer) MarketOrder(context.Context, *MsgMarketOrder) (*MsgMarketOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketOrder not implemented")
}
func (UnimplementedMsgServer) MMOrder(context.Context, *MsgMMOrder) (*MsgMMOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MMOrder not implemented")
}
func (UnimplementedMsgServer) CancelOrder(context.Context, *MsgCancelOrder) (*MsgCancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedMsgServer) CancelAllOrders(context.Context, *MsgCancelAllOrders) (*MsgCancelAllOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAllOrders not implemented")
}
func (UnimplementedMsgServer) CancelMMOrder(context.Context, *MsgCancelMMOrder) (*MsgCancelMMOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelMMOrder not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}
func (UnimplementedMsgServer) testEmbeddedByValue()             {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	// If the following call pancis, it indicates UnimplementedMsgServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePair_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePair(ctx, req.(*MsgCreatePair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdatePairSwapFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdatePairSwapFee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePairSwapFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdatePairSwapFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePairSwapFee(ctx, req.(*MsgUpdatePairSwapFee))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePool(ctx, req.(*MsgCreatePool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateRangedPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateRangedPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateRangedPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateRangedPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateRangedPool(ctx, req.(*MsgCreateRangedPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeposit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Deposit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Deposit(ctx, req.(*MsgDeposit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdraw)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Withdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Withdraw(ctx, req.(*MsgWithdraw))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_LimitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLimitOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LimitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_LimitOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LimitOrder(ctx, req.(*MsgLimitOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MarketOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMarketOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MarketOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MarketOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MarketOrder(ctx, req.(*MsgMarketOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MMOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMMOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MMOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MMOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MMOrder(ctx, req.(*MsgMMOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelOrder(ctx, req.(*MsgCancelOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelAllOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelAllOrders)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelAllOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CancelAllOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelAllOrders(ctx, req.(*MsgCancelAllOrders))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelMMOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelMMOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelMMOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CancelMMOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelMMOrder(ctx, req.(*MsgCancelMMOrder))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mantrachain.liquidity.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreatePair",
			Handler:    _Msg_CreatePair_Handler,
		},
		{
			MethodName: "UpdatePairSwapFee",
			Handler:    _Msg_UpdatePairSwapFee_Handler,
		},
		{
			MethodName: "CreatePool",
			Handler:    _Msg_CreatePool_Handler,
		},
		{
			MethodName: "CreateRangedPool",
			Handler:    _Msg_CreateRangedPool_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _Msg_Deposit_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _Msg_Withdraw_Handler,
		},
		{
			MethodName: "LimitOrder",
			Handler:    _Msg_LimitOrder_Handler,
		},
		{
			MethodName: "MarketOrder",
			Handler:    _Msg_MarketOrder_Handler,
		},
		{
			MethodName: "MMOrder",
			Handler:    _Msg_MMOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Msg_CancelOrder_Handler,
		},
		{
			MethodName: "CancelAllOrders",
			Handler:    _Msg_CancelAllOrders_Handler,
		},
		{
			MethodName: "CancelMMOrder",
			Handler:    _Msg_CancelMMOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mantrachain/liquidity/v1beta1/tx.proto",
}
