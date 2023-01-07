package main

import (
	"golang-mux-gorm-boilerplate/app"
	"golang-mux-gorm-boilerplate/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
