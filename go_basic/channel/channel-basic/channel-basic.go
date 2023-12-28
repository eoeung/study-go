package main

import (
	"fmt"
)

func main() {

	// 예제로 배우는 GO 프로그래밍
	// [Channel]
	//   - 데이터를 주고 받는 통로
	//   - make()로 미리 만들어져 있어야함
	//   - 채널 연산자인 <-를 통해 데이터를 주고 받음
	//   - 상대편이 준비될 때 까지 채널에서 대기 → 별도의 lock없이 데이터를 동기화
	//   - 수신자와 송신자가 서로를 기다리는 속성이 있음
	//   - 수신자와 송신자는 쓰레드(Go에서는 goroutine)라고 생각하면 편함 → 서로 다른 쓰레드간 정보 교환

	// 정수형 채널 생성
	ch := make(chan int, 1) // 어떤 타입의 데이터를 주고 받을지 미리 지정해줘야 함

	go func() {
		// ch <- 123 // 채널에 123을 보냄
		// 위 코드를 주석처리하면 에러 발생 : fatal error: all goroutines are asleep - deadlock!
		// → goroutine에서 데이터를 전송할 때 까지 대기
		// → time.Sleep()이나 fmt.Scanf()같은 goroutine이 끝날 때 까지 기다리는 코드를 적지 않았음
	}()

	ch <- 123 // 채널에 123을 보냄

	var i int = <-ch
	fmt.Println("ch::", i)

	// var i int
	// i = <-ch // 채널로 부터 123을 받음

	// goroutine이 끝날 때 까지 기다리는 기능 구현
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		// done <- true
	}()

	// // 위의 goroutine이 끝날 때 까지 대기
	// // 위에 있는 익명함수를 사용한 goroutine에서 작업이 진행될 때, 메인 루틴은 <-done에서 수신할 때 까지 계속 대기하고 있음
	// // → 익명함수 goroutine이 끝나서 done 채널에 true를 보내면, 수신자 메인 루틴은 이를 받고 프로그램 종료

	// go func() {
	// 	done <- false
	// }()

	go func() {
		for i := 300; i < 500; i++ {
			fmt.Println("####### ", i)
		}
		done <- true
	}()
	// done <- false // 에러 발생 : deadlock : 수신자가 없음

	fmt.Println("@@@@@@@@@@@@@@@@@@")
	<-done // 주석처리하면 메인 고루틴이 쭉 실행되면서, 프로그램이 종료되면 모든 고루틴이 실행되지 않을 수 있음
	fmt.Println("@@@@@@@@@@@@@@@@@@")

	// go by example
	// 채널 생성 : make(chan value-type)
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg) // ping

	numChan := make(chan int)
	// num <- 123 // Error : fatal error: all goroutines are asleep - deadlock! → goroutine을 사용해야하는듯
	x, y := 10, 20
	go add(x, y, numChan)

	t := <-numChan
	fmt.Println(t) // 30
}

func add(x, y int, numChan chan int) {
	numChan <- x + y
}
