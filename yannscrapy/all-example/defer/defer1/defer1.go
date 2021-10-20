package main

import "fmt"

func main() {
	a := f()
	fmt.Printf("a=%v\n", a)
}

func f() (result int) {
	defer func() {
		result++
	}()

	return 0
}
