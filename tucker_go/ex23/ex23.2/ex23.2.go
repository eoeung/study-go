package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("제곱근은 양수여야 합니다. f:%g", f) // fmt.Errorf() 함수로 에러 반환

		// errors 패키지의 New() 함수로 error 생성 가능
		// errors.New("에러 메시지")
	}
	return math.Sqrt(f), nil
}

func main() {
	// [사용자 에러 반환]
	// - 에러를 직접 만들어서 반환

	sqrt, err := Sqrt(-2)
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		return
	}
	fmt.Printf("Sqrt(-2) = %v\n", sqrt)
}
