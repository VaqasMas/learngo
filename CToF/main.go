package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	cel, _ := strconv.ParseFloat(arg, 64)

	Fer := cel*1.8 + 32

	fmt.Printf("%g C x 1.8 + 32 = %g F", cel, Fer)
}
