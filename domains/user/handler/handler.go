package userhandler

import (
	usercore "bankku/domains/user/core"
	"bankku/exceptions"
	"bankku/utils/helpers"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type activityHandler struct {
	service usercore.IServiceUser
}

var validate = validator.New()

func New(service usercore.IServiceUser) *activityHandler {
	return &activityHandler{
		service: service,
	}
}

func (h *activityHandler) Create(c *fiber.Ctx) error {
	request := Request{}

	err := c.BodyParser(&request)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = validate.Struct(&request)
	if err != nil {
		return err
	}

	result := h.service.Create(usercore.Core{
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
		Name:        request.Name,
		Password:    request.Password,
	})

	return c.Status(http.StatusCreated).JSON(helpers.SuccessGetResponseData(Response{
		Id:          result.Id,
		Email:       result.Email,
		Name:        result.Name,
		PhoneNumber: result.PhoneNumber,
		Address:     result.Address,
		Password:    result.Password,
	}))
}

func (h *activityHandler) Verify(c *fiber.Ctx) error {
	request := RequestVerify{}

	err := c.BodyParser(&request)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = validate.Struct(&request)
	if err != nil {
		return err
	}

	result := h.service.Verify(usercore.Core{
		Email:    request.Email,
		Password: request.Password,
	})

	if result.Id > 1 {
		return c.Status(http.StatusOK).JSON(helpers.SuccessActionResponse("account verified"))
	} else {
		return c.Status(http.StatusUnauthorized).JSON(helpers.SuccessActionResponse("account not verified"))
	}
}
