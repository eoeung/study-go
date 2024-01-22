package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg Data) {
	arg.value = 200
	arg.data[100] = 999
}

func main() {
	// 포인터를 사용하지 않는 예제
	var data Data

	ChangeData(data) // int + [200]int 모두 복사 → 1,608 바이트 복사됨
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}
