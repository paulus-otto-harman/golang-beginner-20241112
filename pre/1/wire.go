//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

//func InitializeEvent() Event {
//	wire.Build(NewEvent, NewGreeter, NewMessage)
//	return Event{}
//}

func InitEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
