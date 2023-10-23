package response

import "time"

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type Data struct {
	Player Player `json:"player"`
}
type DetailPlayerResponse struct {
	Status Status `json:"status"`
	Data   Data   `json:"data"`
}

type Player struct {
	ID        uint      `json:"id"`
	Name      string    `json:"fullname"`
	Username  string    `json:"username"`
	Level     string    `json:"level"`
	Account   Account   `json:"account"`
	Wallet    Wallet    `json:"wallet"`
	CreatedAt time.Time `json:"created_at"`
}
type Account struct {
	ID        uint      `json:"id"`
	Grade     string    `json:"grade"`
	Balance   float64   `json:"balance"`
	Bank      string    `json:"bank"`
	CreatedAt time.Time `json:"created_at"`
}
type Wallet struct {
	ID        uint      `json:"id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
