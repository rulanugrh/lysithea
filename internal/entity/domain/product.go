package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Owner       string    `json:"owner"`
	Discount    int       `json:"discount"`
	Description string    `json:"description"`
	ExpireAt    time.Time `json:"expire_at"`
	Stock       int       `json:"stock"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
}

type ProductRequest struct {
	Name        string    `json:"name" validate:"required"`
	Price       int       `json:"price" validate:"required"`
	Owner       string    `json:"owner" validate:"required"`
	Discount    int       `json:"discount" validate:"required"`
	Description string    `json:"description" validate:"required"`
	ExpireAt    time.Time `json:"expire_at" validate:"required"`
	Stock       int       `json:"stock" validate:"required"`
	CategoryID  uint      `json:"category_id" validate:"required"`
}
