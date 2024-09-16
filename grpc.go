package slogging

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log/slog"
)

func GRPCTraceMiddleware(logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			traceId := generateTraceId()
			ctx = ContextWithLogger(ctx, logger.With(StringAttr(xb3traceid, traceId)))
			ctx = context.WithValue(ctx, xb3traceid, traceId)
			return handler(ctx, req)
		}

		traceIds := md.Get(xb3traceid)
		var traceId string
		if len(traceIds) > 0 {
			traceId = traceIds[0]
		} else {
			traceId = generateTraceId()
		}

		ctx = ContextWithLogger(ctx, logger.With(StringAttr(xb3traceid, traceId)))
		ctx = context.WithValue(ctx, xb3traceid, traceId)
		return handler(ctx, req)
	}
}

// example
//func GRPCExampleUsage() {
//	log := NewLogger(
//		SetLevel("debug"),
//		InGraylog("graylog:12201", "debug", "application_name"),
//		SetDefault(true),
//	)
//
//	tracemiddleware := GRPCTraceMiddleware(log)
//
//	srv := grpc.NewServer(
//		grpc.UnaryInterceptor(tracemiddleware),
//	)
//}

func MetadataWithTraceHeaders(ctx context.Context) context.Context {
	traceId, ok := ctx.Value(xb3traceid).(string)
	if !ok || traceId == "" {
		traceId = generateTraceId()
	}

	md := metadata.Pairs(xb3traceid, traceId)
	return metadata.NewOutgoingContext(context.Background(), md)
}

// example
//func exampleMetadataSend(ctx context.Context) {
//	mdctx := MetadataWithTraceHeaders(ctx)
//	client.HelloWorld(ctx, req)
//}
