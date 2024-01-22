package main

import "fmt"

func main() {
	// [Scan()]
	// - 값을 채워넣을 변수들의 메모리 주소를 인수로 받음
	// 성공적으로 입력한 값 개수, 에러 2가지를 반환받음

	var a int
	var b int

	n, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}

	/*
		// 정상 입력
		3 4
		2 3 4 // 입력값 개수, a값, b값

		// 비정상 입력 1
		Hello 4
		0 expected integer // 첫 번째 입력에서 에러 발생 → 두 번째 입력 받지 않음 => 0

		// 비정상 입력 2
		4 Hello
		1 expected integer // 두 번째 입력에서 에러 발생 → 첫 번째 입력은 받았음 => 1
	*/
}
