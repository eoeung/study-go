package main

import "fmt"

func main() {
	// [rune 타입]
	// - UTF-8은 1~3바이트를 사용
	// - 3바이트 정수 타입이 Go에서는 없기 때문에, rune타입을 사용
	// - type rune int32 // rune 타입은 4바이트 정수타입인 int32의 별칭

	var char rune = '한'

	fmt.Printf("%T\n", char) // int32
	fmt.Println(char)        // 54620
	fmt.Printf("%c\n", char) // 한

}
