// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: usage/v1/usage.proto

package v1

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

// UsageServiceClient is the client API for UsageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsageServiceClient interface {
	// ListBilledUsage retrieves all usage for the specified attributionId
	ListBilledUsage(ctx context.Context, in *ListBilledUsageRequest, opts ...grpc.CallOption) (*ListBilledUsageResponse, error)
}

type usageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsageServiceClient(cc grpc.ClientConnInterface) UsageServiceClient {
	return &usageServiceClient{cc}
}

func (c *usageServiceClient) ListBilledUsage(ctx context.Context, in *ListBilledUsageRequest, opts ...grpc.CallOption) (*ListBilledUsageResponse, error) {
	out := new(ListBilledUsageResponse)
	err := c.cc.Invoke(ctx, "/usage.v1.UsageService/ListBilledUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsageServiceServer is the server API for UsageService service.
// All implementations must embed UnimplementedUsageServiceServer
// for forward compatibility
type UsageServiceServer interface {
	// ListBilledUsage retrieves all usage for the specified attributionId
	ListBilledUsage(context.Context, *ListBilledUsageRequest) (*ListBilledUsageResponse, error)
	mustEmbedUnimplementedUsageServiceServer()
}

// UnimplementedUsageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUsageServiceServer struct {
}

func (UnimplementedUsageServiceServer) ListBilledUsage(context.Context, *ListBilledUsageRequest) (*ListBilledUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBilledUsage not implemented")
}
func (UnimplementedUsageServiceServer) mustEmbedUnimplementedUsageServiceServer() {}

// UnsafeUsageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsageServiceServer will
// result in compilation errors.
type UnsafeUsageServiceServer interface {
	mustEmbedUnimplementedUsageServiceServer()
}

func RegisterUsageServiceServer(s grpc.ServiceRegistrar, srv UsageServiceServer) {
	s.RegisterService(&UsageService_ServiceDesc, srv)
}

func _UsageService_ListBilledUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBilledUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageServiceServer).ListBilledUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usage.v1.UsageService/ListBilledUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageServiceServer).ListBilledUsage(ctx, req.(*ListBilledUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsageService_ServiceDesc is the grpc.ServiceDesc for UsageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usage.v1.UsageService",
	HandlerType: (*UsageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBilledUsage",
			Handler:    _UsageService_ListBilledUsage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usage/v1/usage.proto",
}
