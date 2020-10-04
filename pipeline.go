package mediator

import "context"

type PipelineBehaviour interface {
	Process(context.Context, interface{}, Next) error
}

type Behaviour func(context.Context, interface{}, Next) error

type Next func(ctx context.Context) error

type Pipeline func(context.Context, interface{}) error

type PipelineContext struct {
	Behaviours []Behaviour
	Pipeline   Pipeline
}

func NewPipelineContext() PipelineContext { return PipelineContext{} }
