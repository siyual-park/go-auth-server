package main

import (
	"./app"
	"./initializer"
	"log"
)

func main() {
	handlerInitializer := initializer.NewHandlerInitializer()
	runnerInitializer := initializer.NewServerRunnerInitializer(8080)

	app := app.NewApp()
	app.Use(handlerInitializer)
	app.Use(runnerInitializer)

	err := app.Run()
	if err != nil {
		log.Fatal("Listen server", err)
	}
}
