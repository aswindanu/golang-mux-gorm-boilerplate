package main

import (
	"go-todo-rest-api-example/app"
	"go-todo-rest-api-example/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
