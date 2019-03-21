package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// fmt.Println(len(os.Args), os.Args)

	if len(os.Args) != 6 {
		fmt.Println("Please tell me numbers (maximum 5 numbers)")
		return
	}

	var (
		num   [5]float64
		total float64
		sum   float64
		query = os.Args[1:]
	)

	for i, q := range query {
		n, err := strconv.ParseFloat(q, 64)
		if err != nil {
			// fmt.Printf("(%v) : invalid number", err)
			continue
		}
		total++
		sum += n

		num[i] = n

	}

	fmt.Println("Your Numbers : ", num)
	fmt.Printf("Average : %.f", sum/total)
	// args := len(os.Args)

}
