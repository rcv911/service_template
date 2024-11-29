package server

import (
	"context"

	"github.com/rcv911/service_template/pkg/cat_admin_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetCat получение кота
func (s *CatAdminServiceServer) GetCat(ctx context.Context, req *cat_admin_v1.GetCatRequest) (*cat_admin_v1.CatResponse, error) {
	cat, err := s.repo.GetCat(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed get cat: %v", err)
	}

	return &cat_admin_v1.CatResponse{
		Id:    cat.ID,
		Name:  cat.Name,
		Age:   int32(cat.Age),
		Color: cat.Color,
	}, nil
}
