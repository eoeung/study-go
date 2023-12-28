package main

import "fmt"

func main() {
	// 예제로 배우는 GO 프로그래밍
	// [채널 파라미터 (채널 방향)]
	//   - 함수의 파라미터로 송수신을 모두 하는 채널을 전달하지만, 송신 혹은 수신만 할 것인지 여부를 결정할 수 있음
	//     1) 송신 파라미터 : chan<- (p chan<- string)
	//       - 채널로 값을 전송
	//     2) 수신 파라미터 : <-chan (p <-chan string)
	//       - 채널에서 다른 곳으로 값을 전송
	ch := make(chan string, 1)
	sendChan(ch)
	receiveChan(ch) // 예제로 배우는 GO 프로그래밍

	// go by example
	pings := make(chan string, 1)
	ping(pings, "go by example")
	// fmt.Println(<-pings) // Queue에서 pop된 상태라, 이 코드를 주석처리 안하면 pong(pings, pongs)에서의 pings는 값이 없음

	pongs := make(chan string, 1)
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func sendChan(ch chan<- string) {
	ch <- "예제로 배우는 GO 프로그래밍"
	// x := <-ch // Error : invalid operation: cannot receive from send-only channel ch (variable of type chan<- string)
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

// func ping(pings chan string){ // 송수신 모두 가능
func ping(pings chan<- string, msg string) { // 수신용 채널만 전달받음
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
