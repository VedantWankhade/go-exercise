package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 5}
	fmt.Println(a)
	a, _ = remove(a, 2)
	fmt.Println(a)
	fmt.Println(a[:len(a)+1])

	var b []int // slices are initialized to nil
	fmt.Println(b == nil)

	c := []int{}        // initialized using literal - not nil
	d := []int{1, 2, 3} // initialized using literal - not nil
	fmt.Println(c == nil)
	fmt.Println(d == nil)

	e := make([]int, 4) // memory allocation - not nil
	fmt.Println(e == nil)
}

func remove(slice []int, index int) ([]int, error) {
	if index < 0 || index >= len(slice) {
		return nil, fmt.Errorf("Index out of bound")
	}
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1], nil
}
