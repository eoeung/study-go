package main

import (
	"fmt"
	"time"
)

func main() {
	// 예제로 배우는 GO 프로그래밍
	// [Go Select]
	//   - 다중 채널 연산들을 대기할 수 있도록 해줌
	//   - select문 안에 있는 case문에서, 각각 다른 채널을 기다리다가 준비된 채널 case를 실행
	//   - 먼저 도착한 채널의 case문이 실행되고, 준비되지 않으면 계속 대기한다.
	//   - default문이 있으면 case문 채널이 준비되지 않아도, 계속 대기하지 않고 바로 실행한다.
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		// 2개의 고루틴이 모두 실행되기를 기다리고 있음
		select {
		case <-done1: // case를 실행
			fmt.Println("run1 완료")

		case <-done2: // case를 실행
			fmt.Println("run2 완료")
			break EXIT
		}
	}

	// go by example
	// 두 채널에서 select 사용
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// select를 사용해서 동시에 이 값들을 대기
	// 도착하는 대로 각 값을 출력
	for i := 0; i < 2; i++ { // 고루틴이 2개이므로, 고루틴 개수만큼 for문을 돈다.
		fmt.Println("i : ", i)
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	/*
		received one
		received two

		real    0m2.459s // time.Sleep()이 동시에 실행되기 때문에 2초 정도의 시간이 나옴
		user    0m0.000s
		sys     0m0.000s
	*/

	ttt := make(chan string, 3)
	zzz := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 1)
		ttt <- "ttt result 1"
		time.Sleep(time.Second * 2)
		ttt <- "ttt result 2"
		time.Sleep(time.Second * 3)
		ttt <- "ttt result 3"
	}()

	go func() {
		zzz <- "zzz result 1"
	}()

	select {
	case tRes := <-ttt:
		fmt.Println(tRes)
	case zRes := <-zzz:
		fmt.Println(zRes)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout 2")
	}
}

// run1()은 1초간 실행된 후, done1 채널로부터 수신
func run1(done chan bool) {
	time.Sleep(time.Second * 1)
	done <- true
}

// run2()는 2초간 실행된 후, done2 채널로부터 수신
func run2(done chan bool) {
	time.Sleep(time.Second * 2)
	done <- true
}
