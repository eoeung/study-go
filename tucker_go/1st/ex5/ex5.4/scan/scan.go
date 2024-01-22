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
}
