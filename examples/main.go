package main

import (
	"context"
	"log"
	"mediator"
)

func main() {

	m := mediator.NewMediator().
		Use(func(ctx context.Context, cmd interface{}, next mediator.Next) error {
			log.Println("Pre Process!")
			next(ctx)
			log.Println("Post Process")

			return nil
		}).
		RegisterHandler(NewFakeCommandHandler())

	cmd := FakeCommand{
		Name: "Emre",
	}
	ctx := context.Background()
	m.Send(ctx, cmd)

}

type FakeCommand struct {
	Name string
}

type FakeCommandHandler struct{}

func NewFakeCommandHandler() FakeCommandHandler {
	return FakeCommandHandler{}
}

func (handler FakeCommandHandler) Handle(_ context.Context, cmd FakeCommand) error {
	log.Println(cmd.Name)
	return nil
}
