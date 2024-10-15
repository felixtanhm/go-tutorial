package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.golang.org",
	}

	channel := make(chan string)
	for _, url := range urls {
		go checkLink(url, channel)
	}

	for link := range channel {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, channel)
		}(link)
	}
}

func checkLink(url string, channel chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down!")
		channel <- url
		return
	}

	fmt.Println(url, "is up!")
	channel <- url
}
