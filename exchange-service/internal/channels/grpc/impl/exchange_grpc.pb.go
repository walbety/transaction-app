// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: tools/protos/exchange.proto

package impl

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

const (
	ExchangeService_GetExchangeRateByCurrencyAndDate_FullMethodName = "/walbety.exchange.ExchangeService/GetExchangeRateByCurrencyAndDate"
)

// ExchangeServiceClient is the client API for ExchangeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExchangeServiceClient interface {
	GetExchangeRateByCurrencyAndDate(ctx context.Context, in *ExchangeRateRequest, opts ...grpc.CallOption) (*ExchangeRateResponse, error)
}

type exchangeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExchangeServiceClient(cc grpc.ClientConnInterface) ExchangeServiceClient {
	return &exchangeServiceClient{cc}
}

func (c *exchangeServiceClient) GetExchangeRateByCurrencyAndDate(ctx context.Context, in *ExchangeRateRequest, opts ...grpc.CallOption) (*ExchangeRateResponse, error) {
	out := new(ExchangeRateResponse)
	err := c.cc.Invoke(ctx, ExchangeService_GetExchangeRateByCurrencyAndDate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangeServiceServer is the server API for ExchangeService service.
// All implementations must embed UnimplementedExchangeServiceServer
// for forward compatibility
type ExchangeServiceServer interface {
	GetExchangeRateByCurrencyAndDate(context.Context, *ExchangeRateRequest) (*ExchangeRateResponse, error)
	mustEmbedUnimplementedExchangeServiceServer()
}

// UnimplementedExchangeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExchangeServiceServer struct {
}

func (UnimplementedExchangeServiceServer) GetExchangeRateByCurrencyAndDate(context.Context, *ExchangeRateRequest) (*ExchangeRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExchangeRateByCurrencyAndDate not implemented")
}
func (UnimplementedExchangeServiceServer) mustEmbedUnimplementedExchangeServiceServer() {}

// UnsafeExchangeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExchangeServiceServer will
// result in compilation errors.
type UnsafeExchangeServiceServer interface {
	mustEmbedUnimplementedExchangeServiceServer()
}

func RegisterExchangeServiceServer(s grpc.ServiceRegistrar, srv ExchangeServiceServer) {
	s.RegisterService(&ExchangeService_ServiceDesc, srv)
}

func _ExchangeService_GetExchangeRateByCurrencyAndDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).GetExchangeRateByCurrencyAndDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_GetExchangeRateByCurrencyAndDate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).GetExchangeRateByCurrencyAndDate(ctx, req.(*ExchangeRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExchangeService_ServiceDesc is the grpc.ServiceDesc for ExchangeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExchangeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "walbety.exchange.ExchangeService",
	HandlerType: (*ExchangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExchangeRateByCurrencyAndDate",
			Handler:    _ExchangeService_GetExchangeRateByCurrencyAndDate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tools/protos/exchange.proto",
}