package domain

import "gorm.io/gorm"

type Role struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	User        []User `gorm:"many2many:list_users"`
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	NoHP     int    `json:"no_hp"`
	RoleID   uint   `json:"role_id"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleID;references:ID"`
}

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	NoHP     int    `json:"no_hp" validate:"required"`
	RoleID   uint   `json:"role_id"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
