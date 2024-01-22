package main

import "fmt"

func main() {
	// [다중 배열]
	// - 배열을 요소로 가지는 배열

	// Ex) 이중 배열 초기화
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	// b := [2][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}} // → 닫는 중괄호 '}'가 마지막 요소와 같은 줄에 있으면 ',' 생략

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
