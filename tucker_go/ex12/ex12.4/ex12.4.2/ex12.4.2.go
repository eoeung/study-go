package main

import "fmt"

const Y int = 3

func main() {
	// [range]
	//   - 배열로 range를 통해 순회하면 idx, val 2가지 값을 반환한다.

	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i, v := range t {
		fmt.Println(i, v)
	}

	/*
		0 24
		1 25.9
		2 27.8
		3 26.9
		4 26.2
	*/
}
