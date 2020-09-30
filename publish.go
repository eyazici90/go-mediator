package mediator

type publisher interface {
	Publish(msg interface{})
}

func (m *mediator) Publish(msg interface{}) {
	//
}
