package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// Banger Exercise.
	/*
		msg := os.Args[1]
		l := len(msg)
		mark := strings.Repeat("!", l)

		s := mark + msg + mark
		s = strings.ToUpper(s)

		fmt.Println(s)
	*/

	// EX of Strings 1
	/*
	   	path := `c:\program files\duper super\fun.txt
	   c:\program files\really\funny.png`
	   	fmt.Println(path)
	*/

	// EX of Strings 2
	/*
			json := `{
			"Items": [{
				"Item": {
					"name": "Teddy Bear"
				}
			}]
		}`

			fmt.Println(json)
	*/

	// EX of Stings 3
	/*
	   	// uncomment the code below
	   	name := os.Args[1]

	   	// replace and concatenate the `name` variable
	   	// after `hi ` below

	   	msg := `hi ` + name + `!
	   how are you?`

	   	fmt.Println(msg)
	*/

	// EX of strings 4
	/*
		length := utf8.RuneCountInString(os.Args[1])
		fmt.Println(length)
	*/
	// EX of Strings 5
	/*
		msg := os.Args[1]

		s := msg + strings.Repeat("!", utf8.RuneCountInString(msg))

		fmt.Println(s)
	*/

	// EX of Strings 6
	/*
		msg := os.Args[1]
		msg = strings.ToLower(msg)

		fmt.Println(msg)
	*/

	//EX of Strings 7
	/*
	   	msg := `

	   	The weather looks good.
	   I should go and play.
	   	`

	   	fmt.Println(strings.TrimSpace(msg))
	*/

	//EX of Strings 8
	name := "inanç           "

	name = strings.TrimRight(name, " ")
	fmt.Println(name)
	fmt.Println(utf8.RuneCountInString(name))
}
