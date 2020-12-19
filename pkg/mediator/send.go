package mediator

import (
	"context"
)

func (m *mediator) Send(ctx context.Context, request Message) error {
	if m.PipelineContext.Pipeline != nil {
		return m.PipelineContext.Pipeline(ctx, request)
	}
	return m.send(ctx, request)
}

func (m *mediator) send(ctx context.Context, request Message) error {
	key := request.Key()
	handler, ok := m.handlers[key]
	if !ok {
		return ErrHandlerNotFound
	}
	return handler.Handle(ctx, request)
}

func (m *mediator) pipe(call Behaviour) {
	if m.PipelineContext.Pipeline == nil {
		m.PipelineContext.Pipeline = m.send
	}
	seed := m.PipelineContext.Pipeline

	m.PipelineContext.Pipeline = func(ctx context.Context, msg Message) error {
		return call(ctx, msg, func(context.Context) error { return seed(ctx, msg) })
	}
}
