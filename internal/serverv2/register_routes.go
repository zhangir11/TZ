package server

import (
	service "authentication/pkg/servicev2"
	"net/http"
)

func RegisterHandlers(router *http.ServeMux, authService service.Service) {
	h := NewHandler(authService)
	router.HandleFunc("/token", h.Token)     //.Methods("GET")
	router.HandleFunc("/refresh", h.Refresh) //.Methods("POST")
}
