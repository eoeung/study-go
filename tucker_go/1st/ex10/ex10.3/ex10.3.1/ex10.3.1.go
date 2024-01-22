package main

import "fmt"

func main() {
	// [switch문의 다양한 형태]

	// 1) 하나의 case에서 여러 값을 비교하기
	//   - 하나의 case에서 여러 값을 ',(comma)'로 구분해서 검사할 수 있음
	day := "thursday"

	switch day {
	case "monday", "tuesday":
		fmt.Println("월, 화요일은 수업 가는 날입니다.")
	case "wednesday", "thursday", "friday":
		fmt.Println("수, 목, 금요일은 실습 가는 날입니다.")
	}
}
