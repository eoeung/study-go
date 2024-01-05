package main

import "fmt"

func main() {
	// [중첩 for문과 break, 레이블]
	//   - break는 속해있는 for문만 빠져나온다.
	//   - 모든 for문을 나가고 싶을 때, 사용하는 방법

	// 1) bool 변수 사용
	var a, b int = 1, 1
	found := false

	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
