// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonServiceClient interface {
	GetPerson(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonResponse, error)
	WatchPersons(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (PersonService_WatchPersonsClient, error)
}

type personServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServiceClient(cc grpc.ClientConnInterface) PersonServiceClient {
	return &personServiceClient{cc}
}

func (c *personServiceClient) GetPerson(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonResponse, error) {
	out := new(PersonResponse)
	err := c.cc.Invoke(ctx, "/persongo.PersonService/GetPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) WatchPersons(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (PersonService_WatchPersonsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PersonService_ServiceDesc.Streams[0], "/persongo.PersonService/WatchPersons", opts...)
	if err != nil {
		return nil, err
	}
	x := &personServiceWatchPersonsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PersonService_WatchPersonsClient interface {
	Recv() (*PersonResponse, error)
	grpc.ClientStream
}

type personServiceWatchPersonsClient struct {
	grpc.ClientStream
}

func (x *personServiceWatchPersonsClient) Recv() (*PersonResponse, error) {
	m := new(PersonResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PersonServiceServer is the server API for PersonService service.
// All implementations must embed UnimplementedPersonServiceServer
// for forward compatibility
type PersonServiceServer interface {
	GetPerson(context.Context, *PersonRequest) (*PersonResponse, error)
	WatchPersons(*WatchRequest, PersonService_WatchPersonsServer) error
	mustEmbedUnimplementedPersonServiceServer()
}

// UnimplementedPersonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPersonServiceServer struct {
}

func (UnimplementedPersonServiceServer) GetPerson(context.Context, *PersonRequest) (*PersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}
func (UnimplementedPersonServiceServer) WatchPersons(*WatchRequest, PersonService_WatchPersonsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchPersons not implemented")
}
func (UnimplementedPersonServiceServer) mustEmbedUnimplementedPersonServiceServer() {}

// UnsafePersonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonServiceServer will
// result in compilation errors.
type UnsafePersonServiceServer interface {
	mustEmbedUnimplementedPersonServiceServer()
}

func RegisterPersonServiceServer(s grpc.ServiceRegistrar, srv PersonServiceServer) {
	s.RegisterService(&PersonService_ServiceDesc, srv)
}

func _PersonService_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/persongo.PersonService/GetPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).GetPerson(ctx, req.(*PersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_WatchPersons_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PersonServiceServer).WatchPersons(m, &personServiceWatchPersonsServer{stream})
}

type PersonService_WatchPersonsServer interface {
	Send(*PersonResponse) error
	grpc.ServerStream
}

type personServiceWatchPersonsServer struct {
	grpc.ServerStream
}

func (x *personServiceWatchPersonsServer) Send(m *PersonResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PersonService_ServiceDesc is the grpc.ServiceDesc for PersonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "persongo.PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPerson",
			Handler:    _PersonService_GetPerson_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchPersons",
			Handler:       _PersonService_WatchPersons_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "person.proto",
}