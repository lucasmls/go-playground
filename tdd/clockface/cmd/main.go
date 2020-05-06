package main

import (
	"os"
	"time"

	clock "github.com/lucasmls/tdd/clockface/svg"
)

func main() {
	t := time.Now()
	clock.SVGWriter(os.Stdout, t)
}
