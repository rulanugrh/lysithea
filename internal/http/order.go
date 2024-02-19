package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/rulanugrh/lysithea/internal/entity/domain"
	"github.com/rulanugrh/lysithea/internal/entity/web"
	"github.com/rulanugrh/lysithea/internal/middleware"
	"github.com/rulanugrh/lysithea/internal/service"
)

type OrderHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindByUUID(w http.ResponseWriter, r *http.Request)
	FindByUserID(w http.ResponseWriter, r *http.Request)
}

type order struct {
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) OrderHandler {
	return &order{
		service: service,
	}
}

func (o *order) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Order
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.Unmarshal(body, &req)
	data, err := o.service.Create(req)
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

	response, err := json.Marshal(web.Created("success create order", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(201)
	w.Write(response)
}

func (o *order) FindByUUID(w http.ResponseWriter, r *http.Request) {
	uuid := strings.TrimPrefix(r.URL.Path, "/api/v1/order/find/")
	data, err := o.service.FindID(uuid)
	if err != nil {
		response, err := json.Marshal(web.StatusNotFound(err.Error()))
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(404)
		w.Write(response)
		return
	}

	response, err := json.Marshal(web.Success("order found", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(response)
}

func (o *order) FindByUserID(w http.ResponseWriter, r *http.Request) {
	per_page, _ := strconv.Atoi(r.URL.Query().Get("per_page"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	token := r.Header.Get("Authorization")
	claim, err := middleware.CheckToken(token)
	if err != nil {
		response, err := json.Marshal(web.InternalServerError("sorry something error to catch token"))
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(500)
		w.Write(response)
		return
	}

	data, err := o.service.FindByUserID(claim.ID, page, per_page)
	if err != nil {
		response, err := json.Marshal(web.StatusNotFound("order history with this user id not found"))
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(500)
		w.Write(response)
		return
	}

	response, err := json.Marshal(web.Success("printout order history", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}
