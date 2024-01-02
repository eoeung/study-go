package main

import "fmt"

func add(p1, p2, p3 *int) {
	*p3 = *p2 + *p1
}

func main() {
	a := 3
	b := 5
	c := 0

	add(&a, &b, &c)
	fmt.Println(c) // 8 → 함수의 리턴값이 없지만, 포인터 주소를 넘겨서 값을 변경
}
