package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func work(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(ctx, &wg)
	}
	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}
