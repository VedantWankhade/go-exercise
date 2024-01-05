package main

import (
	"bytes"
	"fmt"
)

func main() {
	// fmt.Println(naturalInteger("2134524"))
    // fmt.Println(rune(string([]byte("hello"))[2]))
    // fmt.Println(intsToString([]int{1, 2, 3}))
    fmt.Println(naturalIntegerNR("1213424"))
}

func naturalInteger(num string) string {
    n := len(num)
	if n <= 3 {
		return num
	}
    return naturalInteger(num[:n-3]) + "," + num[n-3:]
}


// uses bytes.Buffer
func intsToString(values []int) string {
    var buf bytes.Buffer
    buf.WriteByte('[')
    for i, v := range values {
        if i > 0 {
            buf.WriteString(", ")
        }
        fmt.Fprintf(&buf, "%d", v)
    }
    buf.WriteByte(']')
    return buf.String()
}

// non recursive impl of naturalInteger using bytes.Buffer
func naturalIntegerNR(num string) string {
    var buf bytes.Buffer
    n := len(num)
    fmt.Printf("[string=%s]\n", num)
    fmt.Printf("[len: %d]\n", n)
    pre, rest := n / 3, n % 3
    if rest == 0 {
        rest = 3
    }
    buf.WriteString(num[:rest])
    fmt.Printf("[pre: %d, rest: %d]\n", pre, rest)
    for i := rest; i < n; i += 3 {
        buf.WriteString(",")
        buf.WriteString(num[i:i+3])
    }
   return buf.String()
}
