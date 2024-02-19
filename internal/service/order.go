package service

import (
	"math"

	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/repository"
)

type OrderService interface {
	Create(req domain.Order) (*web.OrderResponse, error)
	FindID(uuid string) (*web.OrderResponse, error)
	FindByUserID(userID uint, page int, perPage int) (*web.Pagination, error)
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

func (o *order) FindByUserID(userID uint, page int, perPage int) (*web.Pagination, error) {
	data, err := o.repo.FindByUserID(userID, page, perPage)
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
