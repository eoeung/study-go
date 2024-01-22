package main

import (
	"ch16/ex16.2/publicpkg"
	"fmt"
)

func main() {
	// [패키지명과 패키지 외부 공개]
	// - 대문자 : 외부로 공개
	// - 소문자 : 비공개
	fmt.Println("PI :", publicpkg.PI)
	publicpkg.PublicFunc()

	var myint publicpkg.MyInt = 10
	fmt.Println("myint:", myint)

	var mystruct = publicpkg.MyStruct{Age: 18}
	fmt.Println("mystruct:", mystruct)
}
