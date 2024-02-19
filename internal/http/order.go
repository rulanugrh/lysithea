package handler

import "net/http"

type OrderHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}
