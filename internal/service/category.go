package service

import (
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/middleware"
	"github.com/rulanugrh/lysithea/internal/repository"
)

type CategoryService interface {
	Create(req domain.Category) (*web.CategoryCreated, error)
}

type category struct {
	repo     repository.CategoryRepository
	validate middleware.ValidationInterface
}

func NewCategoryService(repo repository.CategoryRepository, validate middleware.ValidationInterface) CategoryService {
	return &category{
		repo:     repo,
		validate: validate,
	}
}

func (c *category) Create(req domain.Category) (*web.CategoryCreated, error) {
	err := c.validate.Validate(req)
	if err != nil {
		return nil, c.validate.ValidationMessage(err)
	}

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
