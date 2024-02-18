package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Product     []Product `json:"product" gorm:"many2many:list_product"`
}
