package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func findPrimeNumbersInRange(num int, ch chan int) {
	for i := 1; i <= num; i++ {
		if isPrime(i) {
			ch <- i
			time.Sleep(time.Millisecond * 500)
		}
	}

	close(ch)
}

func findTheFirstXPrimeNumbers(num int, ch chan int) {
	start := 2
	for i := 0; i < num; i++ {
		for x := start; ; x++ {
			if isPrime(x) {
				ch <- x
				start = x + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}

	close(ch)
}

func main() {
	ch := make(chan int, 30)
	ch2 := make(chan int, 30)

	go findPrimeNumbersInRange(cap(ch), ch)
	go findTheFirstXPrimeNumbers(cap(ch2), ch2)

	for prime := range ch {
		fmt.Printf("[ch] - %d \n", prime)
	}

	for prime := range ch2 {
		fmt.Printf("[ch2] - %d \n", prime)
	}

}
