package slogging

import (
	"context"
	"log/slog"
	"net/http"
)

type HTTPMiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func HTTPTraceMiddleware(logger *slog.Logger) HTTPMiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			traceId := r.Header.Get(xb3traceid)
			if traceId == "" {
				traceId = generateTraceId()
			}

			ctx := ContextWithLogger(r.Context(), logger.With(StringAttr(xb3traceid, traceId)))
			ctx = context.WithValue(ctx, xb3traceid, traceId)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}

func RequestWithTraceHeaders(ctx context.Context, req *http.Request) *http.Request {
	traceId := ctx.Value(xb3traceid).(string)

	req.Header.Set(xb3traceid, traceId)
	return req
}
