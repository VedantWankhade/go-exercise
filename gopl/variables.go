package main

import "fmt"

// package level variables
// go compiler doesnt check if package level variabes are used are not 😑
var a int = 2
var b string
var c = 99.1
var l, z = 22, "hi"

// invalid
// a := 2  // compiler error since this statement is an assignment not declaration
func main() {
	var d float64 // this is declaration
	var e = 2     // this is declartion but also assigns a value
	f := "hello"  // this is an assignment statement (called short declaration)
	var i, j, k int
	// p string := "geel" // invalid (explicit types can onbly be given when declaring (with var))
	var m, n = true, "hello"

	// for local variables short declarations are prefered
	// but var declaration are used when assigned types differe from what we want
	// example - 22.1 literal is float64 by default, what is we want our variable to be float32
	g := 22.1
	var h float32 = 22.1

	// var q int = 22.3 // invalid since float literal cannot be assinged to int variable
	// var q int = int(22.3) // illegal, do the following instead
	Q := 22.3
	var q int = int(Q)

	type day int
	var someday int = 2
	type month int
	var jan month = 1
	var today day
	fmt.Println(today)
	today = 22
	// fmt.Println(today + jan) // invalid
	fmt.Println(today + 2) // valid
	// fmt.Println(today + someday) // invalid
	fmt.Println(today + day(someday)) // valid
	as := today + 2                   // as takes type of day since 2 also takes type day according to the context of expression
}
