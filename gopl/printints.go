package main

import (
	"bytes"
	"fmt"
)

func main() {
	values := []int{1, 4, 6, 7}
	fmt.Println(printints(values))
}

func printints(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteRune(']')
	return buf.String()
}
