package main

import "fmt"

func main() {
	// [맵]
	// - key & value
	// - 언어에 따라서 딕셔너리, 해시테이블, 해시맵으로 불린다.
	// - 리스트나 링과 다르게 container패키지가 아닌 Go 기본 내장 타입

	m := make(map[string]string)
	m["이화랑"] = "서울시 광진구"
	m["송하나"] = "서울시 강남구"
	m["백두산"] = "부산시 사하구"
	m["최번개"] = "전주시 덕진구"

	m["최번개"] = "청주시 상당구"

	fmt.Printf("송하나의 주소는 %s입니다.\n", m["송하나"])
	fmt.Printf("백두산의 주소는 %s입니다.\n", m["백두산"])
}
