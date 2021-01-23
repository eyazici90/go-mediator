package mediator

import "context"

type mediator struct {
	PipelineContext
	handlers map[string]RequestHandler
}

func newMediator(pipelineContext PipelineContext, handlers map[string]RequestHandler) *mediator {
	return &mediator{
		handlers:        handlers,
		PipelineContext: pipelineContext,
	}
}

func (m *mediator) Send(ctx context.Context, req Message) error {
	if m.PipelineContext.Pipeline != nil {
		return m.PipelineContext.Pipeline(ctx, req)
	}
	return m.send(ctx, req)
}

func (m *mediator) send(ctx context.Context, req Message) error {
	key := req.Key()
	handler, ok := m.handlers[key]
	if !ok {
		return ErrHandlerNotFound
	}
	return handler.Handle(ctx, req)
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

func (m *mediator) Publish(msg Message) {
	//
}
