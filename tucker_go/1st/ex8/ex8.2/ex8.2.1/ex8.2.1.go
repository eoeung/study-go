package main

import "fmt"

func main() {
	// [상수 사용 시기]
	// - 변하면 안되는 값에 상수 사용

	// Ex) 파이값을 상수로 지정
	const PI1 float64 = 3.141592
	var PI2 float64 = 3.141592

	// PI1 = 4
	PI2 = 4

	fmt.Printf("원주율 : %f\n", PI1)
	fmt.Printf("원주율 : %f\n", PI2)
}
