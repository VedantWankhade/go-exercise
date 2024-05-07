package main

import "fmt"

func naturals(out chan<- int) {
	for i := 1; i < 20; i++ {
		out <- i
	}
	close(out)
}

func squares(in <-chan int, out chan<- int) {
	for a := range in {
		out <- a * a
	}
	close(out)
}

func printer(in <-chan int) {
	for a := range in {
		fmt.Println(a)
	}
}

func main() {
	nat := make(chan int)
	sq := make(chan int)
	go naturals(nat)
	go squares(nat, sq)
	printer(sq)
}
