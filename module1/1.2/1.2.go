package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)

	go prod(ch)
	go consumer(ch)

	time.Sleep(time.Minute)
}

func prod(ch chan<- int) {
	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {
		ch <- 1
		fmt.Println("prod 1")
	}
}

func consumer(ch <-chan int) {

	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {
		msg := <-ch
		fmt.Println(msg)
	}

}
