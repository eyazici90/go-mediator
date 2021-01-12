package mediator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild_should_return_mediator(t *testing.T) {

	m, err := New().Use(func(context.Context, Message, Next) error {
		return nil
	}).Build()

	assert.NoError(t, err)
	assert.NotNil(t, m)
}
