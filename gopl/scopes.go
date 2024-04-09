package main

import "fmt"

// at package level, the order of declarations do not matter
// even for variable declarations
// hence this is valid
var m = n
var n = 22

func main() {
	// why is this allowed 😭
	// three x variables in different scopes
	x := "hello"
	for _, x := range x {
		var x = x + 'A' - 'a'
		// x := x + 'A' - 'a' // same thing
		fmt.Printf("%c\n", x)
	}
}
