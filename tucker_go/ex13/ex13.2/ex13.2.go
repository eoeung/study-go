package main

import "fmt"

// 집 구조체
type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	// [구조체 변수 초기화]
	//   1) 초기값 생략
	//   2) 모든 필드 초기화
	//   3) 일부 필드 초기화

	// 1) 초기값 생략
	var house11 House

	// 2) 모든 필드 초기화
	//   (1) 한 줄로
	var house21 House = House{"서울시 강동구", 28, 9.8, "아파트"}

	//   (2) 여러 줄로
	var house22 House = House{
		"서울시 강동구",
		28,
		9.8,
		"아파트",
	}

	// 3) 일부 필드 초기화
	// - 필드명 : 값
	// 초기화되지 않은 나머지 값들은 기본값이 할당됨
	//   (1) 한 줄로
	var house31 House = House{Size: 28, Type: "아파트"}

	//   (2) 여러 줄로
	var house32 House = House{
		Size: 28,
		Type: "아파트",
	}

	fmt.Printf("%v ::: \n", house11)
	fmt.Printf("%v ::: \n", house21)
	fmt.Printf("%v ::: \n", house22)
	fmt.Printf("%v ::: \n", house31)
	fmt.Printf("%v ::: \n", house32)
}
