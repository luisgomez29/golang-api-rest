package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint32     `json:"id"`
	FirstName string     `gorm:"not null;size:40" json:"first_name,omitempty"`
	LastName  string     `gorm:"not null;size:40;" json:"last_name,omitempty"`
	FullName  string     `gorm:"-" json:"full_name,omitempty"`
	Email     string     `gorm:"not null;size:60;unique" json:"email,omitempty"`
	Password  string     `gorm:"not null;size:128" json:"password,omitempty"`
	CreatedAt *time.Time `gorm:"not null;default:now()" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"not null;default:now()" json:"updated_at,omitempty"`
	Products  []Product  `gorm:"constraint:OnDelete:CASCADE" json:"products,omitempty"`
}

func (u *User) BeforeSave(*gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) AfterFind(*gorm.DB) error {
	u.FullName = u.FirstName + " " + u.LastName
	return nil
}
