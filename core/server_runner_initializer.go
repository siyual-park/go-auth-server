package core

import (
	"net/http"
)

type ServerRunnerInitializer struct {
	Port    int
	Handler http.Handler
}

func NewServerRunnerInitializer(port int, handler http.Handler) *ServerRunnerInitializer {
	return &ServerRunnerInitializer{Port: port, Handler: handler}
}

func (initializer *ServerRunnerInitializer) Init(context *Context) {
	context.Components.Runner = NewServerRunner(initializer.Port, initializer.Handler)
}
