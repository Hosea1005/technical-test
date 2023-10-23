package request

type DetailPlayerRequest struct {
	ID    string `json:"id"`
	Token string ` validate:"required"`
}
