package main

// ---------------------------------------------------------
// EXERCISE: Convert and Fix
//
//  Fix the code by using the conversion expression.
//
// EXPECTED OUTPUT
//  15.5
// ---------------------------------------------------------
import "fmt"

func main() {
	//a, b := 10, 5.5
	// //a, b := 10, 5.5
	//a = int(b)

	// HINTS
	//   maximum of int8  can be 127
	//   maximum of int16 can be 32767
	//
	// EXPECTED OUTPUT
	//  1127
	// ---------------------------------------------------------

	// DO NOT TOUCH THIS VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	fmt.Println(min, max, max+int16(min))

	//age := 2
	//fmt.Println(7.5 + float64(age))

	//fmt.Println(float64(a) + b)
	//fmt.Println(5.5)
	//fmt.Println(float64(a) + b)

}
