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
		for {
			a, ok := <-naturals
			if ok {
				squares <- a * a
			} else {
				close(squares)
				break
			}
		}
	}()

	for {
		a, ok := <-squares
		if ok {
			fmt.Println(a)
		} else {
			break
		}
	}
}
