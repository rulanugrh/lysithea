package route

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/rulanugrh/lysithea/internal/config"
	handler "github.com/rulanugrh/lysithea/internal/http"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func UserRouter(router *mux.Router, handler handler.UserHandler, conf *config.App) {
	router.HandleFunc("/docs/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s:%s/docs/swagger.json", conf.Server.URL, conf.Server.Port)),
	))
	subrouter := router.PathPrefix("/api/v1/user/").Subrouter()
	subrouter.HandleFunc("/register", handler.Register).Methods("POST")
	subrouter.HandleFunc("/login", handler.Login).Methods("POST")
}
