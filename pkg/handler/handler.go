package handler

import (
	"github.com/gorilla/mux"
	"goLang/pkg/service"
)

type Handler struct {
	services *service.Service
}

func HewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user/{id}", h.getUserById).Methods("GET")
	router.HandleFunc("/users/", h.getUsers).Methods("GET")
	router.HandleFunc("/user/", h.createUser).Methods("POST")
	router.HandleFunc("/user/{id}", h.deleteUser).Methods("DELETE")
	return router
}
