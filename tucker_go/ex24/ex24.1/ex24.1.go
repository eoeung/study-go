package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond) // 1s = 1000ms → 0.3s
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
	go PrintHangul()
	go PrintNumbers()

	time.Sleep(3 * time.Second) // 종료되기 까지 3초 기다림
	// time.Sleep(1 * time.Second) // 1초 후 종료
}
