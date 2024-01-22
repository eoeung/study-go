package main

import "fmt"

type Data struct {
	value int
}

func main() {
	// 인스턴스
	var data Data // Data 타입 값을 저장할 수 있는 메모리 공간 할당 → 할당된 메모리 공간의 실체
	var p *Data = &data
	fmt.Printf("%p\n", p)

	var p1 *Data = &Data{}
	var p2 *Data = p1
	var p3 *Data = p2

	fmt.Printf("%p :: %p :: %p\n", p1, p2, p3)

	// data1 ~ data3은 값은 같지만 서로 다른 인스턴스임 (값이 복사되어 할당됨)
	var data1 *Data
	var data2 *Data = data1
	var data3 *Data = data2
	fmt.Printf("%p :: %p :: %p\n", data1, data2, data3)
}
