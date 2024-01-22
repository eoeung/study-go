package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	var p1 *int = &a // a의 메모리 주소를 p1 포인터 변수에 대입
	var p2 *int = &a
	var p3 *int = &b

	// == 연산으로 같은 메모리 주소값을 가지고 있는지 확인
	fmt.Printf("p1 == p2 : %v\n", p1 == p2) // true
	fmt.Printf("p2 == p3 : %v\n", p2 == p3) // false

}
