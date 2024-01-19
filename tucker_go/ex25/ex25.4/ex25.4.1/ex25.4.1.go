package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// [컨텍스트]
	// - 작업을 지시할 때, 작업 가능 시간/작업 취소 등 조건을 지시 → 작업 명세서
	// - 새로운 고루틴으로 작업을 시작할 때, 일정 시간 동안만 작업 지시/외부에서 작업을 취소할 때 사용
	// - 작업 설정에 관한 데이터 전달도 가능

	// 1) 작업 취소가 가능한 컨텍스트
	//   - 이 컨텍스트를 만들어서 작업자에게 전달하면 작업을 지시한 작업자가 원할 때 작업 취소를 알릴 수 있음

	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background()) // 컨텍스트 생성
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
