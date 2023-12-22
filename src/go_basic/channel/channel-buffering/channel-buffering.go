package main

import "fmt"

func main() {
	// 예제로 배우는 GO 프로그래밍
	// [Go 채널 버퍼링]
	//   - Go는 2가지 채널이 있음
	//       1) Unbuffered Channel
	//            - 채널에서 하나의 수신자가 데이터를 받을 때 까지 송신자가 데이터를 보내는 채널에 묶여있음
	//       2) Buffered Channel
	//            - 수신자가 받을 준비가 되어있지 않아도 지정된 버퍼만큼 데이터를 보내고, 다른 일 작업 가능

	// https://stackoverflow.com/questions/44216442/why-does-a-channel-in-golang-require-a-go-routine
	// ※ 성공적인 전송에는 준비된 수신자 또는 가득 차지 않은 버퍼가 필요

	// 버퍼링되지 않은 채널에서는 다른 루틴이 값을 읽을 때까지 값이 이동할 곳이 없음
	/*
		package main
		import  "fmt"
		func main() {
			c := make(chan int)
			c <- 17
			fmt.Println(<- c)
		}
	*/
	// 단일 스레드, 버퍼링되지 않은 예에서 해당 17은 채널에 기록될 때와 읽을 때 사이에 어디에 있습니까?
	// 버퍼가 없으면 다음 줄이 실행될 때까지 해당 값을 저장할 위치가 없습니다.

	// https://www.educative.io/answers/what-are-buffered-channels-in-golang
	// 기본적으로 Golang의 채널은 버퍼링되지 않으므로 본질적으로 차단 됩니다.
	// 이는 이전에 전송된 값이 수신된 후에만 후속 데이터 값을 전송할 수 있음을 의미합니다.
	// → 버퍼가 없는 경우, 채널은 데이터를 수신하면 무조건 데이터를 내보내야한다. (버퍼 = 저장소)가 없기 때문

	// 에러 코드
	// c := make(chan int)
	// // 메인 루틴에서 채널에 1을 보내면서 상대편 수신자를 기다림 → 준비된 수신자가 없으므로 deadlock
	// c <- 1 // fatal error: all goroutines are asleep - deadlock!)
	// fmt.Println(<-c)

	// Buffered Channel 생성
	// make(chan type, N) : N에 사용할 버퍼 개수 설정
	ch := make(chan int, 1)
	ch <- 101         // 수신자가 없어도 보낼 수 있음
	fmt.Println(<-ch) // 101

	// go by example
	messages := make(chan string, 2) // 최대 2개의 값을 버퍼링할 수 있는 문자열 채널 생성
	messages <- "buffered"
	messages <- "channel"
	// messages <- "abc" // fatal error: all goroutines are asleep - deadlock!
	// 채널이 버퍼링 되었기 때문에, 동시에 실행중인 대응되는 리시버가 없어도 채널에 값 전달 가능
	fmt.Println(<-messages) // buffered
	fmt.Println(<-messages) // channel

}
