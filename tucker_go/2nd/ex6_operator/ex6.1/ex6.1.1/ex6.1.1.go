package main

import "fmt"

func main() {
	// [산술 연산자]
	// 1) 사칙연산(+, -, *, /, %)
	// 2) 비트연산(&, |, ^, &^) → and, or, xor, 비트클리어
	// 3) 시프트연산(<<, >>) → 왼쪽 시프트, 오른쪽 시프트

	var x int32 = 7
	var y int32 = 3

	var s float32 = 3.14
	var t float32 = 5

	// 1-1) 정수 사칙연산
	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x % y = ", x%y)

	// 1-2) 실수 사칙연산
	fmt.Println("s * t = ", s*t) // 15.700001
	fmt.Println("s / t = ", s/t) // 0.628

	// ※ 연산 시 주의사항
	// 1) 연산의 각 항의 타입도 같아야 한다.
	// 2) 연산의 결과 타입도 같다.
}
