package route

import (
	"github.com/gorilla/mux"
	handler "github.com/rulanugrh/lysithea/internal/http"
	"github.com/rulanugrh/lysithea/internal/middleware"
)

func OrderRouter(router *mux.Router, handler handler.OrderHandler) {
	subrouter := router.PathPrefix("/api/v1/order/").Subrouter()
	subrouter.Use(middleware.ValidateToken)
	subrouter.HandleFunc("/create", handler.Create).Methods("POST")
	subrouter.HandleFunc("/history", handler.FindByUserID).Methods("GET")
	subrouter.HandleFunc("/find/:uuid", handler.FindByUUID).Methods("GET")
}
