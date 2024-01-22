package main

import "fmt"

func main() {
	// [슬라이스 요소 접근]
	// - 배열과 동일
	// - slice[index]

	var slice = make([]int, 5)
	fmt.Println(slice) // [0 0 0 0 0]
	slice[1] = 5
	fmt.Println(slice) // [0 5 0 0 0]
}
