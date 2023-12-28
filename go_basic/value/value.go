package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// 문자열에서 range 사용
	// 문자열에서 range는 유니코드 코드 포인트들을 순회
	// 첫번째 값 : rune의 시작 바이트 위치 → index 같은 개념인듯
	// 두번째 값 : rune값 자체
	for i, c := range "GO" {
		fmt.Println(i, c)
	}
	/*
		0 71
		1 79
	*/

	for i, c := range "Hello World" {
		fmt.Println(i, c)
	}
	/*
		0 72
		1 101
		2 108
		3 108
		4 111
		5 32
		6 87
		7 111
		8 114
		9 108
		10 100
	*/
}
