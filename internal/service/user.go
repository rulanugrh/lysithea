package service

import (
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/middleware"
	"github.com/rulanugrh/lysithea/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req domain.UserRequest) (*web.UserRegister, error)
	Login(req domain.UserLogin) (*web.UserLogin, error)
}

type user struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &user{
		repo: repo,
	}
}

func (u *user) Register(req domain.UserRequest) (*web.UserRegister, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return nil, web.StatusBadRequest("cannot generate hash password")
	}

	data, err := u.repo.Register(domain.UserRequest{
		Name:     req.Name,
		Password: string(password),
		Email:    req.Email,
		RoleID:   req.RoleID,
		NoHP:     req.NoHP,
	})

	if err != nil {
		return nil, web.StatusBadRequest("cannot create user")
	}

	response := web.UserRegister{
		Name:  data.Name,
		Email: data.Email,
	}

	return &response, nil
}
func (u *user) Login(req domain.UserLogin) (*web.UserLogin, error) {
	data, err := u.repo.Login(req)
	if err != nil {
		return nil, web.StatusNotFound("email not found")
	}

	compare := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	if compare != nil {
		return nil, web.StatusBadRequest("password not matched")
	}

	generateToken, err := middleware.GenerateToken(domain.UserLogin{
		ID:     data.ID,
		RoleID: data.RoleID,
		Name:   data.Name,
	})

	if err != nil {
		return nil, web.InternalServerError("cannot generate token jwt")
	}

	response := web.UserLogin{
		Token: generateToken,
	}

	return &response, nil
}
