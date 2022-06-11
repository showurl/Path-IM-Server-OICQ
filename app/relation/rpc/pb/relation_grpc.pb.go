// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: relation.proto

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

// RelationServiceClient is the client API for RelationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationServiceClient interface {
	AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*AddFriendResp, error)
	DelFriend(ctx context.Context, in *DelFriendReq, opts ...grpc.CallOption) (*DelFriendResp, error)
	IsFriend(ctx context.Context, in *IsFriendReq, opts ...grpc.CallOption) (*IsFriendResp, error)
	GetFriendModel(ctx context.Context, in *GetFriendModelReq, opts ...grpc.CallOption) (*GetFriendModelResp, error)
	UpdateFriendModel(ctx context.Context, in *UpdateFriendModelReq, opts ...grpc.CallOption) (*UpdateFriendModelResp, error)
	GetFriendIds(ctx context.Context, in *GetFriendIdsReq, opts ...grpc.CallOption) (*GetFriendIdsResp, error)
	AddBlacklist(ctx context.Context, in *AddBlacklistReq, opts ...grpc.CallOption) (*AddBlacklistResp, error)
	DelBlacklist(ctx context.Context, in *DelBlacklistReq, opts ...grpc.CallOption) (*DelBlacklistResp, error)
	IsBlacklist(ctx context.Context, in *IsBlacklistReq, opts ...grpc.CallOption) (*IsBlacklistResp, error)
	GetBlacklist(ctx context.Context, in *GetBlacklistReq, opts ...grpc.CallOption) (*GetBlacklistResp, error)
	GetBlacklistModel(ctx context.Context, in *GetBlacklistModelReq, opts ...grpc.CallOption) (*GetBlacklistModelResp, error)
}

type relationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationServiceClient(cc grpc.ClientConnInterface) RelationServiceClient {
	return &relationServiceClient{cc}
}

