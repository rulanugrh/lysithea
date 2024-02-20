package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/service"
)

type CategoryHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetCategoryBySearch(w http.ResponseWriter, r *http.Request)
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

func (c *category) GetCategoryBySearch(w http.ResponseWriter, r *http.Request) {
	per_page, _ := strconv.Atoi(r.URL.Query().Get("per_page"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	search := r.URL.Query().Get("search")

	var buffer bytes.Buffer
	data, err := c.service.GetCategoryBySearch(page, per_page, search, buffer)
	if err != nil {
		response, err := json.Marshal(web.StatusNotFound("data not found"))
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(404)
		w.Write(response)
		return
	}

	result := web.PaginationElastic{
		Metadata: web.MetadataElastic{
			Page:    page,
			PerPage: per_page,
		},
		Data: data,
	}

	response, err := json.Marshal(web.Success("category founded", result))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}
