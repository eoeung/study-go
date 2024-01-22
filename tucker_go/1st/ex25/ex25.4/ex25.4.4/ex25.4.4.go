package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// [취소도 되면서 값도 설정하는 컨텍스트 만들기]
	// - 컨텍스트를 여러 번 감싸서 설정

	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, 1, "Lilly")
	ctx = context.WithValue(ctx, 2, 9)

	go PrintContextVal(ctx)

	time.Sleep(time.Second * 2)
	cancel()
}

func PrintContextVal(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
			for i := 0; i < 2; i++ {
				fmt.Println(i)
				if v := ctx.Value(i); v != nil {
					fmt.Println(v)
				}
			}
			time.Sleep(time.Second * 1)
		}
	}

}
