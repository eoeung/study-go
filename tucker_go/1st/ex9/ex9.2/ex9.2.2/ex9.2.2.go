package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	// [쇼트서킷]
	// 1) && 연산
	//   - 좌변이 false면 우변을 검사하지 않고 바로 false

	// 2) || 연산
	//   - 좌변이 true면 우변 검사하지 않고 바로 true

	// Ex) 쇼트서킷 생각 못하고 코드 짜면
	if false && IncreaseAndReturn() < 5 { // 함수가 호출되지 않음
		fmt.Println("1 증가")
	}

	if true || IncreaseAndReturn() < 5 { // 함수가 호출되지 않음
		fmt.Println("2 증가")
	}

	fmt.Println(cnt) // 0
}
