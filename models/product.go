package models

import "time"

type Product struct {
	ID        uint32     `json:"id"`
	User      User       `gorm:"-" json:"user"`
	UserID    uint32     `gorm:"not null" json:"user_id"`
	Name      string     `gorm:"not null;size:60" json:"name"`
	CreatedAt *time.Time `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"not null;default:now()" json:"updated_at"`
}
