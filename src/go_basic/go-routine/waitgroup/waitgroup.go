package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// noTimeSleep() // 아무것도 print되지 않음 → 고루틴이 실행되자마자 메인 함수가 종료됨
	// timeSleep()
	waitGroup()
}

// Sleep 없을 때
func noTimeSleep() {
	go helloWorld()
	go goodBye()
}

// Sleep 있을 때
func timeSleep() {
	go helloWorld()
	go goodBye()
	time.Sleep(time.Second * 2) // 메인 함수를 2초 동안 일시중지(Blocking) → 모든 고루틴을 성공적으로 실행
	// → 다만, 실시간으로 진행되는 기능을 Sleep으로 제어할 수는 없음..
}

func waitGroup() {
	var wg sync.WaitGroup
	wg.Add(2) // 대기할 고루틴의 수 → 이 코드에서는 2개의 고루틴 선언 → 내부 카운터 대기 = 2
	go helloWorldWaitGroup(&wg)
	go goodByeWaitGroup(&wg)
	wg.Wait() // 내부 카운터가 0이 될 때 까지, 코드 실행을 차단
}

func helloWorld() {
	fmt.Println("Hello World!")
}

func goodBye() {
	fmt.Println("Good Bye!")
}

func helloWorldWaitGroup(wg *sync.WaitGroup) {
	defer wg.Done() // 내부 카운터 값이 1 감소
	fmt.Println("Hello World!")
}

func goodByeWaitGroup(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Good Bye!")
}
