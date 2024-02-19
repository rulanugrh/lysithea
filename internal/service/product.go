package service

import (
	"math"

	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/repository"
)

type ProductService interface {
	Create(req domain.ProductRequest) (*web.ProductResponse, error)
	FindID(id uint) (*web.ProductResponse, error)
	FindAll(page int, perPage int) (*web.Pagination, error)
}

type product struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &product{
		repo: repo,
	}
}

func (p *product) Create(req domain.ProductRequest) (*web.ProductResponse, error) {
	data, err := p.repo.Create(req)
	if err != nil {
		return nil, web.StatusBadRequest(err.Error())
	}

	response := web.ProductResponse{
		ID:          data.ID,
		Name:        data.Name,
		Discount:    data.Discount,
		Price:       data.Price,
		Stock:       data.Stock,
		ExpireAt:    data.ExpireAt,
		Owner:       data.Owner,
		Category:    data.Category.Name,
		Description: data.Description,
	}

	return &response, nil
}

func (p *product) FindID(id uint) (*web.ProductResponse, error) {
	data, err := p.repo.FindID(id)
	if err != nil {
		return nil, web.StatusNotFound(err.Error())
	}

	response := web.ProductResponse{
		ID:          data.ID,
		Name:        data.Name,
		Discount:    data.Discount,
		Price:       data.Price,
		Stock:       data.Stock,
		ExpireAt:    data.ExpireAt,
		Owner:       data.Owner,
		Category:    data.Category.Name,
		Description: data.Description,
	}

	return &response, nil
}

func (p *product) FindAll(page int, perPage int) (*web.Pagination, error) {
	data, err := p.repo.FindAll(page, perPage)
	if err != nil {
		return nil, web.StatusNotFound(err.Error())
	}

	var response []web.ProductResponse
	for _, rp := range *data {
		result := web.ProductResponse{
			ID:          rp.ID,
			Name:        rp.Name,
			Discount:    rp.Discount,
			Price:       rp.Price,
			Stock:       rp.Stock,
			ExpireAt:    rp.ExpireAt,
			Owner:       rp.Owner,
			Category:    rp.Category.Name,
			Description: rp.Description,
		}

		response = append(response, result)
	}

	count, err := p.repo.CountProduct()
	if err != nil {
		return nil, web.InternalServerError(err.Error())
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
