package main

import "fmt"

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return func(a, b int) int { // 함수 리터럴 사용
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	// [함수 리터럴]
	// - 이름 없는 함수

	fn := getOperator("*") // 함수 리터럴 or nil을 리턴받음
	result := fn(3, 4)     // 함수 타입 변수 사용해서, 함수 호출

	fmt.Println(result)
}
