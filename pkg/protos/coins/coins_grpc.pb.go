// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/protos/coins.proto

package coins

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

// CoinServiceClient is the client API for CoinService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoinServiceClient interface {
	CreateCoin(ctx context.Context, in *CoinRequest, opts ...grpc.CallOption) (*CoinResponse, error)
}

type coinServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoinServiceClient(cc grpc.ClientConnInterface) CoinServiceClient {
	return &coinServiceClient{cc}
}

func (c *coinServiceClient) CreateCoin(ctx context.Context, in *CoinRequest, opts ...grpc.CallOption) (*CoinResponse, error) {
	out := new(CoinResponse)
	err := c.cc.Invoke(ctx, "/CoinService/CreateCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoinServiceServer is the server API for CoinService service.
// All implementations must embed UnimplementedCoinServiceServer
// for forward compatibility
type CoinServiceServer interface {
	CreateCoin(context.Context, *CoinRequest) (*CoinResponse, error)
	mustEmbedUnimplementedCoinServiceServer()
}

// UnimplementedCoinServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCoinServiceServer struct {
}

func (UnimplementedCoinServiceServer) CreateCoin(context.Context, *CoinRequest) (*CoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoin not implemented")
}
func (UnimplementedCoinServiceServer) mustEmbedUnimplementedCoinServiceServer() {}

// UnsafeCoinServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoinServiceServer will
// result in compilation errors.
type UnsafeCoinServiceServer interface {
	mustEmbedUnimplementedCoinServiceServer()
}

func RegisterCoinServiceServer(s grpc.ServiceRegistrar, srv CoinServiceServer) {
	s.RegisterService(&CoinService_ServiceDesc, srv)
}

func _CoinService_CreateCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinServiceServer).CreateCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoinService/CreateCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinServiceServer).CreateCoin(ctx, req.(*CoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoinService_ServiceDesc is the grpc.ServiceDesc for CoinService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoinService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CoinService",
	HandlerType: (*CoinServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoin",
			Handler:    _CoinService_CreateCoin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protos/coins.proto",
}
