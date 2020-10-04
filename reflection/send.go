package reflection

import (
	"context"
	"reflect"

	"github.com/eyazici90/go-mediator"
)

func (m *reflectBasedMediator) Send(ctx context.Context, msg interface{}) error {
	if m.PipelineContext.Pipeline != nil {
		return m.PipelineContext.Pipeline(ctx, msg)
	}
	return m.send(ctx, msg)
}

func (m *reflectBasedMediator) send(ctx context.Context, msg interface{}) error {
	msgType := reflect.TypeOf(msg)
	handler, ok := m.handlers[msgType]
	if !ok {
		return mediator.HandlerNotFound
	}
	handlerFunc, _ := m.handlersFunc[msgType]
	return call(handler, ctx, handlerFunc, msg)
}
