package main

import "fmt"

func main() {
	// [배열 복사]
	// - 대입 연산자 사용

	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()

	// 복사하기 전, b
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	fmt.Println()

	b = a // a배열을 b배열에 복사
	// 복사한 후, b
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	/*
		a[0] = 1
		a[1] = 2
		a[2] = 3
		a[3] = 4
		a[4] = 5

		b[0] = 500
		b[1] = 400
		b[2] = 300
		b[3] = 200
		b[4] = 100

		b[0] = 1
		b[1] = 2
		b[2] = 3
		b[3] = 4
		b[4] = 5
	*/

	x := [5]int{1, 2, 3, 4, 5}
	y := [5]float32{500, 400, 300, 200, 100}

	// y = x // Error : cannot use x (variable of type [5]int) as [5]float32 value in assignment
	// 배열끼리 타입이 다르므로, 복사가 불가능
	// → 배열간 대입 = 전체 배열의 복사
	fmt.Println(x, y)
}
