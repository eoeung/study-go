package main

import "fmt"

func main() {
	// [문자열 순회]
	// 1) 인덱스 사용한 바이트 단위 순회
	// 2) []rune으로 타입 변환 후, 한 글자씩 순회
	// 3) range 키워드 이용

	// 3) range 키워드 이용
	str := "Hello 월드!"
	for _, v := range str { // range를 이용해서 추가적인 메모리 할당 없이 문자열을 한 글자씩 순회
		fmt.Printf("타입: %T 값:%d 문자:%c\n", v, v, v) // → 한 글자 => rune = int32 = 숫자잢
	}

	/*
		타입: int32 값:72 문자:H
		타입: int32 값:101 문자:e
		타입: int32 값:108 문자:l
		타입: int32 값:108 문자:l
		타입: int32 값:111 문자:o
		타입: int32 값:32 문자:
		타입: int32 값:50900 문자:월
		타입: int32 값:46300 문자:드
		타입: int32 값:33 문자:!
	*/
}
