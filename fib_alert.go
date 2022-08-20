package main

import (
	"fmt"
	"time"
)

func main() {
	n := 42

	go alert()
	result := fib(n)
	fmt.Printf("Fib(%d) = %d\n", n, result)
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func alert() {
	for {
		fmt.Printf("demora. Ai!\n")
		time.Sleep(1 * time.Second)
	}
}
