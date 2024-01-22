package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printSquare(ch chan int) {
	for n := range ch {
		fmt.Println(n * n)
	}
	wg.Done()
}

func main() {
	wg.Add(1)

	ch := make(chan int)
	go printSquare(ch)

	for i := 0; i < 5; i++ {
		ch <- i * 2
	}

	close(ch)
	wg.Wait()
}
