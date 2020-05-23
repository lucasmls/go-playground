package main

import "github.com/google/wire"

// InitializeEvent creates an Event. It will error if the Event is staffed with a grumpy greeter.
func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
