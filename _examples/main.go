package main

import (
	"context"
	"errors"
	"log"

	"github.com/eyazici90/go-mediator/internal/must"
	"github.com/eyazici90/go-mediator/mediator"
)

func main() {
	m, err := mediator.New(
		mediator.WithBehaviourFunc(
			func(ctx context.Context, cmd mediator.Message, next mediator.Next) error {
				log.Println("Pre Process - 1!")
				_ = next(ctx)
				log.Println("Post Process - 1")

				return nil
			}), mediator.WithBehaviourFunc(
			func(ctx context.Context, cmd mediator.Message, next mediator.Next) error {
				log.Println("Pre Process!- 2")
				_ = next(ctx)
				log.Println("Post Process - 2")

				return nil
			}),
		mediator.WithHandler(&FakeCommand{}, NewFakeCommandHandler()))

	must.NotFail(err)

	cmd := &FakeCommand{
		Name: "Emre",
	}
	ctx := context.Background()

	_ = m.Send(ctx, cmd)
}

type FakeCommand struct {
	Name string
}

func (*FakeCommand) Key() int { return 1 }

type FakeCommandHandler struct{}

func NewFakeCommandHandler() FakeCommandHandler {
	return FakeCommandHandler{}
}

func (FakeCommandHandler) Handle(_ context.Context, command mediator.Message) error {
	cmd := command.(*FakeCommand)
	if cmd.Name == "" {
		return errors.New("name is empty")
	}
	return nil
}
