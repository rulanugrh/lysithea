package tests

import (
	"testing"

	"github.com/rulanugrh/lysithea/internal/entity/domain"
	repomocks "github.com/rulanugrh/lysithea/internal/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserTest struct {
	suite.Suite
	repo repomocks.UserRepository
}

func NewUserTest() *UserTest {
	return &UserTest{}
}

func (user *UserTest) TestUserRegister() {
	userResult := func(input domain.UserRequest) *domain.User {
		output := &domain.User{}
		output.Name = input.Name
		output.RoleID = input.RoleID
		output.Password = input.Password
		output.Email = input.Email
		output.NoHP = input.NoHP
		return output
	}

	user.repo.On("Register", mock.MatchedBy(func(input domain.UserRequest) bool {
		return input.Name != "" && input.RoleID > 0 && input.Email != ""
	})).Return(userResult, nil)

	data, err := user.repo.Register(domain.UserRequest{
		Name:     "John Doe",
		Email:    "johndoe@john.co.id",
		RoleID:   1,
		NoHP:     8571664532,
		Password: "johndoe123",
	})

	user.Nil(err)
	user.Equal("John Doe", data.Name)
	user.Equal("johndoe@john.co.id", data.Email)
	user.Equal(uint(1), data.RoleID)
}

func (user *UserTest) TestUserLogin() {
	userResult := func(input domain.UserLogin) *domain.User {
		output := &domain.User{}
		output.ID = input.ID
		output.Name = input.Name
		output.RoleID = input.RoleID
		output.Password = input.Password
		output.Email = input.Email
		return output
	}

	user.repo.On("Login", mock.MatchedBy(func(input domain.UserLogin) bool {
		return input.Password != "" && input.Email != ""
	})).Return(userResult, nil)

	data, err := user.repo.Login(domain.UserLogin{
		Email:    "johndoe@john.co.id",
		Password: "johndoe123",
	})

	user.Nil(err)
	user.Equal("johndoe123", data.Password)
	user.Equal("johndoe@john.co.id", data.Email)
}

func TestUser(t *testing.T) {
	suite.Run(t, NewUserTest())
}
