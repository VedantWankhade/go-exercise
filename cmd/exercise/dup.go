package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
    dup4()    
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

func dup4() {
    counts := make(map[string]map[string]int)
    files := os.Args[1:]
    for _, file := range files {
        ofile, err := os.Open(file)
        if err != nil {
            fmt.Println("Error reading file", file)
            continue
        }
        in := bufio.NewScanner(ofile)
        for in.Scan() {
            if counts[in.Text()] == nil {
                counts[in.Text()] = make(map[string]int)
            }
            counts[in.Text()][string(file)]++
        }
    }
    for line, data := range counts {
        totalCount := 0
        for _, count := range data {
            totalCount += count
        }
        if totalCount > 1 {
            fmt.Printf("Line: %s Total Count: %d\n", line, totalCount)
            for filename, count := range data {
                fmt.Printf("File: %v Count: %d\n", filename, count)
            }
        }
    }
}
