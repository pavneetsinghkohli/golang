package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	names := [...]string{
		`Kafka's Revenge`,
		"Stay Golden",
		"Everythingship",
		`Kafka's Revenge 2nd Edition`,
	}

	if len(os.Args) != 2 {
		fmt.Println("Please provide a valid book title")
		return
	}

	search := strings.ToLower(os.Args[1])
	var count = false

	fmt.Println("Search Results :")

	for _, v := range names {
		v = strings.ToLower(v)

		if strings.Contains(v, search) {
			count = true
			fmt.Printf("+ %s\n", v)
		}
	}

	if !count {
		fmt.Printf("We don't have the book: %q\n", os.Args[1])
	}
}
