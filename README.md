
# go-mediator

Simple mediator implementation in go. <br/>

In-process messaging with behaviours.
  

## Commands

      
    
    type CreateOrderCommand struct {
    
    Id string `validate:"required,min=10"`
    
    }
    
    func (CreateOrderCommand) Key() string { return "CreateOrderCommand" }
    
      
    
    type CreateOrderCommandHandler struct {
    
    }
    
    func NewCreateOrderCommandHandler() CreateOrderCommandHandler {
    
    return CreateOrderCommandHandler{}
    
    }
    
    func (CreateOrderCommandHandler) Handle(ctx context.Context, msg mediator.Message) error {
    
    //Do something
    
    return nil
    
    }

## Behaviours

  

***PipelineBehaviour interface implementation usage***

  

    type Logger struct{}
    
    func NewLogger() *Logger { return &Logger{} }
    
    func (l *Logger) Process(ctx context.Context, msg mediator.Message, next mediator.Next) error {
    
    log.Println("Pre Process!")
    
    result := next(ctx)
    
    log.Println("Post Process")
    
    return result
    
    }
    
    m, err := mediator.New(mediator.WithBehaviour(behaviour.NewLogger()))

  

***Func based usage***

    m, err := mediator.New(mediator.WithBehaviourFunc(func(ctx context.Context, msg mediator.Message, next mediator.Next) error {
    
    log.Println("Pre Process!")
    
    next(ctx)
    
    log.Println("Post Process")
    
    return  nil
    
    }))

  

## Usages

  

    m, err := mediator.New(
    mediator.WithBehaviour(behaviour.NewLogger()),
    mediator.WithBehaviour(behaviour.NewValidator()),
    mediator.WithHandler(FakeCommand{}, NewFakeCommandCommandHandler(r)),
    )
    
    cmd := FakeCommand{   
    Name: "Amsterdam",  
    }
    
    ctx := context.Background()
    
    err := m.Send(ctx, cmd)

  

***Func based usage***

    m, err  := mediator.New(mediator.WithBehaviourFunc(func(ctx context.Context, cmd mediator.Message, next mediator.Next) error {
    log.Println("Pre Process - 1!")
    
    next(ctx)
    
    log.Println("Post Process - 1")
    return  nil
    }))

## Examples

[Simple](https://github.com/eyazici90/go-mediator/tree/master/_examples)

  

TBD...
