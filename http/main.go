package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("https://www.youtube.com/watch?v=iGnpyBUsPTk/")
	if err != nil {
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}
