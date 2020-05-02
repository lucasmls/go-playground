package main

import "sync"

// Counter ...
type Counter struct {
	value int
	mu    sync.Mutex
}

// NewCounter ...
func NewCounter() *Counter {
	return &Counter{}
}

// Increment ...
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

// Value ...
func (c *Counter) Value() int {
	return c.value
}
