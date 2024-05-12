package main

import (
	"errors"
	"fmt"
)

type FileError struct{}

func (f FileError) Error() string {
	return "fileerror"
}

type NotFileError struct{}

func (f NotFileError) Error() string {
	return "NOTfileerror"
}

var (
	MyError    FileError    = FileError{}
	NotMyError NotFileError = NotFileError{}
)

func someFunction() error {
	return MyError
}

func notSomeFunction() error {
	return NotMyError
}

func main() {
	err := someFunction()
	fmt.Printf("error string: %s, error type: %T\n", err, err)
	// fmt.Println(err == MyError)
	dynErr, ok := err.(FileError)
	fmt.Println(dynErr, ok)
	fmt.Println(errors.Is(err, FileError{}))

	err1 := notSomeFunction()
	fmt.Printf("error string: %s, error type: %T\n", err1, err1)
	// fmt.Println(err1 == MyError)
	dynErr1, ok := err1.(FileError)
	fmt.Println(dynErr1, ok)
	fmt.Println(errors.Is(err1, FileError{}))
}
