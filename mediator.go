package mediator

type mediator struct {
	PipelineContext
	handlers map[string]RequestHandler
}

func New() Mediator {
	return &mediator{
		handlers:        make(map[string]RequestHandler),
		PipelineContext: NewPipelineContext(),
	}
}
