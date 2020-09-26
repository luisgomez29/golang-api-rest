package models

import (
	"context"
	"github.com/go-pg/pg/v10"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint32    `pg:",pk" json:"id"`
	FirstName string    `pg:"type:varchar(40)" json:"first_name"`
	LastName  string    `pg:"type:varchar(40)" json:"last_name"`
	Email     string    `pg:"type:varchar(60)" pg:",unique" json:"email"`
	Password  string    `pg:"type:varchar(128)" json:"-"`
	CreatedAt time.Time `pg:"default:now()" json:"created_at"`
	UpdatedAt time.Time `pg:"default:now()" json:"updated_at"`
}

var _ pg.BeforeInsertHook = (*User)(nil)

func (u *User) BeforeInsert(ctx context.Context) (context.Context, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashedPassword)
	return ctx, nil
}
