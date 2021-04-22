package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	println(string(bs))
	println(" how many bytes: ", len(bs))
	return len(bs), nil
}
