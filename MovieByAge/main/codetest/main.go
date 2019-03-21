package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	fmt.Println(time.Month(1))

	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}
	n, err := strconv.Atoi(os.Args[1])

	if err != nil || n == 0 {
		fmt.Printf(`%q is not a valid year.`+"\n", os.Args[1])
		return
	} else if n%4 == 0 {
		fmt.Printf("%d is a leap year", n)
	} else {
		fmt.Printf("%d is not a leap year.", n)
	}
}
