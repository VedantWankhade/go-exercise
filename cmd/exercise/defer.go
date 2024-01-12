package main

import (
	"fmt"
	"time"
)

func slowoperation() {
    // the following statement calls the trace function and executes it the retuen value taht is the function returned by trace is hold on and executed after the slooperation funmctgion exitss
    defer trace("slowoperation")() // defer the call to function returned by trace
    time.Sleep(10 * time.Second) // simulate slow function
}

func trace(msg string) func() {
    start := time.Now()
    fmt.Printf("enter %s\n", msg)
    return func() {
        fmt.Printf("exit %s (%s)\n", msg, time.Since(start))
    }
}

func main() {
    slowoperation()
    double(2)
    // demonstrate deferred function runs after panic
    f(3)
}

func double(x int) (res int) {
    // following defered anon function can access the named result (res) after it is evaluated as res = x + x
    defer func() {
        fmt.Printf("double(%d) = %d\n", x, res)
    }()

    return x + x
}

// after panic, deferred functions are called in reverse order - top to bottom of stack
func f(x int) {
    fmt.Printf("f(%d)\n", x+0/x) // panics at x = 0
    defer fmt.Printf("defer %d\n", x)
    f(x-1)
}
