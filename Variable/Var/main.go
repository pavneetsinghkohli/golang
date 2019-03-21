package main

// ---------------------------------------------------------
// EXERCISE: Declare int
//
//  1. Declare and print a variable with an int type
//
//  2. The declared variable's name should be: height
//
// EXPECTED OUTPUT
//  0
// ---------------------------------------------------------
import "fmt"

func main() {
	// var ? ?
	// ?
	var (
		height     int
		isOn       bool
		brightness float64
		vari       string
		s          string

		i         int
		i8        int8
		i16       int16
		i32       int32
		i64       int64
		f32       float32
		f64       float64
		c64       complex64
		c128      complex128
		a         bool
		so        string
		r         rune
		b         byte
		firstName = "Pavneet"
		lastName  = "Singh"
		isLiquid  = "is Liquid"
	)
	var red string
	fmt.Println(red, height)
	fmt.Println(isOn)
	fmt.Println(brightness)
	fmt.Printf("s (%T): %q\n", vari, s)
	fmt.Println(i, i8, i16, i32, i64)
	fmt.Println(f32, f64)
	fmt.Println(c64, c128)
	fmt.Println(a, i64, b, c128, f32)
	fmt.Printf("%q\n", so)
	fmt.Println(r)
	fmt.Println(b)
	fmt.Printf("%q %q\n", firstName, lastName)
	fmt.Println(firstName, lastName)
	_ = isLiquid

}
