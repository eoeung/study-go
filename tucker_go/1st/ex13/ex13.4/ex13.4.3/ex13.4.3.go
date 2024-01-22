package main

import (
	"fmt"
	"unsafe"
)

type User1 struct {
	A int8 // 1바이트
	B int  // 8바이트
	C int8 // 1바이트
	D int  // 8바이트
	E int8 // 1바이트
}

// User1 구조체 필드의 총 바이트는 19바이트
// 하지만, 실제 구조체 크기는 메모리 패딩 때문에 40바이트가 됨
// → 1바이트 변수 A,C,E 모두에 7바이트씩 패딩이 됐기 때문

type User2 struct {
	A int8 // 1바이트
	C int8 // 1바이트
	E int8 // 1바이트
	B int  // 8바이트
	D int  // 8바이트
}

// User2 구조체 필드의 총 바이트도 역시 19바이트
// 실제 구조체 크기는 24바이트

func main() {
	// [구조체 크기]
	// 3) 메모리 패딩을 고려한 필드 배치 방법
	user11 := User1{}
	fmt.Println(unsafe.Sizeof(user11)) // 40

	user21 := User2{}
	fmt.Println(unsafe.Sizeof(user21)) // 24
}
