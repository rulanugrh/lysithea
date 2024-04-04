package tests

import (
	"testing"

	"github.com/rulanugrh/lysithea/internal/entity/domain"
	repomocks "github.com/rulanugrh/lysithea/internal/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type OrderTest struct {
	suite.Suite
	repo repomocks.OrderRepository
}

func NewOrderTest() *OrderTest {
	return &OrderTest{}
}

func (order *OrderTest) TestOrderAddToCart() {
	orderResult := func(input domain.Cart) *domain.Cart {
		output := domain.Cart{}
		output.UserID = input.UserID
		output.ProductID = input.ProductID
		output.Total = input.Total
		return &output
	}

	order.repo.On("AddToCart", mock.MatchedBy(func(input domain.Cart) bool {
		return input.ProductID > 0 && input.Total > 0
	})).Return(orderResult, nil)

	data, err := order.repo.AddToCart(domain.Cart{
		ProductID: 1,
		UserID:    1,
		Total:     10,
	})

	order.Nil(err)
	order.Equal(uint(1), data.ProductID)
	order.Equal(uint(1), data.UserID)
	order.Equal(10, data.Total)
}

func (order *OrderTest) TestOrderBuy() {
	orderResult := func(input domain.Order) *domain.Order {
		output := domain.Order{}
		output.UserID = input.UserID
		output.ProductID = input.ProductID
		output.Total = input.Total
		output.UUID = input.UUID
		output.Status = input.Status
		return &output
	}

	order.repo.On("Buy", mock.MatchedBy(func(input domain.Order) bool {
		return input.ProductID > 0 && input.Total > 0
	})).Return(orderResult, nil)

	data, err := order.repo.Buy(domain.Order{
		ProductID: 1,
		UserID:    1,
		Total:     10,
		UUID:      "72300e52-d62c-42d0-8c53-426da447c798",
		Status:    "Process",
	})

	order.Nil(err)
	order.Equal(uint(1), data.ProductID)
	order.Equal(uint(1), data.UserID)
	order.Equal(10, data.Total)
	order.Equal("72300e52-d62c-42d0-8c53-426da447c798", data.UUID)
	order.Equal("Process", data.Status)
}

func TestOrder(t *testing.T) {
	suite.Run(t, NewOrderTest())
}
