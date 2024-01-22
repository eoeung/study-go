package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "Hello World!"
	str2 := str1

	// 1) *reflect.StringHeader로 변환할 예정
	// 2) string → *reflect.StringHeader로 바로 타입 변환은 불가능
	// 3) unsafe.Pointer()를 이용해서 unsafe.Pointer 타입으로 변환한 다음, *reflect.StringHeader 타입으로 변환
	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1) // &{16840909 12}
	fmt.Println(stringHeader2) // &{16840909 12}

	// ∴ string
}
