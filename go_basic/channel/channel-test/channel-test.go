package main

import "fmt"

func main() {
	// ch := make(chan int)
	// ch <- 123 // fatal error: all goroutines are asleep - deadlock!
	// fmt.Println(<-ch)

	ch := make(chan int)
	go func() {
		ch <- 123
	}()
	fmt.Println(<-ch)
}
