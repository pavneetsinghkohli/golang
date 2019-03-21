package main

// ---------------------------------------------------------
// STORY
//
//  Your boss wants you to create a program that will check
//  whether a person can watch a particular movie or not.
//
//  Imagine that another program provides the age to your
//  program. Depending on what you return, the other program
//  will issue the tickets to the person automatically.
//
// EXERCISE: Movie Ratings
//
//  1. Get the age from the command-line.
//
//  2. Return one of the following messages if the age is:
//     -> Above 17         : "R-Rated"
//     -> Between 13 and 17: "PG-13"
//     -> Below 13         : "PG-Rated"
//
// RESTRICTIONS:
//  1. If age data is wrong or absent let the user know.
//  2. Do not accept negative age.
//
// BONUS:
//  1. BONUS: Use if statements only twice throughout your program.
//  2. BONUS: Use an if statement only once.
//
// EXPECTED OUTPUT
//  go run main.go 18
//    R-Rated
//
//  go run main.go 17
//    PG-13
//
//  go run main.go 12
//    PG-Rated
//
//  go run main.go
//    Requires age
//
//  go run main.go -5
//    Wrong age: "-5"
// ---------------------------------------------------------

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		s, _ = strconv.ParseInt(os.Args[1], 10, 64)
		r    = s <= 0 || s >= 100

		a = s < 100 && s > 17
		b = s <= 17 && s >= 13
		c = s < 13 && s >= 1
	)

	if len(os.Args) != 2 {
		fmt.Println("Requires age number")
		return
	} else if r {
		fmt.Printf("Wrong Age : %d", s)
	} else if c {
		fmt.Println("PG-Rated")
	} else if b {
		fmt.Println("PG-13")
	} else if a {
		fmt.Println("R-Rated")
	}

}
