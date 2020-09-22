package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint32    `gorm:"type:serial;primaryKey;auto_increment" json:"id"`
	FirstName string    `gorm:"not null;size:40" json:"first_name"`
	LastName  string    `gorm:"not null;size:40;" json:"last_name"`
	Email     string    `gorm:"not null;size:60;unique" json:"email"`
	Password  string    `gorm:"not null;size:128" json:"-"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) BeforeSave(*gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
