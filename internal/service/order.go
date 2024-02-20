package service

import (
	"fmt"
	"math"

	"github.com/rulanugrh/lysithea/internal/config"
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/middleware"
	"github.com/rulanugrh/lysithea/internal/repository"
)

type OrderService interface {
	AddToCart(req domain.Cart) (*web.Cart, error)
	FindID(uuid string) (*web.OrderResponse, error)
	History(userID uint, page int, perPage int) (*web.Pagination, error)
	Buy(req domain.Order) (*web.BuyResponse, error)
	Pay(uuid string, userID uint) (*web.PaymentResponse, error)
	Checkout(id uint) (*web.BuyResponse, error)
	Cart(userID uint, page int, perPage int) (*web.Pagination, error)
}

type order struct {
	repo     repository.OrderRepository
	validate middleware.ValidationInterface
	conf     config.App
}

func NewOrderService(repo repository.OrderRepository, validate middleware.ValidationInterface, conf config.App) OrderService {
	return &order{
		repo:     repo,
		validate: validate,
		conf:     conf,
	}
}

func (o *order) AddToCart(req domain.Cart) (*web.Cart, error) {
	err := o.validate.Validate(req)
	if err != nil {
		return nil, o.validate.ValidationMessage(err)
	}

	data, err := o.repo.AddToCart(req)
	if err != nil {
		return nil, web.InternalServerError(err.Error())
	}

	response := web.Cart{
		ProductName:        data.Product.Name,
		ProductOwner:       data.Product.Owner,
		ProductPrice:       data.Product.Price,
		ProductDiscount:    data.Product.Discount,
		ProductDescription: data.Product.Description,
		ProductExpire:      data.Product.ExpireAt,
		TotalBeli:          data.Total,
		URLCheckout:        fmt.Sprintf("%s:%s/api/v1/order/checkout/%d", o.conf.Server.URL, o.conf.Server.Port, data.ID),
	}

	return &response, nil
}

func (o *order) FindID(uuid string) (*web.OrderResponse, error) {
	data, err := o.repo.FindID(uuid)
	if err != nil {
		return nil, web.StatusNotFound(err.Error())
	}

	response := web.OrderResponse{
		UUID:        data.UUID,
		Username:    data.User.Name,
		Name:        data.Product.Name,
		Price:       data.Product.Price,
		Discount:    data.Product.Discount,
		Description: data.Product.Description,
		ExpireAt:    data.Product.CreatedAt,
	}

	return &response, nil
}

func (o *order) History(userID uint, page int, perPage int) (*web.Pagination, error) {
	data, err := o.repo.History(userID, page, perPage)
	if err != nil {
		return nil, web.StatusNotFound("Order not found by this user id")
	}

	var response []web.OrderResponse
	for _, or := range *data {
		result := web.OrderResponse{
			UUID:        or.UUID,
			Username:    or.User.Name,
			Name:        or.Product.Name,
			Price:       or.Product.Price,
			Discount:    or.Product.Discount,
			Description: or.Product.Description,
			ExpireAt:    or.Product.CreatedAt,
			Status:      or.Status,
		}

		response = append(response, result)
	}

	count, err := o.repo.CountOrderByUserID(userID)
	if err != nil {
		return nil, web.InternalServerError("cant count data")
	}

	total := float64(count) / float64(perPage)
	result := web.Pagination{
		Metadata: web.Metadata{
			Page:      page,
			PerPage:   perPage,
			TotalData: int64(count),
			TotalPage: int64(math.Ceil(total)),
		},
		Data: response,
	}

	return &result, nil
}

func (o *order) Buy(req domain.Order) (*web.BuyResponse, error) {
	err := o.validate.Validate(req)
	if err != nil {
		return nil, o.validate.ValidationMessage(err)
	}

	data, err := o.repo.Buy(req)
	if err != nil {
		return nil, web.InternalServerError(err.Error())
	}

	response := web.BuyResponse{
		ProductName:        data.Product.Name,
		ProductPrice:       data.Product.Price,
		ProductDiscount:    data.Product.Discount,
		ProductDescription: data.Product.Description,
		TotalHarga:         ((data.Product.Price - (data.Product.Discount * data.Product.Price)) * req.Total),
		PayURL:             fmt.Sprintf("%s:%s/api/v1/order/pay/%d/%s", o.conf.Server.URL, o.conf.Server.Port, data.UserID, data.UUID),
	}

	return &response, nil
}

func (o *order) Pay(uuid string, userID uint) (*web.PaymentResponse, error) {
	data, err := o.repo.Pay(uuid, userID)
	if err != nil {
		return nil, web.InternalServerError(err.Error())
	}

	err = o.repo.Update(uuid, "paid", *data)
	if err != nil {
		return nil, web.StatusBadRequest(err.Error())
	}

	response := web.PaymentResponse{
		OrderUUID:          data.UUID,
		ProductName:        data.Product.Name,
		ProductPrice:       data.Product.Price,
		ProductOwner:       data.Product.Owner,
		ProductExpire:      data.Product.ExpireAt,
		ProductDiscount:    data.Product.Discount,
		ProductDescription: data.Product.Description,
		TotalHarga:         ((data.Product.Price - (data.Product.Discount * data.Product.Price)) * data.Total),
		Status:             "Paid",
	}

	return &response, nil
}

func (o *order) Checkout(id uint) (*web.BuyResponse, error) {
	data, err := o.repo.Checkout(id)
	if err != nil {
		return nil, web.StatusNotFound(err.Error())
	}

	response := web.BuyResponse{
		ProductName:        data.Product.Name,
		ProductPrice:       data.Product.Price,
		ProductDiscount:    data.Product.Discount,
		ProductDescription: data.Product.Description,
		TotalHarga:         ((data.Product.Price - (data.Product.Discount * data.Product.Price)) * data.Total),
		PayURL:             fmt.Sprintf("%s:%s/api/v1/order/pay/%s", o.conf.Server.URL, o.conf.Server.Port, data.UUID),
	}

	return &response, nil
}

func (o *order) Cart(userID uint, page int, perPage int) (*web.Pagination, error) {
	result, err := o.repo.Cart(userID, page, perPage)
	if err != nil {
		return nil, web.StatusNotFound(err.Error())
	}

	var response []web.Cart
	for _, data := range *result {
		cart := web.Cart{
			ProductName:        data.Product.Name,
			ProductOwner:       data.Product.Owner,
			ProductPrice:       data.Product.Price,
			ProductDiscount:    data.Product.Discount,
			ProductDescription: data.Product.Description,
			ProductExpire:      data.Product.ExpireAt,
			TotalBeli:          data.Total,
			URLCheckout:        fmt.Sprintf("%s:%s/api/v1/order/checkout/%d", o.conf.Server.URL, o.conf.Server.Port, data.ID),
		}

		response = append(response, cart)
	}

	count, err := o.repo.CountCartByUserID(userID)
	if err != nil {
		return nil, web.InternalServerError("cant count data")
	}

	total := float64(count) / float64(perPage)
	res := web.Pagination{
		Metadata: web.Metadata{
			Page:      page,
			PerPage:   perPage,
			TotalData: int64(count),
			TotalPage: int64(math.Ceil(total)),
		},
		Data: response,
	}

	return &res, nil

}
