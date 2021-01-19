package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	pair := make(chan int)
	odd := make(chan int)

	// enviar
	go sender(pair, odd)

	// recibir
	receiver(ctx, pair, odd)
	fmt.Println("-- Time to Execute--", time.Now().Sub(start))
	fmt.Println("-- THE END--")
}

func sender(p, o chan<- int) {
	for i := 1; ; i++ {
		if i%2 == 0 {
			p <- i
		} else {
			o <- i
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func receiver(ctx context.Context, p, o <-chan int) {
	for {
		select {
		case v := <-p:
			fmt.Printf("%v is Pair\n", v)
		case v := <-o:
			fmt.Printf("%v is Odd\n", v)
		case t, ok := <-ctx.Done():
			if !ok {
				fmt.Println(t, ok)
				return
			}
			fmt.Println(t, ok)
		}
	}
}
