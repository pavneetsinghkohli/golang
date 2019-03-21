package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

	words := strings.Fields(corpus)

	// fmt.Printf("\tinput : %q\n\tOuput : %q\n", corpus, c)

	query := os.Args[1:]

	if len(query) < 1 {
		fmt.Printf("Please provide a string")
		return
	}

	filter := [...]string{"and", "or", "was", "the", "since", "very"}

	// s := strings.ToLower(query[0])
	// queryies:
mains:
	for _, q := range query {
		q = strings.ToLower(q)

		for _, v := range filter {
			if q == v {
				continue mains
			}
		}

		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d : %q\n", i+1, q)
				break
			}
		}
	}

}
