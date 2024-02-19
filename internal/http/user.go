package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/service"
)

type UserHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type user struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return &user{
		service: service,
	}
}

func (u *user) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.UserRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.Unmarshal(body, &req)
	data, err := u.service.Register(req)
	if err != nil {
		response, err := json.Marshal(web.StatusBadRequest(err.Error()))
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(400)
		w.Write(response)
		return
	}

	response, err := json.Marshal(web.Created("success create account", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(201)
	w.Write(response)
}

func (u *user) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.UserLogin
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.Unmarshal(body, &req)
	data, err := u.service.Login(req)
	if err != nil {
		response, err := json.Marshal(web.StatusBadRequest(err.Error()))
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(400)
		w.Write(response)
		return
	}

	response, err := json.Marshal(web.Success("success login", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}
