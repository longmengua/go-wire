package wiregen

import (
	"wire-demo/dto"
	"wire-demo/service"
)

func NewMessage() dto.Message {
	return dto.Message("Hi there!")
}

func NewGreeter(m dto.Message) service.Greeter {
	return service.Greeter{Message: m}
}

func NewEvent(g service.Greeter) service.Event {
	return service.Event{Greeter: g}
}
