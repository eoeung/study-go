package main

import "fmt"

func main() {
	// 예제로 배우는 GO 프로그래밍
	var t [3]int
	var tt [5]int
	fmt.Printf("%T, %v\n", t, t)   // [3]int, [0 0 0]
	fmt.Printf("%T, %v\n", tt, tt) // [5]int, [0 0 0 0 0]
	// → 타입이 다른 것을 확인할 수 있음
	// z := [5]int // Error : [5]int (type) is not an expression // 초기값 지정해줘야함
	zz := [5]int{}  // 초기값을 설정하지 않았지만, 배열의 데이터타입인 int의 기본값을 저장
	fmt.Println(zz) // [0 0 0 0 0]

	// 배열 크기를 자동으로 설정
	x := [...]int{}
	xx := [...]int{1, 2, 3, 4, 5}
	fmt.Println(x)  // []
	fmt.Println(xx) // [1 2 3 4 5]

	// go by example
	var a [5]int   // 5개의 정수를 가짐
	fmt.Println(a) // [0 0 0 0 0]

	a[4] = 100
	fmt.Println(a) // [0 0 0 0 100]
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	// 2차원 배열
	var twoD [2][3]int // [row][column]
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD) // [[0 1 2] [1 2 3]]
}
