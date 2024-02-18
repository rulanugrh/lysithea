package repository

import (
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(req domain.UserRequest) (*domain.User, error)
	Login(req domain.UserLogin) (*domain.User, error)
}

type user struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &user{
		db: db,
	}
}

func (u *user) Register(req domain.UserRequest) (*domain.User, error) {
	user_request := domain.User{
		Name:     req.Name,
		NoHP:     req.NoHP,
		Password: req.Password,
		RoleID:   req.RoleID,
		Email:    req.Email,
	}

	findEmail := u.db.Where("email = ?", req.Email).Find(&req)
	if findEmail.RowsAffected != 0 {
		return nil, web.NewInternalServerErrorResponse("cannot request again, because email has been used")
	}

	err := u.db.Create(&user_request).Error
	if err != nil {
		return nil, web.NewInternalServerErrorResponse("cannot create user")
	}

	return &user_request, web.NewCreatedResponse("success create user", user_request)
}

func (u *user) Login(req domain.UserLogin) (*domain.User, error) {
	var user_request domain.User
	findEmail := u.db.Where("email = ?", req.Email).Find(&user_request)
	if findEmail.RowsAffected == 0 {
		return nil, web.NewStatusNotFound("user not found with this email")
	}

	return &user_request, nil
}
