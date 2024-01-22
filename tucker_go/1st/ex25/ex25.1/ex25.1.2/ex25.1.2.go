package main

import "fmt"

func main() {
	// [채널에서 데이터를 가져가지 않아서 프로그램이 멈추는 경우]

	ch := make(chan int) // 크기 0 채널 생성

	ch <- 9
	// 1) 채널에 데이터를 넣었지만, 보관할 곳이 없어서 다른 고루틴에서 데이터를 빼가기를 기다림
	// 2) 어떤 고루틴도 데이터를 빼가지 않음
	// 3) 모든 고루틴이 영원히 대기
	// 4) deadlock 발생
	fmt.Println("Never print")
}
