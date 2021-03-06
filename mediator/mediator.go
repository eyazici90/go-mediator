package mediator

import "context"

type Mediator struct {
	context PipelineContext
}

func newMediator(ctx PipelineContext) *Mediator {
	return &Mediator{
		context: ctx,
	}
}

func (m *Mediator) Send(ctx context.Context, req Message) error {
	if m.context.pipeline != nil {
		return m.context.pipeline(ctx, req)
	}
	return m.send(ctx, req)
}

func (m *Mediator) send(ctx context.Context, req Message) error {
	key := req.Key()
	handler, ok := m.context.handlers[key]
	if !ok {
		return ErrHandlerNotFound
	}
	return handler.Handle(ctx, req)
}

func (m *Mediator) pipe(call Behaviour) {
	if m.context.pipeline == nil {
		m.context.pipeline = m.send
	}
	seed := m.context.pipeline

	m.context.pipeline = func(ctx context.Context, msg Message) error {
		return call(ctx, msg, func(context.Context) error { return seed(ctx, msg) })
	}
}

func (m *Mediator) Publish(msg Message) {
	//
}
