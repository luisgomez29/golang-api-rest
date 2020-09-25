package models

import "time"

type Product struct {
	ID        uint32    `gorm:"type:serial;primaryKey;auto_increment" json:"id"`
	UserID    uint32    `json:"user_id"`
	User      User      `json:"user"`
	Name      string    `gorm:"not null;size:60" json:"name"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
