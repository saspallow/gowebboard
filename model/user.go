package model

import "time"

type User struct {
	ID int
	Username string
	Password string
	CreatedAt time.Time
	UpdatetedAt time.Time
}