package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func getTitles(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)

			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func main() {
	t1 := getTitles("https://www.google.com", "https://www.twitter.com")
	t2 := getTitles("https://www.github.com", "https://www.youtube.com")

	fmt.Println("First ones: ", <-t1, " | ", <-t2)
	fmt.Println("Second ones: ", <-t1, " | ", <-t2)
}
