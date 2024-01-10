package main

import "fmt"

func main() {
    arr := []int{1, 2, 3}
    // fmt.Printf("arr=%v len=%v cap=%v\n", arr, len(arr), cap(arr))
    // fmt.Println(arr[:])
    // fmt.Println(arr[:cap(arr)])
    // arr = append(arr, 99)
    // fmt.Printf("arr=%v len=%v cap=%v\n", arr, len(arr), cap(arr))
    // fmt.Println(arr[:])
    // fmt.Println(arr[:cap(arr)])
    //fmt.Println(appendMany(arr, 99, 98, 97))
    fmt.Println(remove(arr, 1))
}

func _append(arr []int, num int) []int {
    var res []int
    resLen := len(arr) + 1
    if resLen <= cap(arr) {
        // can append without growing the inner array
        res = arr[:resLen]
    } else {
        resCap := resLen
        if resCap < 2 * len(arr) {
            resCap = 2 * len(arr)
        }
        res = make([]int, resLen, resCap)
        copy(res, arr)
    }
    res[len(arr)] = num
    return res
} 

func appendMany(arr []int, nums ...int) []int { 
    var res []int
    resLen := len(arr) + len(nums)
    if resLen <= cap(arr) {
        // can append without growing the inner array
        res = arr[:len(arr)]
    } else {
        resCap := resLen
        res = make([]int, resLen, resCap)
        copy(res, arr)
    }
    copy(res[len(arr):], nums)
    return res 
}

func remove(arr []int, i int) ([]int, int) {
    ret := arr[i]
    copy(arr[i:], arr[i + 1 : ])
    return arr[:len(arr) - 1], ret
}
