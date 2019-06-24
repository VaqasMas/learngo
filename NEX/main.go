package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	//EX of numbers 1
	/*
		fmt.Printf("Sum: %d\n", 50+25)
		fmt.Printf("Diff: %g\n", 50-15.5)
		fmt.Printf("Prod: %g\n", 50*.5)
		fmt.Printf("Divide: %g\n", 50/.5)
		fmt.Printf("Remind: %d\n", 25%3)
		fmt.Printf("Negation: %d\n", -(5 + 2))
	*/

	//EX of numbers 2
	/*
		x := float64(5) / 2
		fmt.Println(x)
	*/

	//EX of numbers 3
	/*
		// This expression should print 20
		fmt.Println(10 + 5 - (5 - 10))

		// This expression should print -16
		fmt.Println((-10 + 0.5) - (1 + 5.5))

		// This expression should print -25
		fmt.Println(5 + 10*(2-5))

		// This expression should print 0.5
		fmt.Println(0.5 * (2 - 1))

		// This expression should print 24
		fmt.Println(((3+1.)/2)*10 + 4)

		// This expression should print 15
		fmt.Println((10 / 2) * (10 % 7))

		// This expression should print 40
		// Note that you may need to use floats to solve this
		fmt.Println(100 / (5. / 2))
	*/

	//EX of numbers 4
	/*
		counter, factor := 45, 0.5

		// TYPE YOUR CODE BELOW
		counter += 5
		factor -= 2

		// LASTLY: REMOVE THE CODE BELOW
		fmt.Println(float64(counter) * factor)
	*/

	//EX of numbers 5
	/*
		var counter int

		// TYPE YOUR CODE HERE
		counter++
		counter--
		counter += 5
		counter *= 10
		counter /= 5

		// DO NOT CHANGE THE CODE BELOW
		fmt.Println(counter)
	*/

	//EX of numbers 6
	/*
		width, height := 10, 2

		width++
		width += height
		width--
		width -= height
		width *= 20
		width /= 25
		width %= 5

		fmt.Println(width)
	*/

	//EX of numbers 7
	/*
		var (
			radius = 10.
			area   float64
			pi     float64
		)

		// ADD YOUR CODE HERE
		pi = math.Pi

		area = pi * (radius * radius)
		// DO NOT TOUCH THIS
		fmt.Printf("radius: %g -> area: %.2f\n", radius, area)
	*/

	//EX of numbers 8
	/*
		var rad, area float64
		radius := os.Args[1]

		rad, _ = strconv.ParseFloat(radius, 64)

		// ADD YOUR CODE HERE

		area = 4 * math.Pi * rad * rad
		// DO NOT TOUCH THIS
		fmt.Printf("radius: %g -> area: %.2f\n", rad, area)
	*/

	//EX of numbers 9

	var radius, vol float64

	// ADD YOUR CODE HERE
	radius, _ = strconv.ParseFloat(os.Args[1], 64)

	vol = (4 * math.Pi * math.Pow(radius, 3)) / 3

	// DO NOT TOUCH THIS
	fmt.Printf("radius: %g -> volume: %.2f\n", radius, vol)
}
