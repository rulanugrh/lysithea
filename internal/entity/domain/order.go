package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UUID      string  `json:"uuid"`
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id"`
	User      User    `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`
}
