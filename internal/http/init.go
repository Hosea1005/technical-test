package http

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"service-auth/domain/usecase"
)

type AuthHandler struct {
	usecase usecase.AuthUseCase
}

func NewDeliveryHttpArea(r *mux.Router, usecase usecase.AuthUseCase) AuthHandler {
	handler := AuthHandler{
		usecase: usecase,
	}

	r.HandleFunc("/register", handler.Register).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/logout", handler.Logout).Methods("GET")
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/list-player", handler.GetListPlayer).Methods("GET")
	api.HandleFunc("/detail-player/{id}", handler.GetDetailPlayer).Methods("GET")
	api.HandleFunc("/account/add", handler.CreateAccount).Methods("POST")
	api.HandleFunc("/wallet/top-up", handler.TopUp).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
	return handler
}
