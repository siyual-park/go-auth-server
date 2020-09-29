package initializer

import (
	"../context"
	"../runner"
)

type ServerRunnerInitializer struct {
	Port int
}

func NewServerRunnerInitializer(port int) *ServerRunnerInitializer {
	return &ServerRunnerInitializer{Port: port}
}

func (initializer *ServerRunnerInitializer) Init(ctx *context.Context) {
	ctx.Components.Runner = runner.NewServerRunner(initializer.Port)
}
