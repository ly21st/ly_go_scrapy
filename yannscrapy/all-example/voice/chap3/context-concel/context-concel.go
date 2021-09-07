package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	go func(ctx context.Context) {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				time.Sleep(2 * time.Second)
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("可以了，通知监控停止")

	cancel()
	wg.Wait() // 等待协程退出
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	// time.Sleep(5 * time.Second)
}
