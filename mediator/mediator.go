package mediator

import "context"

type Mediator struct {
	pipeline *Pipeline
	call     func(ctx context.Context, msg Message, next Next) error
}

func New(opts ...Option) (*Mediator, error) {
	pipeline, err := newPipeline(opts...)
	if err != nil {
		return nil, err
	}

	m := &Mediator{
		pipeline: pipeline,
	}

	call := m.pipeline.behaviors.merge()
	m.call = call

	return m, nil
}

func (m *Mediator) Send(ctx context.Context, req Message) error {
	if len(m.pipeline.behaviors) > 0 {
		return m.call(ctx, req, func(ctx context.Context) error {
			return m.send(ctx, req)
		})
	}
	return m.send(ctx, req)
}

func (m *Mediator) send(ctx context.Context, req Message) error {
	key := req.Key()
	handler, err := m.pipeline.findHandler(key)
	if err != nil {
		return err
	}
	return handler.Handle(ctx, req)
}
