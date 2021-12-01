package model

import "time"

// // User struct
type User struct {
	ID        int64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// // Users struct
type Users struct {
	Users []User `json:"users"`
}