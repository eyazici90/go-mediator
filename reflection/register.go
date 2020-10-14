package reflection

import (
	"reflect"

	"github.com/eyazici90/go-mediator/internal/util"
)

const handleMethodName string = "Handle"

func (m *reflectBasedMediator) RegisterHandlers(handlers ...interface{}) Mediator {
	for _, handler := range handlers {
		m.RegisterHandler(handler)
	}
	return m
}

func (m *reflectBasedMediator) RegisterHandler(handler interface{}) Mediator {
	handlerType := reflect.TypeOf(handler)

	method, ok := handlerType.MethodByName(handleMethodName)

	util.Must(ok, handlerType.String())

	requestType := reflect.TypeOf(method.Func.Interface()).In(2)

	m.handlers[requestType] = handler
	m.handlersFunc[requestType] = method.Func
	return m
}
