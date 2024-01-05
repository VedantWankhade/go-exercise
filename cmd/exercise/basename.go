package main

import (
	"fmt"
	"strings"
)

func main() {
   fmt.Println(basename2("/tmp/dir/file.txt")) 
}

func basename1(s string) string {
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '/' {
            s = s[i+1:]
            break
        }
    }
    for i := len(s) - 1; i >=0; i-- {
        if s[i] == '.' {
            s = s[:i]
            break
        }
    }
    return s
}

func basename2(s string) string {
    slashIndex := strings.LastIndex(s, "/")
    s = s[slashIndex + 1:]
    if dotIndex := strings.LastIndex(s, "."); dotIndex >= 0 {
        s = s[:dotIndex]
    }
    return s
}
