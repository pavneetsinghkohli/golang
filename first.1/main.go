package main

import {
	"fmt"
	"runtime"

} 


// ---------------------------------------------------------
// EXERCISE: Comment out
//
//  Use single and multiline comments to comment Printlns.
//
// EXPECTED OUTPUT
//  You shouldn't see any output after you're done.
// ---------------------------------------------------------

//godoc -src fmt println

func main() {
	fmt.Println(runtime.NumCPU() + "hello")
	godoc -src runtime NumCPU
	// fmt.Println("how") 
	/*

	fmt.Println("are")
	*/

	fmt.Println("you")

	
}


