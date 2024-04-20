package main

import "fmt"

func main() {
	var a map[int]string // initialized to nil
	fmt.Println(a == nil)

	// initialized using literal - not nil
	b := map[int]string{
		1: "hi",
		2: "ajdnb",
	}
	d := map[int]string{}
	fmt.Println(b == nil)
	fmt.Println(d == nil)

	// memory allocation - not nil
	c := make(map[int]string)
	fmt.Println(c == nil)

	// remove element from map
	delete(b, 1)
	fmt.Println(b)
}
