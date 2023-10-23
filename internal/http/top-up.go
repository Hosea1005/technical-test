package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"service-auth/helper"
	request2 "service-auth/internal/http/top_up/request"
)

func (a AuthHandler) TopUp(w http.ResponseWriter, r *http.Request) {
	var (
		req request2.TopUpRequest
	)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, []error{errors.New("unauthorized")})
		return
	}
	defer r.Body.Close()
	auth := r.Header.Get("Authorization")
	req.Token = auth
	res := a.usecase.TopUp(r.Context(), req)
	helper.RespondWithJSON(w, http.StatusOK, res)
}
