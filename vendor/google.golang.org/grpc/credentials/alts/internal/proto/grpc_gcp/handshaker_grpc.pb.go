// Copyright 2018 The gRPC Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The canonical version of this proto can be found at
// https://github.com/grpc/grpc-proto/blob/master/grpc/gcp/handshaker.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: grpc/gcp/handshaker.proto

package grpc_gcp

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
	HandshakerService_DoHandshake_FullMethodName = "/grpc.gcp.HandshakerService/DoHandshake"
)

// HandshakerServiceClient is the client API for HandshakerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HandshakerServiceClient interface {
	// Handshaker service accepts a stream of handshaker request, returning a
	// stream of handshaker response. Client is expected to send exactly one
	// message with either client_start or server_start followed by one or more
	// messages with next. Each time client sends a request, the handshaker
	// service expects to respond. Client does not have to wait for service's
	// response before sending next request.
	DoHandshake(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[HandshakerReq, HandshakerResp], error)
}

type handshakerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHandshakerServiceClient(cc grpc.ClientConnInterface) HandshakerServiceClient {
	return &handshakerServiceClient{cc}
}

func (c *handshakerServiceClient) DoHandshake(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[HandshakerReq, HandshakerResp], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &HandshakerService_ServiceDesc.Streams[0], HandshakerService_DoHandshake_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HandshakerReq, HandshakerResp]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HandshakerService_DoHandshakeClient = grpc.BidiStreamingClient[HandshakerReq, HandshakerResp]

// HandshakerServiceServer is the server API for HandshakerService service.
// All implementations must embed UnimplementedHandshakerServiceServer
// for forward compatibility.
type HandshakerServiceServer interface {
	// Handshaker service accepts a stream of handshaker request, returning a
	// stream of handshaker response. Client is expected to send exactly one
	// message with either client_start or server_start followed by one or more
	// messages with next. Each time client sends a request, the handshaker
	// service expects to respond. Client does not have to wait for service's
	// response before sending next request.
	DoHandshake(grpc.BidiStreamingServer[HandshakerReq, HandshakerResp]) error
	mustEmbedUnimplementedHandshakerServiceServer()
}

// UnimplementedHandshakerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHandshakerServiceServer struct{}

func (UnimplementedHandshakerServiceServer) DoHandshake(grpc.BidiStreamingServer[HandshakerReq, HandshakerResp]) error {
	return status.Errorf(codes.Unimplemented, "method DoHandshake not implemented")
}
func (UnimplementedHandshakerServiceServer) mustEmbedUnimplementedHandshakerServiceServer() {}
func (UnimplementedHandshakerServiceServer) testEmbeddedByValue()                           {}

// UnsafeHandshakerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HandshakerServiceServer will
// result in compilation errors.
type UnsafeHandshakerServiceServer interface {
	mustEmbedUnimplementedHandshakerServiceServer()
}

func RegisterHandshakerServiceServer(s grpc.ServiceRegistrar, srv HandshakerServiceServer) {
	// If the following call panics, it indicates UnimplementedHandshakerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HandshakerService_ServiceDesc, srv)
}

func _HandshakerService_DoHandshake_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HandshakerServiceServer).DoHandshake(&grpc.GenericServerStream[HandshakerReq, HandshakerResp]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HandshakerService_DoHandshakeServer = grpc.BidiStreamingServer[HandshakerReq, HandshakerResp]

// HandshakerService_ServiceDesc is the grpc.ServiceDesc for HandshakerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HandshakerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gcp.HandshakerService",
	HandlerType: (*HandshakerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DoHandshake",
			Handler:       _HandshakerService_DoHandshake_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/gcp/handshaker.proto",
}
