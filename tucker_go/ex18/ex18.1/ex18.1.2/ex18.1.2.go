package main

import "fmt"

func main() {
	// [슬라이스 초기화]
	// 1) {}를 사용해 초기화
	// 2) make를 이용한 초기화

	// 1) {}를 사용해 초기화
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 5: 2, 10: 3} // index : 값
	// → slice2[0] = 1, slice2[5] = 2, slice[10] = 3
	fmt.Println(slice1) // [1 2 3]
	fmt.Println(slice2) // [1 0 0 0 0 2 0 0 0 0 3]

	// ※ 주의점
	var array = [...]int{1, 2, 3}           // 배열 선언
	var slice = []int{1, 2, 3}              // 슬라이스 선언
	fmt.Printf("%T ::: %T\n", array, slice) // [3]int ::: []int

	// 2) make를 이용한 초기화
	// - make(타입, 길이)
	var slice999 = make([]int, 3)
	fmt.Println(slice999) // [0 0 0]
}
