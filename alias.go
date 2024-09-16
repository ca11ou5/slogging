package slogging

import (
	"golang.org/x/exp/constraints"
	"log/slog"
	"time"
)

const (
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
)

type (
	Logger         = slog.Logger
	Level          = slog.Level
	Record         = slog.Record
	Handler        = slog.Handler
	Attr           = slog.Attr
	HandlerOptions = slog.HandlerOptions
)

var (
	New            = slog.New
	NewJSONHandler = slog.NewJSONHandler
	NewTextHandler = slog.NewTextHandler
)

func IntAttr[T constraints.Integer](key string, value T) Attr {
	return slog.Int(key, int(value))
}

func FloatAttr[T constraints.Float](key string, value T) Attr {
	return slog.Float64(key, float64(value))
}

func TimeAttr(key string, time time.Time) Attr {
	return slog.String(key, time.String())
}

func ErrAttr(err error) Attr {
	return slog.String("error", err.Error())
}

func StringAttr(key string, value string) Attr {
	return slog.String(key, value)
}

func AnyAttr(key string, s interface{}) Attr {
	return slog.Any(key, s)
}
