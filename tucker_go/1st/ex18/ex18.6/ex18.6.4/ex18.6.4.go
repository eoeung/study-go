package main

import "fmt"

func main() {
	// 3) 슬라이스 요소 추가
	// - 슬라이스 중간에 요소 추가
	//   (1) 맨 뒤에 요소를 추가
	//   (2) 하나씩 뒤로 복사
	//   (3) 값 수정
	slice := []int{1, 2, 3, 4, 5, 6}
	slice = append(slice, 0) // 맨 뒤에 요소 추가

	idx := 2 // 추가하려는 위치

	for i := len(slice) - idx; i >= idx; i-- {
		slice[i+1] = slice[i]
	}
	fmt.Println(slice) // [1 2 3 3 4 5 6]

	slice[idx] = 100
	fmt.Println(slice) // [1 2 100 3 4 5 6]

	// - append()로 코드 개선하기
	slice = []int{1, 2, 3, 4, 5, 6}
	// slice = append(slice[:idx], 100, slice[idx:]...) → 내가 생각한 직관적인 방법
	slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)
	fmt.Println(slice)

	// - 불필요한 메모리 사용이 없도록 코드 개선
	//   - [100, 3, 4, 5, 6]은 임시 슬라이스 → 불필요한 메모리
	slice = []int{1, 2, 3, 4, 5, 6}
	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	fmt.Println(slice) // [1 2 3 3 4 5 6]

	slice[idx] = 100
	fmt.Println(slice)
}
