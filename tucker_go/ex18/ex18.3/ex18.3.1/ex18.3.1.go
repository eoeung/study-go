package main

import "fmt"

// reflect.SliceHeader
type SliceHeader struct {
	Data uintptr // 실제 배열을 가리키는 포인터
	Len  int     // 요소 개수
	Cap  int     // 실제 배열의 길이
}

func main() {
	// [슬라이스의 내부 동작 원리]
	// 1) make()를 이용한 선언
	var slice = make([]int, 3)
	/*
		Data(포인터) = [0 0 0]
		Len = 3 // 요소 개수
		Cap = 3 // 배열 길이
	*/

	var slice2 = make([]int, 3, 5)
	/*
		Data(포인터) = [0 0 0    ] → 5개 중 3개만 사용, 2개는 비워둠
		Len = 3 // 요소 개수
		Cap = 5 // 배열 길이
	*/

	fmt.Println(slice, len(slice), cap(slice))    // [0 0 0] 3 3
	fmt.Println(slice2, len(slice2), cap(slice2)) // [0 0 0] 3 5
}
