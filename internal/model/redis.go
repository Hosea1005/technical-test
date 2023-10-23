package model

import "time"

type TopUp struct {
	IDUser    int64     `json:"id_user"`
	IDAccount int64     `json:"id_account"`
	IDWallet  int64     `json:"id_wallet"`
	Total     float64   `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
