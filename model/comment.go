package model

import "time"

type Comment struct {
	ID int
	Conent string
	CreatedAt time.Time
	UpdatedAt time.Time
}