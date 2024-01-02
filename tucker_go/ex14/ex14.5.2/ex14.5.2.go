package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) { // 매개변수로 Data 포인터를 받음
	arg.value = 200
	arg.data[100] = 999
}

func main() {
	// 포인터 사용 예제
	// var data *Data // panic: runtime error: invalid memory address or nil pointer dereference
	// → 초기화하지 않아서, 기본값인 nil이 할당됨
	var data Data

	ChangeData(&data) // data의 메모리 주소를 인수로 전달 → 8바이트
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}
