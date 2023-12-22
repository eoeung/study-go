package main

import (
	"fmt"
)

func main() {
	// go by example
	// [비동기 채널 연산]
	//   - 채널의 송수신은 기본적으로 동기적임
	//   - select와 default를 사용해서 non-blocking(비동기) 송수신을 구현 가능
	messages := make(chan string)  // unbuffered channel이라 모두 default case만 수행
	bufMsg := make(chan string, 1) // buffered channel이라 중간에 있는 msg := "hi" → messages <- msg 케이스를 통과해서 "hi"가 출력됨
	signals := make(chan bool)

	testChan(messages, signals)
	/*
		no message received
		no message sent
		no activity
	*/
	testChan(bufMsg, signals)
	/*
		no message received
		sent message hi
		received message hi
	*/
}

func testChan(messages chan string, signals chan bool) {
	// messages에서 값을 수신할 수 있는 경우, select문은 case msg := <-messages에서 값을 가져옴
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("recevied signal", sig)
	default:
		fmt.Println("no activity")
	}
}
