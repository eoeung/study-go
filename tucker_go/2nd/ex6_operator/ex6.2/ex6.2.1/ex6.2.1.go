package main

import "fmt"

func main() {
	// [정수 오버플로]
	// - 정수 타입이 표현할 수 있는 가장 큰 값에서 +1을 한 경우
	// → x < x + 1을 항상 만족하는 것은 아님

	var x int8 = 127
	fmt.Println("x < x + 1", x < x+1) // false
	fmt.Println(x, x+1)               // 127, -128

	fmt.Println("x < x + 2", x < x+2) // false
	fmt.Println(x, x+2)               // 127, -127
}
