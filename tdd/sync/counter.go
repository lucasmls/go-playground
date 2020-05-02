package main

// Counter ...
type Counter struct {
	value int
}

// Increment ...
func (c *Counter) Increment() {
	c.value++
}

// Value ...
func (c *Counter) Value() int {
	return c.value
}
