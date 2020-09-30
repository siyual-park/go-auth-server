package runner

import (
	"../context"
	"fmt"
	"net/http"
)

type ServerRunner struct {
	Port int
}

func NewServerRunner(port int) *ServerRunner {
	return &ServerRunner{Port: port}
}

func (runner *ServerRunner) Run(ctx *context.Context) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", runner.Port), ctx.Components.Handler)
}
