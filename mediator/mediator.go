package mediator

import "context"

var _ Sender = &Mediator{}

type Mediator struct {
	context *PipelineContext
}

func New(opts ...Option) (*Mediator, error) {
	pCtx, err := newPipelineContext(opts...)
	if err != nil {
		return nil, err
	}
	m := &Mediator{
		context: pCtx,
	}

	pCtx.behaviors.reverseApply(m.pipe)
	return m, nil
}

func (m *Mediator) Send(ctx context.Context, req Message) error {
	if m.context.pipeline.empty() {
		return m.send(ctx, req)
	}
	return m.context.pipeline(ctx, req)
}

func (m *Mediator) send(ctx context.Context, req Message) error {
	key := req.Key()
	handler, err := m.context.findHandler(key)
	if err != nil {
		return err
	}
	return handler.Handle(ctx, req)
}

func (m *Mediator) pipe(call Behavior) {
	if m.context.pipeline.empty() {
		m.context.pipeline = m.send
	}
	seed := m.context.pipeline

	m.context.pipeline = func(ctx context.Context, msg Message) error {
		return call(ctx, msg, func(context.Context) error { return seed(ctx, msg) })
	}
}
