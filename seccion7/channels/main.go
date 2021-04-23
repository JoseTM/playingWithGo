package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.es",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)

	}

	for channelResponse := range channel {
		go func(link string) {
			time.Sleep(10 * time.Second)
			checkLink(link, channel)
		}(channelResponse)
	}

}

func checkLink(link string, channel chan string) {
	_, errorResponse := http.Get(link)

	if errorResponse != nil {
		fmt.Println(link, "might be down! (", errorResponse, ")")
		channel <- link
	} else {
		fmt.Println(link, " is up")
		channel <- link
	}

}
