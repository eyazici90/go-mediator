package mediator

import "context"

type Mediator struct {
	pipe *Pipeline
	call func(ctx context.Context, msg Message, next Next) error
}

func New(opts ...Option) (*Mediator, error) {
	pipeline, err := newPipeline(opts...)
	if err != nil {
		return nil, err
	}

	m := &Mediator{
		pipe: pipeline,
	}

	call := m.pipe.bhs.merge()
	m.call = call

	return m, nil
}

func (m *Mediator) Send(ctx context.Context, msg Message) error {
	if len(m.pipe.bhs) > 0 {
		return m.call(ctx, msg, func(ctx context.Context) error {
			return m.send(ctx, msg)
		})
	}
	return m.send(ctx, msg)
}

func (m *Mediator) send(ctx context.Context, msg Message) error {
	key := msg.Key()
	handler, err := m.pipe.findHandler(key)
	if err != nil {
		return err
	}
	return handler.Handle(ctx, msg)
}
