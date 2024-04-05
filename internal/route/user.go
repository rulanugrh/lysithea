package route

import (
	"github.com/gorilla/mux"
	_ "github.com/rulanugrh/lysithea/docs"
	"github.com/rulanugrh/lysithea/internal/config"
	handler "github.com/rulanugrh/lysithea/internal/http"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func UserRouter(router *mux.Router, handler handler.UserHandler, conf *config.App) {
	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	subrouter := router.PathPrefix("/api/v1/user/").Subrouter()
	subrouter.HandleFunc("/register", handler.Register).Methods("POST")
	subrouter.HandleFunc("/login", handler.Login).Methods("POST")
}
