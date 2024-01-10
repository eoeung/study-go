package main

import "fmt"

func main() {
	// [슬라이스를 슬라이싱]
	// - 슬라이싱은 슬라이스에서도 가능하다.

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[1:2]

	fmt.Println("slice1 :", slice1, len(slice1), cap(slice1)) // [1 2 3 4 5] 5 5
	fmt.Println("slice2 :", slice2, len(slice2), cap(slice2)) // [2] 1 4

	// 1) 처음부터 슬라이싱
	slice3 := slice1[0:3]
	slice4 := slice1[:3]
	fmt.Println("slice3 :", slice3, len(slice3), cap(slice3))
	fmt.Println("slice4 :", slice4, len(slice4), cap(slice4))

	// 2) 끝까지 슬라이싱
	slice5 := slice1[2:len(slice1)]
	slice6 := slice1[2:]

	fmt.Println("slice5 :", slice5, len(slice5), cap(slice5))
	fmt.Println("slice6 :", slice6, len(slice6), cap(slice6))

	// 3) 전체 슬라이싱
	slice7 := slice1[:]
	fmt.Println("slice7 :", slice7, len(slice7), cap(slice7))

	// 4) 인덱스 3개로 cap 크기 조절하기
	slice8 := slice1[1:3:4]
	fmt.Println("slice8 :", slice8, len(slice8), cap(slice8)) // [2 3] 2 3
	// slice1[1:3:4] 분석
	//   1) slice1[1:3]까지는 기존과 동일
	//   2) 4의 의미는, 시작 인덱스(1)부터 4번 인덱스까지만 배열을 사용하겠다는 의미
	//     - 그래서 cap값은 4(최대 인덱스) - 1(시작 인덱스) = 3이됨

}
