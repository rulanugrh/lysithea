package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/service"
)

type CategoryHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type category struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) CategoryHandler {
	return &category{
		service: service,
	}
}

func (c *category) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Category
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.Unmarshal(body, &req)
	data, err := c.service.Create(req)
	if err != nil {
		response, err := json.Marshal(web.StatusBadRequest(err.Error()))
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
		w.Write(response)
		return
	}

	response, err := json.Marshal(web.Created("success create category", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}
