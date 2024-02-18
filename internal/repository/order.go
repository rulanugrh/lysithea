package repository

import (
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(req domain.Order) (*domain.Order, error)
	FindID(uuid string) (*domain.Order, error)
	FindByUserID(userID uint) (*[]domain.Order, error)
}

type order struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &order{
		db: db,
	}
}

func (o *order) Create(req domain.Order) (*domain.Order, error) {
	err := o.db.Create(&req).Error
	if err != nil {
		return nil, web.NewInternalServerErrorResponse("cannot create order")
	}

	return &req, nil
}

func (o *order) FindID(uuid string) (*domain.Order, error) {
	var req domain.Order
	err := o.db.Where("uuid = ?", uuid).Find(&req).Error
	if err != nil {
		return nil, web.NewInternalServerErrorResponse("data not found")
	}

	return &req, nil
}

func (o *order) FindByUserID(userID uint) (*[]domain.Order, error) {
	var req []domain.Order
	err := o.db.Where("user_id = ?", userID).Find(&req).Error
	if err != nil {
		return nil, web.NewInternalServerErrorResponse("data not found")
	}

	return &req, nil
}
