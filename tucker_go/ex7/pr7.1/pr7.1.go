package main

import "fmt"

// 함수 이름은 Multiple
// 입력 = int타입 2개, 반환 = int타입 1개
// 두 입력값을 곱한 결과를 반환
func Multiple(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(Multiple(1, 2))
}
