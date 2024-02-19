package middleware

import (
	"encoding/json"
	"net/http"
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

func CheckToken(token string) (*jwtclaim, error) {
	conf := config.GetConfig()
	tokens, _ := jwt.ParseWithClaims(token, &jwtclaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.Server.Secret), web.Forbidden("this is strict page")
	})

	claim, err := tokens.Claims.(*jwtclaim)
	if !err {
		return nil, web.Unauthorized("sorry you not have token")
	}

	return claim, nil
}

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conf := config.GetConfig()
		token := r.Header.Get("Authorization")

		if token == "" {
			response, err := json.Marshal(web.Unauthorized("you dont have token"))
			if err != nil {
				w.WriteHeader(500)
				return
			}

			w.WriteHeader(401)
			w.Write(response)
			return
		}

		tokens, _ := jwt.ParseWithClaims(token, &jwtclaim{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(conf.Server.Secret), web.Forbidden("this is strict page")
		})

		claim, err := tokens.Claims.(*jwtclaim)
		if !err {
			response, err := json.Marshal(web.Unauthorized("you dont have token"))
			if err != nil {
				w.WriteHeader(500)
				return
			}

			w.WriteHeader(401)
			w.Write(response)
			return
		}

		if claim.ExpiresAt.Unix() < time.Now().Unix() {
			response, err := json.Marshal(web.Unauthorized("token expired"))
			if err != nil {
				w.WriteHeader(500)
				return
			}

			w.WriteHeader(401)
			w.Write(response)
			return
		}

		next.ServeHTTP(w, r)

	})
}
