package model

import "time"

type User struct {
	ID        uint64    `gorm:"primary_id;auto_increment" json:"id"`
	Firstname string    `json:"firstname",omitempty`
	Lastname  string    `json:"lastname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}
type UpdateUser struct {
	Firstname *string `json:"firstname",omitempty`
	Lastname  *string `json:"lastname,omitempty"`
	Email     *string `json:"email,omitempty"`
	Password  *string `json:"password,omitempty"`
}
type LoginVar struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
