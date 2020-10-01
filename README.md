

# go-mediator
Simple mediator implementation in golang. <br/>
In-process messaging.

## Commands

     type  CreateOrderCommand  struct { 
		    Id string  `validate:"required,min=10"` 
    }
 
    type  CreateOrderCommandHandler  struct {  
    }
     
    
    func  NewCreateOrderCommandHandler() CreateOrderCommandHandler {
	    return CreateOrderCommandHandler{}
    }
     
    func (handler CreateOrderCommandHandler) Handle(ctx context.Context, cmd CreateOrderCommand) error {
    
	    //Do something
	     return nil
    
    }
    
## Behaviours 

***PipelineBehaviour interface implementation usage***

    type  Logger  struct{}
    
    func  NewLogger() *Logger { return &Logger{} }
    
    
    func (l *Logger) Process(ctx context.Context, cmd interface{}, next mediator.Next) error { 
    
		    log.Println("Pre Process!")
		    
		    result := next(ctx)
		    
		    log.Println("Post Process")
		    
		    return result
    }
    
    m := mediator.New().UseBehaviour(behaviour.NewLogger()).Build()

***Func based usage***

    m := mediator.New().Use(func(ctx context.Context, cmd interface{}, next mediator.Next) error {
				    
					    log.Println("Pre Process!")
					    
					    next(ctx)
					    
					    log.Println("Post Process") 
					    
					    return  nil
				    
				    }).Build()
       

## Usages

    m := mediator.New(). 
			    
			    UseBehaviour(behaviour.NewLogger()).
			    
			    UseBehaviour(behaviour.NewValidator()).
			    
			    RegisterHandlers(command.NewFakeCommandCommandHandler(r)).
			    
			    Build()

    cmd := FakeCommand{
	    Name: "Amsterdam", 
    }
    
    ctx := context.Background()
     
    m.Send(ctx, cmd)
    
## Examples
[Simple](https://github.com/eyazici90/go-mediator/tree/master/examples)

TBD...
