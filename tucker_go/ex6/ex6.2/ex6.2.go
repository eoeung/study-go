package main

import "fmt"

func main() {
	// [정수 오버플로]
	// - 오버플로 : 정수가 정수 타입의 범위르 벗어난 경우에, 값이 비정상으로 변화하는 현상
	// - x < x + 1이 항상 true가 아님

	var x int8 = 127                               // 부호 있는 8비트 정수의 최대값
	fmt.Printf("%d < %d + 1: %v\n", x, x, x < x+1) // 127 < 127 + 1: false
	// → int8 타입은 -128 ~ 127까지의 범위를 가짐
	// => 127인 0111 1111에서 +1을 해서 1000 0000이 됨 ☞ -128이 됨
	// 이런 현상을 오버플로라고함
	fmt.Printf("x = %4d, %08b\n", x, x)         // x =  127, 01111111
	fmt.Printf("x + 1 = %4d, %08b\n", x+1, x+1) // x + 1 = -128, -10000000
	fmt.Printf("x + 2 = %4d, %08b\n", x+2, x+2) // x + 2 = -127, -1111111
	fmt.Printf("x + 3 = %4d, %08b\n", x+3, x+3) // x + 3 = -126, -1111110

	fmt.Println()

	// [정수 언더플로]
	// - 정수 타입이 표현할 수 있는 가장 작은 값에서 -1 연산 진행하면 가장 큰 값으로 바뀜

	var y int8 = -128
	fmt.Printf("%d > %d - 1: %v\n", y, y, y > y-1) // -128 > -128 - 1: false
	fmt.Printf("y = %4d, %08b\n", y, y)            // y = -128, -10000000
	fmt.Printf("y - 1 = %4d, %08b\n", y-1, y-1)    // y - 1 =  127, 01111111
}
