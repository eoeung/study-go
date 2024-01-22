package main

import "fmt"

func main() {
	// [슬라이싱]
	// - 슬라이싱을 하면 슬라이스를 반환
	//   → 새로운 배열이 만들어지는 것이 아니라, 배열의 일부를 포인터로 가리키는 슬라이스를 만들어낸다.
	// - 끝 인덱스 - 1 까지의 값을 가지고 온다.

	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]

	fmt.Println("array :", array)                         // [1 2 3 4 5]
	fmt.Println("slice :", slice, len(slice), cap(slice)) // [2] 1 4
	fmt.Printf("%p\n", slice)                             // 0xc000012338
	// array를 슬라이싱하면, cap 길이는 'len(array) - 슬라이싱 시작 인덱스'가 된다.

	array[1] = 100

	fmt.Println()
	fmt.Println("After change second element")
	fmt.Println("array :", array)                         // [1 100 3 4 5]
	fmt.Println("slice :", slice, len(slice), cap(slice)) // [100] 1 4
	fmt.Printf("%p\n", slice)                             // 0xc000012338

	slice = append(slice, 500)

	fmt.Println()
	fmt.Println("After append 500")
	fmt.Println("array :", array)                         // [1 100 500 4 5]
	fmt.Println("slice :", slice, len(slice), cap(slice)) // [100 500] 2 4
	fmt.Printf("%p\n", slice)                             // 0xc000012338

	// cap 길이 확인
	// - len(array) - 슬라이싱 시작 인덱스 = 5 - 2 = 3
	slice333 := array[2:5]

	fmt.Println()
	fmt.Println("array :", array)                                     // [1 100 500 4 5]
	fmt.Println("slice333 :", slice333, len(slice333), cap(slice333)) // [500 4 5] 3 3
	fmt.Printf("%p\n", slice333)                                      // 0xc000012340
	// 1) 배열의 중간 메모리 주소값을 가지고 있다. (슬라이싱 시작 인덱스)
	// 2) len만큼 값을 가지고 온다.
	// 3) cap은 배열 길이 - 슬라이싱 시작 인덱스

	slice999 := array[4:5]

	fmt.Println()
	fmt.Println("array :", array)                                     // [1 100 500 4 5]
	fmt.Println("slice999 :", slice999, len(slice999), cap(slice999)) // [5] 1 1
	fmt.Printf("%p\n", slice999)                                      // 0xc000012350
	// 1) 배열의 중간 메모리 주소값을 가지고 있다. (슬라이싱 시작 인덱스)
	// 2) len만큼 값을 가지고 온다.
	// 3) cap은 배열 길이 - 슬라이싱 시작 인덱스
}
