package main

import (
	abc "example"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)
	log := abc.NewLogger(
		abc.SetLevel("debug"),
		abc.InGraylog("graylog:12201", "debug", "application_name"),
	)

	//log.Info("hello world")
	for {
		log.Info("hello world")
		time.Sleep(1 * time.Second)
	}
}
