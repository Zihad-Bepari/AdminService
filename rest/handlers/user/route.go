package user

import (
	"admin-service/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoute(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("POST /users",
		manager.With(
			http.HandlerFunc(h.Createusers)))

	mux.Handle("POST /users/changerole",
		manager.With(
			http.HandlerFunc(h.Changerole)))

}
