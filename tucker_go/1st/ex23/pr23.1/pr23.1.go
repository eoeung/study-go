package main

import "fmt"

func Atoi(str string) (int, error) {
	rst := 0
	for _, r := range str {
		if r >= '0' && r <= '9' {
			rst *= 10
			rst += int(r - '0')
		} else {
			return 0, fmt.Errorf("숫자만 입력하세요. 문자 : %c", r)
		}
	}

	return rst, nil
}

func main() {
	n, err := Atoi("34cd")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}

	fmt.Println()
	//////////////////////////////////////////////////////////////////
	fmt.Printf("#### %T ::: %v\n", '3'-'0', '3'-'0') // int32, 3
	// 문자열 상수간의 연산은 해당 문자의 ASCII값 간의 연산 결과를 반환
	// → '3'은 ASCII 값으로 51, '0'은 ASCII 값으로 48
	// => 51 - 48
	fmt.Printf("#### %T ::: %v\n", int('3'-'0'), int('3'-'0')) // int, 3
}
