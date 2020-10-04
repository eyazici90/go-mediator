package main

import (
	"context"
	"errors"
	"log"

	"github.com/eyazici90/go-mediator"
	"github.com/eyazici90/go-mediator/reflection"
)

func main() {

	m := reflection.New().
		Use(func(ctx context.Context, cmd interface{}, next mediator.Next) error {
			log.Println("Pre Process - 1!")
			next(ctx)
			log.Println("Post Process - 1")

			return nil
		}).
		Use(func(ctx context.Context, cmd interface{}, next mediator.Next) error {
			log.Println("Pre Process!- 2")
			next(ctx)
			log.Println("Post Process - 2")

			return nil
		}).
		RegisterHandler(NewFakeCommandHandler()).
		Build()

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
	if cmd.Name == "" {
		return errors.New("Name is empty")
	}
	return nil
}
