package models

import "time"

type Product struct {
	ID        uint32     `json:"id"`
	UserID    uint32     `gorm:"unique" json:"user_id"`
	User      User       `gorm:"constraint:OnDelete:CASCADE" json:"user"`
	Name      string     `gorm:"not null;size:60" json:"name"`
	CreatedAt *time.Time `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"not null;default:now()" json:"updated_at"`
}
