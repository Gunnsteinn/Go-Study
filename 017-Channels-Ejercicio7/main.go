// |****************************************|
// |				EJERCICIO 7				|
// |****************************************|
package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for c := range gen(ctx) {
		fmt.Println(c)
		if c == 100 {
			break
		}
	}
	fmt.Println("1 - NumGoroutine\t", runtime.NumGoroutine())

	cancel()

	time.Sleep(time.Millisecond * 100)
	fmt.Println("2 - NumGoroutine\t", runtime.NumGoroutine())
	fmt.Println("-- THE END --")
}

func gen(ctx context.Context) <-chan int {
	c1 := make(chan int)
	n := 0
	const numOfGoRoutines = 10


	for i := 0; i < numOfGoRoutines; i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case c1 <- n:
					n++
				}
			}
		}()
	}

	return c1
}
