package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"net/http"
	"service-auth/helper"
	"service-auth/internal/http/login/request"
	"service-auth/internal/http/register/response"
)

func (a AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		req request.LoginRequest
	)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, []error{errors.New("unauthorized")})
		return
	}
	defer r.Body.Close()
	validate := validator.New()
	// Use the validator to check if the struct is valid
	if err := validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, validationError := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("Field %s is required", validationError.Field()))
		}
		helper.RespondWithJSON(w, http.StatusBadRequest, response.RegisterResponse{
			Status: response.Status{
				Code:    400,
				Message: errorMessages[0],
			},
		})
		return
	}
	res := a.usecase.Login(r.Context(), req)
	// set token yang ke cookie
	if res.Status.Code == 200 {
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    res.Data.Token,
			HttpOnly: true,
		})
	}
	helper.RespondWithJSON(w, http.StatusOK, res)
}
