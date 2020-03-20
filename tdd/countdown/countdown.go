package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWorld = "Go!"
const countdowStart = 3

// Countdown ...
func Countdown(out io.Writer) {
	for i := countdowStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWorld)
}

func main() {
	Countdown(os.Stdout)
}
