package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 1; i < 20; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	go func() {
		for a := range naturals {
			squares <- a * a
		}
		close(squares)
	}()

	for a := range squares {
		fmt.Println(a)
	}
}
