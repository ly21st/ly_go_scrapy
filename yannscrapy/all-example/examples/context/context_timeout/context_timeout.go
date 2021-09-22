package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	const timeout = 5 *time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := shutdown(ctx)
	fmt.Printf("err=%+v\n", err)
}

func shutdown(ctx context.Context) error {
	ch := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg.Wait()
		close(ch)
	}()

	select {
	case <-ch:
		return nil
	case <-ctx.Done():
		return errors.New("timeout")
	}


}

































