package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	if len(os.Args[1:]) < 2 {
		fmt.Println("Please provide a name and mood")
	}

	mood := [...][4]string{
		{"happy", "Awsome", "Wonderful", "Great"},
		{"Irritating", "very bad", "lausy", "worse"},
	}

	r := rand.Intn(len(mood[0]))
	ask := strings.ToLower(os.Args[2])

	switch ask {
	case "positive":
		fmt.Printf("%s is feeling %s\n", os.Args[1], mood[0][r])
	default:
		fmt.Printf("%s is feeling %s\n", os.Args[1], mood[1][r])

	}

	// // switch r {
	// // case 1:
	// // 	fmt.Printf("%s is feeling %s\n", os.Args[1], mood[0])
	// case 2:

	// case 3:
	// 	fmt.Printf("%s is feeling %s\n", os.Args[1], mood[2])
	// case 4:
	// 	fmt.Printf("%s is feeling %s\n", os.Args[1], mood[3])
	// case 0:
	// 	fmt.Printf("%s is feeling %s\n", os.Args[1], mood[4])
	// }

}
