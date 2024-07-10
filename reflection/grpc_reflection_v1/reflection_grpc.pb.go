// Copyright 2016 The gRPC Authors
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

// Service exported by server reflection.  A more complete description of how
// server reflection works can be found at
// https://github.com/grpc/grpc/blob/master/doc/server-reflection.md
//
// The canonical version of this proto can be found at
// https://github.com/grpc/grpc-proto/blob/master/grpc/reflection/v1/reflection.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: grpc/reflection/v1/reflection.proto

package grpc_reflection_v1

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
	ServerReflection_ServerReflectionInfo_FullMethodName = "/grpc.reflection.v1.ServerReflection/ServerReflectionInfo"
)

// ServerReflectionClient is the client API for ServerReflection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerReflectionClient interface {
	// The reflection service is structured as a bidirectional stream, ensuring
	// all related requests go to a single server.
	ServerReflectionInfo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ServerReflectionRequest, ServerReflectionResponse], error)
}

type serverReflectionClient struct {
	cc grpc.ClientConnInterface
}

func NewServerReflectionClient(cc grpc.ClientConnInterface) ServerReflectionClient {
	return &serverReflectionClient{cc}
}

func (c *serverReflectionClient) ServerReflectionInfo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ServerReflectionRequest, ServerReflectionResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ServerReflection_ServiceDesc.Streams[0], ServerReflection_ServerReflectionInfo_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ServerReflectionRequest, ServerReflectionResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ServerReflection_ServerReflectionInfoClient = grpc.BidiStreamingClient[ServerReflectionRequest, ServerReflectionResponse]

// ServerReflectionServer is the server API for ServerReflection service.
// All implementations should embed UnimplementedServerReflectionServer
// for forward compatibility
type ServerReflectionServer interface {
	// The reflection service is structured as a bidirectional stream, ensuring
	// all related requests go to a single server.
	ServerReflectionInfo(grpc.BidiStreamingServer[ServerReflectionRequest, ServerReflectionResponse]) error
}

// UnimplementedServerReflectionServer should be embedded to have forward compatible implementations.
type UnimplementedServerReflectionServer struct {
}

func (UnimplementedServerReflectionServer) ServerReflectionInfo(grpc.BidiStreamingServer[ServerReflectionRequest, ServerReflectionResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ServerReflectionInfo not implemented")
}

// UnsafeServerReflectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerReflectionServer will
// result in compilation errors.
type UnsafeServerReflectionServer interface {
	mustEmbedUnimplementedServerReflectionServer()
}

func RegisterServerReflectionServer(s grpc.ServiceRegistrar, srv ServerReflectionServer) {
	s.RegisterService(&ServerReflection_ServiceDesc, srv)
}

func _ServerReflection_ServerReflectionInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServerReflectionServer).ServerReflectionInfo(&grpc.GenericServerStream[ServerReflectionRequest, ServerReflectionResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ServerReflection_ServerReflectionInfoServer = grpc.BidiStreamingServer[ServerReflectionRequest, ServerReflectionResponse]

// ServerReflection_ServiceDesc is the grpc.ServiceDesc for ServerReflection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerReflection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.reflection.v1.ServerReflection",
	HandlerType: (*ServerReflectionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerReflectionInfo",
			Handler:       _ServerReflection_ServerReflectionInfo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/reflection/v1/reflection.proto",
}
