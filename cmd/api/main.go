package main

import (
	"Collap/internal/config"
	"Collap/internal/database"
	"Collap/internal/server"
)

func main() {
	conf := config.GetConfig()

	database.InitializeDB(conf.Db)

	server.StartServer()
}
