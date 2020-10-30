package mediator

type mediator struct {
	PipelineContext
	handlers map[string]RequestHandler
}

func newMediator(pipelineContext PipelineContext, handlers map[string]RequestHandler) *mediator {
	return &mediator{
		handlers:        handlers,
		PipelineContext: pipelineContext,
	}
}
