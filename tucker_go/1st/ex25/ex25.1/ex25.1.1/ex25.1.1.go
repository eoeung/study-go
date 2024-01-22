package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// [채널]
	// - 고루틴 간 메시지 전달할 수 있는 메시지 큐

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch) // main() 루틴이 아닌 새로운 고루틴에서 동시에 실행됨
	ch <- 9
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch
	// 채널에서 데이터를 빼오려 시도
	// → 채널에 데이터가 없기 때문에 데이터가 들어올 때 까지 대기
	// => main() 루틴에서 9라는 값을 넣어줌으로써 채널에서 9라는 값을 수신할 수 있게 됨

	time.Sleep(time.Second)
	fmt.Printf("Square : %d\n", n*n)
	wg.Done()
}
