package server

import (
	"context"

	"github.com/rcv911/service_template/pkg/cat_admin_v1"
	"github.com/rcv911/service_template/pkg/cat_v1"
	"google.golang.org/grpc"
)

type CatAdminServiceServer struct {
	cat_admin_v1.UnimplementedCatAdminServiceServer
}

type CatServiceServer struct {
	cat_v1.UnimplementedCatServiceServer
}

func NewGRPCSServer() *grpc.Server {
	grpcServer := grpc.NewServer()

	cat_admin_v1.RegisterCatAdminServiceServer(grpcServer, &CatAdminServiceServer{})
	cat_v1.RegisterCatServiceServer(grpcServer, &CatServiceServer{})

	return grpcServer
}

// GetCat получение кота
func (s *CatAdminServiceServer) GetCat(ctx context.Context, req *cat_admin_v1.GetCatRequest) (*cat_admin_v1.CatResponse, error) {
	return &cat_admin_v1.CatResponse{
		Id:    req.Id,
		Name:  "Барсик",
		Age:   3,
		Color: "Черный",
	}, nil
}
