package server

import (
	"github.com/rcv911/service_template/pkg/cat_admin_v1"
	"github.com/rcv911/service_template/pkg/cat_v1"
	"google.golang.org/grpc"
)

func NewGRPCSServer(catService *CatServiceServer, catAdminService *CatAdminServiceServer) *grpc.Server {
	grpcServer := grpc.NewServer()

	cat_admin_v1.RegisterCatAdminServiceServer(grpcServer, catAdminService)
	cat_v1.RegisterCatServiceServer(grpcServer, catService)

	return grpcServer
}
