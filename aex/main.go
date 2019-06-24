package main

import "path"

// EX. of assignment 1
/*
func main() {
	color := "green"

	color = "blue"

	fmt.Println(color)
}
*/

//EX. of assignment 2
/*
func main() {
	color := "green"

	color = "dark" + " " + color

	fmt.Println(color)
}
*/

//EX. of assignment 3
/*
func main() {
	n := 0.0

	n = 3.14 * 2

	println(n)
}
*/

//EX. of assignment 4
/*
func main() {
	var (
		perimeter     int
		width, height = 5, 6
	)

	perimeter = 2 * (width + height)

	println(perimeter)
}
*/

//EX. of assignment 5
/*
func main() {
	var (
		lang    string
		version int
	)

	lang, version = "go", 2

	println(lang, "version", version)

}
*/

//EX. of assignment 6
/*
func main() {
	var (
		plant  string
		isTrue bool
		temp   float64
	)

	plant, isTrue, temp = "Air is good on Mars", true, 19.5

	println(plant)
	println("It's ", isTrue)
	println("It's ", temp, " degree")
}
*/

//EX. of assignment 7
/*
func main() {

	_, b := multi()

	println(b)
}

func multi() (int, int) {
	return 5, 4
}
*/

//EX. of assignment 8
/*
func main() {

	color, color2 := "red", "blue"

	color, color2 = "orange", "green"

	println("color:", color, "color2:", color2)

}
*/

//EX. of assignment 9
/*
func main() {

	red, blue := "red", "blue"

	red, blue = blue, red

	println("Red:", red, "Blue:", blue)

}
*/

//EX. of assignment 10

func main() {

	dir, _ := path.Split("secret/file.text")

	println(dir)

}
