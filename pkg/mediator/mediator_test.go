package mediator_test

import (
	"context"
	"testing"

	"github.com/eyazici90/go-mediator/pkg/mediator"
	"github.com/stretchr/testify/assert"
)

func TestMediator_should_dispatch_msg_when_send(t *testing.T) {
	cmd := &fakeCommand{
		name: "Amsterdam",
	}
	handler := &fakeCommandHandler{}

	m, _ := mediator.New(
		mediator.WithHandler(&fakeCommand{}, handler),
	)

	err := m.Send(context.Background(), cmd)

	assert.NoError(t, err)
	assert.Equal(t, cmd, handler.captured)
}

func TestMediator_should_execute_behavior_when_send(t *testing.T) {
	var got mediator.Message
	behavior := func(ctx context.Context, msg mediator.Message, next mediator.Next) error {
		got = msg
		return next(ctx)
	}

	cmd := &fakeCommand{
		name: "Amsterdam",
	}
	handler := &fakeCommandHandler{}

	m, _ := mediator.New(
		mediator.WithBehaviourFunc(behavior),
		mediator.WithHandler(&fakeCommand{}, handler),
	)

	err := m.Send(context.Background(), cmd)

	assert.NoError(t, err)
	assert.Equal(t, cmd, got)
}

type fakeCommand struct {
	name string
}

func (*fakeCommand) Key() string { return "fakeCommand" }

type fakeCommandHandler struct {
	captured mediator.Message
}

func (f *fakeCommandHandler) Handle(_ context.Context, msg mediator.Message) error {
	f.captured = msg
	return nil
}
