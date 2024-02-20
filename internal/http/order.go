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
	AddToCart(w http.ResponseWriter, r *http.Request)
	FindID(w http.ResponseWriter, r *http.Request)
	Cart(w http.ResponseWriter, r *http.Request)
	History(w http.ResponseWriter, r *http.Request)
	Pay(w http.ResponseWriter, r *http.Request)
	Buy(w http.ResponseWriter, r *http.Request)
	Checkout(w http.ResponseWriter, r *http.Request)
}

type order struct {
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) OrderHandler {
	return &order{
		service: service,
	}
}

func (o *order) AddToCart(w http.ResponseWriter, r *http.Request) {
	var req domain.Cart
	checkToken, err := middleware.CheckToken(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	req.UserID = checkToken.ID

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.Unmarshal(body, &req)
	data, err := o.service.AddToCart(req)
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

	response, err := json.Marshal(web.Success("adding to cart", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(response)
}

func (o *order) Buy(w http.ResponseWriter, r *http.Request) {
	var req domain.Order
	checkToken, err := middleware.CheckToken(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	req.UserID = checkToken.ID

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.Unmarshal(body, &req)
	data, err := o.service.Buy(req)
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

func (o *order) FindID(w http.ResponseWriter, r *http.Request) {
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

func (o *order) Checkout(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/order/checkout/"))
	data, err := o.service.Checkout(uint(id))
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

	response, err := json.Marshal(web.Success("success checkout", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(response)
}

func (o *order) Pay(w http.ResponseWriter, r *http.Request) {
	uuid := strings.TrimPrefix(r.URL.Path, "/api/v1/order/pay/")
	checkToken, err := middleware.CheckToken(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	data, err := o.service.Pay(uuid, checkToken.ID)
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

	response, err := json.Marshal(web.Success("success pay", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(response)
}

func (o *order) Cart(w http.ResponseWriter, r *http.Request) {
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

	data, err := o.service.Cart(claim.ID, page, per_page)
	if err != nil {
		response, err := json.Marshal(web.StatusNotFound("cart with this user id not found"))
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(500)
		w.Write(response)
		return
	}

	response, err := json.Marshal(web.Success("printout cart", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}

func (o *order) History(w http.ResponseWriter, r *http.Request) {
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

	data, err := o.service.History(claim.ID, page, per_page)
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
