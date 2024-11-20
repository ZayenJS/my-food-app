package main

import (
	"fmt"
	"os"

	"github.com/ZayenJS/config"
	"github.com/ZayenJS/database"
	"github.com/ZayenJS/router"
)

func main() {
	config.Load()
	app := router.Setup()
	database.Setup()

	fmt.Println("Server is running on port", os.Getenv("PORT"))
	err := app.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

	if err != nil {
		panic(err)
	}
}
