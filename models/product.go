package models

import "time"

type Product struct {
	ID        uint32    `pg:",pk" json:"id"`
	Name      string    `pg:"type:varchar(60)" json:"name"`
	UserID    uint32    `pg:"on_delete:CASCADE,unique" json:"user_id"`
	User      *User     `pg:"rel:has-one" json:"user,omitempty"`
	CreatedAt time.Time `pg:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `pg:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
