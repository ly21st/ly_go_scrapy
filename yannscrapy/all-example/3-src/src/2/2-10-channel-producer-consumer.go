// 2-10 生产者、消费者模型
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 生产者
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64)
	go Producer(3, ch) // 生成3的倍数序列
	go Producer(5, ch) // 生成5的倍数序列

	go Consumer(ch)

	//Ctrl +C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
