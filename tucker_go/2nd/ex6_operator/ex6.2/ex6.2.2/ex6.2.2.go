package main

import "fmt"

func main() {
	// [정수 언더플로]
	// - 정수 타입이 표현할 수 있는 가장 작은 값에서 -1을 한 경우
	// → y > y - 1을 항상 만족하는 것은 아님

	var y int8 = -128
	fmt.Println("y > y - 1", y > y-1) // false
	fmt.Println(y, y-1)               // -128, 127

	fmt.Println("y > y - 2", y > y-2) // false
	fmt.Println(y, y-2)               // -128, 126
}
