package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("NumCPU\t", runtime.NumCPU())
	fmt.Println("NumGoroutine\t", runtime.NumGoroutine())
	count := 0
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := count
			v++
			runtime.Gosched()
			count = v
			wg.Done()
		}()
		fmt.Println("NumGoroutine\t", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Cuenta:\t", count)
}
