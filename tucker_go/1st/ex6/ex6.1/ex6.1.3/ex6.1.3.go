package main

import "fmt"

func main() {
	// [왼쪽 시프트 연산자]
	// - 비트가 왼쪽으로 이동되어 빈 자리에는 0이 채워짐
	// - 자릿수를 벗어난 비트는 버려짐
	// - 비트 수를 나타내는 오른쪽 피연산자는 반드시 양의 정수여야 함

	var x int8 = 4
	var y int8 = 64
	var z int8 = -1

	fmt.Printf("x:%08b x<<2: %08b, %d\n", x, x<<2, x<<2)               // x:00000100 x<<2: 00010000, 16
	fmt.Printf("y:%08b y<<2: %08b, %d\n", y, y<<2, y<<2)               // y:01000000 y<<2: 00000000, 0
	fmt.Printf("z:%08b z<<2: %08b, %d\n", uint8(z), uint8(z<<2), z<<2) // z:11111111 z<<2: 11111100, -4
}
