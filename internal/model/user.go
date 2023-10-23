package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Fullname  string    `gorm:"type:varchar(300)" json:"fullname"`
	Username  string    `gorm:"type:varchar(300);unique" json:"username"`
	Password  string    `gorm:"type:varchar(300)" json:"password"`
	Token     string    `gorm:"type:varchar(500)" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
