package logic

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
	"time"
)

type Logic struct {
	Tr trace.Tracer
}

func (l *Logic) Example(ctx context.Context) string {
	ctx, span := l.Tr.Start(ctx, "LogicExample", oteltrace.WithAttributes(attribute.String("id", "123")))
	defer span.End()

	defer func() {
		time.Sleep(100 * time.Millisecond)
	}()
	return l.MoreLogicLevels(ctx)
}

func (l *Logic) MoreLogicLevels(ctx context.Context) string {
	_, span := l.Tr.Start(ctx, "MoreLogicLevels", oteltrace.WithAttributes(attribute.String("hello", "world")))
	defer span.End()

	time.Sleep(50 * time.Millisecond)

	return "Hello, world!"
}
