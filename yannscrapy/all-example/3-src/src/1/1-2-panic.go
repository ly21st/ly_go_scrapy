package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("run in main goroutine")
	go func() {
		fmt.Println("run in child goroutine")
		var ptr *int
		*ptr = 0x12345 // 故意制造崩溃
		go func() {
			fmt.Println("run in grand child goroutine")
			go func() {
				fmt.Println("run in grand grand child goroutine")

			}()
		}()
	}()
	time.Sleep(time.Second)
	fmt.Println("main goroutine will quit")
}
