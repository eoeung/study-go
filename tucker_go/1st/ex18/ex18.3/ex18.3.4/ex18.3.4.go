package main

import "fmt"

func main() {
	// [append() 사용 시, 발생할 수 있는 문제 2]
	// - 슬라이스에 빈 공간이 없는 경우
	//   1) append()는 빈 공간이 충분하지 않은 경우, 새로운 더 큰 배열을 생성
	//     → 보통 기존 배열의 2배
	//   2) 기존 배열의 요소를 모두 새로운 배열에 복사
	//   3) 새로운 배열의 맨 뒤에 값을 추가
	//     - 포인터 : 새로운 배열을 가리키는 슬라이스 구조체 반환
	//     - len : 기존 길이에 추가한 개수만큼 더한 값
	//     - cap : 새로운 배열의 길이

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	fmt.Println("slice1 :", slice1, len(slice1), cap(slice1)) // slice1 : [1 2 3] 3 3
	fmt.Println("slice2 :", slice2, len(slice2), cap(slice2)) // slice2 : [1 2 3 4 5] 5 6
	// 1) append()는 slice1에 빈 공간이 없는 것을 확인 한다.
	// 2) slice1에 새로운 값을 추가하기 위해 새로운 배열을 만든다.
	// 3) append()는 생성한 새로운 배열에 값을 추가하고, 이에 대한 슬라이스 구조체를 반환한다.
	// → slice1과 slice2의 포인터는 다른 배열을 가리킨다.

	slice1[1] = 100

	fmt.Println()
	fmt.Println("After change second element")
	fmt.Println("slice1 :", slice1, len(slice1), cap(slice1)) // slice1 : [1 100 3] 3 3
	fmt.Println("slice2 :", slice2, len(slice2), cap(slice2)) // slice2 : [1 2 3 4 5] 5 6

	slice1 = append(slice1, 500)

	fmt.Println()
	fmt.Println("After append 500")
	fmt.Println("slice1 :", slice1, len(slice1), cap(slice1)) // slice1 : [1 100 3 500] 4 6
	fmt.Println("slice2 :", slice2, len(slice2), cap(slice2)) // slice2 : [1 2 3 4 5] 5 6
}
