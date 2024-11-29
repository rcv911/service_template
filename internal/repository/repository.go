package repository

import (
	"github.com/pkg/errors"
	"github.com/rcv911/service_template/internal/model"
	"github.com/rs/zerolog"
)

type Repository struct {
	logger  zerolog.Logger
	storage map[int64]model.Cat
}

func New(logger zerolog.Logger) *Repository {
	return &Repository{
		logger: logger,
	}
}

func (r *Repository) GetCat(id int64) (*model.Cat, error) {
	cat, ok := r.storage[id]
	if !ok {
		return nil, errors.New("cat not found")
	}

	return &cat, nil
}

func (r *Repository) CreateCat(age int64, name, color string) error {
	newCat := model.Cat{
		Age:   age,
		Name:  name,
		Color: color,
	}

	r.storage[newCat.ID] = newCat

	return nil
}

func (r *Repository) ListCats() []model.Cat {
	var cats []model.Cat

	for _, cat := range r.storage {
		cats = append(cats, cat)
	}

	return cats
}
