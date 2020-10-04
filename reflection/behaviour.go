package reflection

import (
	"context"

	"github.com/eyazici90/go-mediator"
)

func (m *reflectBasedMediator) UseBehaviour(pipelineBehaviour mediator.PipelineBehaviour) Mediator {
	return m.Use(pipelineBehaviour.Process)
}

func (m *reflectBasedMediator) Use(call func(context.Context, interface{}, mediator.Next) error) Mediator {
	m.PipelineContext.Behaviours = append(m.PipelineContext.Behaviours, call)
	return m
}

func (m *reflectBasedMediator) Build() Mediator {
	mediator.ReverseApply(m.PipelineContext.Behaviours, m.pipe)
	return m
}

func (m *reflectBasedMediator) pipe(call mediator.Behaviour) {
	if m.PipelineContext.Pipeline == nil {
		m.PipelineContext.Pipeline = m.send
	}
	seed := m.PipelineContext.Pipeline

	m.PipelineContext.Pipeline = func(ctx context.Context, msg interface{}) error {
		return call(ctx, msg, func(context.Context) error { return seed(ctx, msg) })
	}
}
