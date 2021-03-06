


# go-mediator
Simple mediator implementation in golang. <br/>
In-process messaging.

## Commands

     type  CreateOrderCommand  struct { 
		    Id string  `validate:"required,min=10"` 
    }
    
    func (CreateOrderCommand) Key() string { return "CreateOrderCommand" }

    type  CreateOrderCommandHandler  struct {  
    }
    
    func  NewCreateOrderCommandHandler() CreateOrderCommandHandler {
	    return CreateOrderCommandHandler{}
    }
     
    func (CreateOrderCommandHandler) Handle(ctx context.Context, msg mediator.Message) error {
    
	    //Do something
	     return nil
    
    }
     
  
    
## Behaviours 

***PipelineBehaviour interface implementation usage***

    type  Logger  struct{}
    
    func  NewLogger() *Logger { return &Logger{} }
    
    
    func (l *Logger) Process(ctx context.Context, msg mediator.Message, next mediator.Next) error { 
    
		    log.Println("Pre Process!")
		    
		    result := next(ctx)
		    
		    log.Println("Post Process")
		    
		    return result
    }
    
    m := mediator.NewContext().UseBehaviour(behaviour.NewLogger()).Build()

***Func based usage***

    m := mediator.NewContext().Use(func(ctx context.Context, msg mediator.Message, next mediator.Next) error {
				    
					    log.Println("Pre Process!")
					    
					    next(ctx)
					    
					    log.Println("Post Process") 
					    
					    return  nil
				    
				    }).Build()
       

## Usages

    m := mediator.NewContext().  
		  UseBehaviour(behaviour.NewLogger()). 
		  UseBehaviour(behaviour.NewValidator()). 
		  RegisterHandlers(FakeCommand{}, NewFakeCommandCommandHandler(r)). 
		  Build()

    cmd := FakeCommand{
	    Name: "Amsterdam", 
    }
    
    ctx := context.Background()
     
    m.Send(ctx, cmd)
    

***Func based usage***

    m := mediator.NewContext().Use(func(ctx context.Context, cmd mediator.Message, next mediator.Next) error {
				    
					    log.Println("Pre Process!")
					    
					    next(ctx)
					    
					    log.Println("Post Process") 
					    
					    return  nil
				    
				    }).Build()
       
## Examples
[Simple](https://github.com/eyazici90/go-mediator/tree/master/_examples)

TBD...
