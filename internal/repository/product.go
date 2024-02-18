package repository

import (
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(req domain.ProductRequest) (error, *domain.Product)
	FindID(id uint) (error, *domain.Product)
	FindAll() (error, *[]domain.Product)
	Update(id uint, req domain.Product) (error, *domain.Product)
}

type product struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &product{
		db: db,
	}
}

func (p *product) Create(req domain.ProductRequest) (error, *domain.Product) {
	product := domain.Product{
		Name:        req.Name,
		Discount:    req.Discount,
		ExpireAt:    req.ExpireAt,
		Owner:       req.Owner,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
		Description: req.Description,
	}

	err := p.db.Create(&product).Error
	if err != nil {
		return web.NewInternalServerErrorResponse("cannot create product"), nil
	}

	return web.NewCreatedResponse("success create product", req), &product
}

func (p *product) FindID(id uint) (error, *domain.Product) {
	var product domain.Product
	err := p.db.Where("id = ?", id).Find(&product).Error
	if err != nil {
		return web.NewInternalServerErrorResponse("cannot create product"), nil
	}

	return nil, &product
}

func (p *product) FindAll() (error, *[]domain.Product) {
	var product []domain.Product
	err := p.db.Find(&product).Error

	if err != nil {
		return web.NewInternalServerErrorResponse("cannot create product"), nil
	}

	return nil, &product
}

func (p *product) Update(id uint, req domain.Product) (error, *domain.Product) {
	var update domain.Product
	err := p.db.Model(&req).Where("id = ?", id).Updates(&update).Error

	if err != nil {
		return web.NewInternalServerErrorResponse("cannot create product"), nil
	}

	return nil, &update
}
