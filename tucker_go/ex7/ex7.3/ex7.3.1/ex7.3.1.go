package main

import (
	"fmt"
)

// [멀티 반환 함수]
func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}