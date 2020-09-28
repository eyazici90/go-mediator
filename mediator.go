package mediator

import (
	"context"
	"reflect"
)

type Mediator interface {
	initializer
	sender
	publisher
	pipelineBuilder
}

type reflectBasedMediator struct {
	behaviours   []func(context.Context, interface{}, Next) error
	pipeline     func(context.Context, interface{}) error
	handlers     map[reflect.Type]interface{}
	handlersFunc map[reflect.Type]reflect.Value
}

func New() Mediator {
	return &reflectBasedMediator{
		handlers:     make(map[reflect.Type]interface{}),
		handlersFunc: make(map[reflect.Type]reflect.Value),
		pipeline:     nil,
		behaviours:   nil,
	}
}
