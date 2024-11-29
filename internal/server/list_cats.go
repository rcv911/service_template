package server

import (
	"context"

	"github.com/rcv911/service_template/pkg/cat_admin_v1"
)

func (s *CatAdminServiceServer) ListCats(ctx context.Context, req *cat_admin_v1.ListCatsRequest) (*cat_admin_v1.ListCatsResponse, error) {
	cats := s.repo.ListCats()

	var resp []*cat_admin_v1.CatResponse

	for _, cat := range cats {
		resp = append(resp, &cat_admin_v1.CatResponse{
			Id:    cat.ID,
			Name:  cat.Name,
			Age:   int32(cat.Age),
			Color: cat.Color,
		})
	}

	return &cat_admin_v1.ListCatsResponse{
		Cats:       resp,
		TotalCount: int32(len(cats)),
	}, nil
}
