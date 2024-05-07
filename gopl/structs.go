package main

import (
	"bytes"
	"fmt"
)

func main() {
	type Person struct {
		name string
		age  int
	}

	var p1 Person
	p2 := Person{}
	var p3 *Person
	p4 := &Person{}
	p5 := new(Person)

	fmt.Println(p1, p2, p3, p4, p5)
	fmt.Println(p3 == nil)

	var arr [4]Person
	var slice []Person
	fmt.Println(arr)
	fmt.Println(slice)

	var buf1 bytes.Buffer
	buf2 := bytes.Buffer{}
	var buf3 *bytes.Buffer
	fmt.Println(buf1)
	fmt.Println(buf2)
	fmt.Println(buf3)
}
