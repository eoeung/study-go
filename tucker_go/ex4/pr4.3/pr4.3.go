package main

import "fmt"

func main() {
	var f1 float32 = 123.546789 * 345.678
	var f2 float32 = float32(123.546789) * 345.678

	fmt.Println(f1) // 42707.406
	fmt.Println(f2) // 42707.41

	// f1은 연산을 한 다음 값을 float32로 만들고,
	// f2는 123.546789를 float32로 변환한 다음, 345.678을 구한 값을 다시 float32로 만든다.
	fmt.Println(float32(123.546789)) // 123.54679
}
