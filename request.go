package mediator

import "context"

type RequestHandler interface {
	Handle(context.Context, interface{}) error
}
