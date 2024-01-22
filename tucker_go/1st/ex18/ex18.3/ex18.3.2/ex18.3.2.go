package main

import "fmt"

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	// [슬라이스와 배열의 동작 차이]
	// - 슬라이스 내부 구현이 배열과 다르기 때문에, 사용법이 비슷하다고 해서 똑같이 사용하면 안됨

	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println("array :", array) // array : [1 2 3 4 5]
	fmt.Println("slice :", slice) // slice : [1 2 200 4 5]
	// 1) Go에서 모든 값의 대입 = 복사
	// 2) 함수에 인수로 전달될 때, 다른 변수에 대입, 값의 이동 = 복사
	// 3) 복사 = 타입의 값이 복사됨
	//   - 포인터 = 메모리 주소가 복사됨
	//   - 구조체 = 구조체의 모든 필드 복사
	//   - 배열 = 배열의 모든 값 복사
}
