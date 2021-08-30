package main

import (
	"fmt"
	"runtime"
	"syscall"
	"time"
)

func GetCurrentThreadId() int {
	var user32 *syscall.DLL
	var GetCurrentThreadId *syscall.Proc
	var err error

	user32, err = syscall.LoadDLL("Kernel32.dll")
	if err != nil {
		fmt.Printf("syscall.LoadDLL fail: %v\n", err.Error())
		return 0
	}
	GetCurrentThreadId, err = user32.FindProc("GetCurrentThreadId")
	if err != nil {
		fmt.Printf("user32.FindProc fail: %v\n", err.Error())
		return 0
	}

	var pid uintptr
	pid, _, err = GetCurrentThreadId.Call()

	return int(pid)
}

func main() {
	runtime.GOMAXPROCS(3)
	// 读取当前的线程数
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println("run in main goroutine")
	n := 10
	for i := 0; i < n; i++ {
		go func() {
			fmt.Println("dead loop goroutine start, threadId:", GetCurrentThreadId())
			for {

			} // 死循环
			fmt.Println("dead loop goroutine stop")
		}()
	}

	go func() {
		var count = 0
		for {
			time.Sleep(time.Second)
			count++
			fmt.Println("for goroutine running:", count, "threadId:", GetCurrentThreadId())
		}
	}()
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
	var count = 0
	for {
		time.Sleep(time.Second)
		count++
		fmt.Println("main goroutine running:", count, "threadId:", GetCurrentThreadId())
	}
}
