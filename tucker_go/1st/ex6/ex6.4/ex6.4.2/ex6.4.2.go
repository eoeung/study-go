package main

import "fmt"

func main() {
	// 복잡한 연산

	fmt.Println(3*4^7<<2+3*5 == 7)
	fmt.Println(3*4 ^ 7<<2 + 3*5) // 31
	fmt.Println(12 ^ 7<<2)        // 16
	fmt.Println(7 << 2)           // 28
}
