package core

import (
	"fmt"
	"net/http"
)

type ServerRunner struct {
	Runner
	Port    int
	Handler http.Handler
}

func NewServerRunner(port int, handler http.Handler) *ServerRunner {
	return &ServerRunner{Port: port, Handler: handler}
}

func (runner *ServerRunner) Run(context *Context) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", runner.Port), runner.Handler)
}
