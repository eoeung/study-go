package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	// Ex) 여러 고루틴을 생성해서 문자와 숫자를 출력
	go PrintHangul()  // 새로운 고루틴 생성
	go PrintNumbers() // 새로운 고루틴 생성
	// → 각기 다른 고루틴에서 실행되기 때문에, 동시에 실행됨

	// 고루틴은 3개가 동작
	// 1) 메인 루틴
	// 2) PrintHangul() 루틴
	// 3) PrintNumbers() 루틴
	// → 코어 개수가 3개 이상이 아니면 이 세 고루틴을 동시에 실행시킬 코어가 부족함
	// => 하지만, 동시에 실행되는 것처럼 보임

	time.Sleep(3 * time.Second) // 3초간 대기
}
