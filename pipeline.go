package mediator

import "context"

type Behaviour func(context.Context, Message, Next) error

type Next func(ctx context.Context) error

type Pipeline func(context.Context, Message) error

type PipelineContext struct {
	Behaviours []Behaviour
	Pipeline   Pipeline
}

func NewPipelineContext() PipelineContext { return PipelineContext{} }
