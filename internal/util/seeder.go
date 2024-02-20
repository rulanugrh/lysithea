package util

import (
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	role := []domain.Role{
		{
			ID:          1,
			Name:        "Administrator",
			Description: "this is role administrator",
		},
		{
			ID:          2,
			Name:        "Owner",
			Description: "this is role for owner product",
		},
		{
			ID:          3,
			Name:        "User",
			Description: "this is role user",
		},
	}

	for _, r := range role {
		err := db.Create(&r).Error
		if err != nil {
			return err
		}
	}

	return nil
}
