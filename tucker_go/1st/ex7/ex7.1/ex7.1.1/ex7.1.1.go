package main

import (
	"fmt"
)

// 함수의 위치는 상관없음
// func Add(a int, b int) int {
// 	return a + b
// }

func main() {
	c := Add(3, 6)
	fmt.Println(c)
}

// 함수의 위치는 상관없음
func Add(a int, b int) int {
	return a + b
}
