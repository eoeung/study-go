package main

import "fmt"

func main() {
	// 예제로 배우는 GO 프로그래밍
	// [클로저]
	// - 함수 바깥에 있는 변수를 참조하는 함수값
	// - 함수는 바깥의 변수를 마치 함수 안으로 끌어들인 것 처럼 변수를 읽거나 쓸 수 있게 됨
	next := nextValue() // nextValue() 함수가 리턴하는, 익명함수를 가지고 있게됨
	fmt.Println(next()) // 1 // next()는 익명함수이므로 함수처럼 호출해야함
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3

	anotherNext := nextValue()
	fmt.Println(anotherNext()) // 1 // 다시 시작
	fmt.Println(anotherNext()) // 2

	// go by example
}

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
