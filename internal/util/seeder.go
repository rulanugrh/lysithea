package util

import (
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	err := db.Create(&domain.Role{
		ID:          1,
		Name:        "Administrator",
		Description: "this is role administrator",
	}).Error
	if err != nil {
		return err
	}

	err = db.Create(&domain.Role{
		ID:          2,
		Name:        "User",
		Description: "this is role user",
	}).Error

	if err != nil {
		return err
	}

	return nil
}
