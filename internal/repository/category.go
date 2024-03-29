package repository

import (
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(req domain.Category) (*domain.Category, error)
}

type category struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &category{
		db: db,
	}
}

func (c *category) Create(req domain.Category) (*domain.Category, error) {
	err := c.db.Create(&req).Error
	if err != nil {
		return nil, web.InternalServerError("canot create category")
	}

	return &req, nil
}
