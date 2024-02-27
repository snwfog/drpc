// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: sesamestreet.proto

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

// CookieMonsterClient is the client API for CookieMonster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CookieMonsterClient interface {
	EatCookie(ctx context.Context, in *CookiePool, opts ...grpc.CallOption) (*Crumbs, error)
}

type cookieMonsterClient struct {
	cc grpc.ClientConnInterface
}

func NewCookieMonsterClient(cc grpc.ClientConnInterface) CookieMonsterClient {
	return &cookieMonsterClient{cc}
}

func (c *cookieMonsterClient) EatCookie(ctx context.Context, in *CookiePool, opts ...grpc.CallOption) (*Crumbs, error) {
	out := new(Crumbs)
	err := c.cc.Invoke(ctx, "/sesamestreet.CookieMonster/EatCookie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CookieMonsterServer is the server API for CookieMonster service.
// All implementations must embed UnimplementedCookieMonsterServer
// for forward compatibility
type CookieMonsterServer interface {
	EatCookie(context.Context, *CookiePool) (*Crumbs, error)
	mustEmbedUnimplementedCookieMonsterServer()
}

// UnimplementedCookieMonsterServer must be embedded to have forward compatible implementations.
type UnimplementedCookieMonsterServer struct {
}

func (UnimplementedCookieMonsterServer) EatCookie(context.Context, *CookiePool) (*Crumbs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EatCookie not implemented")
}
func (UnimplementedCookieMonsterServer) mustEmbedUnimplementedCookieMonsterServer() {}

// UnsafeCookieMonsterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CookieMonsterServer will
// result in compilation errors.
type UnsafeCookieMonsterServer interface {
	mustEmbedUnimplementedCookieMonsterServer()
}

func RegisterCookieMonsterServer(s grpc.ServiceRegistrar, srv CookieMonsterServer) {
	s.RegisterService(&CookieMonster_ServiceDesc, srv)
}

func _CookieMonster_EatCookie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CookiePool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CookieMonsterServer).EatCookie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sesamestreet.CookieMonster/EatCookie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CookieMonsterServer).EatCookie(ctx, req.(*CookiePool))
	}
	return interceptor(ctx, in, info, handler)
}

// CookieMonster_ServiceDesc is the grpc.ServiceDesc for CookieMonster service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CookieMonster_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sesamestreet.CookieMonster",
	HandlerType: (*CookieMonsterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EatCookie",
			Handler:    _CookieMonster_EatCookie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sesamestreet.proto",
}
