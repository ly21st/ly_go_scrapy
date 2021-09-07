package main

import "fmt"

func main() {
	a := f()
	fmt.Println(a)
}

func f() (result int) {
	defer func() {
		result++
	}()

	return 0
}
