package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	tests := []string{"", "abc", "1", "12", "123", "1234", "12345", "123456", "1234567", "12345678"}
	for _, t := range tests {
		ans, err := nonRecComma(t)
		if err != nil {
			fmt.Print(err, " ")
		}
		fmt.Println(t, "->", ans)
		// }
	}
}

func notNum(s string) bool {
	return strings.ContainsFunc(s, func(r rune) bool {
		if !unicode.IsDigit(r) {
			return true
		}
		return false
	})
}

func comma(s string) (string, error) {
	s = strings.TrimSpace(s)
	if len(s) < 1 || notNum(s) {
		return "", errors.New("Invalid input")
	}
	if len(s) < 4 {
		return s, nil
	}
	s1, err := comma(s[:len(s)-3])
	if err != nil {
		return "", err
	}
	return s1 + "," + s[len(s)-3:], nil
}

// non recursive version of comma using bytes.Buffer
func nonRecComma(s string) (string, error) {
	s = strings.TrimSpace(s)
	if len(s) < 1 || notNum(s) {
		return "", errors.New("Invalid input")
	}
	if len(s) < 4 {
		return s, nil
	}
	var buf bytes.Buffer
	commaNums := len(s) % 3
	i := commaNums
	for i < len(s) {
		buf.WriteString(s[:i])
		buf.WriteByte(',')
		i += 3
	}

	return buf.String(), nil
}
