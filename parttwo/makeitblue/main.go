package main

// ---------------------------------------------------------
// EXERCISE: Make It Blue
//
//  1. Change `color` variable's value to "blue"
//
//  2. Print it
//
// EXPECTED OUTPUT
//  blue
// ---------------------------------------------------------
import "fmt"

func main() {
	// UNCOMMENT THE CODE BELOW:

	// color := "green"

	// ADD YOUR CODE BELOW:

	//color, nColor := "green", "dark green"

	//color = nColor

	color := "dark" + " green"
	var n float64

	n = 3.14 * 2

	width, height, perimeter := 5, 6, 0

	lang, version := "go", 2

	//var perimeter int

	perimeter = (width + height) * 2
	// EXPECTED OUTPUT
	//  Air is good on Mars
	//  It's true
	//  It is 19.5 degrees

	// UNCOMMENT THE CODE BELOW:

	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, isTrue, temp = "Mars", true, 19.5

	fmt.Println(
		color, n, perimeter,
		lang, "version", version, "\n\n",
		"Air is good on", planet, "\n\n",
		"It's ", isTrue, "\n\n",
		"It is", temp, "degrees",
	)
}
