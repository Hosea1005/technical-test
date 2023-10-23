package model

import "time"

type Bank struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(300)" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
