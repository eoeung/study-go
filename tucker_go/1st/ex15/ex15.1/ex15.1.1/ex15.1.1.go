package main

import "fmt"

func main() {
	// [문자열]
	// - Go에서는 UTF-8 문자코드 사용
	// - 큰 따옴표(") / 백쿼트(`)

	// 1) 특수 문자 처리 여부
	str1 := "Hello\t'World\n"
	str2 := `Go is "awesome"!\nGo is simple and \t'Powerful'`

	fmt.Println(str1) // Hello   'World
	fmt.Println(str2) // Go is "awesome"!\nGo is simple and \t'Powerful'
}
