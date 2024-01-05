package main

import "fmt"

// Celsius and Fahrenheit; even though have same underlying type, they are not of same type themselves
type Celsius float64
type Fahrenheit float64

func main() {
    // var c Celsius = 120.9
    // var f Fahrenheit = 240.9
    // following gives compile error as c and f (Celsius and Fahrenheit are not of same type
    // fmt.Println(c + f)

}

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC Celsius = 0
    BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
    // Fahrenheit() is not a function; it's a conversion
    // but its only changing the meaning (type) explicitly not any value
    return Fahrenheit(c * 9 / 5 + 32)
}

func FToC(f Fahrenheit) Celsius {
    // Celsius() is not a function; it's a conversion
    // but its only changing the meaning (type) explicitly not any value
    return Celsius((f - 32) * 5 / 9)
}
