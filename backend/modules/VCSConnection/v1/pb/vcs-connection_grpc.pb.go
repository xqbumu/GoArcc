// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// VCSConnectionsClient is the client API for VCSConnections service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VCSConnectionsClient interface {
	//ListAllSupportedVCSProviders returns all external providers list supported by alfred
	ListAllSupportedVCSProviders(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListAllSupportedVCSProvidersResponse, error)
	//Authorize provide a URL of external OAuth application
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
	//Callback called by external OAuth application
	Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*AccountVCSConnection, error)
	//ListVCSConnection list all connected service of user's account
	ListVCSConnection(ctx context.Context, in *ListVCSConnectionRequest, opts ...grpc.CallOption) (*ListVCSConnectionResponse, error)
	//GetVCSConnection get a vcs_connection based on id
	GetVCSConnection(ctx context.Context, in *GetVCSConnectionRequest, opts ...grpc.CallOption) (*AccountVCSConnection, error)
}

type vCSConnectionsClient struct {
	cc grpc.ClientConnInterface
}

func NewVCSConnectionsClient(cc grpc.ClientConnInterface) VCSConnectionsClient {
	return &vCSConnectionsClient{cc}
}

func (c *vCSConnectionsClient) ListAllSupportedVCSProviders(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListAllSupportedVCSProvidersResponse, error) {
	out := new(ListAllSupportedVCSProvidersResponse)
	err := c.cc.Invoke(ctx, "/alfred.vcs_connection.v1.VCSConnections/ListAllSupportedVCSProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vCSConnectionsClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, "/alfred.vcs_connection.v1.VCSConnections/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vCSConnectionsClient) Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*AccountVCSConnection, error) {
	out := new(AccountVCSConnection)
	err := c.cc.Invoke(ctx, "/alfred.vcs_connection.v1.VCSConnections/Callback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vCSConnectionsClient) ListVCSConnection(ctx context.Context, in *ListVCSConnectionRequest, opts ...grpc.CallOption) (*ListVCSConnectionResponse, error) {
	out := new(ListVCSConnectionResponse)
	err := c.cc.Invoke(ctx, "/alfred.vcs_connection.v1.VCSConnections/ListVCSConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vCSConnectionsClient) GetVCSConnection(ctx context.Context, in *GetVCSConnectionRequest, opts ...grpc.CallOption) (*AccountVCSConnection, error) {
	out := new(AccountVCSConnection)
	err := c.cc.Invoke(ctx, "/alfred.vcs_connection.v1.VCSConnections/GetVCSConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VCSConnectionsServer is the server API for VCSConnections service.
// All implementations should embed UnimplementedVCSConnectionsServer
// for forward compatibility
type VCSConnectionsServer interface {
	//ListAllSupportedVCSProviders returns all external providers list supported by alfred
	ListAllSupportedVCSProviders(context.Context, *empty.Empty) (*ListAllSupportedVCSProvidersResponse, error)
	//Authorize provide a URL of external OAuth application
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
	//Callback called by external OAuth application
	Callback(context.Context, *CallbackRequest) (*AccountVCSConnection, error)
	//ListVCSConnection list all connected service of user's account
	ListVCSConnection(context.Context, *ListVCSConnectionRequest) (*ListVCSConnectionResponse, error)
	//GetVCSConnection get a vcs_connection based on id
	GetVCSConnection(context.Context, *GetVCSConnectionRequest) (*AccountVCSConnection, error)
}

// UnimplementedVCSConnectionsServer should be embedded to have forward compatible implementations.
type UnimplementedVCSConnectionsServer struct {
}

func (UnimplementedVCSConnectionsServer) ListAllSupportedVCSProviders(context.Context, *empty.Empty) (*ListAllSupportedVCSProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllSupportedVCSProviders not implemented")
}
func (UnimplementedVCSConnectionsServer) Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedVCSConnectionsServer) Callback(context.Context, *CallbackRequest) (*AccountVCSConnection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Callback not implemented")
}
func (UnimplementedVCSConnectionsServer) ListVCSConnection(context.Context, *ListVCSConnectionRequest) (*ListVCSConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVCSConnection not implemented")
}
func (UnimplementedVCSConnectionsServer) GetVCSConnection(context.Context, *GetVCSConnectionRequest) (*AccountVCSConnection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVCSConnection not implemented")
}

// UnsafeVCSConnectionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VCSConnectionsServer will
// result in compilation errors.
type UnsafeVCSConnectionsServer interface {
	mustEmbedUnimplementedVCSConnectionsServer()
}

func RegisterVCSConnectionsServer(s *grpc.Server, srv VCSConnectionsServer) {
	s.RegisterService(&_VCSConnections_serviceDesc, srv)
}

func _VCSConnections_ListAllSupportedVCSProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VCSConnectionsServer).ListAllSupportedVCSProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.vcs_connection.v1.VCSConnections/ListAllSupportedVCSProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VCSConnectionsServer).ListAllSupportedVCSProviders(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VCSConnections_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VCSConnectionsServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.vcs_connection.v1.VCSConnections/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VCSConnectionsServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VCSConnections_Callback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VCSConnectionsServer).Callback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.vcs_connection.v1.VCSConnections/Callback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VCSConnectionsServer).Callback(ctx, req.(*CallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VCSConnections_ListVCSConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVCSConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VCSConnectionsServer).ListVCSConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.vcs_connection.v1.VCSConnections/ListVCSConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VCSConnectionsServer).ListVCSConnection(ctx, req.(*ListVCSConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VCSConnections_GetVCSConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVCSConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VCSConnectionsServer).GetVCSConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.vcs_connection.v1.VCSConnections/GetVCSConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VCSConnectionsServer).GetVCSConnection(ctx, req.(*GetVCSConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VCSConnections_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alfred.vcs_connection.v1.VCSConnections",
	HandlerType: (*VCSConnectionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAllSupportedVCSProviders",
			Handler:    _VCSConnections_ListAllSupportedVCSProviders_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _VCSConnections_Authorize_Handler,
		},
		{
			MethodName: "Callback",
			Handler:    _VCSConnections_Callback_Handler,
		},
		{
			MethodName: "ListVCSConnection",
			Handler:    _VCSConnections_ListVCSConnection_Handler,
		},
		{
			MethodName: "GetVCSConnection",
			Handler:    _VCSConnections_GetVCSConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vcs-connection.proto",
}
