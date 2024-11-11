package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/rcv911/service_template/pkg/cat_admin_v1"
	"google.golang.org/grpc"
)

type CatAdminServiceServer struct {
	cat_admin_v1.UnimplementedCatAdminServiceServer
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

func grpcServerStart(stopCh chan os.Signal) {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	cat_admin_v1.RegisterCatAdminServiceServer(grpcServer, &CatAdminServiceServer{})

	log.Println("Server is running at :8081")

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	<-stopCh
	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()
	log.Println("gRPC server gracefully stopped")
}
