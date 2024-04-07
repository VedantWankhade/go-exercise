package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/vedantwankhade/go-exercise/gopl/random"
)

var OUTPUT = true

func exec(fn func([]string), args []string, name string) {
	// LOG, _ := strconv.ParseBool(os.Getenv("LOG"))
	OUTPUT = false
	log.Default().SetFlags(0)
	start := time.Now().UnixMilli()
	log.Printf("[%s]\n", name)
	fn(args)
	log.Printf("[Took %vms]\n", time.Now().UnixMilli()-start)
}

func main() {
	log.Default().SetFlags(0)
	args := os.Args[1:]
	BENCHMARK, _ := strconv.ParseBool(os.Getenv("BENCHMARK"))

	// benchmark
	if BENCHMARK {
		TESTNUMBER, _ := strconv.Atoi(os.Getenv("TESTNUMBER"))
		args = random.SliceOfRandomStringsWithLength(5, TESTNUMBER)
		exec(echo1, args, "echo1")
		// exec(echo2, args, "echo2")
		// exec(echo3, args, "echo3")
		// exec(echo4, args, "echo4")
		exec(echo5, args, "echo5")
	} else {
		echo1(args)
		// echo2(args)
		// echo3(args)
		// echo4(args)
		echo5(args)
	}
}

// echo1 collects all args into a strings and prints it out at once
func echo1(args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	if OUTPUT {
		fmt.Println(s)
	}
}

// echo2 iterates through the args and prints then one by one
func echo2(args []string) {
	sep := " "
	for _, arg := range args {
		if OUTPUT {
			fmt.Print(arg + sep)
		}
	}
	if OUTPUT {
		fmt.Println()
	}
}

func echo3(args []string) {
	if OUTPUT {
		fmt.Println(args)
	}
}

func echo4(args []string) {
	if OUTPUT {
		fmt.Println(strings.Join(args, " "))
	}
}

// faster version of echo1
func echo5(args []string) {
	var out bytes.Buffer
	for _, arg := range args {
		out.WriteString(arg + " ")
	}
	if OUTPUT {
		fmt.Println(out.String())
	}
}
