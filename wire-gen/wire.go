//go:build wireinject
// +build wireinject

package wiregen

import (
	"github.com/google/wire"
	"wire-demo/service"
)

func InitializeEvent() service.Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return service.Event{}
}
