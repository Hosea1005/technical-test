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
	Name      string    `json:"nama_lengkap"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}
