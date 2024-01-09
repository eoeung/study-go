package main

import (
	"fmt"
	"strings"
)

// [문자열 합산]
// - strings 패키지의 Builder를 이용해 메모리 낭비를 줄이는 방법

// 1) string 합 연산
func ToUpper1(str string) string {
	var rst string
	for _, c := range str { // 문자열을 하나씩 순회 → byte값
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}
	return rst
}

// 2) strings.Builder 사용
func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	var str string = "Hello World"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}
