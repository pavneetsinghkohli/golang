package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var (
		m    = strings.ToLower(os.Args[1])
		year = time.Now().Year()
		leap = year%4 == 0 && (year%100 != 0 || year%400 == 0)
		days int
	)

	if len(os.Args) != 2 {
		fmt.Println("Give me a month name.")
		return
	}

	if m == "april" || m == "june" || m == "september" || m == "november" {
		days = 30
	} else if m == "january" || m == "march" || m == "may" || m == "july" || m == "august" || m == "october" || m == "december" {
		days = 31
	} else if m == "february" && leap {
		days = 29
	} else if m == "february" {
		days = 28
	} else {
		fmt.Printf("%q is not a month or not spelled correctly.\n", m)
		return
	}

	fmt.Printf("%q has %d days.", os.Args[1], days)

}
