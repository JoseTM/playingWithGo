package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	/* 	fmt.Println(resp)
	   	google, err := io.ReadAll(resp.Body)
	   	resp.Body.Close()
	   	if err != nil {
	   		fmt.Println(err)
	   		os.Exit(1)
	   	}

	   	fmt.Printf("%s", google) */

	bs := make([]byte, 99999)

	resp.Body.Read(bs)
	fmt.Println(string(bs))
	resp.Body.Close()

}
