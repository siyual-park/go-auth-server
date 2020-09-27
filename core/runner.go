package core

type Runner interface {
	Run(*Context) error
}
