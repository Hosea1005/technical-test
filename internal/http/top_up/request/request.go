package request

type TopUpRequest struct {
	IDUser    int64   `json:"id_user"`
	IDAccount int64   `json:"id_account"`
	IDWallet  int64   `json:"id_wallet"`
	Total     float64 `json:"total"`
	Token     string  ` validate:"required"`
}
