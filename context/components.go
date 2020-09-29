package context

import (
	"../runner"
	"net/http"
)

type Components struct {
	Runner  runner.Runner
	Handler http.Handler
}

func NewComponents() *Components {
	return &Components{Runner: runner.NewEmptyRunner()}
}
