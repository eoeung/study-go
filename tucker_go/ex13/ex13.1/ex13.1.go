package main

import "fmt"

// 학생 구조체
type Student struct {
	Name  string
	Class int
	No    int
	Score float64
}

// 집 구조체
type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	// [구조체]
	// - 여러 필드를 묶어서 하나의 구조체를 만듦
	// - 서로 다른 타입의 값을 변수 하나로 묶어서 관리

	var a Student
	fmt.Printf("Student ::: %v\n", a)

	var house House
	house.Address = "서울시 강동구"
	house.Size = 28
	house.Price = 9.8
	house.Type = "아파트"

	fmt.Printf("주소:%s\n", house.Address)
	fmt.Printf("크기: %d평\n", house.Size)
	fmt.Printf("가격:%.2f억 원\n", house.Price) // %.2f → 소수점 2자리까지 출력
	fmt.Printf("타입:%s\n", house.Type)
}
