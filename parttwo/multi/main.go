package main

// ---------------------------------------------------------
// EXERCISE: Swapper
//
//  1. Change `color` to "orange"
//     and `color2` to "green" at the same time
//
//     (use multiple-assignment)
//
//  2. Print the variables
//
// EXPECTED OUTPUT
//  orange green
// ---------------------------------------------------------
import (
	"fmt"
	"path"
)

func main() {
	// UNCOMMENT THE CODE BELOW:

	color, color2 := "red", "blue"

	color, color2 = "orange", "green"

	red, blue := "red", "blue"
	red, blue = blue, red
	// ---------------------------------------------------------
	// EXERCISE: Discard The File
	//
	//  1. Print only the directory using `path.Split`
	//
	//  2. Discard the file part
	//
	// RESTRICTION
	//  Use short declaration
	//
	// EXPECTED OUTPUT
	//  secret/
	// ---------------------------------------------------------

	// UNCOMMENT THE CODE BELOW:

	dir, file := path.Split("parttwo/makeitblue/main.go")

	fmt.Println(
		color, color2, red, blue, dir, file,
	)

}
