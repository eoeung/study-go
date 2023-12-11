package main

import "fmt"

func main() {
	var x int8 = 4
	var y int8 = 64

	// %08b
	// %b → 2진수로 표시
	// 08 → 8자리를 채울 건데, 빈 자리는 0으로 채운다.
	// << 왼쪽 쉬프트 연산
	fmt.Printf("x:%08b x<<2:%08b %d\n", x, x<<2, x<<2)
	fmt.Printf("y:%08b y<<2:%08b %d\n", y, y<<2, y<<2)

	// >> 오른쪽 쉬프트 연산
	var a int8 = 16
	var b int8 = -128
	var z int8 = -1
	var w uint8 = 128

	fmt.Printf("a:%08b a>>2:%08b %d\n", a, a>>2, a>>2)
	fmt.Printf("b:%08b b>>2:%08b %d\n", uint8(b), uint8(b>>2), b>>2)
	fmt.Printf("z:%08b z>>2:%08b %d\n", uint8(z), uint8(z>>2), z>>2)
	fmt.Printf("w:%08b w>>2:%08b %d\n", w, w>>2, w>>2)
}
