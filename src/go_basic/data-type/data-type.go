package main

import "fmt"

func main() {
	// Raw String Literal
	// - 복수라인 가능 / ``(Backtick 사용)
	rawLiteral := `아리랑\n
	
	아리랑\n
	아라리요`

	// Interpreted String Literal
	// - 해석된 문자열 리터럴
	// - 복수라인 불가능 → +로 이어서 작성가능
	interLiteral := "아리랑아리랑\n아라리요"
	interLiteral2 := "아리랑아리랑\n" +
		"아라리요"

	fmt.Println(rawLiteral)
	fmt.Println(interLiteral)
	fmt.Println(interLiteral2)

	// 데이터 타입 변환 T(v)
	// - T: 변환할 타입, v: 값
	// - 명시적으로 무조건 지정해줘야함
	str := `ABC
	DEFG
	  
	HIJK`
	bytes := []byte(str)
	str2 := string(bytes)
	fmt.Println(bytes)
	fmt.Println(str2)
}
