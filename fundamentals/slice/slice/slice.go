package main

import (
	"fmt"
	"reflect"
)

func main() {
	firstArray := [3]int{1, 2, 3} // It's a array
	firstSlice := []int{1, 2, 3}  // It's a slice

	fmt.Println(firstArray, firstSlice)
	fmt.Println(reflect.TypeOf(firstArray), reflect.TypeOf(firstSlice))

	// We can have multiple slices of a array, but we're still
	// poiting to the same array
	// If we modify a slice, we will be mutating the true array

	secondArray := [5]int{1, 2, 3, 4, 5}
	secondSlice := secondArray[1:3] // Get indexes 1 and 2

	fmt.Println(secondSlice)

	thirdSlice := secondArray[:2] // Get indexes 0 and 1
	fmt.Println(thirdSlice)

}
