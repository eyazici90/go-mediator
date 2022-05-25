package mediator

import "context"

type Mediator struct {
	pipeline *Pipeline
}

func New(opts ...Option) (*Mediator, error) {
	pipeline, err := newPipeline(opts...)
	if err != nil {
		return nil, err
	}

	m := &Mediator{
		pipeline: pipeline,
	}

	pipe := m.pipeline.behaviors.merge()
	m.pipeline.call = func(ctx context.Context, msg Message) error {
		return pipe(ctx, msg, func(ctx context.Context) error {
			return m.send(ctx, msg)
		})
	}
	return m, nil
}

func (m *Mediator) Send(ctx context.Context, req Message) error {
	if len(m.pipeline.behaviors) > 0 {
		return m.pipeline.call(ctx, req)
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
