package main

import "fmt"

// 별칭으로 함수 정의 줄여쓰기
type opFunc func(int, int) int

func main() {
	// [함수 타입 변수]
	// - 함수를 값으로 갖는 변수

	// 1) 컴퓨터는 0, 1로 나타낼 수 있는 숫자값만 가진다.
	// 2) 포인터는 숫자로 나타낼 수 있는 메모리 주소를 값으로 가진다.
	// → 함수를 숫자로 표현하면 된다!

	// - CPU 내부에는 프로그램 카운터가 존재
	//   → 다음 실행할 라인을 나타내는 레지스터
	//   => 함수의 시작 포인트 = 함수 포인터

	// var operator func(int, int) int // 함수 타입 변수를 선언 → 초기화를 안했기 때문에 기본값인 nil을 가짐
	var operator opFunc
	operator = getOperator("*") // operator값을 getOperator 호출 결과값으로 복사

	var result = operator(3, 4)
	fmt.Println(result)
}

// [함수의 타입을 정의]
// - 함수명과 코드블록을 제외한 함수 정의
func add(a, b int) int {
	return a + b
}

// → 함수명/코드블록 제외 = 함수 정의
// func (int, int) int

func mul(a, b int) int {
	return a * b
}

// 함수 정의에서, 매개변수명은 생략이 가능하다.
// func getOperator(op string) func(int, int) int {
// func getOperator(op string) func(a int, b int) int {

func getOperator(op string) opFunc {
	// 1) getOperator()함수는 op값을 인수로 받음
	// 2) 인수로 오는 op값이 +면 add()함수 반환
	// 3) *면 mul()함수를 반환
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
		// 함수 타입 역시 유효하지 않은 메모리 주소인 nil을 값으로 가질 수 있음
		// → 어떤 함수도 가리키지 않음을 나타냄
	}
}
