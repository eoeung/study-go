package main

import "fmt"

const Y int = 3

func main() {
	// [배열 요소 개수는 무조건 상수]

	x := 5
	// a := [x]int{1, 2, 3, 4, 5} // Error : invalid array length x
	b := [Y]int{1, 2, 3}
	var c [10]int

	fmt.Println(x, b, c)
}
