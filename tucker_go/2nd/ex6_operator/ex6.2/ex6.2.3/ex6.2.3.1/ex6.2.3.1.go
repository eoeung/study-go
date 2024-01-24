package main

import "fmt"

func main() {
	// [실수 비교]
	// - 실수 연산은 우리가 생각한대로 일치한 결과가 나오지 않음

	var a, b, c float64 = 0.1, 0.2, 0.3

	fmt.Printf("%f + %f == %f ::: %v\n", a, b, c, a+b == c)
	fmt.Println(a + b) // 0.30000000000000004
}
