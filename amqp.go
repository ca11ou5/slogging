package slogging

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"log/slog"
)

type AMQPMiddlewareFunc func(amqp091.Delivery) (context.Context, amqp091.Delivery)

func AMQPTraceMiddleware(logger *slog.Logger) AMQPMiddlewareFunc {
	return func(msg amqp091.Delivery) (context.Context, amqp091.Delivery) {
		traceId, ok := msg.Headers[xb3traceid].(string)
		if !ok || traceId == "" {
			traceId = generateTraceId()
		}

		ctx := ContextWithLogger(context.Background(), logger.With(StringAttr(xb3traceid, traceId)))
		ctx = context.WithValue(ctx, "traceId", traceId)

		return ctx, msg
	}
}

// example
//func ExampleAMQPTracing() {
//	log := NewLogger(
//		SetLevel("debug"),
//		InGraylog("graylog:12201", "debug", "application_name"),
//		SetDefault(true),
//	)
//
//	msgs := make(<-chan amqp091.Delivery)
//	amqpTraceMiddleware := AMQPTraceMiddleware(log)
//
//	go func() {
//		for msg := range msgs {
//			ProcessMessage(amqpTraceMiddleware(msg))
//		}
//	}()
//
//}
//
//func ProcessMessage(ctx context.Context, msg amqp091.Delivery) bool {
//	// SOME LOGIC
//	return true
//}

func AMQPTableWithTraceHeaders(ctx context.Context, table amqp091.Table) amqp091.Table {
	traceId, ok := ctx.Value("traceId").(string)
	if !ok || traceId == "" {
		traceId = generateTraceId()
	}

	table[xb3traceid] = traceId
	return table
}
