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

func (order *OrderTest) TestOrderCheckout() {
	orderResult := func(id uint) *domain.Order {
		output := domain.Order{}
		return &output
	}

	order.repo.On("Checkout", mock.MatchedBy(func(id uint) bool {
		return id > 0
	})).Return(orderResult, nil)

	data, err := order.repo.Checkout(1)

	order.Nil(err)
	order.Equal(&domain.Order{}, data)
}

func (order *OrderTest) TestOrderPay() {
	orderResult := func(uuid string, userID uint) *domain.Order {
		output := domain.Order{}
		output.UUID = uuid
		output.UserID = userID
		return &output
	}

	order.repo.On("Pay", mock.AnythingOfType("string"), mock.AnythingOfType("uint")).Return(orderResult, nil)

	data, err := order.repo.Pay("72300e52-d62c-42d0-8c53-426da447c798", 1)

	order.Nil(err)
	order.Equal(uint(1), data.UserID)
	order.Equal("72300e52-d62c-42d0-8c53-426da447c798", data.UUID)
}

func (order *OrderTest) TestOrderUpdate() {
	order.repo.On("Update", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("domain.Order")).Return(nil)

	err := order.repo.Update("72300e52-d62c-42d0-8c53-426da447c798", "success", domain.Order{
		Total: 10,
	})

	order.Nil(err)
}

func (order *OrderTest) TestOrderFindID() {
	orderResult := func(uuid string) *domain.Order {
		output := domain.Order{}
		output.UUID = uuid
		return &output
	}

	order.repo.On("FindID", mock.MatchedBy(func(uuid string) bool {
		return uuid != ""
	})).Return(orderResult, nil)

	data, err := order.repo.FindID("72300e52-d62c-42d0-8c53-426da447c798")

	order.Nil(err)
	order.Equal("72300e52-d62c-42d0-8c53-426da447c798", data.UUID)
}

func (order *OrderTest) TestOrderHistory() {
	orderHistory :=  func(userID uint, page int, perPage int) *[]domain.Order {
		output := &[]domain.Order{}
		for _, v := range *output {
			v.UserID = userID
		}

		return output
	}

	order.repo.On("History", mock.AnythingOfType("uint"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(orderHistory, nil)

	data, err := order.repo.History(1, 10, 10)

	for _, v := range *data {
		order.Equal(uint(1), v.UserID)
	}
	order.Nil(err)
}

func (order *OrderTest) TestOrderCart() {
	orderHistory :=  func(userID uint, page int, perPage int) *[]domain.Cart {
		output := &[]domain.Cart{}
		for _, v := range *output {
			v.UserID = userID
		}

		return output
	}

	order.repo.On("Cart", mock.AnythingOfType("uint"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(orderHistory, nil)

	data, err := order.repo.Cart(1, 10, 10)

	for _, v := range *data {
		order.Equal(uint(1), v.UserID)
	}
	order.Nil(err)
}
func TestOrder(t *testing.T) {
	suite.Run(t, NewOrderTest())
}
