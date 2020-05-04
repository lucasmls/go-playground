package main

import (
	"os"
	"time"

	"github.com/lucasmls/tdd/clock"
)

func main() {
	t := time.Now()
	clock.SVGWriter(os.Stdout, t)
}
