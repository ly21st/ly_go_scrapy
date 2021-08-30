// 2-5 通道写安全
package main

import "fmt"

func send(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)
}

func recv(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

// 确保通道写安全的最好方式是由负责写通道的协程自己来关闭通道，读通道的协程不要去关闭通道。
func main() {
	var ch = make(chan int, 1)
	go send(ch)
	recv(ch)
}
