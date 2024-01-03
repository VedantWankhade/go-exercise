package main

import (
    "fmt"
    "os"
)

func main() {
    // echo1()
    // echo2()
    betterEcho()
}

func echo1() {
    for i := 1; i < len(os.Args); i++ {
        fmt.Print(os.Args[i])
        if i != len(os.Args) - 1 {
            fmt.Print(" ")
        } else {
            fmt.Println()
        }
    }
}

func echo2() {
    for i, arg := range(os.Args[1:]) {
        fmt.Print(arg)
        // an extra -1 cuz its ranging over 1 less element than original os.Args
        if i != len(os.Args) - 1 -1 {
            fmt.Print(" ")
        } else {
            fmt.Println()
        }
    }
}

// betterEcho is a better implementation than above functions as it accumulates strings first and 
// prints everything in one go in the end instead of printing one word per loop iteration
// but uses more memory
func betterEcho() {
    res := ""
    postfix := " "
    for i, arg := range(os.Args[1:]) {
       res += arg + postfix
       if i == len(os.Args) - 3 {
            postfix = ""
       }
    }
    fmt.Println(res)
}
