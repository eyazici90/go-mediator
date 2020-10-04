package mediator

import (
	"reflect"
	"strings"
)

const handlerName = "Handler"

func (m *mediator) RegisterHandlers(handlers ...RequestHandler) Mediator {
	for _, handler := range handlers {
		m.RegisterHandler(handler)
	}
	return m
}

func (m *mediator) RegisterHandler(handler RequestHandler) Mediator {
	handlerName := reflect.TypeOf(handler).Name()
	requestType := strings.ReplaceAll(handlerName, handlerName, "")

	m.handlers[requestType] = handler
	return m
}
