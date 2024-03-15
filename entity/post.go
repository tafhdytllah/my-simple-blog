package entity

import "time"

type Post struct {
	ID         int
	UserID     int
	Title      string
	Content    string
	PictureUrl *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
