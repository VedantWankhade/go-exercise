package main

import "fmt"

func main() {
    fmt.Println(anagram("listen", "silent"))
    fmt.Println(anagram("sad", "das"))
    fmt.Println(anagram("y", "yes"))
    fmt.Println(anagram("yet", "yes"))
}

func anagram(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    data := make(map[rune]int)
    for i:= 0; i < len(s1); i++ {
        data[rune(s1[i])]++
    }
for i:= 0; i < len(s1); i++ {
        data[rune(s2[i])]--
    }

    fmt.Println(data)
    for _, v := range data {
        if v != 0 {
            return false
        }
    } 

    return true
}
