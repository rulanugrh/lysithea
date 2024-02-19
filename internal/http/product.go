package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/service"
)

type ProductHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindID(w http.ResponseWriter, r *http.Request)
}

type product struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) ProductHandler {
	return &product{
		service: service,
	}
}

func (p *product) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.ProductRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.Unmarshal(body, &req)

	data, err := p.service.Create(req)
	if err != nil {
		response, _ := json.Marshal(web.StatusBadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	response, err := json.Marshal(web.Created("success create product", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(201)
	w.Write(response)
}

func (p *product) FindID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/v1/product/find/")
	converID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	data, err := p.service.FindID(uint(converID))
	if err != nil {
		response, err := json.Marshal(web.StatusNotFound("data not found with this id"))
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(404)
		w.Write(response)
		return
	}

	response, err := json.Marshal(web.Success("data found", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}