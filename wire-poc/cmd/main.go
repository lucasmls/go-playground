package main

import "fmt"

// Message ...
type Message string

// NewMessage ...
func NewMessage() Message {
	return Message("Hi there!")
}

// Greeter ...
type Greeter struct {
	Message Message
}

// NewGreeter ...
func NewGreeter(m Message) Greeter {
	return Greeter{
		Message: m,
	}
}

// Greet ...
func (g Greeter) Greet() Message {
	return g.Message
}

// Event ...
type Event struct {
	Greeter Greeter
}

// NewEvent ...
func NewEvent(g Greeter) Event {
	return Event{
		Greeter: g,
	}
}

// Start ...
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
}
