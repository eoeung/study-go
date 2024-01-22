package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	// 1) for range는 데이터를 계속 기다림
	// 2) 채널이 닫힌 상태라면, for문을 종료한다.
	for n := range ch {
		fmt.Printf("Square : %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	// [ex25.1.3 예제를 수정 → 채널을 닫는 부분 생성]

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) // 채널을 닫음
	wg.Wait()
}
