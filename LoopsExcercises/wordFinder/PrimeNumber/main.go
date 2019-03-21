package main

// ---------------------------------------------------------
// EXERCISE: Crunch the primes
//
//  1. Get numbers from the command-line.
//  2. `for range` over them.
//  4. Convert each one to an int.
//  5. If one of the numbers isn't an int, just skip it.
//  6. Print the ones that are only the prime numbers.
//
// RESTRICTION
//  The user can run the program with any number of
//  arguments.
//
// HINT
//  Find whether a number is prime using this algorithm:
//  https://stackoverflow.com/a/1801446/115363
//
// EXPECTED OUTPUT
//  go run main.go 2 4 6 7 a 9 c 11 x 12 13
//    2 7 11 13
//
//  go run main.go 1 2 3 5 7 A B C
//    2 3 5 7
// ---------------------------------------------------------

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	const msg = "Please provide a few valid positive numbers and not zero "

	query := os.Args[1:]

	if len(query) < 1 {
		fmt.Println(msg)
	}
main:
	for _, q := range query {

		q, err := strconv.Atoi(q)

		if err != nil || q <= 1 {
			continue
		}

		if q == 2 || q == 3 {
			fmt.Printf("%d ", q)
			continue

		} else if q%2 == 0 || q%3 == 0 {
			// fmt.Printf("%d ", q)
			continue
		}

		i, w := 5, 2

		for i*i <= q {
			if q%i == 0 {
				// fmt.Printf("%d ", q)
				continue main
			}
			i += w
			w = 6 - w

		}

		fmt.Printf("%d ", q)
	}

}
