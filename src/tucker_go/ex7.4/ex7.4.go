package main

import "fmt"

func printNo(n int) {
	// 재귀 호출 탈출 조건
	if n == 0 {
		return
	}

	fmt.Println(n)
	printNo(n - 1)          // 재귀 호출
	fmt.Println("After", n) // 재귀 호출이 모두 끝나면 출력됨
}

func main() {
	printNo(3)
	/*
		3
		2
		1
		After 1
		After 2
		After 3
	*/
	// 재귀 호출을 할 때 마다 Stack이 쌓임
	// Stack이 무한정 쌓이면 프로그램이 종료될 수도 있음
	// Go는 자동 증가하는 Stack을 사용하고 있음
}
