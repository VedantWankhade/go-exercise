package main

import "fmt"

func main() {
    f := squares()
    fmt.Printf("squares: %T, %v\n", squares, squares)
    fmt.Printf("return value of squares - f: %T, %v\n", f, f)
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())

    fmt.Println("---------------------")
    fmt.Println(squares()())
    fmt.Println(squares()())
    fmt.Println(squares()())
    fmt.Println(squares()())
}

func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}
