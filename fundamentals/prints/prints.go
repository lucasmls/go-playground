package main

import "fmt"

func main() {
	fmt.Print("Hey \n")

	fmt.Println("Ho")

	x := 3.1415
	xs := fmt.Sprint(x)

	fmt.Println("The value of X is: " + xs)
	fmt.Println("The value of X is:", x)

	fmt.Printf("The value of X is: %f \n", x)
	fmt.Printf("The value of X is: %.2f \n", x)

	a := 1      // Integer = %d
	b := 1.9999 // Float = %f
	c := false  // Boolean = %t
	d := "Go!"  // String = %s

	fmt.Printf("%d %f %.1f %t %s \n", a, b, b, c, d)
	fmt.Printf("%v %v %v %v", a, b, c, d)
}
