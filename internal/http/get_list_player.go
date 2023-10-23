package http

import (
	"net/http"
	"service-auth/helper"
	"service-auth/internal/http/list_player/request"
)

func (a AuthHandler) GetListPlayer(w http.ResponseWriter, r *http.Request) {
	var (
		req request.ListPlayerRequest
	)
	auth := r.Header.Get("Authorization")
	req.Token = auth
	search := r.FormValue("searchByLevel")
	req.Search = search
	res := a.usecase.GetListPlayer(r.Context(), req)
	helper.RespondWithJSON(w, http.StatusOK, res)
}
