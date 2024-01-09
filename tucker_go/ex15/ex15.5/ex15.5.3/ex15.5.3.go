package main

import "fmt"

func main() {
	// [문자열 대소 비교하기]
	// >, >=, <, <=
	// - 첫 글자부터 하나씩 값을 비교 → 유니코드 값이 다를 경우, 대소를 반환
	// - 값이 같을 경우, 다음 글자 비교

	str1 := "BBB"
	str2 := "aaaaAAA"
	str3 := "BBAD"
	str4 := "ZZZ"

	fmt.Printf("%s > %s : %v\n", str1, str2, str1 > str2)
	fmt.Printf("%s < %s : %v\n", str1, str3, str1 < str3)
	fmt.Printf("%s <= %s : %v\n", str1, str4, str1 <= str4)

	// 문자열 길이에 상관없이, 앞 글자부터 비교
}
