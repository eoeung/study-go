package main

import "fmt"

func main() {
	// [배열]
	//   - 같은 타입의 데이터들로 이루어짐
	//   - 최대 요소 개수가 고정됨

	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}

	/*
		24 // 포맷을 지정하지 않으면, 실수는 최소 소수점 자리수로 표시됨
		25.9
		27.8
		26.9
		26.2
	*/
}
