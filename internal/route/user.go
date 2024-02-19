package route

import (
	"github.com/gorilla/mux"
	handler "github.com/rulanugrh/lysithea/internal/http"
)

func UserRouter(router *mux.Router, handler handler.UserHandler) {
	subrouter := router.PathPrefix("/api/v1/user/").Subrouter()
	subrouter.HandleFunc("/register", handler.Register).Methods("POST")
	subrouter.HandleFunc("/login", handler.Login).Methods("POST")
}
