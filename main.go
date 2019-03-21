package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

func main() {

	// Digit character       : █
	// Separator character   : ░

	zero := [5]string{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := [5]string{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := [5]string{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := [5]string{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := [5]string{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := [5]string{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := [5]string{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := [5]string{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := [5]string{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := [5]string{
		"███",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	sep := [5]string{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	// sep1 := [5]string{
	// 	"   ",
	// 	"   ",
	// 	"   ",
	// 	"   ",
	// 	"   ",
	// }

	digits := [10][5]string{zero, one, two, three, four, five, six, seven, eight, nine}

	for {
		screen.Clear()
		screen.MoveTopLeft()

		now := time.Now()

		hour, minute, second := now.Hour(), now.Minute(), now.Second()

		clock := [8][5]string{
			digits[hour/10], digits[hour%10],
			sep,
			digits[minute/10], digits[minute%10],
			sep,
			digits[second/10], digits[second%10],
		}

		// clock1 := [8][5]string{
		// 	digits[hour/10], digits[hour%10],
		// 	sep1,
		// 	digits[minute/10], digits[minute%10],
		// 	sep1,
		// 	digits[second/10], digits[second%10],
		// }
		for line := range digits[0] {
			// Print a line for each placeholder in digits
			for i, digit := range clock {

				next := clock[i][line]

				if digit == sep && second%2 == 0 {
					next = "   "
				}
				fmt.Printf("%v   ", next)

			}
			fmt.Println()
		}

		time.Sleep(time.Second)

	}

}
