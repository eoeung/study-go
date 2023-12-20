package main

func main() {
	// 예제로 배우는 GO 프로그래밍
	// [익명 함수]
	// 1) 함수 전체를 변수에 할당
	// 2) 다른 함수의 파라미터에 직접 정의되어 사용됨
	// 3) 함수명을 선언하지 않음
	sum := func(n ...int) int { // 익명함수 정의
		s := 0
		for _, value := range n {
			s += value
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5)
	println(result)

	// [일급 함수]
	// 프로그래밍 언어가 해당 언어의 함수들이 다른 변수처럼 다루어지는 경우, 일급 함수를 가진다고 한다.
	// Go에서, 함수는 Go의 기본 타입과 동일하게 취급됨
	add := func(i, j int) int { // 변수 add에 익명함수 할당
		return i + j
	}

	// add 함수 전달
	r1 := calc(add, 10, 20)
	println(r1) // 30

	r2 := calc(func(x, y int) int { return x - y }, 10, 20)
	println(r2) // -10

	// [type]을 이용한 함수 원형 정의
	r3 := calc2(add, 10, 20)
	println(r3) // 30
}

// [일급 함수]
// calc의 인자는 3개이다.
// f func(int, int) int
// a int
// b int
func calc(f func(int, int) int, a, b int) int {
	result := f(a, b) // a,b를 f의 인자로 넣은 결과값을 반환
	return result
}

// [type]을 이용한 함수 원형 정의
// 원형 정의
// - Delegate : 함수의 원형을 정의하고, 함수를 타 메서드에 전달하고 리턴받는 기능
type calculator func(int, int) int

func calc2(f calculator, a, b int) int {
	result := f(a, b)
	return result
}
