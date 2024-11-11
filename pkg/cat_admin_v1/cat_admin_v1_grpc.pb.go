// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: cat_admin_v1/cat_admin_v1.proto

package cat_admin_v1

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
	CatAdminService_CreateCat_FullMethodName = "/cat_admin_v1.CatAdminService/CreateCat"
	CatAdminService_GetCat_FullMethodName    = "/cat_admin_v1.CatAdminService/GetCat"
	CatAdminService_UpdateCat_FullMethodName = "/cat_admin_v1.CatAdminService/UpdateCat"
	CatAdminService_DeleteCat_FullMethodName = "/cat_admin_v1.CatAdminService/DeleteCat"
	CatAdminService_ListCats_FullMethodName  = "/cat_admin_v1.CatAdminService/ListCats"
)

// CatAdminServiceClient is the client API for CatAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис CRUD операций для работы с котами
type CatAdminServiceClient interface {
	// Создать нового кота
	CreateCat(ctx context.Context, in *CreateCatRequest, opts ...grpc.CallOption) (*CatResponse, error)
	// Получить информацию о коте по его ID
	GetCat(ctx context.Context, in *GetCatRequest, opts ...grpc.CallOption) (*CatResponse, error)
	// Обновить информацию о коте
	UpdateCat(ctx context.Context, in *UpdateCatRequest, opts ...grpc.CallOption) (*CatResponse, error)
	// Удалить кота по ID
	DeleteCat(ctx context.Context, in *DeleteCatRequest, opts ...grpc.CallOption) (*DeleteCatResponse, error)
	// Получить список всех котов
	ListCats(ctx context.Context, in *ListCatsRequest, opts ...grpc.CallOption) (*ListCatsResponse, error)
}

type catAdminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatAdminServiceClient(cc grpc.ClientConnInterface) CatAdminServiceClient {
	return &catAdminServiceClient{cc}
}

func (c *catAdminServiceClient) CreateCat(ctx context.Context, in *CreateCatRequest, opts ...grpc.CallOption) (*CatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CatResponse)
	err := c.cc.Invoke(ctx, CatAdminService_CreateCat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catAdminServiceClient) GetCat(ctx context.Context, in *GetCatRequest, opts ...grpc.CallOption) (*CatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CatResponse)
	err := c.cc.Invoke(ctx, CatAdminService_GetCat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catAdminServiceClient) UpdateCat(ctx context.Context, in *UpdateCatRequest, opts ...grpc.CallOption) (*CatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CatResponse)
	err := c.cc.Invoke(ctx, CatAdminService_UpdateCat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catAdminServiceClient) DeleteCat(ctx context.Context, in *DeleteCatRequest, opts ...grpc.CallOption) (*DeleteCatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCatResponse)
	err := c.cc.Invoke(ctx, CatAdminService_DeleteCat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catAdminServiceClient) ListCats(ctx context.Context, in *ListCatsRequest, opts ...grpc.CallOption) (*ListCatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCatsResponse)
	err := c.cc.Invoke(ctx, CatAdminService_ListCats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatAdminServiceServer is the server API for CatAdminService service.
// All implementations must embed UnimplementedCatAdminServiceServer
// for forward compatibility.
//
// Сервис CRUD операций для работы с котами
type CatAdminServiceServer interface {
	// Создать нового кота
	CreateCat(context.Context, *CreateCatRequest) (*CatResponse, error)
	// Получить информацию о коте по его ID
	GetCat(context.Context, *GetCatRequest) (*CatResponse, error)
	// Обновить информацию о коте
	UpdateCat(context.Context, *UpdateCatRequest) (*CatResponse, error)
	// Удалить кота по ID
	DeleteCat(context.Context, *DeleteCatRequest) (*DeleteCatResponse, error)
	// Получить список всех котов
	ListCats(context.Context, *ListCatsRequest) (*ListCatsResponse, error)
	mustEmbedUnimplementedCatAdminServiceServer()
}

// UnimplementedCatAdminServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCatAdminServiceServer struct{}

func (UnimplementedCatAdminServiceServer) CreateCat(context.Context, *CreateCatRequest) (*CatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCat not implemented")
}
func (UnimplementedCatAdminServiceServer) GetCat(context.Context, *GetCatRequest) (*CatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCat not implemented")
}
func (UnimplementedCatAdminServiceServer) UpdateCat(context.Context, *UpdateCatRequest) (*CatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCat not implemented")
}
func (UnimplementedCatAdminServiceServer) DeleteCat(context.Context, *DeleteCatRequest) (*DeleteCatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCat not implemented")
}
func (UnimplementedCatAdminServiceServer) ListCats(context.Context, *ListCatsRequest) (*ListCatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCats not implemented")
}
func (UnimplementedCatAdminServiceServer) mustEmbedUnimplementedCatAdminServiceServer() {}
func (UnimplementedCatAdminServiceServer) testEmbeddedByValue()                         {}

// UnsafeCatAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatAdminServiceServer will
// result in compilation errors.
type UnsafeCatAdminServiceServer interface {
	mustEmbedUnimplementedCatAdminServiceServer()
}

func RegisterCatAdminServiceServer(s grpc.ServiceRegistrar, srv CatAdminServiceServer) {
	// If the following call pancis, it indicates UnimplementedCatAdminServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CatAdminService_ServiceDesc, srv)
}

func _CatAdminService_CreateCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatAdminServiceServer).CreateCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatAdminService_CreateCat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatAdminServiceServer).CreateCat(ctx, req.(*CreateCatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatAdminService_GetCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatAdminServiceServer).GetCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatAdminService_GetCat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatAdminServiceServer).GetCat(ctx, req.(*GetCatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatAdminService_UpdateCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatAdminServiceServer).UpdateCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatAdminService_UpdateCat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatAdminServiceServer).UpdateCat(ctx, req.(*UpdateCatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatAdminService_DeleteCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatAdminServiceServer).DeleteCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatAdminService_DeleteCat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatAdminServiceServer).DeleteCat(ctx, req.(*DeleteCatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatAdminService_ListCats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatAdminServiceServer).ListCats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatAdminService_ListCats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatAdminServiceServer).ListCats(ctx, req.(*ListCatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CatAdminService_ServiceDesc is the grpc.ServiceDesc for CatAdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatAdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cat_admin_v1.CatAdminService",
	HandlerType: (*CatAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCat",
			Handler:    _CatAdminService_CreateCat_Handler,
		},
		{
			MethodName: "GetCat",
			Handler:    _CatAdminService_GetCat_Handler,
		},
		{
			MethodName: "UpdateCat",
			Handler:    _CatAdminService_UpdateCat_Handler,
		},
		{
			MethodName: "DeleteCat",
			Handler:    _CatAdminService_DeleteCat_Handler,
		},
		{
			MethodName: "ListCats",
			Handler:    _CatAdminService_ListCats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cat_admin_v1/cat_admin_v1.proto",
}
