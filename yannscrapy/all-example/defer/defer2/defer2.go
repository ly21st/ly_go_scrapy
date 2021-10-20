package main

import "fmt"

func main() {
	a := f2()
	fmt.Printf("a=%v\n", a)

	a = f3()
	fmt.Printf("a=%v\n", a)
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
