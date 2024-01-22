package main

import "fmt"

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b
	var d float32 = c * 3

	fmt.Println(a) // 1234.523
	fmt.Println(b) // 3456.123
	fmt.Println(c) // 4.266663e+06 → 원래 1234.523 x 3456.123 연산의 결과값은 4266663.334329이다.
	fmt.Println(d) // 1.2799989e+07 → 원래 4266663.334329 x 3 연산의 결과값은 12799990.002987이다.
	// 오차가 점점 커지고 있음을 알 수 있음
}
