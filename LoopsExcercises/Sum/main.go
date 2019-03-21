package main

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers
//
//  1. By using a loop, sum the numbers between 1 and 10.
//  2. Print the sum.
//
// EXPECTED OUTPUT
//  Sum: 55
// ---------------------------------------------------------
import (
	"fmt"
	"os"
	"strconv"
)

const msg = "Input : [Min number] [Max number]"

func main() {

	if len(os.Args) != 3 {
		fmt.Println(msg)
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil || min > max {
		fmt.Println("Provide correct " + msg)
		return
	}

	var sum int

	for i := min; ; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Print(i)

		if i != max {
			fmt.Print(" + ")
		}
		sum += i
		if i == max {
			break
		}

	}
	fmt.Printf(" = %d\n", sum)
}
