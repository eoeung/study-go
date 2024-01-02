package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func main() {
	// 바로 포인터 변수에 구조체 생성 → 주소를 초기값으로 대입
	// 1) 기존 방식
	var data Data       // Data타입 구조체 변수 data를 선언
	var p *Data = &data // Data타입의 포인터 변수인 p를 선언, data 변수 주소를 반환

	// 2) 구조체를 생성해 초기화하는 방식
	var t *Data = &Data{} // *Data타입 구조체 변수 t를 선언, Data구조체를 만들어 주소를 반환
	fmt.Printf("## p %p ::: ## t %p", p, t)
}
