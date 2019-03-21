package main

// ---------------------------------------------------------
// EXERCISE: Count the Arguments
//
//  Print the count of the command-line arguments
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 names.
// ---------------------------------------------------------
import (
	"fmt"
	"os"
)

// func main()
//  {
// UNCOMMENT & FIX THIS CODE
// count := len(os.Args) - 1
// myprogram := os.Args[0]
// name := os.Args[1]

// UNCOMMENT IT & THEN DO NOT TOUCH THIS CODE
/*fmt.Printf("There are %d names.\n", count)
fmt.Println(
	"\n\nBelow is a path from our program \n\n", myprogram,
	"\n\n Hi", name, "\n\n How are you?",
)*/
// RESTRICTIONS
//  1. Be sure to match to the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store the all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// TYPE YOUR CODE HERE
	count, name := len(os.Args)-1, os.Args
	fmt.Println(
		"There are", count, "people! \n\n",
		"Hello great ", name, "!\n\n",
		// "Hello great ", name[2], "!\n\n",
		// "Hello great ", name[3], "!\n\n",
		"Nice to meet you All.",
	)
	// BONUS #1:
	// Observe the error if you pass less then 3 arguments.
	// Search on the web how to solve that.

}
