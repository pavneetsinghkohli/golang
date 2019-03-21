package main

import "fmt"

func main() {
	richter := 2.5

	switch r := richter; {
	case r < 2:
		fallthrough
	case r >= 2 && r < 3:
		fallthrough
	case r >= 5 && r < 6:
		fmt.Println("Not important", r)
	case r >= 6 && r < 7:
		fallthrough
	case r >= 7:
		fmt.Println("Destructive")
	}
	fmt.Println(richter)
}
