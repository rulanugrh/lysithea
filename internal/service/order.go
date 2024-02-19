package service

import (
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/repository"
)

type OrderService interface {
	Create(req domain.Order) (*web.OrderResponse, error)
	FindID(uuid string) (*web.OrderResponse, error)
}

type order struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &order{
		repo: repo,
	}
}

func (o *order) Create(req domain.Order) (*web.OrderResponse, error) {
	data, err := o.repo.Create(req)
	if err != nil {
		return nil, web.InternalServerError(err.Error())
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
