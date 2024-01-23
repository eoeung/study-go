package main

import "fmt"

func main() {
	// [변수]
	// 변수 : 값을 저장하는 메모리 공간을 가리킴
	// - 변수를 이용해서 쉽고 효과적으로 메모리에 있는 데이터를 조작
	// - 프로그램은 컴퓨터 입장에서, 메모리에 있는 데이터를 언제/어떻게 변경할지를 나타낸 문서

	var a int = 10                    // a 변수 선언
	var msg string = "Hello Variable" // msg 변수 선언

	// 변수를 이용해서 메모리에 있는 데이터를 조작
	a = 20
	msg = "Good Morning"

	fmt.Println(msg, a)
}
