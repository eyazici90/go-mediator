package mediator

import "context"

func reverseApply(behaviours []func(context.Context, interface{}, Next) error,
	action func(func(context.Context, interface{}, Next) error)) {
	for i := len(behaviours) - 1; i >= 0; i-- {
		action(behaviours[i])
	}
}
