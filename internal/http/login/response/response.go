package response

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type Data struct {
	Token string `json:"token"`
}
type LoginResponse struct {
	Status Status `json:"status"`
	Data   Data   `json:"data"`
}
