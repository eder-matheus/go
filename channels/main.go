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
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://ufrgs.br",
	}

	// creating new channel
	c := make(chan string)

	for _, link := range links {
		// go keyword creates new routines ("threads") that executes the
		// function it is calling
		go checkLink(link, c)
	}

	// the for below whatches the channel c
	// when something is sent through c, the for loop starts the go routine
	for l := range c {
		go func(link string) {
			time.Sleep(30 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
