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
		num [5]float64

		query = os.Args[1:]
	)

	for i, q := range query {
		n, err := strconv.ParseFloat(q, 64)
		if err != nil {
			// fmt.Printf("(%v) : invalid number", err)
			continue
		}

		num[i] = n

	}
	for range num {
		for i := range num {
			if i < len(num)-1 && num[i] > num[i+1] {
				num[i], num[i+1] = num[i+1], num[i]

				// fmt.Println(i, num[i])
			}
		}
	}

	fmt.Println("Your Numbers : ", num)
	// fmt.Printf("leanth: %d", len(num)-1)
	// // args := len(os.Args)

}
