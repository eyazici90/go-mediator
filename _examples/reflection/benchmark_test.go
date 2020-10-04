package main

import (
	"context"
	"testing"

	"github.com/eyazici90/go-mediator/reflection"
)

func BenchmarkReflectMediator(b *testing.B) {
	m := reflection.New().RegisterHandler(NewFakeCommandHandler()).Build()

	cmd := FakeCommand{Name: "Emre"}
	ctx := context.TODO()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Send(ctx, cmd)
	}
}
func BenchmarkHandler(b *testing.B) {
	handler := NewFakeCommandHandler()
	cmd := FakeCommand{Name: "Emre"}
	ctx := context.TODO()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler.Handle(ctx, cmd)
	}
}
