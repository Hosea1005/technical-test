package response

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type CreateAccountResponse struct {
	Status Status `json:"status"`
}
