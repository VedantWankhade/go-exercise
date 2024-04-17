package main

import "fmt"

func main() {
	var a byte = 5
	var b uint8 = 5
	fmt.Printf("%08b %08b\n", a, b)
	fmt.Printf("%d %d\n", a, b)
	fmt.Println(a + b)
}
