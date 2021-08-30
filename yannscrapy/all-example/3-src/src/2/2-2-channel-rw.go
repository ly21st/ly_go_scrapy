// 2.2 读写通道

package main

import "fmt"

func main() {
	var ch chan int = make(chan int, 4)
	for i := 0; i < cap(ch); i++ {
		ch <- i // 写通道
	}
	for len(ch) > 0 {
		var value int = <-ch // 读通道
		fmt.Println(value)
	}
}
