package main

import "fmt"

func main() {
	// [슬라이스 요소 추가]
	// - 배열은 길이가 정해지면 늘릴 수 없었지만, 슬라이스는 요소를 추가해서 길이를 늘릴 수 있음
	// - append()
	//   - append()는 요소가 추가된 새로운 슬라이스를 반환

	var slice = []int{1, 2, 3}

	slice2 := append(slice, 4) // append(인수를 추가하고 싶은 슬라이스, 요소)

	fmt.Println(slice)
	fmt.Println(slice2)

	// 여러 값 추가
	slice = append(slice, 3, 4, 5, 6, 7)
	fmt.Println(slice)
}
