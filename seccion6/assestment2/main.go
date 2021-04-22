package main

import (
	"io"
	"log"
	"os"
)

type logWriter struct{}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("crash, boom, bang")
		os.Exit(0)
	}

	filePath := os.Args[1]
	file, errorPath := os.Open(filePath)
	if errorPath != nil {
		log.Fatal("no fichero baby mio")
		os.Exit(0)
	}

	lw := logWriter{}
	io.Copy(lw, file)
	//io.Copy(os.Stdout, file)

}

func (logWriter) Write(bs []byte) (int, error) {
	println(string(bs))
	println(" how many bytes: ", len(bs))
	return len(bs), nil
}
