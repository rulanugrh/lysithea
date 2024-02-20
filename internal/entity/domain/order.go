package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UUID      string  `json:"uuid"`
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id" validate:"required"`
	Status    string  `json:"status"`
	Total     int     `json:"total_beli" validate:"required"`
	User      User    `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`
}

type Cart struct {
	gorm.Model
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id" validate:"required"`
	Total     int     `json:"total_beli" validate:"required"`
	User      User    `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`
}
