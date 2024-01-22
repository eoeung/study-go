package main

import "fmt"

func main() {
	// [실수 비교]
	// - 오차가 존재하므로 실수의 == 연산은 위험할 수 있음

	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c) // 0.100000 + 0.200000 == 0.300000 : false
	fmt.Println(a + b)                                    // 0.30000000000000004

}
