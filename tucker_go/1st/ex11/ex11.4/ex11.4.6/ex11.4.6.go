package main

import "fmt"

func find45(a int) (int, bool) {
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true
		}
	}

	return 0, false
}

func main() {
	// [중첩 for문과 break, 레이블 → 클린 코드로 지향]
	//   - 플래그 변수나 레이블 사용을 최소화
	//   - 중첩된 내부 로직을 함수로 묶어 중첩을 줄여야함

	var a, b int = 1, 0

	for ; a <= 9; a++ {
		var found bool
		if b, found = find45(a); found {
			break
		}
	}

	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
