package main

import "fmt"

func main() {
	str := "Hello 월드"
	runes := []rune(str)

	fmt.Println(runes)                          // [72 101 108 108 111 32 50900 46300]
	fmt.Printf("len(str) = %d\n", len(str))     // 12 → 영어+공백 = 1x6개 = 6바이트 + 한글 = 3x2개 = 12바이트
	fmt.Printf("len(runes) = %d\n", len(runes)) // 8 → 모든 글자 수 = 8개
}
