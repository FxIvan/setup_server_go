package domain

import (
	"time"
)

type UserRole string

const (
	Admin   UserRole = "admin"
	Cashier UserRole = "user"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}
