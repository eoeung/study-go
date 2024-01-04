package main

import "fmt"

func main() {
	// [상수]
	// - 절대 변하면 안되는 값
	// - 기본 자료형만 가능
	//   (int, float, string, complex, bool, rune)

	const C int = 10

	var b int = C * 20
	C = 20             // cannot assign to C (상수는 대입문 좌변에 올 수 없음)
	fmt.Println(b, &C) // invalid operation: cannot take address of C (상수의 메모리 주소값에는 접근할 수 없음)

}
