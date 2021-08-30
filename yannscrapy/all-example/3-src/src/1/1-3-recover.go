// 1.3 协程异常处理-recover

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("run in main goroutine")
	go func() {

		fmt.Println("run in child goroutine")
		go func() {
			fmt.Println("run in grand child goroutine")
			go func() {
				defer func() { // 要在对应的协程里执行
					fmt.Println("执行defer:")
					if err := recover(); err != nil {
						fmt.Println("捕获error:", err)
					}
				}()
				fmt.Println("run in grand grand child goroutine")
				var ptr *int
				*ptr = 0x12345 // 故意制造崩溃
			}()
		}()
	}()
	time.Sleep(time.Second)
	fmt.Println("main goroutine will quit")
}
