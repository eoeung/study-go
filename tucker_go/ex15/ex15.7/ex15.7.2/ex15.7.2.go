package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// [실제 주소값 출력]
	// - 문자열과 []byte의 주소값이 달랐다.

	var str string = "Hello World"
	var slice []byte = []byte(str)

	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceheader := (*reflect.StringHeader)(unsafe.Pointer(&slice))

	// %x : 16진수 변환 (a~f는 소문자로)
	fmt.Printf("str:\t%x\n", stringheader)  // str:    &{3c157a b}
	fmt.Printf("slice:\t%x\n", sliceheader) // slice:  &{c000112080 b}
	// 1) 서로 다른 주소를 가리키고 있음
	// 2) slice로 타입 변환을 할 때, 문자열을 복사
	// 3) 새로운 메모리 공간을 만들어서 슬라이스가 가리키도록 한다.
	// → 문자열 불변 원칙을 지킬 수 있음
}
