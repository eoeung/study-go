package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select { // select문에서 ch와 quit채널 모두를 기다림
		case n := <-ch:
			fmt.Printf("Square : %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}

func main() {
	// [select문]
	// - 채널에 데이터가 들어오기를 대기하는 상황에서, 데이터가 들어오지 않으면?
	//   1) 다른 작업을 수행
	//   2) 여러 채널을 동시에 대기하고 싶을 때

	// - select문은 여러 채널을 동시에 기다릴 수 있음
	// - 어떤 채널이라도 데이터를 읽어오면 select문이 종료됨
	//   → 하나의 case만 처리되면 종료됨

	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- true
	wg.Wait()
}
