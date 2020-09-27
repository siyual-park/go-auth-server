package core

type App struct {
	context *Context
}

func NewApp() *App {
	return &App{context: NewContext()}
}

func (app *App) Use(initializer Initializer) {
	initializer.Init(app.context)
}

func (app *App) Run() error {
	return app.context.Components.Runner.Run(app.context)
}
