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
type CountdownSleeper struct{}

// Sleep ...
func (d CountdownSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWorld = "Go!"
const countdowStart = 3

// Countdown ...
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdowStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWorld)
}

func main() {
	sleeper := &CountdownSleeper{}
	Countdown(os.Stdout, sleeper)
}
