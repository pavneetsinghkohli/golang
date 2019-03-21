package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	query := os.Args[1:]

	if len(query) < 1 {
		fmt.Printf("Please provide a string")
		return
	}

	env := filepath.SplitList(os.Getenv("path"))

	// env := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))

	for _, q := range query {

		for i, e := range env {
			e, q = strings.ToLower(e), strings.ToLower(q)

			if strings.Contains(e, q) {
				fmt.Printf("#%-2d : %q\n", i+1, e)
				// break
			}
		}
	}

}
