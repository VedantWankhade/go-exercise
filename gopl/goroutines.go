package main

import "fmt"

func print(s string) {
	for {
		fmt.Println(s)
	}
}

func main() {
	go print("thread #1")
	go print("thread #2")
	// an empty infinte loop to let the other go routunines run
	// otherwise the main go routine will return shutting down other go routines
	for {

	}
}
