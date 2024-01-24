package main

import "fmt"

func main() {
	// [Scan()]
	// func Scan(a ...any) (n int, err error)
	// - 값을 채워넣을 변수들의 메모리 주소를 인자로 받음
	// - 한 번에 여러 값을 입력하면, 변수 사이를 공란으로 구분 ('Enter'도 공란으로 인식)
	// - n: 성공적으로 입력한 값 개수

	var a, b int
	n, err := fmt.Scan(&a, &b)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
