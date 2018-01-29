package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s URL\n", os.Args[0])
		os.Exit(1)
	}
	response, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		f, err := os.Create("curl.txt")
		if err != nil {
			log.Fatalln("could not create file")
		}
		defer f.Close()

		_, err = io.Copy(f, response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	word := os.Args[2]
	b, err := os.Open("curl.txt")
	if err != nil {
		log.Fatalln("could not open file")
	}
	defer b.Close()
	counts := WordCount(b)
	fmt.Printf("Number of %s:", word)
	fmt.Println(counts[word])
}

func WordCount(rdr io.Reader) map[string]int {
	counts := map[string]int{}
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)

		counts[word]++
	}
	return counts
}
