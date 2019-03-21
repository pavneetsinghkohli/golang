package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	m, err := strconv.ParseFloat(os.Args[1], 64)

	if len(os.Args) != 2 || err != nil {
		fmt.Printf("%q : Wrong Input - Give me the magnitude of the earthquake.\n", os.Args[1])
		return
	}

	d:= "description"
	

	switch m {
	case m < 2.0 
		d ="micro"
	case  m < 3.0 && m >= 2.0 : 
		d ="very minor"
	case  m < 4.0 && m >= 3.0 : 
		d ="minor"
	case  m < 5.0 && m >= 4.0 : 
		d ="light"
	case  m < 6.0 && m >= 5.0 : 
		d ="moderate"
	case  m < 7.0 && m >= 6.0 : 
		d ="strong"
	case  m < 8.0 && m >= 7.0 : 
		d ="major"
	case  m < 9.0 && m >= 8.0 : 
		d ="great"
	case  m >= 10.0 
		d ="massive"
	}
}
