package mediator

import (
	"reflect"
)

type Mediator interface {
	initializer
	sender
	publisher
	pipelineBuilder
}

type mediator struct {
	behaviours   []Behaviour
	pipeline     Pipeline
	handlers     map[reflect.Type]interface{}
	handlersFunc map[reflect.Type]reflect.Value
}

func New() Mediator {
	return &mediator{
		handlers:     make(map[reflect.Type]interface{}),
		handlersFunc: make(map[reflect.Type]reflect.Value),
		pipeline:     nil,
		behaviours:   nil,
	}
}
