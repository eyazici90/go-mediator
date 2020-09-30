package mediator

func reverseApply(behaviours []Behaviour,
	action func(Behaviour)) {
	for i := len(behaviours) - 1; i >= 0; i-- {
		action(behaviours[i])
	}
}

func must(ok bool, handlerType string) {
	if !ok {
		panic("handle method does not exists for the typeOf" + handlerType)
	}
}
