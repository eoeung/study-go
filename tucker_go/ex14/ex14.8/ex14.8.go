package main

import "fmt"

type Data struct {
	value int
}

func main() {
	// 인스턴스를 만들면서, 포인터 값을 초기화하는 방법
	// 1) & 사용
	p1 := &Data{}

	// 2) new 사용
	var p2 = new(Data)

	fmt.Printf("%p ::: %p", p1, p2)
}

func TestFunc() {
	d := &Data{} // d 포인터 변수 선언, Data 인스턴스 생성
	d.value = 100
	fmt.Println(d)
}
