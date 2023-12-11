package main

import "fmt"

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return // 이미 위에서 return할 값을 정해줬기 때문에, return 키워드만 적어도 됨
	}

	result = a / b
	success = true
	return result, success // 이미 위에서 return할 값을 정해줬기 때문에, return 키워드만 적어도 됨
	// 작동에는 문제가 없으나 result, success는 불필요한 코드임
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)

	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
