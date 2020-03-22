package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper ...
type Sleeper interface {
	Sleep()
}

// CountdownSleeper ...
type CountdownSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep ...
func (c *CountdownSleeper) Sleep() {
	c.sleep(c.duration)
}

const finalWorld = "Go!"
const countdownStart = 3

// Countdown ...
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWorld)
}

func main() {
	sleeper := &CountdownSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
