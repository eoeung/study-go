package main

import "fmt"

func main() {
	// [[]rune 타입 변환으로 글자 수 알아내기]
	// string 타입, rune 슬라이스 타입인 []rune 타입 → 상호 타입 변환이 가능

	str := "Hello World"

	// 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd' 문자코드 출력
	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	fmt.Println(str)
	fmt.Println(string(runes))
}
