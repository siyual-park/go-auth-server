package core

type Context struct {
	Components *Components
}

func NewContext() *Context {
	return &Context{Components: NewComponents()}
}
