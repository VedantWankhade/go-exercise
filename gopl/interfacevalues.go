package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buf1 bytes.Buffer
	do(&buf1)
	var buf2 *bytes.Buffer
	fmt.Println(buf2 == nil)
	do(buf2)
	var buf3 io.Writer
	// buf3 = new(bytes.Buffer)
	fmt.Println(buf3 == nil)
	do(buf3)
}

func do(w io.Writer) {
	fmt.Printf("%T\n", w)
	fmt.Println("in do", w == nil)
}
