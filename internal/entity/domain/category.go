package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Product     []Product `json:"product" gorm:"many2many:list_product"`
}
