package main

import "fmt"

func main() {
	// [문자열은 불변]
	// - 불변 : 'string타입이 가리키는 문자열'의 일부분을 수정할 수 없음

	var str string = "Hello World"
	str = "How are you?"
	// str[2] = 'Z' // Error : cannot assign to str[2]

	var slice []byte = []byte(str)
	slice[2] = 'Z'

	fmt.Println(str)          // How are you?
	fmt.Printf("%s\n", slice) // HoZ are you?
	// 1) 문자열은 결국 byte 배열 → []byte로 타입 변환하면, 같은 메모리 공간을 가리킬 것?
	// 2) 하지만 str과 slice의 결과는 다르게 나왔다.
}
