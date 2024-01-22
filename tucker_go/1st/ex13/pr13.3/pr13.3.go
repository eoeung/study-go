package main

import (
	"fmt"
	"unsafe"
)

// 기존 Padding 구조체
// type Padding struct {
// 	A int8    // 1바이트
// 	B int     // 8바이트
// 	C float64 // 8바이트
// 	D uint16  // 2바이트
// 	E int     // 8바이트
// 	F float32 // 4바이트
// 	G int8    // 1바이트
// }

type Padding struct {
	A int8    // 1바이트
	G int8    // 1바이트
	D uint16  // 2바이트
	F float32 // 4바이트
	B int     // 8바이트
	C float64 // 8바이트
	E int     // 8바이트
}

func main() {
	padding := Padding{}
	fmt.Println(unsafe.Sizeof(padding))
	// 32바이트
}
