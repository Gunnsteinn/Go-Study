package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("chqueo de error 1:", ctx.Err())
	fmt.Println("num gorutinas 1:", runtime.NumGoroutine())

	for value := range generate(ctx) {
		fmt.Println(value)
		if value == 5 {
			break
		}
	}


	fmt.Println("chequeo de error:", ctx.Err())
	fmt.Println("num gorutinas 2:", runtime.NumGoroutine())

	fmt.Println("A punto de cancelar context.")
	cancel()
	fmt.Println("context cancelado.")

	time.Sleep(time.Millisecond * 100)
	fmt.Println("chequeo de error 3:", ctx.Err())
	fmt.Println("num gorutinas 3:", runtime.NumGoroutine())
}

func generate(ctx context.Context) <-chan int {
	c := make(chan int)
	n := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case c <- n:
				n++
			}
		}
	}()
	return c
}
