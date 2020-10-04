package mediator

import (
	"reflect"
	"strings"
)

const handlerNamePrefix string = `Handler`

func (m *mediator) RegisterHandlers(handlers ...RequestHandler) Mediator {
	for _, handler := range handlers {
		m.RegisterHandler(handler)
	}
	return m
}

func (m *mediator) RegisterHandler(handler RequestHandler) Mediator {
	handlerName := reflect.TypeOf(handler).Name()
	requestType := strings.ReplaceAll(handlerName, handlerNamePrefix, "")

	m.handlers[requestType] = handler
	return m
}
