package mediator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMediator_should_dispath_msg_when_send(t *testing.T) {
	cmd := &fakeCommand{
		name: "Amsterdam",
	}
	handler := &fakeCommandHandler{}

	m, err := New().
		RegisterHandler(&fakeCommand{}, handler).
		Build()

	err = m.Send(context.Background(), cmd)

	assert.NoError(t, err)
	assert.Equal(t, cmd, handler.captured)
}

func TestMediator_should_execute_behaviour_when_send(t *testing.T) {

	var got Message
	behaviour := func(ctx context.Context, msg Message, next Next) error {
		got = msg
		next(ctx)

		return nil
	}

	cmd := &fakeCommand{
		name: "Amsterdam",
	}
	handler := &fakeCommandHandler{}

	m, err := New().
		Use(behaviour).
		RegisterHandler(&fakeCommand{}, handler).
		Build()

	err = m.Send(context.Background(), cmd)

	assert.NoError(t, err)
	assert.Equal(t, cmd, got)
}

type fakeCommand struct {
	name string
}

func (*fakeCommand) Key() string { return "fakeCommand" }

type fakeCommandHandler struct {
	captured Message
}

func (f *fakeCommandHandler) Handle(_ context.Context, msg Message) error {
	f.captured = msg
	return nil
}
