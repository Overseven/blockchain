// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package node

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

// NoderClient is the client API for Noder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoderClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectReply, error)
	GetListOfNodes(ctx context.Context, in *ListOfNodesRequest, opts ...grpc.CallOption) (*ListOfNodesReply, error)
	AddTransaction(ctx context.Context, in *AddTransactionRequest, opts ...grpc.CallOption) (*AddTransactionReply, error)
	PushBlock(ctx context.Context, in *PushBlockRequest, opts ...grpc.CallOption) (*PushBlockReply, error)
	GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (*GetBlocksReply, error)
}

type noderClient struct {
	cc grpc.ClientConnInterface
}

func NewNoderClient(cc grpc.ClientConnInterface) NoderClient {
	return &noderClient{cc}
}

func (c *noderClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectReply, error) {
	out := new(ConnectReply)
	err := c.cc.Invoke(ctx, "/pnode.Noder/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noderClient) GetListOfNodes(ctx context.Context, in *ListOfNodesRequest, opts ...grpc.CallOption) (*ListOfNodesReply, error) {
	out := new(ListOfNodesReply)
	err := c.cc.Invoke(ctx, "/pnode.Noder/GetListOfNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noderClient) AddTransaction(ctx context.Context, in *AddTransactionRequest, opts ...grpc.CallOption) (*AddTransactionReply, error) {
	out := new(AddTransactionReply)
	err := c.cc.Invoke(ctx, "/pnode.Noder/AddTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noderClient) PushBlock(ctx context.Context, in *PushBlockRequest, opts ...grpc.CallOption) (*PushBlockReply, error) {
	out := new(PushBlockReply)
	err := c.cc.Invoke(ctx, "/pnode.Noder/PushBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noderClient) GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (*GetBlocksReply, error) {
	out := new(GetBlocksReply)
	err := c.cc.Invoke(ctx, "/pnode.Noder/GetBlocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoderServer is the server API for Noder service.
// All implementations must embed UnimplementedNoderServer
// for forward compatibility
type NoderServer interface {
	Connect(context.Context, *ConnectRequest) (*ConnectReply, error)
	GetListOfNodes(context.Context, *ListOfNodesRequest) (*ListOfNodesReply, error)
	AddTransaction(context.Context, *AddTransactionRequest) (*AddTransactionReply, error)
	PushBlock(context.Context, *PushBlockRequest) (*PushBlockReply, error)
	GetBlocks(context.Context, *GetBlocksRequest) (*GetBlocksReply, error)
	mustEmbedUnimplementedNoderServer()
}

// UnimplementedNoderServer must be embedded to have forward compatible implementations.
type UnimplementedNoderServer struct {
}

func (UnimplementedNoderServer) Connect(context.Context, *ConnectRequest) (*ConnectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedNoderServer) GetListOfNodes(context.Context, *ListOfNodesRequest) (*ListOfNodesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListOfNodes not implemented")
}
func (UnimplementedNoderServer) AddTransaction(context.Context, *AddTransactionRequest) (*AddTransactionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTransaction not implemented")
}
func (UnimplementedNoderServer) PushBlock(context.Context, *PushBlockRequest) (*PushBlockReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushBlock not implemented")
}
func (UnimplementedNoderServer) GetBlocks(context.Context, *GetBlocksRequest) (*GetBlocksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlocks not implemented")
}
func (UnimplementedNoderServer) mustEmbedUnimplementedNoderServer() {}

// UnsafeNoderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoderServer will
// result in compilation errors.
type UnsafeNoderServer interface {
	mustEmbedUnimplementedNoderServer()
}

func RegisterNoderServer(s grpc.ServiceRegistrar, srv NoderServer) {
	s.RegisterService(&Noder_ServiceDesc, srv)
}

func _Noder_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoderServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pnode.Noder/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoderServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Noder_GetListOfNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOfNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoderServer).GetListOfNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pnode.Noder/GetListOfNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoderServer).GetListOfNodes(ctx, req.(*ListOfNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Noder_AddTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoderServer).AddTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pnode.Noder/AddTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoderServer).AddTransaction(ctx, req.(*AddTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Noder_PushBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoderServer).PushBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pnode.Noder/PushBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoderServer).PushBlock(ctx, req.(*PushBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Noder_GetBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoderServer).GetBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pnode.Noder/GetBlocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoderServer).GetBlocks(ctx, req.(*GetBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Noder_ServiceDesc is the grpc.ServiceDesc for Noder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Noder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pnode.Noder",
	HandlerType: (*NoderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _Noder_Connect_Handler,
		},
		{
			MethodName: "GetListOfNodes",
			Handler:    _Noder_GetListOfNodes_Handler,
		},
		{
			MethodName: "AddTransaction",
			Handler:    _Noder_AddTransaction_Handler,
		},
		{
			MethodName: "PushBlock",
			Handler:    _Noder_PushBlock_Handler,
		},
		{
			MethodName: "GetBlocks",
			Handler:    _Noder_GetBlocks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol/node/node.proto",
}
