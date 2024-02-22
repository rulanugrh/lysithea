package repository

import (
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/util"
	"gorm.io/gorm"
)

type OrderRepository interface {
	AddToCart(req domain.Cart) (*domain.Cart, error)
	FindID(uuid string) (*domain.Order, error)
	Cart(userID uint, page int, perPage int) (*[]domain.Cart, error)
	CountOrderByUserID(userID uint) (int64, error)
	Buy(req domain.Order) (*domain.Order, error)
	Checkout(id uint) (*domain.Order, error)
	Update(uuid string, status string, model domain.Order) error
	Pay(uuid string, userID uint) (*domain.Order, error)
	History(userID uint, page int, perPage int) (*[]domain.Order, error)
	CountCartByUserID(userID uint) (int64, error)
}

type order struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &order{
		db: db,
	}
}

func (o *order) AddToCart(req domain.Cart) (*domain.Cart, error) {
	err := o.db.Create(&req).Error
	if err != nil {
		return nil, web.InternalServerError("cannot create order")
	}

	err = o.db.Preload("Product").Preload("User").Find(&req).Error
	if err != nil {
		return nil, web.InternalServerError(err.Error())
	}

	return &req, nil
}

func (o *order) FindID(uuid string) (*domain.Order, error) {
	var req domain.Order
	err := o.db.Where("uuid = ?", uuid).Preload("Product").Preload("User").Find(&req).Error
	if err != nil {
		return nil, web.InternalServerError("data not found")
	}

	return &req, nil
}

func (o *order) Cart(userID uint, page int, perPage int) (*[]domain.Cart, error) {
	var req []domain.Cart
	err := o.db.Scopes(util.PaginationSet(page, perPage)).Where("user_id = ?", userID).Preload("Product").Preload("User").Find(&req).Error
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

func (o *order) Buy(req domain.Order) (*domain.Order, error) {
	req.UUID = util.GenerateUUID()
	req.Status = "not paid"
	err := o.db.Create(&req).Error
	if err != nil {
		return nil, web.InternalServerError("cant buy product")
	}

	err = o.db.Preload("Product").Preload("User").Find(&req).Error
	if err != nil {
		return nil, web.InternalServerError(err.Error())
	}

	return &req, nil
}

func (o *order) Update(uuid string, status string, model domain.Order) error {
	err := o.db.Model(&model).Where("uuid = ?", uuid).Preload("Product").Preload("User").Update("status", status).Error
	if err != nil {
		return web.InternalServerError(err.Error())
	}

	return nil
}

func (o *order) Pay(uuid string, userID uint) (*domain.Order, error) {
	var update domain.Order
	err := o.db.Where("uuid = ?", uuid).Where("user_id = ?", userID).Preload("Product").Preload("User").Find(&update).Error
	if err != nil {
		return nil, web.InternalServerError(err.Error())
	}

	return &update, nil
}

func (o *order) Checkout(id uint) (*domain.Order, error) {
	var cart domain.Cart
	err := o.db.Where("id = ?", id).Find(&cart).Error
	if err != nil {
		return nil, web.StatusNotFound(err.Error())
	}

	var req domain.Order
	req.ProductID = cart.ProductID
	req.UUID = util.GenerateUUID()
	req.UserID = cart.UserID
	req.Total = cart.Total

	err = o.db.Create(&req).Error
	if err != nil {
		return nil, web.InternalServerError(err.Error())
	}

	errFind := o.db.Preload("Product").Preload("User").Find(&req).Error
	if errFind != nil {
		return nil, web.InternalServerError(errFind.Error())
	}

	return &req, nil
}

func (o *order) History(userID uint, page int, perPage int) (*[]domain.Order, error) {
	var req []domain.Order
	err := o.db.Scopes(util.PaginationSet(page, perPage)).Where("user_id = ?", userID).Where("status = ?", "paid").Preload("Product").Preload("User").Find(&req).Error
	if err != nil {
		return nil, web.InternalServerError("data not found")
	}

	return &req, nil
}

func (o *order) CountCartByUserID(userID uint) (int64, error) {
	var model domain.Cart
	var count int64

	err := o.db.Model(&model).Where("user_id = ?", userID).Count(&count).Error
	if err != nil {
		return 0, web.InternalServerError("cannot count data by this user id")
	}

	return count, nil
}
