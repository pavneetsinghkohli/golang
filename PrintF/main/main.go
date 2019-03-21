package main

// ---------------------------------------------------------
// EXERCISE: Print Your Age
//
//  Print your age using Prinft
//
// EXPECTED OUTPUT
//  I'm 30 years old.
//
// NOTE
//  You should change 30 to your age, of course.
// ---------------------------------------------------------

import (
	"fmt"
	"os"
)

func main() {
	// var (
	// 	first = "Pavneet"
	// 	last  = "Singh"
	// )
	// trfa := "These are %t claims."

	fmt.Printf("Your name is %s and your lastname is %s.\n", os.Args[1], os.Args[2])

}
