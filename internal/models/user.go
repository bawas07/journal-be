package models

import (
	"time"
)

type User struct {
	ID           int64      `json:"id" db:"id"`
	Email        string     `json:"email" db:"email"`
	Username     string     `json:"username" db:"username"`
	Password     string     `json:"-" db:"password"`
	ReminderTime *time.Time `json:"reminder_time" db:"reminder_time"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}
