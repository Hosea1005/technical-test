package model

import "time"

type Account struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	IDUser    int64     `gorm:"type:bigint;unique" json:"id_user"`
	IDBank    int64     `gorm:"type:bigint" json:"id_bank"`
	Grade     string    `gorm:"type:varchar(300)" json:"grade"`
	Balance   float64   `gorm:"type:float" json:"balance"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
