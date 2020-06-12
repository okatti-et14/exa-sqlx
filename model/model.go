package model

import "time"

// Users model for users table
type Users struct {
	UserID     *string    `db:"user_id"`
	Password   *string    `db:"password"`
	InsertDate *time.Time `db:"insert_date"`
	UpdateDate *time.Time `db:"update_date"`
	UserNames  UserNames  `db:"user_names"`
}

// UserNames model for user_name table
type UserNames struct {
	UserID   *string `db:"user_id"`
	UserName *string `db:"user_name"`
}

// UsersUnionUseNames model for users table join user_names table
t!
}
