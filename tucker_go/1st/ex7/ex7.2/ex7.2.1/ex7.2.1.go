package main

import "fmt"

func main() {
	// 코드가 반복된다. → 함수를 사용하면 편리하다.
	math := 80
	eng := 74
	history := 95
	fmt.Println("김일등 님 평균 점수는", (math+eng+history)/3, "입니다.")

	math = 88
	eng = 92
	history = 53
	fmt.Println("송이등 님 평균 점수는", (math+eng+history)/3, "입니다.")

	math = 78
	eng = 73
	history = 78
	fmt.Println("박삼등 님 평균 점수는", (math+eng+history)/3, "입니다.")
}
