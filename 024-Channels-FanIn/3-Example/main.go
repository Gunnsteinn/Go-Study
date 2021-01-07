// |****************************************|
// |				CHANNELS				|
// |****************************************|
// Utilizando el patron FanIn generamos un canal general el cual se va completando con datos de diferentes canales
// consideremos que el canal que vamos llenando es de tipo string.
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	c := fanIn(prime(), pairOddGenerator(), nameGenerator("Gun"))

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("--THE END--")
}

func prime() <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%d - %v", i+1, CheckPrime(rand.Intn(100-1+1)+1))
		}
	}()
	return c
}

func CheckPrime(number int) bool {
	if number == 0 || number == 1 {
		return false
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				return false
			}
		}
		return true
	}
}

func pairOddGenerator() <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			if (rand.Intn(100-1+1)+1)%2 == 0 {
				c <- fmt.Sprintf("%d - Pair", i+1)
			} else {
				c <- fmt.Sprintf("%d - Odd", i+1)
			}
		}
	}()

	return c
}

func nameGenerator(value string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", value, i)
		}
	}()
	return c
}

func fanIn(input1, input2, input3 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	go func() {
		for {
			c <- <-input3
		}
	}()

	return c
}
