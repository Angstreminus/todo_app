package main

import (
	"todo_app/config"
	"todo_app/server"
)

func main() {
	config := config.InitConfig()
	dbHandler := server.InitDatabase(config)
	Server := server.InitHttpServer(config, dbHandler)
	Server.Start()
}
