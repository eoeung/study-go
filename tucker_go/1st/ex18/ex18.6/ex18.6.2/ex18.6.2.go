package main

import "fmt"

func main() {
	// 1-3) copy() 함수를 이용해 슬라이스 복사
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3, 10) // len: 3, cap: 10
	slice3 := make([]int, 10)    // len, cap: 10

	cnt1 := copy(slice2, slice1)
	cnt2 := copy(slice3, slice1)

	fmt.Println(cnt1, slice2) // 3 [1 2 3]
	fmt.Println(cnt2, slice3) // 5 [1 2 3 4 5 0 0 0 0 0]
	// len값이 다른 경우, 작은 값만 복사한다.

	// slice1을 그대로 복사하기
	sliceT := make([]int, len(slice1)) // len, cap: 5
	copy(sliceT, slice1)

	fmt.Println(sliceT)
}
