package response

import "time"

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type Data struct {
	Player []Player `json:"player"`
}
type ListPlayerResponse struct {
	Status Status `json:"status"`
	Data   Data   `json:"data"`
}

type Player struct {
	ID        uint      `json:"id"`
	Name      string    `json:"fullname"`
	Username  string    `json:"username"`
	Level     string    `json:"level"`
	CreatedAt time.Time `json:"created_at"`
}
