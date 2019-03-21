package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	// maxTurns = 5 // less is more difficult
	usage = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
(Provide -v flag to see the picked numbers.)
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args, maxTurns := os.Args[1:], 5

	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	var verbose bool

	if args[0] == "-v" {
		verbose = true
	}

	guess, err := strconv.Atoi(args[len(args)-1])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	min := 10

	if guess > 10 {
		min = guess
		maxTurns = rand.Intn(guess)
		// fmt.Println(min, maxTurns)

	}

	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(min + 1)

		if verbose {
			fmt.Printf("%d ", n)
		}

		if n == guess {
			if turn == 1 {
				fmt.Println("🥇 FIRST TIME WINNER!!!")
			} else {
				fmt.Println("🎉  YOU WON!")
			}
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}