func (c *relationServiceClient) AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*AddFriendResp, error) {
	out := new(AddFriendResp)
	err := c.cc.Invoke(ctx, "/relation.relationService/AddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) DelFriend(ctx context.Context, in *DelFriendReq, opts ...grpc.CallOption) (*DelFriendResp, error) {
	out := new(DelFriendResp)
	err := c.cc.Invoke(ctx, "/relation.relationService/DelFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) IsFriend(ctx context.Context, in *IsFriendReq, opts ...grpc.CallOption) (*IsFriendResp, error) {
	out := new(IsFriendResp)
	err := c.cc.Invoke(ctx, "/relation.relationService/IsFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetFriendModel(ctx context.Context, in *GetFriendModelReq, opts ...grpc.CallOption) (*GetFriendModelResp, error) {
	out := new(GetFriendModelResp)
	err := c.cc.Invoke(ctx, "/relation.relationService/GetFriendModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) UpdateFriendModel(ctx context.Context, in *UpdateFriendModelReq, opts ...grpc.CallOption) (*UpdateFriendModelResp, error) {
	out := new(UpdateFriendModelResp)
	err := c.cc.Invoke(ctx, "/relation.relationService/UpdateFriendModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetFriendIds(ctx context.Context, in *GetFriendIdsReq, opts ...grpc.CallOption) (*GetFriendIdsResp, error) {
	out := new(GetFriendIdsResp)
	err := c.cc.Invoke(ctx, "/relation.relationService/GetFriendIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) AddBlacklist(ctx context.Context, in *AddBlacklistReq, opts ...grpc.CallOption) (*AddBlacklistResp, error) {
	out := new(AddBlacklistResp)
	err := c.cc.Invoke(ctx, "/relation.relationService/AddBlacklist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) DelBlacklist(ctx context.Context, in *DelBlacklistReq, opts ...grpc.CallOption) (*DelBlacklistResp, error) {
	out := new(DelBlacklistResp)
	err := c.cc.Invoke(ctx, "/relation.relationService/DelBlacklist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) IsBlacklist(ctx context.Context, in *IsBlacklistReq, opts ...grpc.CallOption) (*IsBlacklistResp, error) {
	out := new(IsBlacklistResp)
	err := c.cc.Invoke(ctx, "/relation.relationService/IsBlacklist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetBlacklist(ctx context.Context, in *GetBlacklistReq, opts ...grpc.CallOption) (*GetBlacklistResp, error) {
	out := new(GetBlacklistResp)
	err := c.cc.Invoke(ctx, "/relation.relationService/GetBlacklist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetBlacklistModel(ctx context.Context, in *GetBlacklistModelReq, opts ...grpc.CallOption) (*GetBlacklistModelResp, error) {
	out := new(GetBlacklistModelResp)
	err := c.cc.Invoke(ctx, "/relation.relationService/GetBlacklistModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationServiceServer is the server API for RelationService service.
// All implementations must embed UnimplementedRelationServiceServer
// for forward compatibility
type RelationServiceServer interface {
	AddFriend(context.Context, *AddFriendReq) (*AddFriendResp, error)
	DelFriend(context.Context, *DelFriendReq) (*DelFriendResp, error)
	IsFriend(context.Context, *IsFriendReq) (*IsFriendResp, error)
	GetFriendModel(context.Context, *GetFriendModelReq) (*GetFriendModelResp, error)
	UpdateFriendModel(context.Context, *UpdateFriendModelReq) (*UpdateFriendModelResp, error)
	GetFriendIds(context.Context, *GetFriendIdsReq) (*GetFriendIdsResp, error)
	AddBlacklist(context.Context, *AddBlacklistReq) (*AddBlacklistResp, error)
	DelBlacklist(context.Context, *DelBlacklistReq) (*DelBlacklistResp, error)
	IsBlacklist(context.Context, *IsBlacklistReq) (*IsBlacklistResp, error)
	GetBlacklist(context.Context, *GetBlacklistReq) (*GetBlacklistResp, error)
	GetBlacklistModel(context.Context, *GetBlacklistModelReq) (*GetBlacklistModelResp, error)
	mustEmbedUnimplementedRelationServiceServer()
}

// UnimplementedRelationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRelationServiceServer struct {
}

func (UnimplementedRelationServiceServer) AddFriend(context.Context, *AddFriendReq) (*AddFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedRelationServiceServer) DelFriend(context.Context, *DelFriendReq) (*DelFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelFriend not implemented")
}
func (UnimplementedRelationServiceServer) IsFriend(context.Context, *IsFriendReq) (*IsFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFriend not implemented")
}
func (UnimplementedRelationServiceServer) GetFriendModel(context.Context, *GetFriendModelReq) (*GetFriendModelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendModel not implemented")
}
func (UnimplementedRelationServiceServer) UpdateFriendModel(context.Context, *UpdateFriendModelReq) (*UpdateFriendModelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFriendModel not implemented")
}
func (UnimplementedRelationServiceServer) GetFriendIds(context.Context, *GetFriendIdsReq) (*GetFriendIdsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendIds not implemented")
}
func (UnimplementedRelationServiceServer) AddBlacklist(context.Context, *AddBlacklistReq) (*AddBlacklistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlacklist not implemented")
}
func (UnimplementedRelationServiceServer) DelBlacklist(context.Context, *DelBlacklistReq) (*DelBlacklistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelBlacklist not implemented")
}
func (UnimplementedRelationServiceServer) IsBlacklist(context.Context, *IsBlacklistReq) (*IsBlacklistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsBlacklist not implemented")
}
func (UnimplementedRelationServiceServer) GetBlacklist(context.Context, *GetBlacklistReq) (*GetBlacklistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlacklist not implemented")
}
func (UnimplementedRelationServiceServer) GetBlacklistModel(context.Context, *GetBlacklistModelReq) (*GetBlacklistModelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlacklistModel not implemented")
}
func (UnimplementedRelationServiceServer) mustEmbedUnimplementedRelationServiceServer() {}

// UnsafeRelationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationServiceServer will
// result in compilation errors.
type UnsafeRelationServiceServer interface {
	mustEmbedUnimplementedRelationServiceServer()
}

func RegisterRelationServiceServer(s grpc.ServiceRegistrar, srv RelationServiceServer) {
	s.RegisterService(&RelationService_ServiceDesc, srv)
}

func _RelationService_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.relationService/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).AddFriend(ctx, req.(*AddFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_DelFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).DelFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.relationService/DelFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).DelFriend(ctx, req.(*DelFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_IsFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).IsFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.relationService/IsFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).IsFriend(ctx, req.(*IsFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetFriendModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendModelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetFriendModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.relationService/GetFriendModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetFriendModel(ctx, req.(*GetFriendModelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_UpdateFriendModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFriendModelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).UpdateFriendModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.relationService/UpdateFriendModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).UpdateFriendModel(ctx, req.(*UpdateFriendModelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetFriendIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetFriendIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.relationService/GetFriendIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetFriendIds(ctx, req.(*GetFriendIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_AddBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).AddBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.relationService/AddBlacklist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).AddBlacklist(ctx, req.(*AddBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_DelBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).DelBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.relationService/DelBlacklist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).DelBlacklist(ctx, req.(*DelBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_IsBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).IsBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.relationService/IsBlacklist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).IsBlacklist(ctx, req.(*IsBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.relationService/GetBlacklist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetBlacklist(ctx, req.(*GetBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetBlacklistModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlacklistModelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetBlacklistModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.relationService/GetBlacklistModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetBlacklistModel(ctx, req.(*GetBlacklistModelReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationService_ServiceDesc is the grpc.ServiceDesc for RelationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "relation.relationService",
	HandlerType: (*RelationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFriend",
			Handler:    _RelationService_AddFriend_Handler,
		},
		{
			MethodName: "DelFriend",
			Handler:    _RelationService_DelFriend_Handler,
		},
		{
			MethodName: "IsFriend",
			Handler:    _RelationService_IsFriend_Handler,
		},
		{
			MethodName: "GetFriendModel",
			Handler:    _RelationService_GetFriendModel_Handler,
		},
		{
			MethodName: "UpdateFriendModel",
			Handler:    _RelationService_UpdateFriendModel_Handler,
		},
		{
			MethodName: "GetFriendIds",
			Handler:    _RelationService_GetFriendIds_Handler,
		},
		{
			MethodName: "AddBlacklist",
			Handler:    _RelationService_AddBlacklist_Handler,
		},
		{
			MethodName: "DelBlacklist",
			Handler:    _RelationService_DelBlacklist_Handler,
		},
		{
			MethodName: "IsBlacklist",
			Handler:    _RelationService_IsBlacklist_Handler,
		},
		{
			MethodName: "GetBlacklist",
			Handler:    _RelationService_GetBlacklist_Handler,
		},
		{
			MethodName: "GetBlacklistModel",
			Handler:    _RelationService_GetBlacklistModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relation.proto",
}