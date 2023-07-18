// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/pb/order.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderManagementClient is the client API for OrderManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderManagementClient interface {
	SaveCartItems(ctx context.Context, in *CartItemsSaveRequest, opts ...grpc.CallOption) (*CartItemsSaveResponse, error)
	CashOnDelivery(ctx context.Context, in *CashOnDeliveryRequest, opts ...grpc.CallOption) (*CashOnDeliveryResponse, error)
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	OrderDispacth(ctx context.Context, in *OrderDispacthRequest, opts ...grpc.CallOption) (*OrderDispacthResponse, error)
	OrderCancel(ctx context.Context, in *OrderCancelRequest, opts ...grpc.CallOption) (*OrderCancelResponse, error)
	AdminViewAllOrders(ctx context.Context, in *AdminViewAllOrdersRequest, opts ...grpc.CallOption) (*AdminViewAllOrdersResponse, error)
	ViewOrderById(ctx context.Context, in *ViewOrderByIdRequest, opts ...grpc.CallOption) (*ViewOrderByIdResponse, error)
}

type orderManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderManagementClient(cc grpc.ClientConnInterface) OrderManagementClient {
	return &orderManagementClient{cc}
}

func (c *orderManagementClient) SaveCartItems(ctx context.Context, in *CartItemsSaveRequest, opts ...grpc.CallOption) (*CartItemsSaveResponse, error) {
	out := new(CartItemsSaveResponse)
	err := c.cc.Invoke(ctx, "/order.OrderManagement/SaveCartItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) CashOnDelivery(ctx context.Context, in *CashOnDeliveryRequest, opts ...grpc.CallOption) (*CashOnDeliveryResponse, error) {
	out := new(CashOnDeliveryResponse)
	err := c.cc.Invoke(ctx, "/order.OrderManagement/CashOnDelivery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, "/order.OrderManagement/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) OrderDispacth(ctx context.Context, in *OrderDispacthRequest, opts ...grpc.CallOption) (*OrderDispacthResponse, error) {
	out := new(OrderDispacthResponse)
	err := c.cc.Invoke(ctx, "/order.OrderManagement/OrderDispacth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) OrderCancel(ctx context.Context, in *OrderCancelRequest, opts ...grpc.CallOption) (*OrderCancelResponse, error) {
	out := new(OrderCancelResponse)
	err := c.cc.Invoke(ctx, "/order.OrderManagement/OrderCancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) AdminViewAllOrders(ctx context.Context, in *AdminViewAllOrdersRequest, opts ...grpc.CallOption) (*AdminViewAllOrdersResponse, error) {
	out := new(AdminViewAllOrdersResponse)
	err := c.cc.Invoke(ctx, "/order.OrderManagement/AdminViewAllOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) ViewOrderById(ctx context.Context, in *ViewOrderByIdRequest, opts ...grpc.CallOption) (*ViewOrderByIdResponse, error) {
	out := new(ViewOrderByIdResponse)
	err := c.cc.Invoke(ctx, "/order.OrderManagement/ViewOrderById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderManagementServer is the server API for OrderManagement service.
// All implementations must embed UnimplementedOrderManagementServer
// for forward compatibility
type OrderManagementServer interface {
	SaveCartItems(context.Context, *CartItemsSaveRequest) (*CartItemsSaveResponse, error)
	CashOnDelivery(context.Context, *CashOnDeliveryRequest) (*CashOnDeliveryResponse, error)
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	OrderDispacth(context.Context, *OrderDispacthRequest) (*OrderDispacthResponse, error)
	OrderCancel(context.Context, *OrderCancelRequest) (*OrderCancelResponse, error)
	AdminViewAllOrders(context.Context, *AdminViewAllOrdersRequest) (*AdminViewAllOrdersResponse, error)
	ViewOrderById(context.Context, *ViewOrderByIdRequest) (*ViewOrderByIdResponse, error)
	mustEmbedUnimplementedOrderManagementServer()
}

// UnimplementedOrderManagementServer must be embedded to have forward compatible implementations.
type UnimplementedOrderManagementServer struct {
}

func (UnimplementedOrderManagementServer) SaveCartItems(context.Context, *CartItemsSaveRequest) (*CartItemsSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveCartItems not implemented")
}
func (UnimplementedOrderManagementServer) CashOnDelivery(context.Context, *CashOnDeliveryRequest) (*CashOnDeliveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CashOnDelivery not implemented")
}
func (UnimplementedOrderManagementServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderManagementServer) OrderDispacth(context.Context, *OrderDispacthRequest) (*OrderDispacthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDispacth not implemented")
}
func (UnimplementedOrderManagementServer) OrderCancel(context.Context, *OrderCancelRequest) (*OrderCancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderCancel not implemented")
}
func (UnimplementedOrderManagementServer) AdminViewAllOrders(context.Context, *AdminViewAllOrdersRequest) (*AdminViewAllOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminViewAllOrders not implemented")
}
func (UnimplementedOrderManagementServer) ViewOrderById(context.Context, *ViewOrderByIdRequest) (*ViewOrderByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewOrderById not implemented")
}
func (UnimplementedOrderManagementServer) mustEmbedUnimplementedOrderManagementServer() {}

// UnsafeOrderManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderManagementServer will
// result in compilation errors.
type UnsafeOrderManagementServer interface {
	mustEmbedUnimplementedOrderManagementServer()
}

func RegisterOrderManagementServer(s grpc.ServiceRegistrar, srv OrderManagementServer) {
	s.RegisterService(&OrderManagement_ServiceDesc, srv)
}

func _OrderManagement_SaveCartItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemsSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).SaveCartItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderManagement/SaveCartItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).SaveCartItems(ctx, req.(*CartItemsSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_CashOnDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CashOnDeliveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).CashOnDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderManagement/CashOnDelivery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).CashOnDelivery(ctx, req.(*CashOnDeliveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderManagement/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_OrderDispacth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderDispacthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).OrderDispacth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderManagement/OrderDispacth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).OrderDispacth(ctx, req.(*OrderDispacthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_OrderCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderCancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).OrderCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderManagement/OrderCancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).OrderCancel(ctx, req.(*OrderCancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_AdminViewAllOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminViewAllOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).AdminViewAllOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderManagement/AdminViewAllOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).AdminViewAllOrders(ctx, req.(*AdminViewAllOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_ViewOrderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewOrderByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).ViewOrderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderManagement/ViewOrderById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).ViewOrderById(ctx, req.(*ViewOrderByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderManagement_ServiceDesc is the grpc.ServiceDesc for OrderManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderManagement",
	HandlerType: (*OrderManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveCartItems",
			Handler:    _OrderManagement_SaveCartItems_Handler,
		},
		{
			MethodName: "CashOnDelivery",
			Handler:    _OrderManagement_CashOnDelivery_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _OrderManagement_CreateOrder_Handler,
		},
		{
			MethodName: "OrderDispacth",
			Handler:    _OrderManagement_OrderDispacth_Handler,
		},
		{
			MethodName: "OrderCancel",
			Handler:    _OrderManagement_OrderCancel_Handler,
		},
		{
			MethodName: "AdminViewAllOrders",
			Handler:    _OrderManagement_AdminViewAllOrders_Handler,
		},
		{
			MethodName: "ViewOrderById",
			Handler:    _OrderManagement_ViewOrderById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/order.proto",
}