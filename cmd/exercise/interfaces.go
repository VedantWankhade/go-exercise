package main

import "fmt"

type ByteCounter int

// satisfies io.Writer
func (b *ByteCounter) Write(bytes []byte) (n int, err error) {
	*b += ByteCounter(len(bytes))
	return len(bytes), nil
}

// satisfies string interface - the receiver is not pointer as we dont want to change it
// if we set the receiver as pointer like
// func (b *ByteCounter) String() string
// then when calling any of Sprintf, Println etc we have to pass the address of bytecounter obj like so
// fmt.Println(&bc) otherwise it will just print the vanilla string
func (b ByteCounter) String() string {
    return fmt.Sprintf("[bytecouter: %d]", int(b))
}

func main() {
	var bc ByteCounter = 0
	fmt.Println(bc)
	bc.Write([]byte{1, 2, 3})
	fmt.Println(bc)

	fmt.Fprintf(&bc, "random strings", bc)
	fmt.Println(bc)
}
