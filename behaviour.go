package mediator

import "context"

func (m *mediator) UseBehaviour(pipelineBehaviour PipelineBehaviour) Mediator {
	return m.Use(pipelineBehaviour.Process)
}

func (m *mediator) Use(call func(context.Context, interface{}, Next) error) Mediator {
	m.PipelineContext.Behaviours = append(m.PipelineContext.Behaviours, call)
	return m
}

func (m *mediator) Build() Mediator {
	ReverseApply(m.PipelineContext.Behaviours, m.pipe)
	return m
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
