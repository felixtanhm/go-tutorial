package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	lw := logWriter{}

	io.Copy(lw, res.Body)
}

func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
