package main

import "fmt"

func main() {
	// 9까지 홀수의 제곱값 출력
	// 1 * 1 = 1
	// 3 * 3 = 9
	// 5 * 5 = 25
	// 7 * 7 = 49
	// 9 * 9 = 81

	for i := 1; i <= 9; i += 2 {
		fmt.Printf("%d * %d = %d\n", i, i, i*i)
	}
}
