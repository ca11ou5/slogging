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
	test := &TestStruct{
		A: "sadness",
		B: 1337,
		C: true,
	}

	for {
		log.Info("hello world",
			abc.StructAttr("application", test))
		time.Sleep(1 * time.Second)
	}
}

type TestStruct struct {
	A string
	B int
	C bool
}
