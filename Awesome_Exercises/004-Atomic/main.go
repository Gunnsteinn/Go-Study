package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("NumCPU\t", runtime.NumCPU())
	fmt.Println("NumGoroutine\t", runtime.NumGoroutine())

	var count int64
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&count,1)
			runtime.Gosched()
			fmt.Println("Count:", atomic.LoadInt64(&count))
			wg.Done()
		}()
		fmt.Println("NumGoroutine\t", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Cuenta:\t", count)
}
