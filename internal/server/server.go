package server

import (
	"github.com/rcv911/service_template/internal/model"
	"github.com/rcv911/service_template/pkg/cat_admin_v1"
	"github.com/rcv911/service_template/pkg/cat_v1"
)

type Repository interface {
	ListCats() []model.Cat
	CreateCat(age int64, name, color string) error
	GetCat(id int64) (*model.Cat, error)
}

type CatAdminServiceServer struct {
	cat_admin_v1.UnimplementedCatAdminServiceServer
	repo Repository
}

type CatServiceServer struct {
	cat_v1.UnimplementedCatServiceServer
	repo Repository
}

func NewCatServiceServer(repo Repository) *CatServiceServer {
	return &CatServiceServer{repo: repo}
}

func NewCatAdminServiceServer(repo Repository) *CatAdminServiceServer {
	return &CatAdminServiceServer{repo: repo}
}
