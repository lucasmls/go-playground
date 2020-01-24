package main

import (
	"concurrency/html"
	"fmt"
	"time"
)

func fastest(url1, url2, url3 string) string {
	ch1 := html.GetTitles(url1)
	ch2 := html.GetTitles(url2)
	ch3 := html.GetTitles(url3)

	select {
	case res1 := <-ch1:
		return res1
	case res2 := <-ch2:
		return res2
	case res3 := <-ch3:
		return res3
	case <-time.After(1000 * time.Millisecond):
		return "None of them..."
		// default:
		// 	return "None of them..."
	}
}

func main() {
	fastestOne := fastest(
		"https://www.google.com",
		"https://www.twitter.com",
		"https://www.youtube.com",
	)

	fmt.Println(fastestOne)
}
