package main

import "fmt"

func main() {
	// [슬라이싱, append() 활용하기]
	// 1) 슬라이스 복제
	// 2) 슬라이스 요소 삭제
	// 3) 슬라이스 요소 추가

	// 슬라이스가 다른 배열을 가리키게 만들기 위한 방법 : '복사'

	// 1-1) 슬라이스 복제
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1)) // slice1과 같은 길이의 슬라이스 생성

	for i, v := range slice1 {
		slice2[i] = v // slice1의 모든 요소를 복사
	}

	slice1[1] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)

	// 1-2) append() 함수로 코드 개선하기
	slice3 := append([]int{}, slice1...) // 배열이나 슬라이스 뒤에 ...을 하면, 모든 요소값을 의미
	// → append([]int{}, slice1[0], slice1[1], slice1[2], slice1[3], slice1[4])
	fmt.Println(slice3)
}
