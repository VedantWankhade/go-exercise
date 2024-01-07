package main

import "fmt"

func main() {
    arr := [4]int{1, 2, 3, 4}
    // fmt.Println(arr)
    //do(&arr)
    //fmt.Println(arr)
    //i := 5
    //d(&i)
    fmt.Println(arr)
    reverse(&arr)
    fmt.Println(arr)
}

func do(arr *[4]int) {
    (*arr)[2] = 99
    fmt.Println(arr)
    fmt.Println(*arr)
    fmt.Printf("%T", arr)
    fmt.Printf("\n%T", *arr)
}

func d(i *int) {
    fmt.Println(i)
    fmt.Println(*i)
}

func reverse(arr *[4]int) {
    for i, j := 0, len(arr) - 1; i < j; i, j = i + 1, j - 1 {
        arr[i], arr[j] = arr[j], arr[i]    
    }
}
