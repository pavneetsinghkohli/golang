package main

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	const (
		msg     = "Please provide a valid Guess, this programm runs %d times"
		maxTurn = 5
	)

	var (
		m       string
		greater int
	)

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf(msg, maxTurn)
		fmt.Println()
		return
	}

	g, err1 := strconv.Atoi(args[0])
	e, err2 := strconv.Atoi(args[1])

	if err1 != nil || err2 != nil || e <= 0 || g <= 0 {
		fmt.Println(msg)
		return
	}

	if g < e {
		greater = e
	} else {
		greater = g
	}

	fmt.Println("Greater Value : ", greater)

	// t := time.Now()
	rand.Seed(time.Now().UnixNano())

	for turn := 0; turn < maxTurn; turn++ {

		n := rand.Intn(greater + 1)
		fmt.Println(n)

		if n == g || n == e {
			if turn == 0 {

				switch rand.Intn(3) {
				case 0:
					m = "First Attempt"
				case 1:
					m = "Congratulations for the correct Guess in first Attempt"
				case 2:
					m = "Balle Balle"
				}

				fmt.Println(m)
			} else {
				switch rand.Intn(3) {
				case 0:
					m = "Good Job but not the first Guess"
				case 1:
					m = "Not a first Attempt though"
				case 2:
					m = "Balle Balle - not that first"
				}
				fmt.Println(m)
			}
			return
		}

	}
	switch rand.Intn(3) {
	case 0:
		m = "You Loose "
	case 1:
		m = "Try Again"
	case 2:
		m = "Looser"
	}
	fmt.Println(m)
	return
}
