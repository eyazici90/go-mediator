package mediator

import "context"

const maxSize = 64

type Option func(m *Mediator) error

type (
	Handler interface {
		Handle(context.Context, Message) error
	}
	PipelineBehaviour interface {
		Process(context.Context, Message, Next) error
	}
	Message interface {
		Key() int
	}
)

type Mediator struct {
	pipe *pipeline
	call func(ctx context.Context, msg Message, next Next) error
}

func New(opts ...Option) (*Mediator, error) {
	p := pipeline{
		handlers: make([]Handler, maxSize),
	}
	m := &Mediator{
		pipe: &p,
	}

	for _, opt := range opts {
		if err := opt(m); err != nil {
			return nil, err
		}
	}

	call := m.pipe.bhs.merge()
	m.call = call

	return m, nil
}

func WithBehaviour(behavior PipelineBehaviour) Option {
	return func(m *Mediator) error {
		return m.pipe.useBehavior(behavior)
	}
}

func WithBehaviourFunc(fn func(context.Context, Message, Next) error) Option {
	return func(m *Mediator) error {
		return m.pipe.use(fn)
	}
}

func WithHandler(req Message, rh Handler) Option {
	return func(m *Mediator) error {
		return m.pipe.registerHandler(req, rh)
	}
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
