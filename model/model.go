package model

import "time"

// Users model for users table
type Users struct {
	UserId     *string    `db:"user_id"`
	Password   *string    `db:"password"`
	InsertDate *time.Time `db:"insert_date"`
	UpdateDate *time.Time `db:"update_date"`
}
