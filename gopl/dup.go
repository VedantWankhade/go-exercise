package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func main() {
	// dup1()
	BENCHMARK, err := strconv.ParseBool(os.Getenv("BENCHMARK"))
	if err != nil {
		log.Fatal("error parsing BENCHMARK: ", err)
	}
	if BENCHMARK {
		start := time.Now().UnixMilli()
		// dup2()
		// dup3()
		dup4()
		fmt.Printf("[Took %v]\n", time.Now().UnixMilli()-start)
	} else {
		dup2()
		dup3()
		dup4()
	}
}

// dup1 reads from stdin and prints repeated words
func dup1() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	hash := make(map[string]int)
	for scanner.Scan() {
		hash[scanner.Text()]++
	}
	for k, v := range hash {
		if v > 1 {
			fmt.Println(k + " -> " + strconv.Itoa(v))
		}
	}
}

// dup2 reads from files and prints repeated words or reads from stdin if no files given
func dup2() {
	hash := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		count(os.Stdin, hash)
	} else {
		for _, file := range files {
			fileReader, err := os.Open(file)
			if err != nil {
				fmt.Println("skipping file:", err)
			}
			count(fileReader, hash)
		}
	}
	for k, v := range hash {
		if v > 1 {
			fmt.Println(k + " -> " + strconv.Itoa(v))
		}
	}
}

func count(reader io.Reader, hash map[string]int) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords) // split words (on space) default would split on newline
	for scanner.Scan() {
		hash[scanner.Text()]++
	}
}

// same as dup2 but reads all files firsdt then counts
func dup3() {
	files := os.Args[1:]
	hash := make(map[string]int)
	var content bytes.Buffer
	for _, file := range files {
		bytes, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("skipping file ", err)
		}
		_, err = content.Write(bytes)
		if err != nil {
			fmt.Println("err writing to buffer ", err)
		}
	}
	fmt.Println(content.String())
	for _, token := range strings.Fields(content.String()) {
		hash[token]++
	}
	fmt.Println(hash)
}

// same as dup2 but also prints filename where duplicate words occures
func dup4() {
	hash := make(map[string]map[string]int)
	files := os.Args[1:]
	for _, file := range files {
		openFile, err := os.Open(file)
		if err != nil {
			fmt.Println("skipping file: error opening : ", err)
		}
		scanner := bufio.NewScanner(openFile)
		for scanner.Scan() {
			text := scanner.Text()
			if hash[text] == nil {
				hash[text] = make(map[string]int)
			}
			hash[text][file]++
		}
	}
	// fmt.Println(hash)
	for w, m := range hash {
		fmt.Print(w + " -> ")
		total := 0
		for f, n := range m {
			fmt.Print(path.Base(f) + "(" + strconv.Itoa(n) + ") ")
			total += n
		}
		fmt.Println("total(" + strconv.Itoa(total) + ")")
	}
}
