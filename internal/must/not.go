package must

func NotFalse(ok bool, handlerType string) {
	if !ok {
		panic("handle method does not exists for the typeOf" + handlerType)
	}
}
