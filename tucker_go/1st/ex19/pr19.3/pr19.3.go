package main

import (
	"fmt"
	"strings"
)

type myString string

func (m myString) ToLower() myString {
	str := strings.ToLower(string(m))
	return myString(str)
}

func (m myString) ToUpper() myString {
	str := strings.ToUpper(string(m))
	return myString(str)
}

func main() {
	msg := myString("hello Go World")

	msg2 := msg.ToLower()
	fmt.Println(msg2)

	msg3 := msg.ToUpper()
	fmt.Println(msg3)
	// 1) 값 타입 메서드라 메서드 내부에서 복사된 myString 타입을 가진다.
	// 2) ToLower() / ToUpper() 함수를 적용한 결과값 자체를 myString으로 타입 변환해서 보낸다.
	// 3) ∴ 함수가 적용된 결과값이 반환된다.
}
