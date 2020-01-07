package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	fmt.Println("Typeof 32000:", reflect.TypeOf(32000))

	// Without signal (only positives) => uint8, uint16, uint32, uint64

	// With signals (positives and negatives) => int8, int16, int32, int64

	maxInt := math.MaxInt64
	fmt.Println("The maximum value of int is:", maxInt)
	fmt.Println("maxInt type:", reflect.TypeOf(maxInt))

	fmt.Println("========")

	const unicodeInt rune = 'a' // Respresents a a mapping to the Unicode table (int32)
	fmt.Println("Typeof rune:", reflect.TypeOf(unicodeInt))
	fmt.Println(unicodeInt)

	fmt.Println("========")

	const x float32 = 49.99
	fmt.Println("Typeof X:", reflect.TypeOf(x))
	fmt.Println("Literal type of 49.00 is:", reflect.TypeOf(49.00)) // float64

	fmt.Println("=======")

	bo := true // or false
	fmt.Println("Typeof bo:", reflect.TypeOf(bo))
	fmt.Println(!bo)

	fmt.Println("=======")

	firstString := "Hey, my type is string!"
	fmt.Println(firstString + "!!")
	fmt.Println("Typeof firstString:", reflect.TypeOf(firstString))
	fmt.Println("firstString length is:", len(firstString))

	multiLineString := `hey!
	i'm
	breaking
	lines!
	`
	fmt.Println("multiLineString:", multiLineString)

	character := 'a'
	fmt.Println("Character:", character)                        // There's no char type in go, this variable represents a Unicode value (int32)
	fmt.Println("Typeof character:", reflect.TypeOf(character)) // There's no char type in go, this variable represents a Unicode value (int32)
}
