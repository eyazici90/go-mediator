package reflection

import (
	"reflect"

	"github.com/eyazici90/go-mediator"
)

type reflectBasedMediator struct {
	mediator.PipelineContext
	handlers     map[reflect.Type]interface{}
	handlersFunc map[reflect.Type]reflect.Value
}

func New() Mediator {
	return &reflectBasedMediator{
		handlers:        make(map[reflect.Type]interface{}),
		handlersFunc:    make(map[reflect.Type]reflect.Value),
		PipelineContext: mediator.NewPipelineContext(),
	}
}
