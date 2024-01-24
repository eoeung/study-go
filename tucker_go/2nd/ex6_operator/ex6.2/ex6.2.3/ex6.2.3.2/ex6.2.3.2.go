package main

import "fmt"

const epsilon = 0.000001 // 아주 작은 값

func equal(a, b float64) bool {
	if a > b {
		if a-b <= epsilon {
			return true
		} else {
			return false
		}
	} else {
		if b-a <= epsilon {
			return true
		} else {
			return false
		}
	}
}

func main() {
	// [작은 오차 무시하기]
	// - 실수값을 정확하게 표현할 수 없음
	// → 오차가 생길 수 밖에 없음

	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f ::: %v\n", a+b, c, equal(a+b, c))
}
