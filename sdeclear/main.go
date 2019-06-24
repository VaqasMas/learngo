package main

import "fmt"

//EX. of short declear 1
/*
func main() {
	i, f, s, b := 10, 64.5, "Hello", true

	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
	)
}
*/

//EX. of short declear 2
/*
func main() {
	a := 14
	b := true

	fmt.Println(a, b)
}
*/

//EX. of short declear 3
/*
func main() {
	a, c := 42, "good"

	fmt.Println(a, c)
}
*/

// EX. of short declear 4
/*
func main() {
	sum := 27 + 3.5

	fmt.Println(sum)
}
*/

// EX. of short declear 5
/*
func main() {
	on, off := true, true

	_ = off

	fmt.Println(on)
}
*/

//EX. of short declear 6

func main() {
	age, yourAge := 15, 20

	age, ratio := 42, 3.14

	fmt.Println(age, yourAge, ratio)
}
