package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.com",
		"http://amazon.com",
		"http://stackoverflow.com",
	}
	c := make(chan string)

	for _, link := range links {
		go getSiteStatus(link, c)

	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			getSiteStatus(link, c)
		}(l)

	}

}

func getSiteStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- link
	} else {
		fmt.Println(link, "is up")
		c <- link
	}
}
