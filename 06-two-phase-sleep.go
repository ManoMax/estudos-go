package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)

	joinChan := make(chan int)
	defer close(joinChan)

	var wg = sync.WaitGroup{}

	channels := make([]chan int, n)
	for i := 0; i < n; i++ {
		channels[i] = make(chan int, 1)
	}

	for i := 0; i < n; i++ {
		in_channel := ((i + n) % n)
		out_channel := ((i + 1) % n)
		wg.Add(1)
		go func_random(joinChan, channels[in_channel], channels[out_channel], &wg)
	}

	for i := 0; i < n; i++ {
		<-joinChan
	}

	fmt.Printf("No total foram executadas %d goroutines\n", n)
}

func func_random(joinChan chan int, inChan <-chan int, outChan chan<- int, wg *sync.WaitGroup) {

	st := rand.Intn(5)
	time.Sleep(time.Second * time.Duration(st))
	out := rand.Intn(10)
	outChan <- out

	wg.Done()
	wg.Wait()

	in := <-inChan

	fmt.Printf("Esperando por %d segundos\n", in)
	time.Sleep(time.Second * time.Duration(in))

	joinChan <- 0
}
