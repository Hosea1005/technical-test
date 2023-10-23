package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"service-auth/helper"
	"service-auth/internal/http/create_account/request"
)

func (a AuthHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var (
		req request.CreateAccountRequest
	)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, []error{errors.New("unauthorized")})
		return
	}
	defer r.Body.Close()
	auth := r.Header.Get("Authorization")
	req.Token = auth
	res := a.usecase.CreateAccount(r.Context(), req)
	helper.RespondWithJSON(w, http.StatusOK, res)
}
