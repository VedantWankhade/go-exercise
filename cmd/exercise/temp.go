package main

import (
	"flag"
	"fmt"

	"github.com/VedantWankhade/go-exercise/cmd/exercise/temp"
)

func main() {
	// var c Celsius = 120.9
	// var f Fahrenheit = 240.9
	// following gives compile error as c and f (Celsius and Fahrenheit are not of same type
	// fmt.Println(c + f)

	var temp = temp.CelsiusFlag("temp", 19.0, "the temperature")
	flag.Parse()
	fmt.Println(*temp)
}
