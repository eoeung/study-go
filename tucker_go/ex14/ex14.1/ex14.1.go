package main

import "fmt"

type User struct{}

func main() {
	// 1. 포인터 변수 선언
	var p *int // 포인터 변수는 데이터 타입 앞에 *를 붙여 선언
	var pp *float64
	var ppp *User

	fmt.Printf("%T :: %T :: %T\n", p, pp, ppp) // *int :: *float64 :: *main.User

	// 2. 포인터 변수에 값 대입
	var a int
	var t *int
	t = &a  // a의 메모리 주소를, 포인터 변수 t에 대입
	*t = 20 // 포인터 변수 t가 가리키는 메모리 공간의 값을 20으로 변경
	fmt.Printf("%T :: %v\n", t, t)
}
