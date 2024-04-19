package main

import "fmt"

func main() {
	var arr [3]int // all values initialized to zero
	fmt.Println("len", len(arr), "cap", cap(arr), arr)

	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println("len", len(arr1), "cap", cap(arr1), arr1)

	arr2 := [4]int{1, 3, 4}
	fmt.Println("len", len(arr2), "cap", cap(arr2), arr2)

	arr3 := [...]int{1, 3, 4}
	fmt.Println("len", len(arr3), "cap", cap(arr3), arr3)

	// following is invalid
	// n := 2
	// arr4 := [n]int{1, 3}
	// have to use const instead
	const n = 2
	arr4 := [n]int{1, 2}
	fmt.Println("len", len(arr4), "cap", cap(arr4), arr4)

	// array literal - we can give by index
	// arr5 := [3]string{2: "he", 0: "llo", 1: "ts"}

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	// d := [3]int{1, 2}
	// fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int

	x := []int{1, 2, 3}
	y := []int{1, 2, 3, 4}
	fmt.Println("len", len(x), "cap", cap(x), x[:3])
	fmt.Println("len", len(y), "cap", cap(y), y[:3])

	underlyingArr := [5]int{1, 2, 3, 4, 5}
	x1 := underlyingArr[:2]
	y1 := underlyingArr[:3]
	fmt.Println("len", len(x1), "cap", cap(x1), x1[:5])
	fmt.Println("len", len(y1), "cap", cap(y1), y1[:3])
}
