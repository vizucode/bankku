package usermodel

import (
	usercore "bankku/domains/user/core"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string
	Name        string
	PhoneNumber string
	Address     string
	Password    string
	IsActive    bool
}

func ToCore(model User) usercore.Core {
	return usercore.Core{
		Id:          model.ID,
		Email:       model.Email,
		Name:        model.Name,
		PhoneNumber: model.PhoneNumber,
		Address:     model.Address,
		Password:    model.Password,
		IsActive:    model.IsActive,
	}
}

func ToModel(Core usercore.Core) User {
	return User{
		Email:       Core.Email,
		Name:        Core.Name,
		PhoneNumber: Core.PhoneNumber,
		Address:     Core.Address,
		Password:    Core.Password,
	}
}
