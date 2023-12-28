package main

import (
	"fmt"
	"time"
)

func main() {
	// go by example
	// [채널 동기화]
	//   - 고루틴간 실행을 동기화
	//   - 고루틴이 끝날 때 까지 대기하기 위해 블로킹 리시브 사용

	done := make(chan bool, 1)
	go worker(done)

	fmt.Println("mmmmmmmmmmmm")
	<-done // done이 값을 수신하기 전까지 모든 고루틴은 여기서 대기(Blocking)
	fmt.Println("tttttttttttt")
}

// 고루틴에서 실행하기 위한 함수
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true // 만약 이 코드가 주석처리되면, worker가 실행되기도 전에 main()이 종료되므로 프로그램이 종료될 수 있음
}

/*
	mmmmmmmmmmmm
	working...
	done
	tttttttttttt
*/
