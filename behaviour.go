package mediator

import "context"

type PipelineBehaviour interface {
	Process(context.Context, interface{}, Next) error
}

type pipelineBuilder interface {
	Build() Mediator
	UseBehaviour(PipelineBehaviour) Mediator
	Use(call func(context.Context, interface{}, Next) error) Mediator
}

func (m *reflectBasedMediator) UseBehaviour(pipelineBehaviour PipelineBehaviour) Mediator {
	return m.Use(pipelineBehaviour.Process)
}

func (m *reflectBasedMediator) Use(call func(context.Context, interface{}, Next) error) Mediator {
	m.behaviours = append(m.behaviours, call)
	return m
}

func (m *reflectBasedMediator) Build() Mediator {
	reverseApply(m.behaviours, m.pipe)
	return m
}

func (m *reflectBasedMediator) pipe(call func(context.Context, interface{}, Next) error) {
	if m.pipeline == nil {
		m.pipeline = m.send
	}
	seed := m.pipeline

	m.pipeline = func(ctx context.Context, msg interface{}) error {
		return call(ctx, msg, func(context.Context) error { return seed(ctx, msg) })
	}
}
