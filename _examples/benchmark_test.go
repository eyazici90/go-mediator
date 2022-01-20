package main

import (
	"context"
	"testing"

	"github.com/eyazici90/go-mediator/mediator"
)

func BenchmarkMediator(b *testing.B) {
	m, _ := mediator.New(mediator.WithHandler(&FakeCommand{}, NewFakeCommandHandler()))

	cmd := &FakeCommand{Name: "Emre"}
	ctx := context.TODO()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Send(ctx, cmd)
	}
}

func BenchmarkHandler(b *testing.B) {
	handler := NewFakeCommandHandler()

	cmd := &FakeCommand{Name: "Emre"}
	ctx := context.TODO()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = handler.Handle(ctx, cmd)
	}
}
