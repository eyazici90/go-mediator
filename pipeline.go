package mediator

import "context"

type Pipeline func(context.Context, interface{}) error
