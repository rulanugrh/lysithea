package service

import (
	"bytes"
	"encoding/json"
	"log"
	"math"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/middleware"
	"github.com/rulanugrh/lysithea/internal/repository"
)

type ProductService interface {
	Create(req domain.ProductRequest) (*web.ProductResponse, error)
	FindID(id uint) (*web.ProductResponse, error)
	FindAll(page int, perPage int) (*web.Pagination, error)
	FindAllByCategoryID(categoryID uint, page int, perPage int) (*web.Pagination, error)
	GetProductBySearch(page int, perPage int, search string, buf bytes.Buffer) (map[string]interface{}, error)
}

type product struct {
	repo     repository.ProductRepository
	validate middleware.ValidationInterface
	es       *elasticsearch.Client
}

func NewProductService(repo repository.ProductRepository, validation middleware.ValidationInterface, es *elasticsearch.Client) ProductService {
	return &product{
		repo:     repo,
		validate: validation,
		es:       es,
	}
}

func (p *product) Create(req domain.ProductRequest) (*web.ProductResponse, error) {
	err := p.validate.Validate(req)
	if err != nil {
		return nil, p.validate.ValidationMessage(err)
	}

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

func (p *product) FindAllByCategoryID(categoryID uint, page int, perPage int) (*web.Pagination, error) {
	data, err := p.repo.FindByCategoryID(page, perPage, categoryID)
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

	count, err := p.repo.CountProductByCategoryID(categoryID)
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

func (p *product) GetProductBySearch(page int, perPage int, search string, buf bytes.Buffer) (map[string]interface{}, error) {
	query := map[string]interface{}{
		"from":  page,
		"limit": perPage,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []interface{}{
					map[string]interface{}{
						"match": map[string]interface{}{
							"name": map[string]interface{}{
								"query":         search,
								"fuzziness":     "AUTO",
								"prefix_length": 1,
							},
						},
					},
					map[string]interface{}{
						"match": map[string]interface{}{
							"expire_at": map[string]interface{}{
								"query":         search,
								"fuzziness":     "AUTO",
								"prefix_length": 1,
							},
						},
					},
				},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(&query); err != nil {
		log.Fatalf("Encoding error, %v", err.Error())
	}

	var resposeBind map[string]interface{}
	res, err := p.es.Search(
		p.es.Search.WithBody(&buf),
		p.es.Search.WithPretty(),
		p.es.Search.WithIndex("products"),
		p.es.Search.WithTrackTotalHits(false),
		p.es.Search.WithFilterPath("hits.hits._source.name",
			"hits.hits._source.expire_at", "hits.hits._source.id"),
	)

	if err != nil {
		log.Fatalf("Error getting response, %v", err.Error())
	}

	if err = json.NewDecoder(res.Body).Decode(&resposeBind); err != nil {
		log.Fatalf("Error Decoder, %v", err.Error())
	}

	return resposeBind, nil

}
