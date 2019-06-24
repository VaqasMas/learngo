package main

import (
	"fmt"
	"os"
)

func main() {

	//EX of Printf 1
	/*
		var age = 34

		fmt.Printf("I'm %d years old.\n", age)
	*/

	//EX of Printf 2
	/*
		fname := "vaqas"
		lname := "masood"

		fmt.Printf("My name is %s and my lastname is %s.", fname, lname)
	*/

	//EX of Printf 3
	/*
		tf := false
		fmt.Printf("These are %t claims.", tf)
	*/

	//EX of Printf 4
	/*
		fmt.Printf(" Temperature is %.1f degrees.", 29.51)
	*/

	//EX of Printf 5
	/*
		fmt.Printf("%q", "hello world")
	*/

	//EX of Printf 6
	/*
		fmt.Printf("Type of %d is %[1]T", 3)
	*/

	//EX of Printf 7
	/*
		fmt.Printf("Type of %.2f is %[1]T", 3.14)
	*/

	//EX of Printf 8
	/*
		fmt.Printf("Type of %s is %[1]T", "hello")
	*/

	//EX of Printf 9
	/*
		fmt.Printf("Type of %t is %[1]T", true)
	*/

	//EX of Printf 10
	fname := os.Args[1]
	lname := os.Args[2]

	fmt.Printf("Your first name is %[1]s and your last name is %[2]s.",
		fname, lname)

}
