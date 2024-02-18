package util

import (
	"github.com/rulanugrh/lysithea/internal/config"
	"github.com/rulanugrh/lysithea/internal/entity/domain"
)

func Migrate() {
	_, db := config.NewConnection()
	db.AutoMigrate(&domain.Role{}, &domain.Role{}, &domain.Category{}, &domain.Order{}, &domain.User{})
}
