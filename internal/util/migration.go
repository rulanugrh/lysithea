package util

import (
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&domain.Role{}, &domain.Role{}, &domain.Category{}, &domain.Cart{}, &domain.Order{}, &domain.User{})
	if err != nil {
		return web.InternalServerError("cannot migrate data")
	}

	return nil
}
