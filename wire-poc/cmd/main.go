package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// Message ...
type Message string

// NewMessage ...
func NewMessage() Message {
	return Message("Hi there!")
}

// Greeter ...
type Greeter struct {
	Message Message
	Grumpy  bool
}

// NewGreeter ...
func NewGreeter(m Message) Greeter {
	grumpy := false

	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}

	return Greeter{
		Message: m,
		Grumpy:  grumpy,
	}
}

// Greet ...
func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("go away!")
	}

	return g.Message
}

// Event ...
type Event struct {
	Greeter Greeter
}

// NewEvent ...
func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}

	return Event{
		Greeter: g,
	}, nil
}

// Start ...
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	e, err := InitializeEvent()
	if err != nil {
		fmt.Printf("failed to create event: %s \n", err)
		os.Exit(2)
	}

	e.Start()
}
