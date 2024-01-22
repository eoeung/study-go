package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// [문자열 합산]

	var str string = "Hello"
	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	addr1 := stringheader.Data

	str += " World!"
	addr2 := stringheader.Data

	str += " Welcome!"
	addr3 := stringheader.Data

	fmt.Println(str)                  // Hello World! Welcome!
	fmt.Printf("addr1:\t%x\n", addr1) // addr1:  b50aef
	fmt.Printf("addr2:\t%x\n", addr2) // addr2:  c00000a0e0
	fmt.Printf("addr3:\t%x\n", addr3) // addr3:  c000010120
	// 1) 기존 문자열 메모리 공간을 건드리지 않고, 새로운 메모리 공간을 만들어서 문자열을 합친다.
	// 2) string 합 연산 이후, 주소값이 변경된다.
	// 3) 문자열 불변 원칙이 준수된다.
	// → 하지만, 문자열 합 연산을 자주하게 된다면 메모리가 낭비된다.

}
