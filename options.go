package slogging

import (
	"github.com/Graylog2/go-gelf/gelf"
	"log"
)

func InGraylog(graylogURL, level, containerName string) LoggerOption {
	w, err := gelf.NewWriter(graylogURL)
	if err != nil {
		log.Fatal(err)
	}

	var l Level
	err = l.UnmarshalText([]byte(level))
	if err != nil {
		log.Fatal(err)
	}

	return func(o *LoggerConfig) {
		o.InGraylog = &gelfData{
			w:             w,
			level:         l,
			containerName: containerName,
		}
	}
}

func SetLevel(l string) LoggerOption {
	return func(o *LoggerConfig) {
		o.Level.UnmarshalText([]byte(l))
	}
}

func WithSource(with bool) LoggerOption {
	return func(o *LoggerConfig) {
		o.WithSource = with
	}
}

func SetJSONFormat(set bool) LoggerOption {
	return func(o *LoggerConfig) {
		o.IsJSON = set
	}
}

func SetDefault(set bool) LoggerOption {
	return func(o *LoggerConfig) {
		o.SetDefault = set
	}
}
