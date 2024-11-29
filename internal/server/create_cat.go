package server

import (
	"context"

	"github.com/rcv911/service_template/pkg/cat_admin_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CatAdminServiceServer) CreateCat(ctx context.Context, req *cat_admin_v1.CreateCatRequest) (*cat_admin_v1.CatResponse, error) {
	err := s.repo.CreateCat(int64(req.GetAge()), req.GetName(), req.GetColor())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &cat_admin_v1.CatResponse{}, nil
}
