package entity

import "time"

type User struct {
	ID        int
	Email     string
	Password  string
	Name      string
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
