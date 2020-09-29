package initializer

import (
	"../context"
	"../handler"
	"net/http"
)

type HandlerInitializer struct {
}

func NewHandlerInitializer() *HandlerInitializer {
	return &HandlerInitializer{}
}

func (initializer *HandlerInitializer) Init(context *context.Context) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handler.Ping)

	context.Components.Handler = mux
}
