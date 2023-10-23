package request

type CreateAccountRequest struct {
	IDUser  int64   `json:"id_user"`
	Grade   string  `json:"grade"`
	IDBank  int64   `json:"id_bank"`
	Balance float64 `json:"balance"`
	Token   string  ` validate:"required"`
}
