package main

import "fmt"

func main() {
	str1 := "안녕하세요. 한글 문자열입니다."
	str2 := str1

	fmt.Printf(str1) // 안녕하세요. 한글 문자열입니다.
	fmt.Printf("\n")
	fmt.Printf(str2) // 안녕하세요. 한글 문자열입니다.

	// 1) string은 reflect 패키지의 StringHeader 구조체를 통해 추측 가능
	// 2) Data, Len이라는 필드를 가짐
	// 3) 구조체는 크기만큼 복사된다. → 각 필드값만 복사됨 (Data 포인터 값, Len값)
	// 4) 복사된 필드값을 str2에 저장
	// 5) str1, str2 모두 같은 메모리를 가리킴 → Data 포인터 값이 실제 문자열을 가리키고 있으므로
}

// reflect 패키지의 StringHeader 구조체
type StringHeader struct {
	Data uintptr // 문자열의 데이터가 있는 메모리 주소 → 포인터
	Len  int
}
