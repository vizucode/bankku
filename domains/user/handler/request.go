package userhandler

type Request struct {
	Email       string `json:"email" form:"email" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Address     string `json:"address" form:"address" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
}

type RequestVerify struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
