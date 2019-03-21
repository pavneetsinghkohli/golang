package main

// ---------------------------------------------------------
// EXERCISE: Short Declare
//
//  Declare and then print four variables using
//  the short declaration statement.
//
// EXPECTED OUTPUT
// ---------------------------------------------------------
import "fmt"

func main() {
	// ADD YOUR DECLARATIONS HERE
	//

	// THEN UNCOMMENT THE CODE BELOW
	//	i: 314 f: 3.14 s: Hello b: true
	i, f, s, b := 314, 3.14, "Hello", true

	sum := 27 + 3.5

	a, c := 42, "good"

	roc := true

	roc, pap := false, true

	_ = pap

	age, yourAge := 41, 20

	age, ratio := 42, 3.14

	age = 33

	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
		a, c,
		sum, roc,
		age, yourAge, ratio,
	)
}
