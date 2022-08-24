package main

import (
	"fmt"
	"math/rand"
)

type Item struct {
	id int
}

type Bid struct {
	item      Item
	bidValue  int
	bidFailed bool
}

func itemsStream() chan Item {
	num_items := 10
	ch_items := make(chan Item, num_items)
	go func() {
		for i := 0; i < num_items; i++ {
			item := Item{id: i}
			ch_items <- item
		}
		close(ch_items)
	}()
	return ch_items
}

func bid(item Item) Bid {
	randonNum := rand.Intn(10)
	bid := Bid{item: item, bidValue: randonNum, bidFailed: false}
	return bid
}

func handle(nServers int) chan Bid {
	ch_out := make(chan Bid, nServers)
	items := itemsStream()
	join_ch := make(chan int, nServers)

	for j := 0; j < nServers; j++ {
		go func() {
			for item := range items {
				ch_out <- bid(item)
			}
			join_ch <- 1
		}()
	}

	go func() {
		for j := 0; j < nServers; j++ {
			<-join_ch
		}
		close(ch_out)
	}()

	return ch_out
}

func main() {
	ch_bid := handle(5)
	for bid := range ch_bid {
		fmt.Println("Bid recived: ", bid)
	}
}
