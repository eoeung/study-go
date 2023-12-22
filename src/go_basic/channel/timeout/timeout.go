package main

import (
	"fmt"
	"time"
)

func main() {
	// go by example
	// [Timeout]
	//   - 외부 리소스를 연결하거나, 실행 시간을 제한할 때 사용

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2) // 2초 후에 c1 채널에 결과값 반환하는 외부 호출 실행
		c1 <- "result 1"
	}()

	// select문 : 여러 채널 중에서 먼저 준비된 채널의 동작을 실행
	select {
	case res := <-c1: // 결과값을 대기
		fmt.Println(res)
	case <-time.After(time.Second * 1): // 1초의 타임아웃 후에 전달되는 값을 대기
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

	/*
		timeout 1
		result 2
	*/
}
