package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com/",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://www.mercadolibre.com.ar",
	}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for {
		go func(c chan string) {
			time.Sleep(time.Second * 2)
			checkLink(<-c, c)
		}(c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Unable to reach", link)
		c <- link
		return
	}
	fmt.Println(link, "seems to be ok")
	c <- link
}
