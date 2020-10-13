package mediator

import (
	"context"
	"reflect"
)

func (m *mediator) Send(ctx context.Context, msg interface{}) error {
	if m.PipelineContext.Pipeline != nil {
		return m.PipelineContext.Pipeline(ctx, msg)
	}
	return m.send(ctx, msg)
}

func (m *mediator) send(ctx context.Context, request interface{}) error {
	requestType := reflect.TypeOf(request)
	handler, ok := m.handlers[requestType]
	if !ok {
		return HandlerNotFound
	}
	return handler.Handle(ctx, request)
}

func (m *mediator) pipe(call Behaviour) {
	if m.PipelineContext.Pipeline == nil {
		m.PipelineContext.Pipeline = m.send
	}
	seed := m.PipelineContext.Pipeline

	m.PipelineContext.Pipeline = func(ctx context.Context, msg interface{}) error {
		return call(ctx, msg, func(context.Context) error { return seed(ctx, msg) })
	}
}
