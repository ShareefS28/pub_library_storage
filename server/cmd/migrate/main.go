package main

import (
	"server/config"
	"server/database"
)

func main() {
	cfg := config.Load()

	database.Connect(cfg.DBDSN)

	database.Migrate()
}
