package core

import "errors"

type EmptyRunner struct {
}

func NewEmptyRunner() *EmptyRunner {
	return &EmptyRunner{}
}

func (runner *EmptyRunner) Run(context *Context) error {
	return errors.New("Runner is empty")
}
