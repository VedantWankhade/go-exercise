package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
    dup2()    
}

func dup1() {
    counts := make(map[string]int)
    in := bufio.NewScanner(os.Stdin)
    for in.Scan() {
        counts[in.Text()]++
    }
    for lineIndex, count := range counts {
        if count > 1 {
            fmt.Println(lineIndex, count)
        }
    }
}

func dup2() {
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func countLines(f *os.File, counts map[string]int) {
    in := bufio.NewScanner(f)
    for in.Scan() {
        counts[in.Text()]++
    }
}

// dup3 reads whole content as one then splits into lines
func dup3() {
    counts := make(map[string]int)
    for _, filename := range  os.Args[1:] {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
        for _, line := range strings.Split(string(data), "\n") {
            counts[line]++
        }
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Println(n, line)
        }
    }
}
