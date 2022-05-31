package main

import (
	"golang-simple/app"
	"golang-simple/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
