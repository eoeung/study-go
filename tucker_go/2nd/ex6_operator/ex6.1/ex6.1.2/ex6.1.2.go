package main

import "fmt"

func main() {
	// [산술 연산자]
	// 1) 사칙연산(+, -, *, /, %)
	// 2) 비트연산(&, |, ^, &^) → and, or, xor, 비트클리어
	// 3) 시프트연산(<<, >>) → 왼쪽 시프트, 오른쪽 시프트

	// 2) 비트연산
	var a, b int = 10, 34
	fmt.Println("a & b = ", a&b) // 2 ## and 연산
	fmt.Println("a | b = ", a|b) // 42 ## or 연산
	fmt.Println("a ^ b = ", a^b) // 40 ## xor 연산
	// ^를 단독으로 사용하면, 비트 반전
	fmt.Println("^a = ", ^a) // -11 ## 비트 반전
	fmt.Println("^b = ", ^b) // -35 ## 비트 반전

	var c, d int = 10, 2
	fmt.Println("c &^ d = ", c&^d) // 8 ## 비트 클리어 → 오른쪽에 있는 비트를 클리어하는 연산
}
