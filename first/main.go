package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		n   int
		err error
	)

	if n, err = strconv.Atoi("10"); err != nil {
		fmt.Printf("error: %s (n: %d)", err, n)
		return
	}

	fmt.Println(n)
}
