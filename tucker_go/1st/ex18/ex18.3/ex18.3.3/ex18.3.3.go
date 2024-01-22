package main

import "fmt"

func main() {
	// [append() 사용 시, 발생할 수 있는 문제 1]
	// - 슬라이스에 빈 공간이 있는 경우
	//   1) 슬라이스에 값을 추가할 수 있는 빈 공간이 있는지 확인 → cap - len (배열 길이 - 요소 개수)
	//   2) 남은 빈 공간의 개수 >= 추가하는 값의 개수 → 배열 뒷 부분에 값을 추가한 뒤 len값 증가

	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5)

	fmt.Println("slice1 :", slice1, len(slice1), cap(slice1)) // slice1 : [0 0 0] 3 5
	fmt.Println("slice2 :", slice2, len(slice2), cap(slice2)) // slice2 : [0 0 0 4 5] 5 5
	// 1) append()는 '새로운 슬라이스를 반환'한다. → slice1은 아무런 영향도 없음
	// 2) append()는 slice1에 4, 5를 추가할 만한 공간이 있는지를 찾는다.
	// 3) append()는 slice1에 값을 추가할 공간이 있으면, 값을 추가하고 len을 증가시킨 slice 구조체를 반환한다.
	// 4) 이 slice 구조체는 포인터, len, cap값이 다음과 같다.
	//   - 포인터 : slice1의 실제 배열을 가리키는 포인터
	//     → 어떻게 같은 배열을 가리키는데 가져오는 값이 다를까?
	//     => len값을 이용해서 실제 배열에서 인덱스만큼 잘라서 가지고 오는 것 같음
	//   - len : 5 (4, 5 라는 값을 추가했으므로)
	//   - cap : 5

	slice1[1] = 100
	fmt.Println()

	fmt.Println("slice1 :", slice1, len(slice1), cap(slice1)) // slice1 : [0 100 0] 3 5
	fmt.Println("slice2 :", slice2, len(slice2), cap(slice2)) // slice2 : [0 100 0 4 5] 5 5

	slice1 = append(slice1, 500)

	fmt.Println()

	fmt.Println("After append 500")
	fmt.Println("slice1 :", slice1, len(slice1), cap(slice1)) // slice1 : [0 100 0 500] 4 5
	fmt.Println("slice2 :", slice2, len(slice2), cap(slice2)) // slice2 : [0 100 0 500 5] 5 5
	// 1) append()는 첫 번째 인자인 slice1에서 빈 공간을 찾는다.
	// 2) slice1에는 빈 공간이 2개 있다. (cap - len = 5 - 3)
	// 3) 3번 인덱스가 비어있으므로 500을 추가한다.
	// 4) slice1과 slice2는 같은 배열을 가리키는 포인터를 가지고 있다.
}
