package http

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"service-auth/helper"
	"service-auth/internal/http/detail_player/request"
)

func (a AuthHandler) GetDetailPlayer(w http.ResponseWriter, r *http.Request) {
	log.Println("MASUK")
	var (
		req request.DetailPlayerRequest
	)
	vars := mux.Vars(r)
	id := vars["id"]
	req.ID = id
	auth := r.Header.Get("Authorization")
	req.Token = auth
	res := a.usecase.GetDetailPlayer(r.Context(), req)
	helper.RespondWithJSON(w, http.StatusOK, res)
}
