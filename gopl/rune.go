package main

import "fmt"

func main() {
	ascii := 'a'
	unicode := 'ǣ'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	// "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 D 'D'"
	fmt.Printf("%d %[1]q\n", newline)
	// "10 '\n'"
}
