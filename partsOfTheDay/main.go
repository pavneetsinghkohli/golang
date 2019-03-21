package main

import (
	"fmt"
	"time"
)

func main() {

	// t = 2

	// fmt.Println(t)

	switch t := time.Now().Hour(); {
	case t >= 18:
		fmt.Println("Good Evening")
	case t >= 12:
		fmt.Println("Good Afternoon")
	case t >= 6:
		fmt.Println("Good Morning")
		// case t >= 23:
		// 	fmt.Println("Good Night")
	default:
		fmt.Println("Good Night")
	}

}
