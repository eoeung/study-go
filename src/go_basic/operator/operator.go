package main

import "fmt"

func main() {
	a := 1
	// 전위연산자는 지원하지 않음
	// ++a
	// a = ++a
	a++
	fmt.Println(a)

	// 포인터연산자
	k := 10    // var k int = 10
	var p = &k // k의 메모리 주소 할당
	var s = &p

	// * = 메모리 주소가 가리키는 값 출력 or (구조체, 슬라이스 등) 타입의 포인터를 가리키는 포인터 변수
	// & = 메모리 주소
	// k의 값, k의 메모리 주소
	fmt.Println(k, &k) // 10 0xc000096090
	// fmt.Println(*k) // k는 10이라는 값이므로, 메모리 주소가 아님

	// p = k의 메모리 주소, &p = p를 가리키는 다른 주소, *p = p가 가리키는 값
	fmt.Println(p, &p, *p) // 0xc000096090 0xc0000aa020 10

	// s = p의 메모리 주소, &s = s를 가리키는 다른 주소, *s = s가 가리키는 값 → p가 가리키는 값 → k의 메모리 주소
	fmt.Println(s, &s, *s) // 0xc0000aa020 0xc0000aa028 0xc000096090
}
