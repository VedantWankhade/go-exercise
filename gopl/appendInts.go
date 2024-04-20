package main

import "fmt"

// be carefull while creating slices off of an array.
// modifying the slice will also modify the underlying array
// whenever possible just use a := []int{} or make([]int, len) to create an isolated slice
func main() {
	arr := [10]int{1, 2, 3}
	a := arr[:5]
	fmt.Println(a)
	b := appendInts(a, 99) // this may also modify a
	fmt.Println("a", a)
	fmt.Println("b", b)
}

func appendInts(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[2] = 22 // to check if the original slice is modified
	z[len(x)] = y
	return z
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
