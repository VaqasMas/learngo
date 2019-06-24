package main

import (
	"fmt"
	"os"
)

func main() {

	//EX. of ARGS 1
	/*
		count := len(os.Args)

		fmt.Printf("There are %d names.\n", count)
	*/
	/*
		count := len(os.Args) - 1

		fmt.Printf("There are %d names.\n", count)
	*/

	//EX. of ARGS 2
	/*
		path := os.Args

		fmt.Printf("Path: ", path)
	*/
	//EX of ARGS 3
	/*
		name := os.Args[1]

		fmt.Println("Hi ", name)
		fmt.Println("How are you")
	*/

	//EX of ARGS 4
	/*
		nameO, nameT, nameTH := os.Args[1], os.Args[2], os.Args[3]

		count := len(os.Args) - 1

		fmt.Println("There are ", count, " people!")
		fmt.Println("Hello great", nameO, "!")
		fmt.Println("Hello great", nameT, "!")
		fmt.Println("Hello great", nameTH, "!")
		fmt.Println("Nice to meet you all.", "!")
	*/

	//EX of ARGS 5

	fmt.Println("There are ", len(os.Args)-1, "people!")
	fmt.Println("Hello great ", os.Args[1], "!")
	fmt.Println("Hello great ", os.Args[2], "!")
	fmt.Println("Hello great ", os.Args[3], "!")
	fmt.Println("Hello great ", os.Args[4], "!")
	fmt.Println("Hello great ", os.Args[5], "!")
	fmt.Println("Nice to meet you all")

}
