package service

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/middleware"
	"github.com/rulanugrh/lysithea/internal/repository"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type CategoryService interface {
	Create(req domain.Category) (*web.CategoryCreated, error)
	GetCategoryBySearch(page int, perPage int, search string, buf bytes.Buffer) (map[string]interface{}, error)
}

type category struct {
	repo     repository.CategoryRepository
	validate middleware.ValidationInterface
	es       *elasticsearch.Client
	trace    trace.Tracer
	meter    metric.MeterProvider
}

func NewCategoryService(repo repository.CategoryRepository, validate middleware.ValidationInterface, es *elasticsearch.Client, trace trace.Tracer, meter metric.MeterProvider) CategoryService {
	return &category{
		repo:     repo,
		validate: validate,
		es:       es,
		trace:    trace,
		meter:    meter,
	}
}

func (c *category) Create(req domain.Category) (*web.CategoryCreated, error) {
	ctx, span := c.trace.Start(context.Background(), "create-category")
	defer span.End()

	meter := c.meter.Meter("meter-create-category")
	counter, _ := meter.Float64Counter("metric_called")
	counter.Add(ctx, 1)

	err := c.validate.Validate(req)
	if err != nil {
		return nil, c.validate.ValidationMessage(err)
	}

	data, err := c.repo.Create(req)
	if err != nil {
		return nil, web.InternalServerError("cannot create category")
	}

	response := web.CategoryCreated{
		Name:        data.Name,
		Description: data.Description,
	}

	return &response, nil
}

func (c *category) GetCategoryBySearch(page int, perPage int, search string, buf bytes.Buffer) (map[string]interface{}, error) {
	ctx, span := c.trace.Start(context.Background(), "get-category-by-search")
	defer span.End()

	meter := c.meter.Meter("meter--category")
	counter, _ := meter.Float64Counter("metric_called")
	counter.Add(ctx, 1)

	query := map[string]interface{}{
		"collapse": map[string]interface{}{
			"field": "name.keyword",
		},
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
							"product_name": map[string]interface{}{
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
	res, err := c.es.Search(
		c.es.Search.WithBody(&buf),
		c.es.Search.WithPretty(),
		c.es.Search.WithIndex("categories"),
		c.es.Search.WithTrackTotalHits(false),
		c.es.Search.WithFilterPath("hits.hits._source.name",
			"hits.hits._source.category_id"),
	)

	if err != nil {
		log.Fatalf("Error getting response, %v", err.Error())
	}

	if err = json.NewDecoder(res.Body).Decode(&resposeBind); err != nil {
		log.Fatalf("Error Decoder, %v", err.Error())
	}

	return resposeBind, nil

}
