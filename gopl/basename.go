package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "/home/vedant/workspace/temp.sh.bck"
	s1 := "temp"
	basename1(s)
	basename2(s)
	basename1(s1)
	basename2(s1)
}

func basename1(s string) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	fmt.Println(s)
}

func basename2(s string) {
	slashIndex := strings.LastIndex(s, "/")
	if slashIndex != -1 {
		s = s[slashIndex+1:]
	}
	firstIndex := strings.Index(s, ".")
	if firstIndex != -1 {
		s = s[:firstIndex]
	}
	fmt.Println(s)
}
