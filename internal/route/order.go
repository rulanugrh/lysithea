package route

import (
	"github.com/gorilla/mux"
	handler "github.com/rulanugrh/lysithea/internal/http"
	"github.com/rulanugrh/lysithea/internal/middleware"
)

func OrderRouter(router *mux.Router, handler handler.OrderHandler) {
	subrouter := router.PathPrefix("/api/v1/order/").Subrouter()
	subrouter.Use(middleware.ValidateToken)
	subrouter.HandleFunc("/add", handler.AddToCart).Methods("POST")
	subrouter.HandleFunc("/buy", handler.Buy).Methods("POST")
	subrouter.HandleFunc("/cart", handler.Cart).Methods("GET")
	subrouter.HandleFunc("/pay/:uuid", handler.Pay).Methods("PUT")
	subrouter.HandleFunc("/history", handler.History).Methods("GET")
	subrouter.HandleFunc("/find/:uuid", handler.FindID).Methods("GET")
	subrouter.HandleFunc("/checkout/:id", handler.Checkout).Methods("POST")
}
