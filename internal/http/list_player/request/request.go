package request

type ListPlayerRequest struct {
	Token  string `validate:"required"`
	Search string
}
