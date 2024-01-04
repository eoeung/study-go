package main

import "fmt"

// [iota]
// - iota로 간편하게 열거값 사용
// - iota는 소괄호를 벗어나면 초기화됨

// 1) 기본적인 iota사용
const (
	Red   int = iota // 0
	Blue  int = iota // 1
	Green int = iota // 2
	// Green int // Error : missing init expr for Green
)

// 2) 똑같은 규칙이면 타입, iota 생략 가능
const (
	C1 uint = iota + 1
	C2
	C3 int = iota
)

const (
	BitFlag1 uint = 1 << iota
	BitFlag2
	BitFlag3
	BitFlag4
)

func main() {
	fmt.Println(C1, C2, C3)                             // 1 2 2
	fmt.Println(BitFlag1, BitFlag2, BitFlag3, BitFlag4) // 1 2 4 8
}
