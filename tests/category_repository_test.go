package tests

import (
	"testing"

	"github.com/rulanugrh/lysithea/internal/entity/domain"
	repomocks "github.com/rulanugrh/lysithea/internal/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CategoryTest struct {
	suite.Suite
	repo repomocks.CategoryRepository
}

func NewCategoryTest() *CategoryTest {
	return &CategoryTest{}
}

func (category *CategoryTest) TestCategoryCreate() {
	categoryResult := func(input domain.Category) *domain.Category {
		output := &domain.Category{}
		output.Name = input.Name
		output.Description = input.Description
		return output
	}

	category.repo.On("Create", mock.MatchedBy(func(input domain.Category) bool {
		return input.Name != "" && input.Description != ""
	})).Return(categoryResult, nil)

	data, err := category.repo.Create(domain.Category{
		Name:        "Electronic",
		Description: "Simple Electronic",
	})

	category.Nil(err)
	category.Equal("Electronic", data.Name)
	category.Equal("Simple Electronic", data.Description)
}

func TestCategory(t *testing.T) {
	suite.Run(t, NewCategoryTest())
}
