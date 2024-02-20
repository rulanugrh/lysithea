package service

import (
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/repository"
)

type CategoryService interface {
	Create(req domain.Category) (*web.CategoryCreated, error)
}

type category struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &category{
		repo: repo,
	}
}

func (c *category) Create(req domain.Category) (*web.CategoryCreated, error) {
	data, err := c.repo.Create(req)
	if err != nil {
		return nil, web.InternalServerError("cannot create category")
	}

	response := web.CategoryCreated{
		Name:        data.Name,
		Description: data.Description,
	}

	return &response, nil
}
