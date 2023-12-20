package main

import "fmt"

func main() {
	i := 2
	max := 7

	// t := 1
	// if t { // non-boolean condition in if statement // 무조건 bool만 가능
	// 	fmt.Println(true)
	// }

	// Optional Statement
	if val := i * 2; val < max {
		fmt.Println(i, val, max) // 2 4 7
	}

	// 같은 구문
	if (i * 2) < max {
		fmt.Println(i, i*2, max) // 2 4 7
	}

	// 변수 여러 개도 선언할 수 있음
	if a, b := 1, 2; a < b {
		fmt.Println(a, b) // 1 2
	}

	// val++ // undefined: val
}
