package tests

import (
	"testing"
	"time"

	"github.com/rulanugrh/lysithea/internal/entity/domain"
	repomocks "github.com/rulanugrh/lysithea/internal/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ProductTest struct {
	suite.Suite
	repo repomocks.ProductRepository
}

func NewProductTest() *ProductTest {
	return &ProductTest{}
}

func (product *ProductTest) TestProductCreate() {
	productResult := func(input domain.ProductRequest) *domain.Product {
		output := &domain.Product{}
		output.Name = input.Name
		output.CategoryID = input.CategoryID
		output.Price = input.Price
		output.Owner = input.Owner
		output.Description = input.Description
		output.Stock = input.Stock
		output.Discount = input.Discount
		output.ExpireAt = input.ExpireAt
		return output
	}

	product.repo.On("Create", mock.MatchedBy(func(input domain.ProductRequest) bool {
		return input.Name != "" && input.Owner != ""
	})).Return(productResult, nil)

	data, err := product.repo.Create(domain.ProductRequest{
		Name:        "Macbook 15 Pro",
		Owner:       "John Doe",
		CategoryID:  1,
		Discount:    10,
		Stock:       20,
		Price:       10000000,
		ExpireAt:    time.Now(),
		Description: "New series Laptop Macbook",
	})

	product.Nil(err)
	product.Equal("Macbook 15 Pro", data.Name)
	product.Equal(10, data.Discount)
	product.Equal(uint(1), data.CategoryID)
	product.Equal(20, data.Stock)
	product.Equal("New series Laptop Macbook", data.Description)
}

func (product *ProductTest) TestProductFindByID() {
	productResult := func(id uint) *domain.Product {
		output := &domain.Product{}
		output.ID = id
		return output
	}

	product.repo.On("FindID", mock.MatchedBy(func(id uint) bool {
		return id > 0
	})).Return(productResult, nil)

	data, err := product.repo.FindID(1)

	product.Nil(err)
	product.Equal(uint(1), data.ID)
}

func (product *ProductTest) TestProductFindAll() {
	productResult := func(page int, perPage int) *[]domain.Product {
		output := &[]domain.Product{}
		return output
	}

	product.repo.On("FindAll", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(productResult, nil)

	data, err := product.repo.FindAll(10, 10)

	product.Nil(err)
	product.Equal(&[]domain.Product{}, data)
}

func (product *ProductTest) TestProductUpdate() {
	productResult := func(id uint, input domain.Product) *domain.Product {
		output := &domain.Product{}
		output.ID = id
		output.Stock = input.Stock
		return output
	}

	product.repo.On("Update", mock.AnythingOfType("uint"), mock.AnythingOfType("domain.Product")).Return(productResult, nil)

	data, err := product.repo.Update(1, domain.Product{
		Stock: 10,
	})

	product.Nil(err)
	product.Equal(uint(1), data.ID)
	product.Equal(10, data.Stock)
}

func TestProduct(t *testing.T) {
	suite.Run(t, NewProductTest())
}
