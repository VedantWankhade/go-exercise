// better fizzbuzz
package main

import (
	"fmt"
	"strconv"
)

type fiz map[int]string

func (b *fiz) with(n int, s string) *fiz {
	(*b)[n] = s
	return b
}

func main() {
	// game := (&fiz{}).with(3, "Fizz").with(5, "Buzz")
	game := (&fiz{}).with(3, "Fizz").with(5, "Buzz").with(7, "Foo").with(9, "Bar")
	for i := 1; i <= 100; i++ {
		out := ""
		for k, v := range *game {
			if i%k == 0 {
				out += v
			}
		}
		if out == "" {
			out = strconv.Itoa(i)
		}
		fmt.Println(out)
	}
}
