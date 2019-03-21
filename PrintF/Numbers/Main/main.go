package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// fmt.Println(50 + 25)
	// fmt.Println(50 - 15.5)
	// fmt.Println(50 * .5)
	// fmt.Println(50 / .5)
	// fmt.Println(25 % 3)
	// fmt.Println(-(5 + 2))

	// // x := 5. / 2
	// // fmt.Println(x)

	// // This expression should print 20
	// fmt.Println(10 + 5 - (5 - 10))

	// // This expression should print -16
	// fmt.Println(-10 + (0.5 - (1 + 5.5)))

	// // This expression should print -25
	// fmt.Println((5+10)*2 - 5)

	// // This expression should print 0.5
	// fmt.Println(0.5 * (2 - 1))

	// // This expression should print 24
	// fmt.Println(((3+1)/2)*10 + 4)

	// // This expression should print 15
	// fmt.Println((10 / 2) * (10 % 7))

	// // This expression should print 40
	// // Note that you may need to use floats to solve this
	// fmt.Println(100 / (5. / 2))

	// DO NOT TOUCH THIS
	// counter, factor := 45, 0.5

	// // TYPE YOUR CODE BELOW
	// // ...
	// counter++
	// counter++
	// counter++
	// counter++
	// counter++

	// factor--
	// factor--

	// fmt.Println(float64(counter) * factor)

	// // LASTLY: REMOVE THE CODE BELOW
	// //_, _ = counter, factor
	// width, height := 10, 2

	// width += 1
	// width += height
	// width -= 1
	// width -= height
	// width *= 20
	// width /= 25
	// width %= 5

	// fmt.Println(width)

	// radius, area := 10., 0.

	// area = math.Pi * math.Pow(radius, 2)

	// fmt.Printf("radius: %g -> area: %f\n", radius, area)

	// var (
	// 	radius float64
	// 	area   float64
	// )

	// radius, _ = strconv.ParseFloat(os.Args[1], 64)

	// area = 4 * math.Pi * math.Pow(radius, 2)

	// fmt.Printf("radius: %g -> area: %.2f\n", radius, area)

	//v=4/3 * pi * r**3

	var (
		radius float64
		vol    float64
	)

	radius, _ = strconv.ParseFloat(os.Args[1], 64)

	vol = (4 * math.Pi * math.Pow(radius, 3)) / 3

	fmt.Printf("radius: %g -> Volume: %.2f\n", radius, vol)

}
