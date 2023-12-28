package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 예제로 배우는 GO 프로그래밍
	// [Go Routine]
	//   - Go 런타임이 관리
	//   - 가벼운 쓰레드
	//   - OS 쓰레드와 1:1로 대응되지 않음 (multiplexing) → 다중화 : 두 개 이상의 저수준의 채널을 하나의 고수준의 채널로 통합
	//   - Go Channel을 통해 Go Routine간 통신을 쉽게 할 수 있도록 함

	// [다중 CPU 처리]
	//   - Go는 기본으로 1개의 CPU를 사용
	//   - Concurrent 처리 : 여러 개의 Go Routine을 만들어도, 1개의 CPU에서 작업을 시분할하여 처리
	//   - Parallel 처리 : 복수개의 CPU를 가진 경우에 Go 프로그램을 다중 CPU에서 병렬 처리
	// Concurrent : 독립적으로 실행되는 프로세스의 구성 // 한 번에 많은 것을 다룸
	// Parallel : 계산을 동시에 실행하는 것 // 한 번에 많은 작업을 진행

	// 4개의 CPU를 사용
	runtime.GOMAXPROCS(4)

	// 함수를 동기적으로 실행
	say("Sync")

	// 함수를 비동기적으로 실행
	// 별도의 Go Routine으로 동작
	// 메인 루틴은 계속 다음 문장 실행 → time.Sleep()
	go say("Async1")
	go say("Async2")
	go say("Async3")

	// 3초 대기
	time.Sleep(time.Second * 1)

	// [익명함수 Go Routine]
	// sync.WaitGroup 생성
	// → 여러 Go Routine이 끝날 때 까지 기다리는 역할
	var wait sync.WaitGroup
	wait.Add(2) // 2개의 Go Routine을 가짐 (Go Routine개수를 wait.Add(개수)로 지정해준다.)

	// 익명함수를 사용한 goroutine
	go func() {
		defer wait.Done() // 끝나면 wait.Done() 호출
		fmt.Println("Hello")
	}()

	// 익명함수에 파라미터 전달
	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("Hi")

	wait.Wait() // Go Routine이 모두 끝날 때 까지 대기

	// go by example
	// 블로킹된 호출 → 동기적으로 호출된 함수
	f("direct")

	// 인터리브된 2개의 고루틴 결과값 출력
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, "***", i)
	}
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
