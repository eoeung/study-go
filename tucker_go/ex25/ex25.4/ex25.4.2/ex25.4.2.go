package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 2) 작업 시간을 설정한 컨텍스트
	//   - 일정 시간 동안만 작업을 지시할 수 있는 컨텍스트
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // 2번째 인자로 보낸 시간이 지나면 ctx.Done() 채널에 시그널을 보냄
	go PrintEverySecond(ctx)
	time.Sleep(time.Second * 5) // 메인 루틴이 종료되는 것을 방지
	cancel()                    // 취소
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done(): // cancel()이 실행되면, ctx.Done() 채널에 시그널을 보냄
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}

	}
}
