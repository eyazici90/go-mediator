package must

func NotFalse(ok bool, handlerType string) {
	if !ok {
		panic("handle method does not exists for the typeOf" + handlerType)
	}
}

func NotFail(err error) {
	if err != nil {
		panic(err)
	}
}
