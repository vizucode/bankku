package usercore

import "time"

type Core struct {
	Id          uint
	Email       string
	Name        string
	PhoneNumber string
	Address     string
	Password    string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
