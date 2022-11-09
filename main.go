package main

import (
	"bankku/config"
	"bankku/exceptions"
	"bankku/routes"
	utils "bankku/utils/database/mysql"

	jsongoccy "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.GetConfig()
	db := utils.InitDB(cfg)
	app := fiber.New(fiber.Config{
		JSONEncoder:  jsongoccy.Marshal,
		JSONDecoder:  jsongoccy.Unmarshal,
		ErrorHandler: exceptions.CustomErrorHandling,
	})
	app.Use(recover.New())

	routes.InitRoutes(app, db)

	if err := app.Listen(":3030"); err != nil {
		panic(err)
	}
}
