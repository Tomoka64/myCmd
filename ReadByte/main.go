package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type NewWriter struct{}

func main() {
	src := os.Args[1]
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "<Usage> %s URL\n", os.Args[0])
		os.Exit(1)
	}
	url, err := http.Get(src)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//
	nw := NewWriter{}
	io.Copy(nw, url.Body)

}

func (NewWriter) Write(bs []byte) (int, error) {
	fmt.Printf("bytes: %v\n", len(bs))
	return len(bs), nil
}
