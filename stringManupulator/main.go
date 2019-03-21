package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate the
//  given strings to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------

func main() {

	if len(os.Args) != 3 {
		fmt.Printf("Unknown command  or parameters passed: %q.\n", os.Args[1])
		return
	}

	op, s := os.Args[1], os.Args[2]
	switch op {
	case "lower":
		fmt.Println(strings.ToLower(s))
	case "upper":
		fmt.Println(strings.ToUpper(s))
	case "title":
		fmt.Println(strings.Title(s))
	default:
		fmt.Printf("Unknown command : %q\n", op)

	}
}
