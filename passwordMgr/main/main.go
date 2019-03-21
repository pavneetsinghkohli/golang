package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	u, p := args[1], args[2]

	if u != "jack" && u != "inanc" {
		fmt.Printf("Access denied for %q.\n", u)
	} else if (u == "jack" && p != "1888") || (u == "inanc" && p != "1879") {
		fmt.Printf("Invalid password for %q.\n", u)
	} else {
		fmt.Printf("Access granted to %q.\n", u)
	}
}
