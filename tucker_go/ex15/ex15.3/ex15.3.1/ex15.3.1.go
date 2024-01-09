package main

import "fmt"

func main() {
	// [len()으로 문자열 크기 알아내기]
	// - 문자 수가 아니라, 문자열이 차지하는 메모리 크기를 반환

	str1 := "가나다라마"
	str2 := "abcde"

	fmt.Printf("len(str1) = %d\n", len(str1)) // 15
	fmt.Printf("len(str2) = %d\n", len(str2)) // 5
}
