package route

import (
	"github.com/gorilla/mux"
	handler "github.com/rulanugrh/lysithea/internal/http"
	"github.com/rulanugrh/lysithea/internal/middleware"
)

func CategoryRouter(router *mux.Router, handler handler.CategoryHandler) {
	subrouter := router.PathPrefix("/api/v1/category/").Subrouter()
	subrouter.Use(middleware.ValidateToken)
	subrouter.HandleFunc("/create", handler.Create).Methods("POST")
	subrouter.HandleFunc("/search", handler.GetCategoryBySearch).Methods("GET")
}
