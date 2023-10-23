package request

type LogoutRequest struct {
	ID int64 `json:"id_user" validate:"required"`
}
