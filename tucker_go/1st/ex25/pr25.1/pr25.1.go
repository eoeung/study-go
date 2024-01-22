package main

import (
	"fmt"
	"time"
)

func printSquare(ch chan int) {
	n := <-ch
	fmt.Println(n * n)
}

func main() {
	ch := make(chan int)
	printSquare(ch)
	// go printSquare(ch)

	ch <- 5
	time.Sleep(time.Second * 5)
}

// ＊printSquare(ch)로 코드를 작성 시
//   1) printSquare() 함수 안에서 ch에서 값 읽기를 시도
//   2) 채널에 아무런 값이 없어서 무한 대기 → 데드락

//////////////////////////////////////////////////
// func main() {
// 	ch := make(chan int)
// 	ch <- 5
// 	printSquare(ch)

// 	time.Sleep(time.Second * 5)
// }

// 1) 이 코드는 에러
// → 채널이 비동기적으로 값을 받을 수 있는 버퍼가 없는 블로킹 채널이기 때문

//////////////////////////////////////////////////
// func main() {
// 	ch := make(chan int, 1)
// 	ch <- 5
// 	printSquare(ch)

// 	time.Sleep(time.Second * 5)
// }

// 2) 이 코드는 동작
// → buffered 채널로 만들었기 때문에, 동시성 프로그래밍에서 사용하는 고루틴이 없어도 값 저장이 가능
