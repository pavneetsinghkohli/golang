package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	const corpus = "the quick brown fox jumps over the lazy little dog"
	words := strings.Fields(corpus)

	// fmt.Printf("\tinput : %q\n\tOuput : %q\n", corpus, c)

	query := os.Args[1:]

	if len(query) < 1 {
		fmt.Printf("Please provide a string")
		return
	}

	// s := strings.ToLower(query[0])
	// queryies:
	for _, q := range query {
		qs := strings.ToLower(q)
		for i, w := range words {
			if qs == w {
				fmt.Printf("#%-2d : %q\n", i+1, q)
				break
			}

		}

	}

}
