package main

import "fmt"

func main() {
	// [슬라이스 선언]
	// - 배열과 다르게 []만 선언
	// - 슬라이스를 초기화하지 않으면, 길이가 0인 슬라이스가 만들어짐

	var slice []int

	if len(slice) == 0 { // 슬라이스의 길이
		fmt.Println("slice is empty", slice)
	}

	slice[1] = 10 // Error : index out of range [1] with length 0
	fmt.Println(slice)
}
