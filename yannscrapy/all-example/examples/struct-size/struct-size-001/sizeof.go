package main

import (
	"fmt"
	"unsafe"
)

type AA struct {
	name string
	age  int
	buf  [5]byte
	addr string
	buf2 [5]byte
}

func main() {
	fmt.Printf("sizeof(AA)=%0b\n", unsafe.Sizeof(AA{}))

	fmt.Printf("0x%0x\n", uintptr(-int(unsafe.Sizeof(AA{}))&15))

	// x + uintptr(-int(x)&7)
	// x + uintptr(-int(x)&7)
}
