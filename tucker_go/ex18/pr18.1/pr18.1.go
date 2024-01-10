package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5} // [1 2 3 4 5] 5 5
	slice := array[1:3]
	fmt.Println(slice, len(slice), cap(slice)) // [2 3] 2 4 // cap은 배열 전체를 사용

	slice = append(slice, 100)                 // [2 3 100] 3 4
	fmt.Println(slice, len(slice), cap(slice)) // [2 3 100] 3 4 // cap은 배열 전체를 사용
	// 슬라이싱은 같은 배열을 가리킨다.
	fmt.Println(array) // [1 2 3 100 5]
}
