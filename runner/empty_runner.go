package runner

import (
	"../context"
	"errors"
)

type EmptyRunner struct {
}

func NewEmptyRunner() *EmptyRunner {
	return &EmptyRunner{}
}

func (runner *EmptyRunner) Run(context *context.Context) error {
	return errors.New("Runner is empty")
}
