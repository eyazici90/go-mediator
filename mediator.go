package mediator

import (
	"reflect"
)

type mediator struct {
	PipelineContext
	handlers map[reflect.Type]RequestHandler
}

func newMediator(pipelineContext PipelineContext, handlers map[reflect.Type]RequestHandler) *mediator {
	return &mediator{
		handlers:        handlers,
		PipelineContext: pipelineContext,
	}
}
