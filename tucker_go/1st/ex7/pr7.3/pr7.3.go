package main

import "fmt"

// N번째 피보나치 수열 값 F(N)
// F(0) = 0
// F(1) = 1
// F(N) = F(N-2) + F(N-1)

func F(n int) int {
	// if n == 0 {
	// 	return 0
	// }
	// if n == 1 {
	// 	return 1
	// }
	if n < 2 {
		return n
	}
	return F(n-2) + F(n-1)
}

func main() {
	// 피보나치 수열 9번째 값 출력
	fmt.Println(F(9))
}
