package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { // for range는 데이터를 계속 기다림
		fmt.Printf("Square : %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done() // 실행되지 않음 → for range 구문은 채널에 데이터가 들어오기를 계속 기다리기 때문
}

func main() {
	// [고루틴에서 데이터를 계속 기다리면서, 데이터가 들어오면 작업을 수행]
	// - 좀비 루틴(고루틴 릭)
	// - 채널을 제때 닫아주지 않아서 고루틴에서 데이터를 기다리며 무한 대기하는 경우

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}

/*
	Square : 0
	Square : 4
	Square : 16
	Square : 36
	Square : 64
	Square : 100
	Square : 144
	Square : 196
	Square : 256
	Square : 324
	fatal error: all goroutines are asleep - deadlock!
*/

// for range 구문은 채널에 데이터가 계속 들어오기를 기다리는데, wg.Done()이 실행되지 않아서 모든 고루틴이 멈추게 됨 → deadlock
