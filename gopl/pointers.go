package main

import "fmt"

func main() {
	x := 2
	y := 2
	y = x
	p := &x
	q := &y

	fmt.Println(p, q)

	var a = f()
	fmt.Println(a, "->", *a)
	var b = f()
	fmt.Println(b, "->", *b)

	// f() will always create a distinct value
	fmt.Println(f() == f())

	// new function creates a variable of type T and returns its address
	t := new(int)
	fmt.Println(*t)
	*t = 2
	fmt.Println(*t)

	// theres no pointer to pointer
	// xx := &p // illegal
}

func f() *int {
	v := 1
	fmt.Println("in f", &v)
	// value of v will still be in memory even when f() is removed from stack
	return &v
}
