package app

import (
	"../context"
	"../initializer"
)

type App struct {
	context *context.Context
}

func NewApp() *App {
	return &App{context: context.NewContext()}
}

func (app *App) Use(initializer initializer.Initializer) {
	initializer.Init(app.context)
}

func (app *App) Run() error {
	return app.context.Components.Runner.Run(app.context)
}
