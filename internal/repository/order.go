package repository

import (
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/util"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(req domain.Order) (*domain.Order, error)
	FindID(uuid string) (*domain.Order, error)
	FindByUserID(userID uint, page int, perPage int) (*[]domain.Order, error)
	CountOrderByUserID(userID uint) (int64, error)
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
	req.UUID = util.GenerateUUID()
	err := o.db.Create(&req).Error
	if err != nil {
		return nil, web.InternalServerError("cannot create order")
	}

	return &req, nil
}

func (o *order) FindID(uuid string) (*domain.Order, error) {
	var req domain.Order
	err := o.db.Where("uuid = ?", uuid).Find(&req).Error
	if err != nil {
		return nil, web.InternalServerError("data not found")
	}

	return &req, nil
}

func (o *order) FindByUserID(userID uint, page int, perPage int) (*[]domain.Order, error) {
	var req []domain.Order
	err := o.db.Scopes(util.PaginationSet(page, perPage)).Where("user_id = ?", userID).Find(&req).Error
	if err != nil {
		return nil, web.InternalServerError("data not found")
	}

	return &req, nil
}

func (o *order) CountOrderByUserID(userID uint) (int64, error) {
	var model domain.Order
	var count int64

	err := o.db.Model(&model).Where("user_id = ?", userID).Count(&count).Error
	if err != nil {
		return 0, web.InternalServerError("cannot count data by this user id")
	}

	return count, nil
}
