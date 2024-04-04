package tests

import (
	repomocks "github.com/rulanugrh/lysithea/internal/mocks"
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

}
