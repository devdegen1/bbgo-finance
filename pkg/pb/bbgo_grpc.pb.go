// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// MarketDataServiceClient is the client API for MarketDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketDataServiceClient interface {
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (MarketDataService_SubscribeClient, error)
	QueryKLines(ctx context.Context, in *QueryKLinesRequest, opts ...grpc.CallOption) (*QueryKLinesResponse, error)
}

type marketDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketDataServiceClient(cc grpc.ClientConnInterface) MarketDataServiceClient {
	return &marketDataServiceClient{cc}
}

func (c *marketDataServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (MarketDataService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &MarketDataService_ServiceDesc.Streams[0], "/bbgo.MarketDataService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &marketDataServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MarketDataService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type marketDataServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *marketDataServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *marketDataServiceClient) QueryKLines(ctx context.Context, in *QueryKLinesRequest, opts ...grpc.CallOption) (*QueryKLinesResponse, error) {
	out := new(QueryKLinesResponse)
	err := c.cc.Invoke(ctx, "/bbgo.MarketDataService/QueryKLines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketDataServiceServer is the server API for MarketDataService service.
// All implementations must embed UnimplementedMarketDataServiceServer
// for forward compatibility
type MarketDataServiceServer interface {
	Subscribe(*SubscribeRequest, MarketDataService_SubscribeServer) error
	QueryKLines(context.Context, *QueryKLinesRequest) (*QueryKLinesResponse, error)
	mustEmbedUnimplementedMarketDataServiceServer()
}

// UnimplementedMarketDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMarketDataServiceServer struct {
}

func (UnimplementedMarketDataServiceServer) Subscribe(*SubscribeRequest, MarketDataService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedMarketDataServiceServer) QueryKLines(context.Context, *QueryKLinesRequest) (*QueryKLinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryKLines not implemented")
}
func (UnimplementedMarketDataServiceServer) mustEmbedUnimplementedMarketDataServiceServer() {}

// UnsafeMarketDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketDataServiceServer will
// result in compilation errors.
type UnsafeMarketDataServiceServer interface {
	mustEmbedUnimplementedMarketDataServiceServer()
}

func RegisterMarketDataServiceServer(s grpc.ServiceRegistrar, srv MarketDataServiceServer) {
	s.RegisterService(&MarketDataService_ServiceDesc, srv)
}

func _MarketDataService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MarketDataServiceServer).Subscribe(m, &marketDataServiceSubscribeServer{stream})
}

type MarketDataService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type marketDataServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *marketDataServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MarketDataService_QueryKLines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryKLinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketDataServiceServer).QueryKLines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bbgo.MarketDataService/QueryKLines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketDataServiceServer).QueryKLines(ctx, req.(*QueryKLinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MarketDataService_ServiceDesc is the grpc.ServiceDesc for MarketDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarketDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bbgo.MarketDataService",
	HandlerType: (*MarketDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryKLines",
			Handler:    _MarketDataService_QueryKLines_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _MarketDataService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/pb/bbgo.proto",
}

// UserDataServiceClient is the client API for UserDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserDataServiceClient interface {
	// should support streaming
	SubscribeUserData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UserDataService_SubscribeUserDataClient, error)
}

type userDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserDataServiceClient(cc grpc.ClientConnInterface) UserDataServiceClient {
	return &userDataServiceClient{cc}
}

func (c *userDataServiceClient) SubscribeUserData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UserDataService_SubscribeUserDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserDataService_ServiceDesc.Streams[0], "/bbgo.UserDataService/SubscribeUserData", opts...)
	if err != nil {
		return nil, err
	}
	x := &userDataServiceSubscribeUserDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserDataService_SubscribeUserDataClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type userDataServiceSubscribeUserDataClient struct {
	grpc.ClientStream
}

func (x *userDataServiceSubscribeUserDataClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserDataServiceServer is the server API for UserDataService service.
// All implementations must embed UnimplementedUserDataServiceServer
// for forward compatibility
type UserDataServiceServer interface {
	// should support streaming
	SubscribeUserData(*Empty, UserDataService_SubscribeUserDataServer) error
	mustEmbedUnimplementedUserDataServiceServer()
}

// UnimplementedUserDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserDataServiceServer struct {
}

func (UnimplementedUserDataServiceServer) SubscribeUserData(*Empty, UserDataService_SubscribeUserDataServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeUserData not implemented")
}
func (UnimplementedUserDataServiceServer) mustEmbedUnimplementedUserDataServiceServer() {}

// UnsafeUserDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserDataServiceServer will
// result in compilation errors.
type UnsafeUserDataServiceServer interface {
	mustEmbedUnimplementedUserDataServiceServer()
}

func RegisterUserDataServiceServer(s grpc.ServiceRegistrar, srv UserDataServiceServer) {
	s.RegisterService(&UserDataService_ServiceDesc, srv)
}

func _UserDataService_SubscribeUserData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserDataServiceServer).SubscribeUserData(m, &userDataServiceSubscribeUserDataServer{stream})
}

type UserDataService_SubscribeUserDataServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type userDataServiceSubscribeUserDataServer struct {
	grpc.ServerStream
}

func (x *userDataServiceSubscribeUserDataServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

// UserDataService_ServiceDesc is the grpc.ServiceDesc for UserDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bbgo.UserDataService",
	HandlerType: (*UserDataServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeUserData",
			Handler:       _UserDataService_SubscribeUserData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/pb/bbgo.proto",
}

// TradingServiceClient is the client API for TradingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TradingServiceClient interface {
	// request-response
	SubmitOrder(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*SubmitOrderResponse, error)
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
	QueryOrder(ctx context.Context, in *QueryOrderRequest, opts ...grpc.CallOption) (*QueryOrderResponse, error)
	QueryOrders(ctx context.Context, in *QueryOrdersRequest, opts ...grpc.CallOption) (*QueryOrdersResponse, error)
	QueryTrades(ctx context.Context, in *QueryTradesRequest, opts ...grpc.CallOption) (*QueryTradesResponse, error)
}

type tradingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTradingServiceClient(cc grpc.ClientConnInterface) TradingServiceClient {
	return &tradingServiceClient{cc}
}

func (c *tradingServiceClient) SubmitOrder(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*SubmitOrderResponse, error) {
	out := new(SubmitOrderResponse)
	err := c.cc.Invoke(ctx, "/bbgo.TradingService/SubmitOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, "/bbgo.TradingService/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingServiceClient) QueryOrder(ctx context.Context, in *QueryOrderRequest, opts ...grpc.CallOption) (*QueryOrderResponse, error) {
	out := new(QueryOrderResponse)
	err := c.cc.Invoke(ctx, "/bbgo.TradingService/QueryOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingServiceClient) QueryOrders(ctx context.Context, in *QueryOrdersRequest, opts ...grpc.CallOption) (*QueryOrdersResponse, error) {
	out := new(QueryOrdersResponse)
	err := c.cc.Invoke(ctx, "/bbgo.TradingService/QueryOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingServiceClient) QueryTrades(ctx context.Context, in *QueryTradesRequest, opts ...grpc.CallOption) (*QueryTradesResponse, error) {
	out := new(QueryTradesResponse)
	err := c.cc.Invoke(ctx, "/bbgo.TradingService/QueryTrades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradingServiceServer is the server API for TradingService service.
// All implementations must embed UnimplementedTradingServiceServer
// for forward compatibility
type TradingServiceServer interface {
	// request-response
	SubmitOrder(context.Context, *SubmitOrderRequest) (*SubmitOrderResponse, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	QueryOrder(context.Context, *QueryOrderRequest) (*QueryOrderResponse, error)
	QueryOrders(context.Context, *QueryOrdersRequest) (*QueryOrdersResponse, error)
	QueryTrades(context.Context, *QueryTradesRequest) (*QueryTradesResponse, error)
	mustEmbedUnimplementedTradingServiceServer()
}

// UnimplementedTradingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTradingServiceServer struct {
}

func (UnimplementedTradingServiceServer) SubmitOrder(context.Context, *SubmitOrderRequest) (*SubmitOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitOrder not implemented")
}
func (UnimplementedTradingServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedTradingServiceServer) QueryOrder(context.Context, *QueryOrderRequest) (*QueryOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOrder not implemented")
}
func (UnimplementedTradingServiceServer) QueryOrders(context.Context, *QueryOrdersRequest) (*QueryOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOrders not implemented")
}
func (UnimplementedTradingServiceServer) QueryTrades(context.Context, *QueryTradesRequest) (*QueryTradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTrades not implemented")
}
func (UnimplementedTradingServiceServer) mustEmbedUnimplementedTradingServiceServer() {}

// UnsafeTradingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TradingServiceServer will
// result in compilation errors.
type UnsafeTradingServiceServer interface {
	mustEmbedUnimplementedTradingServiceServer()
}

func RegisterTradingServiceServer(s grpc.ServiceRegistrar, srv TradingServiceServer) {
	s.RegisterService(&TradingService_ServiceDesc, srv)
}

func _TradingService_SubmitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServiceServer).SubmitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bbgo.TradingService/SubmitOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServiceServer).SubmitOrder(ctx, req.(*SubmitOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradingService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bbgo.TradingService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradingService_QueryOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServiceServer).QueryOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bbgo.TradingService/QueryOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServiceServer).QueryOrder(ctx, req.(*QueryOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradingService_QueryOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServiceServer).QueryOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bbgo.TradingService/QueryOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServiceServer).QueryOrders(ctx, req.(*QueryOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradingService_QueryTrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServiceServer).QueryTrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bbgo.TradingService/QueryTrades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServiceServer).QueryTrades(ctx, req.(*QueryTradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TradingService_ServiceDesc is the grpc.ServiceDesc for TradingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TradingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bbgo.TradingService",
	HandlerType: (*TradingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitOrder",
			Handler:    _TradingService_SubmitOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _TradingService_CancelOrder_Handler,
		},
		{
			MethodName: "QueryOrder",
			Handler:    _TradingService_QueryOrder_Handler,
		},
		{
			MethodName: "QueryOrders",
			Handler:    _TradingService_QueryOrders_Handler,
		},
		{
			MethodName: "QueryTrades",
			Handler:    _TradingService_QueryTrades_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/bbgo.proto",
}
