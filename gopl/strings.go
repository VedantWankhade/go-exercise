package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello world"
	fmt.Printf("s=%s typeof s=%[1]T\ns[0]=%s typeof s[1]=%[2]T\n", s, string(s[0]))

	x := "h྅"
	fmt.Println(x)
	fmt.Println(len(x))                    // number of bytes
	fmt.Println(utf8.RuneCountInString(x)) // numebr of characters / runes
	fmt.Println(string(x[0]), " ", string(x[1]), " ", string(x[2]), " ", string(x[3]))
	fmt.Printf("%b %b %b %b\n", x[0], x[1], x[2], x[3])

	for _, c := range x {
		fmt.Print(c, " ")
	}
	fmt.Println()

	for _, c := range x {
		fmt.Print(string(c), " ")
	}
	fmt.Println()

	y := "hello world"
	y1 := y[1:5]
	fmt.Println(y1)

	for _, c := range y1 {
		fmt.Print(c, " ")
	}
	fmt.Println()
	for _, c := range y1 {
		fmt.Print(string(c), " ")
	}
	fmt.Println()
	for _, c := range y1 {
		fmt.Printf("%08b ", c)
	}
	fmt.Println()
	// raw stings - taken literaly
	z := `hello ྅ \n\t 
		yes
	`
	fmt.Println(z)
}
