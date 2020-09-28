package mediator

import (
	"context"
	"reflect"
)

type sender interface {
	Send(context.Context, interface{}) error
}

func (m *reflectBasedMediator) Send(ctx context.Context, msg interface{}) error {
	if m.pipeline != nil {
		return m.pipeline(ctx, msg)
	}
	return m.send(ctx, msg)
}

func (m *reflectBasedMediator) send(ctx context.Context, msg interface{}) error {
	msgType := reflect.TypeOf(msg)
	handler, ok := m.handlers[msgType]
	if !ok {
		return HandlerNotFound
	}
	handlerFunc, _ := m.handlersFunc[msgType]
	return call(handler, ctx, handlerFunc, msg)
}
