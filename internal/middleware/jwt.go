package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rulanugrh/lysithea/internal/config"
	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
)

type jwtclaim struct {
	ID     uint
	Name   string
	RoleID uint
	jwt.RegisteredClaims
}

func GenerateToken(user domain.UserLogin) (string, error) {
	conf := config.GetConfig()
	time := jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	claims := &jwtclaim{
		ID:     user.ID,
		Name:   user.Name,
		RoleID: user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: time,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(conf.Server.Secret))
	if err != nil {
		return "", web.StatusBadRequest("cannot create token")
	}

	return tokenString, nil
}
