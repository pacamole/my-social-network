package models

import "time"

type User struct {
	ID         uint16    `gorm:"primary_key,autoIncrement" json:"id"`
	Email      string    `db:"email" json:"email"`
	Name       string    `db:"name" json:"name"`
	Password   string    `db:"password" json:"password,omitempty"`
	Created_at time.Time `db:"created_at" json:"created_at"`
	Updated_at time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
