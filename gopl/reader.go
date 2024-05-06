package main

import "fmt"

type MyReader interface {
	Read() string
}

type StringReader string

func (sr *StringReader) Read() string {
	return string(*sr)
}

func getMyReader(s string) MyReader {
	h := StringReader(s)
	return &h
}

func read(r MyReader) {
	fmt.Println(r.Read())
}

func main() {
	r := getMyReader("hello")
	read(r)

	s := LimitReader{
		n:      20,
		reader: getMyReader("world"),
	}
	s.Read()
}

type LimitReader struct {
	reader MyReader
	n      int
}

func (lr *LimitReader) Read() {
	lr.reader.Read()
	fmt.Println("printed with limitreader")
}
