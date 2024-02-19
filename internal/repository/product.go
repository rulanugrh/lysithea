package repository

import (
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/util"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(req domain.ProductRequest) (*domain.Product, error)
	FindID(id uint) (*domain.Product, error)
	FindAll(page int, perPage int) (*[]domain.Product, error)
	Update(id uint, req domain.Product) (*domain.Product, error)
	CountProduct() (int64, error)
	CountProductByCategoryID(categoryID uint) (int64, error)
}

type product struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &product{
		db: db,
	}
}

func (p *product) Create(req domain.ProductRequest) (*domain.Product, error) {
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
		return nil, web.InternalServerError("cannot create product")
	}

	return &product, nil
}

func (p *product) FindID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := p.db.Where("id = ?", id).Find(&product).Error
	if err != nil {
		return nil, web.InternalServerError("cannot create product")
	}

	return &product, nil
}

func (p *product) FindAll(page int, perPage int) (*[]domain.Product, error) {
	var product []domain.Product
	err := p.db.Scopes(util.PaginationSet(page, perPage)).Find(&product).Error

	if err != nil {
		return nil, web.StatusNotFound("data not found")
	}

	return &product, nil
}

func (p *product) FindByCategoryID(page int, perPage int, categoryID uint) (*[]domain.Product, error) {
	var product []domain.Product
	err := p.db.Scopes(util.PaginationSet(page, perPage)).Where("category_id = ?", categoryID).Find(&product).Error

	if err != nil {
		return nil, web.StatusNotFound("data not found")
	}

	return &product, nil
}

func (p *product) Update(id uint, req domain.Product) (*domain.Product, error) {
	var update domain.Product
	err := p.db.Model(&req).Where("id = ?", id).Updates(&update).Error

	if err != nil {
		return nil, web.InternalServerError("cannot create product")
	}

	return &update, nil
}

func (p *product) CountProduct() (int64, error) {
	var model domain.Product
	var count int64

	err := p.db.Model(&model).Count(&count).Error
	if err != nil {
		return 0, web.InternalServerError("cannot count product data")
	}

	return count, nil
}

func (p *product) CountProductByCategoryID(categoryID uint) (int64, error) {
	var model domain.Product
	var count int64

	err := p.db.Model(&model).Where("category_id = ?", categoryID).Count(&count).Error
	if err != nil {
		return 0, web.InternalServerError("cannot count product data")
	}

	return count, nil
}
