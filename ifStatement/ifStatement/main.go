package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://en.oxforddictionaries.com/explore/is-the-letter-y-a-vowel-or-a-consonant/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.*(

// ---------------------------------------------------------

func main() {

	var (
		args = os.Args
		l    = len(args) - 1
		c    = len(os.Args[1])
		c1   = os.Args[1]
	)

	// fmt.Printf("%s,%d, %d", args, l, c)

	if l == 0 || c > 2 {
		fmt.Printf("Give me a letter, instead of %s", c1)

	} else if strings.IndexAny(c1, "aeiouwy") == -1 {
		fmt.Printf(`"%s" is a consonant`, c1)
	} else if strings.IndexAny(c1, "wy") != -1 {
		fmt.Printf(`"%s" is sometimes a vowel, sometimes not.`, c1)
	} else if strings.IndexAny(c1, "aeiou") != -1 {
		fmt.Printf(`"%s" is a vowel`, c1)
	} else {
		fmt.Println("Not a correct unicode")
	}

}
