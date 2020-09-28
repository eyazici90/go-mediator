package mediator

import "reflect"

type initializer interface {
	RegisterHandler(handler interface{}) Mediator
	RegisterHandlers(handlers ...interface{}) Mediator
}

func (m *reflectBasedMediator) RegisterHandlers(handlers ...interface{}) Mediator {
	for _, handler := range handlers {
		m.RegisterHandler(handler)
	}
	return m
}

func (m *reflectBasedMediator) RegisterHandler(handler interface{}) Mediator {
	handlerType := reflect.TypeOf(handler)

	method, ok := handlerType.MethodByName(handleMethodName)
	must(handlerType.String(), ok)

	cType := reflect.TypeOf(method.Func.Interface()).In(2)

	m.handlers[cType] = handler
	m.handlersFunc[cType] = method.Func
	return m
}

func must(desc string, ok bool) {
	if !ok {
		panic("handle method does not exists for the typeOf" + desc)
	}
}
