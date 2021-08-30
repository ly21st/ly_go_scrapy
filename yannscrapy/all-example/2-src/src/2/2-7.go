// 2.7 interface类型转换

package main

import "fmt"

func main() {
	var x interface{}

	s := "darren"
	x = s
	y, ok := x.(int)
	z, ok1 := x.(string)
	fmt.Println(y, ok)
	fmt.Println(z, ok1)
}

//0 false
//darren true
