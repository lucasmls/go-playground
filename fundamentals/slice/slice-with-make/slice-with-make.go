package main

import "fmt"

func main() {
	firstSlice := make([]int, 10)
	firstSlice[9] = 12
	fmt.Println(firstSlice)

	// in this case, we're creating a slice of 10 positions
	// but the internal array that the slice is based has 20 positions
	secondSlice := make([]int, 10, 20)
	fmt.Println(secondSlice)
	fmt.Println(len(secondSlice), cap(secondSlice))

	secondSlice = append(secondSlice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(secondSlice, len(secondSlice), cap(secondSlice))

	// In this case, that we append the item 21, in a slice with cap
	// of 20, the slice automatically will double this cap
	secondSlice = append(secondSlice, 123)
	fmt.Println(secondSlice, len(secondSlice), cap(secondSlice))
}
