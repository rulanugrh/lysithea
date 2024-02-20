package route

import (
	"github.com/gorilla/mux"
	handler "github.com/rulanugrh/lysithea/internal/http"
	"github.com/rulanugrh/lysithea/internal/middleware"
)

func ProductRoute(router *mux.Router, handler handler.ProductHandler) {
	subRouter := router.PathPrefix("/api/v1/order").Subrouter()
	subRouter.Use(middleware.ValidateToken)
	subRouter.HandleFunc("/create", handler.Create).Methods("POST")
	subRouter.HandleFunc("/category/:id", handler.FindAllByCategoryID).Methods("GET")
	subRouter.HandleFunc("/find/:id", handler.FindID).Methods("GET")
	subRouter.HandleFunc("/find", handler.FindAll).Methods("GET")
	subRouter.HandleFunc("/search", handler.FindBySearch).Methods("GET")
}
