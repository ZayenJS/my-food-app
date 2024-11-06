package main

import (
	"github.com/ZayenJS/config"
	"github.com/ZayenJS/database"
	"github.com/ZayenJS/router"
)

func main() {
	config.Load()
	app := router.Setup()
	database.Setup()
	app.Run() // listen and serve on 0.0.0.0:8080
}
