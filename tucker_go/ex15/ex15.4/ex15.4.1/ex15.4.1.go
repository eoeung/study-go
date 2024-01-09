package main

import "fmt"

func main() {
	// [문자열 순회]
	// 1) 인덱스 사용한 바이트 단위 순회
	// 2) []rune으로 타입 변환 후, 한 글자씩 순회
	// 3) range 키워드 이용

	// 1) 인덱스 사용한 '바이트 단위' 순회
	str := "Hello 월드!"

	for i := 0; i < len(str); i++ {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i])
	}

	/*
		타입:uint8 값:72 문자값:H
		타입:uint8 값:101 문자값:e
		타입:uint8 값:108 문자값:l
		타입:uint8 값:108 문자값:l
		타입:uint8 값:111 문자값:o
		타입:uint8 값:32 문자값:
		타입:uint8 값:236 문자값:ì
		타입:uint8 값:155 문자값:
		타입:uint8 값:148 문자값:
		타입:uint8 값:235 문자값:ë
		타입:uint8 값:147 문자값:
		타입:uint8 값:156 문자값:
		타입:uint8 값:33 문자값:!
	*/
	// 1바이트인 영문, 공백은 잘 나오지만 3바이트인 한글은 깨져서 나옴
}
