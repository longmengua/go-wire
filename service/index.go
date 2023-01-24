package service

import (
	"fmt"
	"wire-demo/dto"
)

type Greeter struct {
	Message dto.Message // <- adding a Message field
}

func (g Greeter) Greet() dto.Message {
	return g.Message
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
