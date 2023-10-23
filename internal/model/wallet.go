package model

import "time"

type Wallet struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	IDUser    int64     `gorm:"type:bigint" json:"id_user"`
	Balance   float64   `gorm:"type:float" json:"balance"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
