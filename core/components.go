package core

type Components struct {
	Runner Runner
}

func NewComponents() *Components {
	return &Components{Runner: NewEmptyRunner()}
}
