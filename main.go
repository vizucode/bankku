package main

import (
	"bankku/config"
	"bankku/route"
	database "bankku/utils/database/mysql"
)

func main() {
	cfg := config.GetConfig()
	db := database.InitDB(cfg)
	route.Route(db)
}
