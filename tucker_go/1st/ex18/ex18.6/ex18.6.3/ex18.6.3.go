package main

import "fmt"

func main() {
	// 2) 슬라이스 요소 삭제
	// - 슬라이스 중간 요소 삭제
	//   (1) 삭제할 인덱스에 다음 인덱스 값을 채우는 작업
	//   (2) 마지막 요소 잘라내기
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2 // 삭제할 인덱스

	// 마지막 인덱스가 오면, for문 조건이 성립되지 않으므로 루프를 돌지 않는다.
	for i := idx + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}

	fmt.Println(slice) // [1 2 4 5 6 6]

	slice = slice[:len(slice)-1]
	fmt.Println(slice) // [1 2 4 5 6]

	// - append()로 코드 개선하기
	slice = []int{1, 2, 3, 4, 5, 6}
	slice = append(slice[:idx], slice[idx+1:]...) // 슬라이싱할 때, 마지막 인덱스는 포함하지 않는 것을 이용
	fmt.Println(slice)
}
