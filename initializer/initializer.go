package initializer

import (
	"../context"
)

type Initializer interface {
	Init(ctx *context.Context)
}
