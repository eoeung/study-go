package main

import "fmt"

func main() {
	// 예제로 배우는 GO 프로그래밍
	s := make(chan int, 2)
	s <- 1
	s <- 2
	close(s)
	// 1) 채널에서 값을 수신 : 값, 성공여부로 for, if문 사용해서 채널로부터 받은 값 출력하기
	for {
		if val, success := <-s; success {
			fmt.Println(val, success)
		} else {
			break
		}
	}

	// 2) for range문 사용 : 값만 출력
	zzz := make(chan int, 2)
	zzz <- 555
	zzz <- 999
	close(zzz)
	// close(zzz)하지 않으면 에러 발생 - fatal error: all goroutines are asleep - deadlock!
	// 채널이 종료됨을 감지해야, for range를 쓸 수 있음
	// for val, _ := range zzz{ // range over s (variable of type chan int) permits only one iteration variable
	for i := range zzz {
		fmt.Println("zzz :: ", i)
	}

	// go by example
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// 채널을 닫아도, 채널에 남아있는 값은 수신할 수 있음
	for elem := range queue {
		fmt.Println(elem)
	}
}
