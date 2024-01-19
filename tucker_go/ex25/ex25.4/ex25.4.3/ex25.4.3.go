package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 3) 특정 값을 설정한 컨텍스트
	//   - 별도 지시사항이 있을 때 처럼 컨텍스트에 특정 키로 값을 읽어올 수 있도록 설정

	wg.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9) // key-val로 값 추가
	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil { // Value의 반환 타입은 any
		n := v.(int)
		fmt.Printf("Square :%d\n", n*n)
	}
	wg.Done()
}
