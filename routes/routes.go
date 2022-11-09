package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	userhandler "bankku/domains/user/handler"
	userrepo "bankku/domains/user/repository"
	userservice "bankku/domains/user/service"
)

func InitRoutes(ctx *fiber.App, db *gorm.DB) {
	/*
		Dependency Injection
	*/

	userRepo := userrepo.New(db)
	userService := userservice.New(userRepo)
	userHandler := userhandler.New(userService)

	/*
		Routes
	*/

	ctx.Post("/register", userHandler.Create)
	ctx.Put("/verify", userHandler.Verify)
}
